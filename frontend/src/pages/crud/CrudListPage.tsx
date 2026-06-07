import { FormEvent, useEffect, useMemo, useState } from 'react';
import type { ReactNode } from 'react';
import { useParams } from 'react-router-dom';
import { createResource, listResource, updateResource } from '../../api/crudApi';
import { getApiErrorMessage } from '../../api/http';
import { useAuth } from '../../app/AuthContext';
import { DataTable } from '../../components/common/DataTable';
import { Pagination } from '../../components/common/Pagination';
import { resources, type ResourceConfig, type ResourceFieldConfig, type ResourceSelectOption } from '../../config/resources';

interface ModalState {
  mode: 'create' | 'update';
  rowID?: number | string;
}

type FormValues = Record<string, string>;
type OptionMap = Record<string, ResourceSelectOption[]>;

function valueToFormString(value: unknown, field?: ResourceFieldConfig): string {
  if (value === null || value === undefined) {
    return '';
  }

  if (field?.type === 'datetime-local') {
    const raw = String(value);
    if (raw.includes('T')) {
      return raw.replace('Z', '').slice(0, 16);
    }
    return raw;
  }

  if (field?.type === 'date') {
    return String(value).slice(0, 10);
  }

  return String(value);
}

function toApiDateTime(value: string): string {
  if (!value) {
    return value;
  }

  if (value.endsWith('Z')) {
    return value;
  }

  if (/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}$/.test(value)) {
    return `${value}:00Z`;
  }

  if (/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$/.test(value)) {
    return `${value}Z`;
  }

  return value;
}

function formStringToPayloadValue(value: string, field: ResourceFieldConfig): unknown {
  const trimmed = typeof value === 'string' ? value.trim() : value;

  if (trimmed === '') {
    return field.nullable ? null : '';
  }

  if (field.type === 'number') {
    const parsed = Number(trimmed);
    return Number.isFinite(parsed) ? parsed : trimmed;
  }

  if (field.type === 'datetime-local') {
    return toApiDateTime(trimmed);
  }

  if (field.type === 'select') {
    const staticOptions = field.select?.staticOptions ?? [];
    const option = staticOptions.find((item) => String(item.value) === trimmed);

    if (typeof option?.value === 'number') {
      return Number(trimmed);
    }

    if (/^\d+$/.test(trimmed) && !field.select?.staticOptions) {
      return Number(trimmed);
    }

    return trimmed;
  }

  return trimmed;
}

function buildInitialFormValues(resource: ResourceConfig, source?: Record<string, unknown>): FormValues {
  const values: FormValues = {};
  const fields = resource.formFields ?? [];
  const template = resource.createTemplate ?? {};

  fields.forEach((field) => {
    const sourceValue = source?.[field.name] ?? template[field.name];
    values[field.name] = valueToFormString(sourceValue, field);
  });

  return values;
}

