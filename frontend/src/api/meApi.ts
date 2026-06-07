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

export interface MeCargoType {
  id: number;
  name: string;
}

export interface CreateMeCargoRequest {
  group_member_id?: number;
  cargo_type_id: number;
  item_number: string;
  places_count: number;
  weight_kg: number;
  volumetric_weight_kg: number;
}

export interface MeCargoCreateResponse {
  cargo_statement_id: number;
  statement_status: string;
  cargo_item_id: number;
  item_number: string;
  places_count: number;
  weight_kg: number;
  volumetric_weight_kg: number;
  cargo_type_id: number;
  cargo_type_name: string;
}

export async function getMyCargoTypes(): Promise<MeCargoType[]> {
  const { data } = await http.get('/api/v1/me/cargo-types');
  const items = data?.items;
  return Array.isArray(items) ? items : [];
}

export async function createMyCargo(payload: CreateMeCargoRequest): Promise<MeCargoCreateResponse> {
  const { data } = await http.post('/api/v1/me/cargo', payload);
  return data;
}

export interface MeIdentityDocument {
  id: number;
  tourist_id: number;
  document_type: string;
  document_series: string;
  document_number: string;
  issue_date: string;
  expiration_date?: string | null;
  issued_by: string;
  citizenship: string;
}

export interface CreateMeIdentityDocumentRequest {
  document_type: string;
  document_series: string;
  document_number: string;
  issue_date: string;
  expiration_date?: string | null;
  issued_by: string;
  citizenship: string;
}

export interface UpdateMeIdentityDocumentRequest {
  document_type?: string;
  document_series?: string;
  document_number?: string;
  issue_date?: string;
  expiration_date?: string | null;
  issued_by?: string;
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

export interface MeTravelRequest {
  tourist_id: number;
  first_name: string;
  last_name: string;
  middle_name?: string | null;
  sex: string;
  birth_date: string;
  desired_hotel_id?: number | null;
  desired_hotel_name?: string | null;
  desired_hotel_address?: string | null;
  has_document: boolean;
  group_count: number;
}

export interface UpdateMeTravelRequestRequest {
  desired_hotel_id?: number | null;
}

export async function getMyTravelRequest(): Promise<MeTravelRequest> {
  const { data } = await http.get('/api/v1/me/travel-request');
  return data;
}

export async function updateMyTravelRequest(
  payload: UpdateMeTravelRequestRequest,
): Promise<MeTravelRequest> {
  const { data } = await http.put('/api/v1/me/travel-request', payload);
  return data;
}
