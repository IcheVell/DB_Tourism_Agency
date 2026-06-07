import { FormEvent, useEffect, useMemo, useState } from 'react';
import { listResource } from '../../api/crudApi';
import { getApiErrorMessage } from '../../api/http';
import {
  getMyTravelRequest,
  updateMyTravelRequest,
  type MeTravelRequest,
} from '../../api/meApi';

interface HotelOption {
  id: number;
  name: string;
  address: string;
}

export function MyTravelRequestPage() {
  const [request, setRequest] = useState<MeTravelRequest | null>(null);
  const [hotels, setHotels] = useState<HotelOption[]>([]);
  const [desiredHotelID, setDesiredHotelID] = useState('');
  const [loading, setLoading] = useState(true);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  const selectedHotel = useMemo(() => {
    if (!desiredHotelID) {
      return null;
    }

    return hotels.find((hotel) => String(hotel.id) === desiredHotelID) ?? null;
  }, [desiredHotelID, hotels]);

  async function loadData() {
    setLoading(true);
    setError(null);

    try {
      const [requestResult, hotelsResult] = await Promise.all([
        getMyTravelRequest(),
        listResource('/api/v1/hotels', 1, 100),
      ]);

      const hotelItems = Array.isArray(hotelsResult.items)
        ? hotelsResult.items.map((item) => ({
            id: Number(item.id),
            name: String(item.name ?? ''),
            address: String(item.address ?? ''),
          })).filter((item) => Number.isFinite(item.id) && item.id > 0)
        : [];

      setRequest(requestResult);
      setHotels(hotelItems);
      setDesiredHotelID(requestResult.desired_hotel_id ? String(requestResult.desired_hotel_id) : '');
    } catch (caughtError) {
      setRequest(null);
      setHotels([]);
      setError(getApiErrorMessage(caughtError));
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    void loadData();
  }, []);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    if (!request?.has_document) {
      setError('Сначала заполните документ в разделе «Мои документы»');
      return;
    }

    setSaving(true);
    setError(null);
    setSuccess(null);

    try {
      const hotelID = desiredHotelID ? Number(desiredHotelID) : null;

      if (hotelID !== null && (!Number.isFinite(hotelID) || hotelID <= 0)) {
        setError('Выберите корректную гостиницу');
        return;
      }

      const updatedRequest = await updateMyTravelRequest({
        desired_hotel_id: hotelID,
      });

      setRequest(updatedRequest);
      setDesiredHotelID(updatedRequest.desired_hotel_id ? String(updatedRequest.desired_hotel_id) : '');
      setSuccess('Заявка обновлена. Менеджер увидит выбранную гостиницу при добавлении вас в группу.');
    } catch (caughtError) {
      setError(getApiErrorMessage(caughtError));
    } finally {
      setSaving(false);
    }
  }

  return (
    <section className="page-section">
      <h1>Моя заявка на поездку</h1>

      <div className="card">
        <h2>Как это работает</h2>
        <p>
          Здесь турист указывает желаемую гостиницу. Это не фактическое расселение:
          менеджер сначала добавляет туриста в группу, затем оформляет визу, расселение,
          экскурсии и грузовые операции.
        </p>
      </div>

      {loading && <div className="page-state">Загрузка...</div>}
      {error && <div className="error-box">{error}</div>}
      {success && <div className="success-box">{success}</div>}

      {!loading && request && (
        <>
          <div className="dashboard-grid">
            <div className="card">
              <h2>Турист</h2>
              <p><strong>ФИО:</strong> {request.last_name} {request.first_name} {request.middle_name ?? ''}</p>
              <p><strong>Пол:</strong> {request.sex === 'male' ? 'мужской' : 'женский'}</p>
              <p><strong>Дата рождения:</strong> {request.birth_date}</p>
            </div>

            <div className="card">
              <h2>Статус оформления</h2>
              <p><strong>Документ:</strong> {request.has_document ? 'заполнен' : 'не заполнен'}</p>
              <p><strong>Групп:</strong> {request.group_count}</p>
              <p>
                <strong>Желаемая гостиница:</strong>{' '}
                {request.desired_hotel_name
                  ? `${request.desired_hotel_name} — ${request.desired_hotel_address ?? ''}`
                  : 'не выбрана'}
              </p>
            </div>
          </div>

          {!request.has_document && (
            <div className="warning-box">
              Документ ещё не заполнен. До записи на экскурсии и полноценного оформления поездки нужно заполнить раздел «Мои документы».
            </div>
          )}

          {request.group_count === 0 && (
            <div className="warning-box">
              Вы ещё не добавлены в туристическую группу. После выбора гостиницы менеджер сможет добавить вас в группу и оформить поездку.
            </div>
          )}

          <form className="card profile-form" onSubmit={handleSubmit}>
            <label>
              Желаемая гостиница
              <select
                value={desiredHotelID}
                onChange={(event) => setDesiredHotelID(event.target.value)}
              >
                <option value="">Не выбрана</option>
                {hotels.map((hotel) => (
                  <option key={hotel.id} value={hotel.id}>
                    {hotel.name} — {hotel.address}
                  </option>
                ))}
              </select>
            </label>

            {selectedHotel && (
              <div className="info-box">
                Вы выбрали: <strong>{selectedHotel.name}</strong>, {selectedHotel.address}.
                Менеджер сможет использовать эту гостиницу как пожелание при добавлении вас в группу.
              </div>
            )}

            <button type="submit" className="primary-button" disabled={saving || !request.has_document}>
              {saving ? 'Сохранение...' : 'Сохранить заявку'}
            </button>
          </form>
        </>
      )}
    </section>
  );
}