function buildPayload(resource: ResourceConfig, values: FormValues): Record<string, unknown> {
  const payload: Record<string, unknown> = {};
  const fields = resource.formFields ?? [];

  fields.forEach((field) => {
    payload[field.name] = formStringToPayloadValue(values[field.name] ?? '', field);
  });

  return payload;
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

function getRowID(row: Record<string, unknown>): number | string | null {
  const id = row.id ?? row.ID;
  if (typeof id === 'number' || typeof id === 'string') {
    return id;
  }
  return null;
}

function parseComparableDate(value: string): number | null {
  if (!value) {
    return null;
  }

  const parsed = Date.parse(value);
  return Number.isFinite(parsed) ? parsed : null;
}

function validateDateOrder(values: FormValues, from: string, to: string, message: string): string | null {
  const fromValue = parseComparableDate(values[from] ?? '');
  const toValue = parseComparableDate(values[to] ?? '');

  if (fromValue === null || toValue === null) {
    return null;
  }

  if (fromValue >= toValue) {
    return message;
  }

  return null;
}

function validateForm(resource: ResourceConfig, values: FormValues): string | null {
  const fields = resource.formFields ?? [];

  for (const field of fields) {
    const value = values[field.name] ?? '';
    if (field.required && !field.nullable && value.trim() === '') {
      return `Заполните поле: ${field.label}`;
    }
  }

  const checks: Record<string, Array<[string, string, string]>> = {
    touristGroups: [['arrival_date', 'departure_date', 'Дата вылета должна быть позже даты прилёта']],
    accommodations: [['check_in_at', 'check_out_at', 'Дата выезда должна быть позже даты заселения']],
    excursionSchedules: [['start_time', 'end_time', 'Время окончания экскурсии должно быть позже времени начала']],
    identityDocuments: [['issue_date', 'expiration_date', 'Дата окончания действия документа должна быть позже даты выдачи']],
    visas: [
      ['submitted_at', 'decision_at', 'Дата решения по визе должна быть позже даты подачи'],
      ['decision_at', 'issued_at', 'Дата выдачи визы должна быть позже даты решения'],
      ['valid_from', 'valid_until', 'Дата окончания действия визы должна быть позже даты начала'],
    ],
  };

  for (const [from, to, message] of checks[resource.key] ?? []) {
    const error = validateDateOrder(values, from, to, message);
    if (error) {
      return error;
    }
  }

  if (resource.key === 'visas' && values.status === 'issued') {
    if (!values.number || !values.issued_at || !values.valid_from || !values.valid_until) {
      return 'Для выданной визы нужны номер, дата выдачи, начало и окончание действия';
    }
  }

  return null;
}

export function CrudListPage() {
  const { resourceKey } = useParams();
  const { hasPermission } = useAuth();

  const resource = resourceKey ? resources[resourceKey] : undefined;

  const [items, setItems] = useState<Record<string, unknown>[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [modal, setModal] = useState<ModalState | null>(null);
  const [formValues, setFormValues] = useState<FormValues>({});
  const [formError, setFormError] = useState<string | null>(null);
  const [saving, setSaving] = useState(false);
  const [selectOptions, setSelectOptions] = useState<OptionMap>({});

  const title = useMemo(() => resource?.label ?? 'Раздел не найден', [resource]);
  const canCreate = Boolean(resource?.createPermission && hasPermission(resource.createPermission));
  const canUpdate = Boolean(resource?.updatePermission && hasPermission(resource.updatePermission));

  async function loadData() {
    if (!resource) {
      setItems([]);
      setTotal(0);
      setLoading(false);
      return;
    }

    setLoading(true);
    setError(null);

    try {
      const result = await listResource(resource.endpoint, page, limit);
      setItems(Array.isArray(result.items) ? result.items : []);
      setTotal(Number.isFinite(result.total) ? result.total : 0);
    } catch (caughtError) {
      setItems([]);
      setTotal(0);
      setError(getApiErrorMessage(caughtError));
    } finally {
      setLoading(false);
    }
  }

  async function loadOptions(currentResource: ResourceConfig) {
    const fields = currentResource.formFields ?? [];
    const optionEntries = await Promise.all(
      fields.map(async (field): Promise<[string, ResourceSelectOption[]]> => {
        if (field.type !== 'select') {
          return [field.name, []];
        }

        if (field.select?.staticOptions) {
          return [field.name, field.select.staticOptions];
        }

        if (!field.select?.endpoint) {
          return [field.name, []];
        }

        try {
          const result = await listResource(field.select.endpoint, 1, 1000);
          const valueField = field.select.valueField ?? 'id';
          const labelFields = field.select.labelFields ?? ['name'];
          const options = result.items
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
            .filter((item): item is ResourceSelectOption => item !== null);

          return [field.name, options];
        } catch {
          return [field.name, []];
        }
      }),
    );

    setSelectOptions(Object.fromEntries(optionEntries));
  }

  useEffect(() => {
    setPage(1);
    setModal(null);
    setFormError(null);
  }, [resourceKey]);

  useEffect(() => {
    void loadData();
  }, [resource, page, limit]);

  useEffect(() => {
    if (!resource) {
      setSelectOptions({});
      return;
    }

    void loadOptions(resource);
  }, [resource]);

  function openCreateForm() {
    if (!resource) {
      return;
    }

    setFormValues(buildInitialFormValues(resource));
    setFormError(null);
    setModal({ mode: 'create' });
  }

  function openUpdateForm(row: Record<string, unknown>) {
    if (!resource) {
      return;
    }

    const rowID = getRowID(row);
    if (rowID === null) {
      setError('Нельзя изменить запись без id');
      return;
    }

    setFormValues(buildInitialFormValues(resource, row));
    setFormError(null);
    setModal({ mode: 'update', rowID });
  }

  function closeModal() {
    if (saving) {
      return;
    }

    setModal(null);
    setFormError(null);
  }

  function updateField(fieldName: string, value: string) {
    setFormValues((current) => ({
      ...current,
      [fieldName]: value,
    }));
  }

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    if (!resource || !modal) {
      return;
    }

    const validationError = validateForm(resource, formValues);
    if (validationError) {
      setFormError(validationError);
      return;
    }

    setSaving(true);
    setFormError(null);

    try {
      const payload = buildPayload(resource, formValues);

      if (modal.mode === 'create') {
        await createResource(resource.endpoint, payload);
      } else {
        if (modal.rowID === undefined) {
          setFormError('Не найден id изменяемой записи');
          return;
        }
        await updateResource(resource.endpoint, modal.rowID, payload);
      }

      setModal(null);
      await loadData();
    } catch (caughtError) {
      setFormError(getApiErrorMessage(caughtError));
    } finally {
      setSaving(false);
    }
  }

  function renderField(field: ResourceFieldConfig): ReactNode {
    const value = formValues[field.name] ?? '';

    if (field.type === 'select') {
      const options = selectOptions[field.name] ?? [];

      return (
        <label key={field.name} className="form-field">
          {field.label}
          <select
            value={value}
            onChange={(event) => updateField(field.name, event.target.value)}
            required={field.required && !field.nullable}
          >
            {field.nullable && <option value="">Не выбрано</option>}
            {!field.nullable && value === '' && <option value="">Выберите значение</option>}
            {options.map((option) => (
              <option key={`${field.name}-${option.value}`} value={String(option.value)}>
                {option.label}
              </option>
            ))}
          </select>
        </label>
      );
    }

    return (
      <label key={field.name} className="form-field">
        {field.label}
        <input
          type={field.type === 'number' ? 'number' : field.type}
          value={value}
          onChange={(event) => updateField(field.name, event.target.value)}
          required={field.required && !field.nullable}
        />
      </label>
    );
  }

  const rows = items.map((item) => {
    if (!canUpdate || getRowID(item) === null) {
      return item;
    }

    return {
      ...item,
      actions: (
        <button type="button" className="secondary-button" onClick={() => openUpdateForm(item)}>
          Изменить
        </button>
      ),
    };
  });

  if (!resource) {
    return (
      <section className="page-section">
        <h1>Раздел не найден</h1>
        <div className="error-box">Неизвестный ресурс: {resourceKey}</div>
      </section>
    );
  }

  return (
    <section className="page-section">
      <div className="page-header">
        <div>
          <h1>{title}</h1>
          {resource.helperText && <p className="muted-text">{resource.helperText}</p>}
        </div>

        {canCreate && (
          <button type="button" className="primary-button" onClick={openCreateForm}>
            Добавить
          </button>
        )}
      </div>

      {loading && <div className="page-state">Загрузка...</div>}
      {error && <div className="error-box">{error}</div>}

      {!loading && !error && (
        <>
          <DataTable rows={rows} />
          <Pagination page={page} limit={limit} total={total} onPageChange={setPage} />
        </>
      )}

      {modal && (
        <div className="modal-backdrop">
          <div className="modal-card">
            <div className="modal-header">
              <h2>{modal.mode === 'create' ? 'Добавить' : 'Изменить'}: {resource.label}</h2>
              <button type="button" className="icon-button" onClick={closeModal}>×</button>
            </div>

            <form className="resource-form" onSubmit={handleSubmit}>
              {(resource.formFields ?? []).map(renderField)}

              {formError && <div className="error-box">{formError}</div>}

              <div className="modal-actions">
                <button type="button" onClick={closeModal} disabled={saving}>Отмена</button>
                <button type="submit" className="primary-button" disabled={saving}>
                  {saving ? 'Сохранение...' : modal.mode === 'create' ? 'Создать' : 'Сохранить'}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </section>
  );
}
