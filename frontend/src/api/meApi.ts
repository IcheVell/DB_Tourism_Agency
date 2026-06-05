import { http } from './http';
import type { PaginatedResponse } from '../types/common';

export async function getMyTours(page = 1, limit = 20): Promise<PaginatedResponse<Record<string, unknown>>> {
  const { data } = await http.get('/api/v1/me/tours', { params: { page, limit } });
  return data;
}

export async function getMyVisas(page = 1, limit = 20): Promise<PaginatedResponse<Record<string, unknown>>> {
  const { data } = await http.get('/api/v1/me/visas', { params: { page, limit } });
  return data;
}

export async function getMyAccommodations(page = 1, limit = 20): Promise<PaginatedResponse<Record<string, unknown>>> {
  const { data } = await http.get('/api/v1/me/accommodations', { params: { page, limit } });
  return data;
}

export async function getMyExcursions(page = 1, limit = 20): Promise<PaginatedResponse<Record<string, unknown>>> {
  const { data } = await http.get('/api/v1/me/excursions', { params: { page, limit } });
  return data;
}

export async function getMyCargo(page = 1, limit = 20): Promise<PaginatedResponse<Record<string, unknown>>> {
  const { data } = await http.get('/api/v1/me/cargo', { params: { page, limit } });
  return data;
}

export interface MeIdentityDocument {
    id: number;
    tourist_id: number;
    document_type: string;
    document_series?: string | null;
    document_number: string;
    citizenship: string;
}

export interface CreateMeIdentityDocumentRequest {
    document_type: string;
    document_series?: string | null;
    document_number: string;
    citizenship: string;
}

export interface UpdateMeIdentityDocumentRequest {
    document_type?: string;
    document_series?: string | null;
    document_number?: string;
    citizenship?: string;
}

export interface CreateMeExcursionBookingRequest {
    excursion_schedule_id: number;
    group_member_id?: number;
}

export async function getMyIdentityDocument(): Promise<MeIdentityDocument> {
    const { data } = await http.get('/api/v1/me/identity-document');
    return data;
}

export async function createMyIdentityDocument(
    payload: CreateMeIdentityDocumentRequest,
): Promise<MeIdentityDocument> {
    const { data } = await http.post('/api/v1/me/identity-document', payload);
    return data;
}

export async function updateMyIdentityDocument(
    payload: UpdateMeIdentityDocumentRequest,
): Promise<MeIdentityDocument> {
    const { data } = await http.put('/api/v1/me/identity-document', payload);
    return data;
}

export async function createMyExcursionBooking(
    payload: CreateMeExcursionBookingRequest,
): Promise<void> {
    await http.post('/api/v1/me/excursion-bookings', payload);
}
