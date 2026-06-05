import { FormEvent, useEffect, useState } from 'react';
import axios from 'axios';
import {
    createMyIdentityDocument,
    getMyIdentityDocument,
    updateMyIdentityDocument,
    type MeIdentityDocument,
} from '../../api/meApi';
import { getApiErrorMessage } from '../../api/http';

interface FormState {
    document_type: string;
    document_series: string;
    document_number: string;
    citizenship: string;
}

const initialForm: FormState = {
    document_type: 'passport',
    document_series: '',
    document_number: '',
    citizenship: 'Россия',
};

export function MyDocumentPage() {
    const [document, setDocument] = useState<MeIdentityDocument | null>(null);
    const [form, setForm] = useState<FormState>(initialForm);
    const [loading, setLoading] = useState(true);
    const [missing, setMissing] = useState(false);
    const [saving, setSaving] = useState(false);
    const [error, setError] = useState<string | null>(null);
    const [success, setSuccess] = useState<string | null>(null);

    useEffect(() => {
        async function loadDocument() {
            setLoading(true);
            setError(null);

            try {
                const result = await getMyIdentityDocument();

                setDocument(result);
                setMissing(false);
                setForm({
                    document_type: result.document_type,
                    document_series: result.document_series ?? '',
                    document_number: result.document_number,
                    citizenship: result.citizenship,
                });
            } catch (caughtError) {
                if (axios.isAxiosError(caughtError) && caughtError.response?.status === 404) {
                    setDocument(null);
                    setMissing(true);
                    setForm(initialForm);
                    return;
                }

                setError(getApiErrorMessage(caughtError));
            } finally {
                setLoading(false);
            }
        }

        void loadDocument();
    }, []);

    function updateField<K extends keyof FormState>(field: K, value: FormState[K]) {
        setForm((current) => ({
            ...current,
            [field]: value,
        }));
    }

    async function handleSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();

        setSaving(true);
        setError(null);
        setSuccess(null);

        const payload = {
            document_type: form.document_type.trim(),
            document_series: form.document_series.trim() || null,
            document_number: form.document_number.trim(),
            citizenship: form.citizenship.trim(),
        };

        try {
            const result = missing
                ? await createMyIdentityDocument(payload)
                : await updateMyIdentityDocument(payload);

            setDocument(result);
            setMissing(false);
            setForm({
                document_type: result.document_type,
                document_series: result.document_series ?? '',
                document_number: result.document_number,
                citizenship: result.citizenship,
            });

            setSuccess(missing ? 'Документ создан' : 'Документ обновлён');
        } catch (caughtError) {
            setError(getApiErrorMessage(caughtError));
        } finally {
            setSaving(false);
        }
    }

    return (
        <section className="page-section">
            <h1>Мои документы</h1>

            {loading && <div className="page-state">Загрузка...</div>}

            {!loading && missing && (
                <div className="warning-box">
                    Документ ещё не заполнен. Без документа запись на экскурсии недоступна.
                </div>
            )}

            {error && <div className="error-box">{error}</div>}
            {success && <div className="success-box">{success}</div>}

            {!loading && (
                <form className="card profile-form" onSubmit={handleSubmit}>
                    <label>
                        Тип документа
                        <select
                            value={form.document_type}
                            onChange={(event) => updateField('document_type', event.target.value)}
                            required
                        >
                            <option value="passport">Паспорт</option>
                            <option value="international_passport">Заграничный паспорт</option>
                            <option value="birth_certificate">Свидетельство о рождении</option>
                        </select>
                    </label>

                    <label>
                        Серия
                        <input
                            value={form.document_series}
                            onChange={(event) => updateField('document_series', event.target.value)}
                        />
                    </label>

                    <label>
                        Номер
                        <input
                            value={form.document_number}
                            onChange={(event) => updateField('document_number', event.target.value)}
                            required
                        />
                    </label>

                    <label>
                        Гражданство
                        <input
                            value={form.citizenship}
                            onChange={(event) => updateField('citizenship', event.target.value)}
                            required
                        />
                    </label>

                    <button type="submit" className="primary-button" disabled={saving}>
                        {saving ? 'Сохранение...' : missing ? 'Создать документ' : 'Сохранить'}
                    </button>
                </form>
            )}

            {document && (
                <div className="card">
                    <h2>Текущий документ</h2>
                    <p><strong>ID:</strong> {document.id}</p>
                    <p><strong>Тип:</strong> {document.document_type}</p>
                    <p><strong>Серия:</strong> {document.document_series ?? '—'}</p>
                    <p><strong>Номер:</strong> {document.document_number}</p>
                    <p><strong>Гражданство:</strong> {document.citizenship}</p>
                </div>
            )}
        </section>
    );
}