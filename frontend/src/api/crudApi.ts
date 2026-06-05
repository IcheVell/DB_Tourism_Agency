import { http } from './http';

export interface CrudListResponse {
    items: Record<string, unknown>[];
    total: number;
    page: number;
    limit: number;
}

function normalizeListResponse(data: any, fallbackPage: number, fallbackLimit: number): CrudListResponse {
    if (Array.isArray(data)) {
        return {
            items: data,
            total: data.length,
            page: fallbackPage,
            limit: fallbackLimit,
        };
    }

    const rawItems = data?.items ?? data?.Items ?? data?.data ?? data?.rows ?? [];
    const items = Array.isArray(rawItems) ? rawItems : [];

    const total = Number(data?.total ?? data?.Total ?? items.length);
    const page = Number(data?.page ?? data?.Page ?? fallbackPage);
    const limit = Number(data?.limit ?? data?.Limit ?? fallbackLimit);

    return {
        items,
        total: Number.isFinite(total) && total >= 0 ? total : items.length,
        page: Number.isFinite(page) && page > 0 ? page : fallbackPage,
        limit: Number.isFinite(limit) && limit > 0 ? limit : fallbackLimit,
    };
}

export async function listResource(
    endpoint: string,
    page = 1,
    limit = 20,
): Promise<CrudListResponse> {
    const { data } = await http.get(endpoint, {
        params: {
            page,
            limit,
        },
    });

    return normalizeListResponse(data, page, limit);
}

export async function getResourceByID(
    endpoint: string,
    id: number | string,
): Promise<Record<string, unknown>> {
    const { data } = await http.get(`${endpoint}/${id}`);
    return data;
}

export async function createResource(
    endpoint: string,
    payload: Record<string, unknown>,
): Promise<Record<string, unknown>> {
    const { data } = await http.post(endpoint, payload);
    return data;
}

export async function updateResource(
    endpoint: string,
    id: number | string,
    payload: Record<string, unknown>,
): Promise<Record<string, unknown>> {
    const { data } = await http.put(`${endpoint}/${id}`, payload);
    return data;
}

export async function deleteResource(endpoint: string, id: number | string): Promise<void> {
    await http.delete(`${endpoint}/${id}`);
}