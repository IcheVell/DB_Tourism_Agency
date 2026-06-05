import { useEffect, useMemo, useState } from 'react';
import { getApiErrorMessage } from '../../api/http';
import {
    getMyAccommodations,
    getMyCargo,
    getMyExcursions,
    getMyTours,
    getMyVisas,
} from '../../api/meApi';
import { DataTable } from '../../components/common/DataTable';

type MeListType = 'tours' | 'visas' | 'accommodations' | 'excursions' | 'cargo';

interface MeListPageProps {
    type: MeListType;
}

const pageConfig = {
    tours: {
        title: 'Мои поездки',
        loader: getMyTours,
    },
    visas: {
        title: 'Мои визы',
        loader: getMyVisas,
    },
    accommodations: {
        title: 'Моё расселение',
        loader: getMyAccommodations,
    },
    excursions: {
        title: 'Мои экскурсии',
        loader: getMyExcursions,
    },
    cargo: {
        title: 'Мой груз',
        loader: getMyCargo,
    },
} satisfies Record<MeListType, {
    title: string;
    loader: (page: number, limit: number) => Promise<{
        items?: Record<string, unknown>[] | null;
        total?: number | null;
        page?: number | null;
        limit?: number | null;
    }>;
}>;

export function MeListPage({ type }: MeListPageProps) {
    const config = pageConfig[type];

    const [items, setItems] = useState<Record<string, unknown>[]>([]);
    const [total, setTotal] = useState(0);
    const [page, setPage] = useState(1);
    const [limit] = useState(20);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const totalPages = useMemo(() => {
        return Math.max(1, Math.ceil(total / limit));
    }, [limit, total]);

    useEffect(() => {
        async function loadData() {
            setLoading(true);
            setError(null);

            try {
                const result = await config.loader(page, limit);

                setItems(Array.isArray(result.items) ? result.items : []);
                setTotal(Number(result.total ?? 0));
            } catch (caughtError) {
                setItems([]);
                setTotal(0);
                setError(getApiErrorMessage(caughtError));
            } finally {
                setLoading(false);
            }
        }

        void loadData();
    }, [config, limit, page]);

    return (
        <section className="page-section">
            <h1>{config.title}</h1>

            {loading && <div className="page-state">Загрузка...</div>}

            {error && <div className="error-box">{error}</div>}

            {!loading && !error && (
                <>
                    <DataTable rows={items} />

                    <div className="pagination">
                        <button
                            type="button"
                            disabled={page <= 1}
                            onClick={() => setPage((current) => Math.max(1, current - 1))}
                        >
                            Назад
                        </button>

                        <span>
              Страница {page} из {totalPages}
            </span>

                        <button
                            type="button"
                            disabled={page >= totalPages}
                            onClick={() => setPage((current) => Math.min(totalPages, current + 1))}
                        >
                            Вперёд
                        </button>
                    </div>
                </>
            )}
        </section>
    );
}