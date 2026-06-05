import { FormEvent, useEffect, useMemo, useState } from 'react';
import { useParams } from 'react-router-dom';
import { createResource, listResource } from '../../api/crudApi';
import { getApiErrorMessage } from '../../api/http';
import { useAuth } from '../../app/AuthContext';
import { DataTable } from '../../components/common/DataTable';
import { Pagination } from '../../components/common/Pagination';
import { resources } from '../../config/resources';

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

    const [createOpen, setCreateOpen] = useState(false);
    const [createJson, setCreateJson] = useState('{}');
    const [createError, setCreateError] = useState<string | null>(null);
    const [creating, setCreating] = useState(false);

    const title = useMemo(() => {
        return resource?.label ?? 'Раздел не найден';
    }, [resource]);

    const canCreate = Boolean(
        resource?.createPermission && hasPermission(resource.createPermission),
    );

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

    useEffect(() => {
        setPage(1);
    }, [resourceKey]);

    useEffect(() => {
        void loadData();
    }, [resource, page, limit]);

    function openCreateForm() {
        const template = resource?.createTemplate ?? {};
        setCreateJson(JSON.stringify(template, null, 2));
        setCreateError(null);
        setCreateOpen(true);
    }

    function closeCreateForm() {
        if (creating) {
            return;
        }

        setCreateOpen(false);
        setCreateError(null);
    }

    async function handleCreateSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();

        if (!resource) {
            return;
        }

        setCreating(true);
        setCreateError(null);

        try {
            const payload = JSON.parse(createJson);

            if (!payload || typeof payload !== 'object' || Array.isArray(payload)) {
                setCreateError('JSON должен быть объектом');
                return;
            }

            await createResource(resource.endpoint, payload as Record<string, unknown>);

            setCreateOpen(false);
            setCreateJson('{}');

            await loadData();
        } catch (caughtError) {
            if (caughtError instanceof SyntaxError) {
                setCreateError('Некорректный JSON');
            } else {
                setCreateError(getApiErrorMessage(caughtError));
            }
        } finally {
            setCreating(false);
        }
    }

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
                <h1>{title}</h1>

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
                    <DataTable rows={items} />

                    <Pagination
                        page={page}
                        limit={limit}
                        total={total}
                        onPageChange={setPage}
                    />
                </>
            )}

            {createOpen && (
                <div className="modal-backdrop">
                    <div className="modal-card">
                        <div className="modal-header">
                            <h2>Добавить: {resource.label}</h2>

                            <button type="button" className="icon-button" onClick={closeCreateForm}>
                                ×
                            </button>
                        </div>

                        <form onSubmit={handleCreateSubmit}>
                            <label>
                                JSON данные
                                <textarea
                                    className="json-textarea"
                                    value={createJson}
                                    onChange={(event) => setCreateJson(event.target.value)}
                                    rows={14}
                                    spellCheck={false}
                                />
                            </label>

                            {createError && <div className="error-box">{createError}</div>}

                            <div className="modal-actions">
                                <button type="button" onClick={closeCreateForm} disabled={creating}>
                                    Отмена
                                </button>

                                <button type="submit" className="primary-button" disabled={creating}>
                                    {creating ? 'Создание...' : 'Создать'}
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
        </section>
    );
}