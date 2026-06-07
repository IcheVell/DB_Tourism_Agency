import { FormEvent, useEffect, useMemo, useState } from 'react';
import {
  createMyCargo,
  getMyCargo,
  getMyCargoTypes,
  getMyTours,
  type MeCargoType,
} from '../../api/meApi';
import { getApiErrorMessage } from '../../api/http';
import { DataTable } from '../../components/common/DataTable';
import { Pagination } from '../../components/common/Pagination';

interface CargoFormState {
  group_member_id: string;
  cargo_type_id: string;
  item_number: string;
  places_count: string;
  weight_kg: string;
  volumetric_weight_kg: string;
}

const initialCargoForm: CargoFormState = {
  group_member_id: '',
  cargo_type_id: '',
  item_number: '',
  places_count: '1',
  weight_kg: '',
  volumetric_weight_kg: '0',
};

export function MyCargoPage() {
  const [items, setItems] = useState<Record<string, unknown>[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [loading, setLoading] = useState(true);

  const [cargoTypes, setCargoTypes] = useState<MeCargoType[]>([]);
  const [tours, setTours] = useState<Record<string, unknown>[]>([]);
  const [form, setForm] = useState<CargoFormState>(initialCargoForm);
  const [creating, setCreating] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const hasSeveralTours = tours.length > 1;

  const selectedCargoTypeID = useMemo(() => {
    if (form.cargo_type_id) {
      return form.cargo_type_id;
    }

    return cargoTypes.length > 0 ? String(cargoTypes[0].id) : '';
  }, [cargoTypes, form.cargo_type_id]);

  async function loadData() {
    setLoading(true);
    setError(null);

    try {
      const [cargoResult, cargoTypeResult, tourResult] = await Promise.all([
        getMyCargo(page, limit),
        getMyCargoTypes(),
        getMyTours(1, 100),
      ]);

      const loadedCargoTypes = Array.isArray(cargoTypeResult) ? cargoTypeResult : [];
      const loadedTours = Array.isArray(tourResult.items) ? tourResult.items : [];

      setItems(Array.isArray(cargoResult.items) ? cargoResult.items : []);
      setTotal(Number(cargoResult.total ?? 0));
      setCargoTypes(loadedCargoTypes);
      setTours(loadedTours);

      setForm((current) => ({
        ...current,
        cargo_type_id: current.cargo_type_id || (loadedCargoTypes[0] ? String(loadedCargoTypes[0].id) : ''),
        group_member_id: current.group_member_id || (loadedTours.length === 1 ? String(loadedTours[0].group_member_id ?? '') : ''),
      }));
    } catch (caughtError) {
      setItems([]);
      setTotal(0);
      setError(getApiErrorMessage(caughtError));
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    void loadData();
  }, [page, limit]);

  function updateField<K extends keyof CargoFormState>(field: K, value: CargoFormState[K]) {
    setForm((current) => ({
      ...current,
      [field]: value,
    }));
  }

  async function handleCreateCargo(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    setCreating(true);
    setError(null);
    setSuccess(null);

    try {
      const cargoTypeID = Number(selectedCargoTypeID);
      const placesCount = Number(form.places_count);
      const weightKg = Number(form.weight_kg);
      const volumetricWeightKg = Number(form.volumetric_weight_kg || 0);
      const groupMemberID = form.group_member_id ? Number(form.group_member_id) : undefined;

      if (!Number.isFinite(cargoTypeID) || cargoTypeID <= 0) {
        setError('Выберите вид груза');
        return;
      }

      if (!Number.isFinite(placesCount) || placesCount <= 0) {
        setError('Количество мест должно быть больше нуля');
        return;
      }

      if (!Number.isFinite(weightKg) || weightKg <= 0) {
        setError('Вес должен быть больше нуля');
        return;
      }

      if (!Number.isFinite(volumetricWeightKg) || volumetricWeightKg < 0) {
        setError('Объёмный вес не может быть отрицательным');
        return;
      }

      await createMyCargo({
        group_member_id: groupMemberID,
        cargo_type_id: cargoTypeID,
        item_number: form.item_number.trim(),
        places_count: placesCount,
        weight_kg: weightKg,
        volumetric_weight_kg: volumetricWeightKg,
      });

      setSuccess('Груз добавлен. Менеджер сможет взвесить, промаркировать, упаковать и отправить его.');
      setForm((current) => ({
        ...initialCargoForm,
        cargo_type_id: current.cargo_type_id,
        group_member_id: current.group_member_id,
      }));
      await loadData();
    } catch (caughtError) {
      setError(getApiErrorMessage(caughtError));
    } finally {
      setCreating(false);
    }
  }

  return (
    <section className="page-section">
      <h1>Мой груз</h1>

      <div className="card">
        <h2>Добавить груз</h2>
        <p className="muted-text">
          Турист добавляет сведения о грузе. Дальше менеджер представительства оформляет ведомость,
          маркирует, взвешивает, упаковывает и отправляет груз.
        </p>

        {tours.length === 0 && !loading && (
          <div className="warning-box">
            Вы ещё не добавлены в туристическую группу. Груз можно добавить после назначения в группу.
          </div>
        )}

        <form className="profile-form" onSubmit={handleCreateCargo}>
          {hasSeveralTours && (
            <label>
              Поездка / группа
              <select
                value={form.group_member_id}
                onChange={(event) => updateField('group_member_id', event.target.value)}
                required
              >
                <option value="">Выберите группу</option>
                {tours.map((tour) => (
                  <option key={String(tour.group_member_id)} value={String(tour.group_member_id)}>
                    {String(tour.group_name ?? `Группа #${tour.group_id ?? tour.group_member_id}`)}
                  </option>
                ))}
              </select>
            </label>
          )}

          <label>
            Вид груза
            <select
              value={selectedCargoTypeID}
              onChange={(event) => updateField('cargo_type_id', event.target.value)}
              required
            >
              {cargoTypes.length === 0 && <option value="">Нет видов груза</option>}
              {cargoTypes.map((cargoType) => (
                <option key={cargoType.id} value={cargoType.id}>
                  {cargoType.name}
                </option>
              ))}
            </select>
          </label>

          <label>
            Номер / описание места
            <input
              value={form.item_number}
              onChange={(event) => updateField('item_number', event.target.value)}
              placeholder="Например: сумка 1, коробка с одеждой"
            />
          </label>

          <label>
            Количество мест
            <input
              type="number"
              min="1"
              step="1"
              value={form.places_count}
              onChange={(event) => updateField('places_count', event.target.value)}
              required
            />
          </label>

          <label>
            Вес, кг
            <input
              type="number"
              min="0.001"
              step="0.001"
              value={form.weight_kg}
              onChange={(event) => updateField('weight_kg', event.target.value)}
              required
            />
          </label>

          <label>
            Объёмный вес, кг
            <input
              type="number"
              min="0"
              step="0.001"
              value={form.volumetric_weight_kg}
              onChange={(event) => updateField('volumetric_weight_kg', event.target.value)}
            />
          </label>

          <button type="submit" className="primary-button" disabled={creating || tours.length === 0 || cargoTypes.length === 0}>
            {creating ? 'Добавление...' : 'Добавить груз'}
          </button>
        </form>
      </div>

      {success && <div className="success-box">{success}</div>}
      {error && <div className="error-box">{error}</div>}
      {loading && <div className="page-state">Загрузка...</div>}

      {!loading && (
        <>
          <DataTable rows={items} />
          <Pagination page={page} limit={limit} total={total} onPageChange={setPage} />
        </>
      )}
    </section>
  );
}
