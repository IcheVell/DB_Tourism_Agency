import { FormEvent, useEffect, useState } from 'react';
import { getReport } from '../../api/reportsApi';
import { getApiErrorMessage } from '../../api/http';
import { listResource } from '../../api/crudApi';
import { useAuth } from '../../app/AuthContext';
import { DataTable } from '../../components/common/DataTable';
import {
  reportConfigs,
  type ReportConfig,
  type ReportParamConfig,
  type ReportSelectOption,
} from '../../config/reports';
import type { JsonObject } from '../../types/common';

type ReportParams = Record<string, string>;
type ReportOptionMap = Record<string, ReportSelectOption[]>;

function toRows(value: unknown): JsonObject[] {
  if (Array.isArray(value)) {
    return value as JsonObject[];
  }

  if (value && typeof value === 'object') {
    const objectValue = value as Record<string, unknown>;

    const firstArray = Object.values(objectValue).find(Array.isArray);
    if (Array.isArray(firstArray)) {
      return firstArray as JsonObject[];
    }

    return [objectValue];
  }

  return [];
}

function buildReportPath(config: ReportConfig, params: ReportParams): string {
  if (config.key === 'tourist-info') {
    return `${config.path}/${params.tourist_id}`;
  }

  return config.path;
}

function buildOptionLabel(item: Record<string, unknown>, labelFields: string[]): string {
  const parts = labelFields
    .map((field) => item[field])
    .filter((value) => value !== null && value !== undefined && String(value).trim() !== '')
    .map((value) => String(value));

  if (parts.length === 0) {
    return String(item.id ?? 'Без названия');
  }

  return parts.join(' · ');
}

function emptyParamsFor(report?: ReportConfig): ReportParams {
  if (!report) {
    return {};
  }

  return Object.fromEntries(report.params.map((param) => [param.name, '']));
}

export function ReportsPage() {
  const { hasPermission } = useAuth();
  const availableReports = reportConfigs.filter((report) => hasPermission(report.permission));

  const [selectedKey, setSelectedKey] = useState(availableReports[0]?.key ?? '');
  const [params, setParams] = useState<ReportParams>(() => emptyParamsFor(availableReports[0]));
  const [options, setOptions] = useState<ReportOptionMap>({});
  const [result, setResult] = useState<unknown>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [optionsLoading, setOptionsLoading] = useState(false);

  const selectedReport = availableReports.find((report) => report.key === selectedKey);

  useEffect(() => {
    if (selectedKey || availableReports.length === 0) {
      return;
    }

    setSelectedKey(availableReports[0].key);
    setParams(emptyParamsFor(availableReports[0]));
  }, [availableReports, selectedKey]);

  useEffect(() => {
    async function loadOptions(report: ReportConfig) {
      setOptionsLoading(true);
      setOptions({});

      const entries = await Promise.all(
        report.params.map(async (param): Promise<[string, ReportSelectOption[]]> => {
          if (param.type !== 'select') {
            return [param.name, []];
          }

          if (param.select?.staticOptions) {
            return [param.name, param.select.staticOptions];
          }

          if (!param.select?.endpoint) {
            return [param.name, []];
          }

          try {
            const response = await listResource(param.select.endpoint, 1, 1000);
            const valueField = param.select.valueField ?? 'id';
            const labelFields = param.select.labelFields ?? ['name'];
            const loadedOptions = response.items
              .map((item) => {
                const value = item[valueField];
                if (value === null || value === undefined) {
                  return null;
                }

                return {
                  value: typeof value === 'number' ? value : String(value),
                  label: buildOptionLabel(item, labelFields),
                };
              })
              .filter((item): item is ReportSelectOption => item !== null);

            return [param.name, loadedOptions];
          } catch {
            return [param.name, []];
          }
        }),
      );

      setOptions(Object.fromEntries(entries));
      setOptionsLoading(false);
    }

    if (!selectedReport) {
      setOptions({});
      return;
    }

    void loadOptions(selectedReport);
  }, [selectedReport]);

  function updateParam(name: string, value: string) {
    setParams((current) => ({
      ...current,
      [name]: value,
    }));
  }

  function validateRequiredParams(report: ReportConfig): string | null {
    for (const param of report.params) {
      if (param.required && !params[param.name]) {
        return `Заполните поле: ${param.label}`;
      }
    }

    if (params.from && params.to) {
      const from = Date.parse(params.from);
      const to = Date.parse(params.to);

      if (Number.isFinite(from) && Number.isFinite(to) && from > to) {
        return 'Дата начала не должна быть позже даты окончания';
      }
    }

    return null;
  }

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    if (!selectedReport) {
      return;
    }

    const validationError = validateRequiredParams(selectedReport);
    if (validationError) {
      setError(validationError);
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const path = buildReportPath(selectedReport, params);
      const data = await getReport(path, params);
      setResult(data);
    } catch (caughtError) {
      setResult(null);
      setError(getApiErrorMessage(caughtError));
    } finally {
      setLoading(false);
    }
  }

  function renderParamInput(param: ReportParamConfig) {
    const value = params[param.name] ?? '';

    if (param.type === 'select') {
      const paramOptions = options[param.name] ?? [];

      return (
        <label key={param.name} className="form-field">
          {param.label}
          <select
            value={value}
            required={param.required}
            disabled={optionsLoading}
            onChange={(event) => updateParam(param.name, event.target.value)}
          >
            <option value="">{param.required ? 'Выберите значение' : 'Все'}</option>
            {paramOptions.map((option) => (
              <option key={`${param.name}-${option.value}`} value={String(option.value)}>
                {option.label}
              </option>
            ))}
          </select>
        </label>
      );
    }

    return (
      <label key={param.name} className="form-field">
        {param.label}
        <input
          type={param.type}
          value={value}
          required={param.required}
          onChange={(event) => updateParam(param.name, event.target.value)}
        />
      </label>
    );
  }

  return (
    <section className="page-section">
      <h1>Отчёты</h1>

      {availableReports.length === 0 && <div className="empty-state">Нет доступных отчётов</div>}

      {availableReports.length > 0 && (
        <form className="report-form card" onSubmit={handleSubmit}>
          <label className="form-field">
            Отчёт
            <select
              value={selectedKey}
              onChange={(event) => {
                const nextKey = event.target.value;
                const nextReport = availableReports.find((report) => report.key === nextKey);
                setSelectedKey(nextKey);
                setParams(emptyParamsFor(nextReport));
                setResult(null);
                setError(null);
              }}
            >
              {availableReports.map((report) => (
                <option key={report.key} value={report.key}>
                  {report.title}
                </option>
              ))}
            </select>
          </label>

          <div className="report-params-grid">
            {selectedReport?.params.map(renderParamInput)}
          </div>

          {optionsLoading && <div className="page-state">Загрузка списков выбора...</div>}

          <button type="submit" className="primary-button" disabled={loading || optionsLoading}>
            {loading ? 'Формирование...' : 'Сформировать'}
          </button>
        </form>
      )}

      {error && <div className="error-box">{error}</div>}
      {result !== null && <DataTable rows={toRows(result)} />}
    </section>
  );
}
