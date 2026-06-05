import { FormEvent, useState } from 'react';
import { getReport } from '../../api/reportsApi';
import { getApiErrorMessage } from '../../api/http';
import { useAuth } from '../../app/AuthContext';
import { DataTable } from '../../components/common/DataTable';
import { reportConfigs, type ReportConfig } from '../../config/reports';
import type { JsonObject } from '../../types/common';

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

function buildReportPath(config: ReportConfig, params: Record<string, string>): string {
  if (config.key === 'tourist-info') {
    return `${config.path}/${params.tourist_id}`;
  }

  return config.path;
}

export function ReportsPage() {
  const { hasPermission } = useAuth();
  const availableReports = reportConfigs.filter((report) => hasPermission(report.permission));

  const [selectedKey, setSelectedKey] = useState(availableReports[0]?.key ?? '');
  const [params, setParams] = useState<Record<string, string>>({});
  const [result, setResult] = useState<unknown>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const selectedReport = availableReports.find((report) => report.key === selectedKey);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    if (!selectedReport) return;

    for (const param of selectedReport.params) {
      if (param.required && !params[param.name]) {
        setError(`Заполните поле: ${param.label}`);
        return;
      }
    }

    setLoading(true);
    setError(null);

    try {
      const path = buildReportPath(selectedReport, params);
      const data = await getReport(path, params);
      setResult(data);
    } catch (caughtError) {
      setError(getApiErrorMessage(caughtError));
    } finally {
      setLoading(false);
    }
  }

  return (
    <section className="page-section">
      <h1>Отчёты</h1>

      {availableReports.length === 0 && <div className="empty-state">Нет доступных отчётов</div>}

      {availableReports.length > 0 && (
        <form className="report-form" onSubmit={handleSubmit}>
          <label>
            Отчёт
            <select
              value={selectedKey}
              onChange={(event) => {
                setSelectedKey(event.target.value);
                setParams({});
                setResult(null);
              }}
            >
              {availableReports.map((report) => (
                <option key={report.key} value={report.key}>{report.title}</option>
              ))}
            </select>
          </label>

          {selectedReport?.params.map((param) => (
            <label key={param.name}>
              {param.label}
              <input
                type={param.type}
                value={params[param.name] ?? ''}
                required={param.required}
                onChange={(event) => setParams((current) => ({ ...current, [param.name]: event.target.value }))}
              />
            </label>
          ))}

          <button type="submit" disabled={loading}>{loading ? 'Формирование...' : 'Сформировать'}</button>
        </form>
      )}

      {error && <div className="error-box">{error}</div>}
      {result !== null && <DataTable rows={toRows(result)} />}
    </section>
  );
}
