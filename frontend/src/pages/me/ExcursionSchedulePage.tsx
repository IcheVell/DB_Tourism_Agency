import { useEffect, useMemo, useState } from 'react';
import axios from 'axios';
import { listResource } from '../../api/crudApi';
import { getApiErrorMessage } from '../../api/http';
import {
  createMyExcursionBooking,
  getMyIdentityDocument,
  getMyTours,
} from '../../api/meApi';
import { useAuth } from '../../app/AuthContext';
import { DataTable } from '../../components/common/DataTable';
import { Pagination } from '../../components/common/Pagination';
import { resources } from '../../config/resources';

export function ExcursionSchedulePage() {
  const { hasPermission } = useAuth();

  const [items, setItems] = useState<Record<string, unknown>[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [limit] = useState(20);
  const [loading, setLoading] = useState(true);
  const [bookingID, setBookingID] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const [tours, setTours] = useState<Record<string, unknown>[]>([]);
  const [selectedGroupMemberID, setSelectedGroupMemberID] = useState('');
  const [hasDocument, setHasDocument] = useState(false);
  const [documentChecked, setDocumentChecked] = useState(false);

  const canBook = hasPermission('own_excursions.create');

  const bookingBlockedReason = useMemo(() => {
    if (!canBook) {
      return 'У вашей роли нет права на запись на экскурсии.';
    }

    if (!documentChecked) {
      return 'Проверяется документ.';
    }

    if (!hasDocument) {
      return 'Перед записью на экскурсию заполните раздел «Мои документы».';
    }

    if (tours.length === 0) {
      return 'Запись доступна после того, как менеджер добавит вас в туристическую группу.';
    }

    if (tours.length > 1 && !selectedGroupMemberID) {
      return 'Выберите поездку / группу для записи.';
    }

    return null;
  }, [canBook, documentChecked, hasDocument, selectedGroupMemberID, tours.length]);

  async function loadData() {
    setLoading(true);
    setError(null);

    try {
      const [scheduleResult, tourResult] = await Promise.all([
        listResource(resources.excursionSchedules.endpoint, page, limit),
        getMyTours(1, 100),
      ]);

      const loadedTours = Array.isArray(tourResult.items) ? tourResult.items : [];

      setItems(Array.isArray(scheduleResult.items) ? scheduleResult.items : []);
      setTotal(Number.isFinite(scheduleResult.total) ? scheduleResult.total : 0);
      setTours(loadedTours);
      setSelectedGroupMemberID((current) => {
        if (current) {
          return current;
        }

        return loadedTours.length === 1 ? String(loadedTours[0].group_member_id ?? '') : '';
      });
    } catch (caughtError) {
      setItems([]);
      setTotal(0);
      setTours([]);
      setError(getApiErrorMessage(caughtError));
    } finally {
      setLoading(false);
    }
  }

  async function checkDocument() {
    setDocumentChecked(false);

    try {
      await getMyIdentityDocument();
      setHasDocument(true);
    } catch (caughtError) {
      if (axios.isAxiosError(caughtError) && caughtError.response?.status === 404) {
        setHasDocument(false);
        return;
      }

      setError(getApiErrorMessage(caughtError));
    } finally {
      setDocumentChecked(true);
    }
  }

  useEffect(() => {
    void loadData();
  }, [page, limit]);

  useEffect(() => {
    void checkDocument();
  }, []);

  async function handleBook(scheduleID: number) {
    if (bookingBlockedReason) {
      setError(bookingBlockedReason);
      return;
    }

    setBookingID(scheduleID);
    setError(null);
    setSuccess(null);

    try {
      const groupMemberID = selectedGroupMemberID ? Number(selectedGroupMemberID) : undefined;

      await createMyExcursionBooking({
        excursion_schedule_id: scheduleID,
        group_member_id: groupMemberID,
      });

      setSuccess('Вы записаны на экскурсию. Запись появится в разделе «Мои экскурсии».');
    } catch (caughtError) {
      setError(getApiErrorMessage(caughtError));
    } finally {
      setBookingID(null);
    }
  }

  const rows = items.map((item) => ({
    ...item,
    action: canBook ? (
      <button
        type="button"
        className="primary-button"
        disabled={Boolean(bookingBlockedReason) || bookingID === Number(item.id)}
        onClick={() => handleBook(Number(item.id))}
      >
        {bookingID === Number(item.id) ? 'Запись...' : 'Записаться'}
      </button>
    ) : '—',
  }));

  return (
    <section className="page-section">
      <h1>Расписание экскурсий</h1>

      <div className="card">
        <h2>Запись на экскурсию</h2>
        <p>
          Записаться можно только после заполнения документа и добавления туриста в группу.
          Если у вас несколько поездок, выберите нужную группу перед записью.
        </p>

        {tours.length > 1 && (
          <label>
            Поездка / группа
            <select
              value={selectedGroupMemberID}
              onChange={(event) => setSelectedGroupMemberID(event.target.value)}
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
      </div>

      {bookingBlockedReason && <div className="warning-box">{bookingBlockedReason}</div>}
      {success && <div className="success-box">{success}</div>}
      {error && <div className="error-box">{error}</div>}
      {loading && <div className="page-state">Загрузка...</div>}

      {!loading && (
        <>
          <DataTable rows={rows} />

          <Pagination
            page={page}
            limit={limit}
            total={total}
            onPageChange={setPage}
          />
        </>
      )}
    </section>
  );
}
