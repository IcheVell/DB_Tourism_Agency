import { http } from './http';

export async function getReport<T = unknown>(path: string, params: Record<string, string | number | undefined>): Promise<T> {
  const cleanParams = Object.fromEntries(
    Object.entries(params).filter(([, value]) => value !== undefined && value !== ''),
  );

  const { data } = await http.get(path, { params: cleanParams });
  return data;
}
