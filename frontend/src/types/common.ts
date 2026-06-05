export interface ApiErrorResponse {
  message: string;
}

export interface PaginatedResponse<T> {
  items: T[];
  total: number;
  page: number;
  limit: number;
}

export type JsonObject = Record<string, unknown>;
