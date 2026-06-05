import { useEffect, useState } from 'react';
import { listResource } from '../../api/crudApi';
import { getApiErrorMessage } from '../../api/http';
import { createMyExcursionBooking } from '../../api/meApi';
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

    const canBook = hasPermission('own_excursions.create');

    async function loadData() {
        setLoading(true);
        setError(null);

        try {
            const result = await listResource(resources.excursionSchedules.endpoint, page, limit);

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

    useEffect(() => {
        void loadData();
    }, [page, limit]);

    async function handleBook(scheduleID: number) {
        setBookingID(scheduleID);
        setError(null);
        setSuccess(null);

        try {
            await createMyExcursionBooking({
                excursion_schedule_id: scheduleID,
            });

            setSuccess('Вы записаны на экскурсию');
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
                disabled={bookingID === Number(item.id)}
                onClick={() => handleBook(Number(item.id))}
            >
                {bookingID === Number(item.id) ? 'Запись...' : 'Записаться'}
            </button>
        ) : '—',
    }));

    return (
        <section className="page-section">
            <h1>Расписание экскурсий</h1>

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