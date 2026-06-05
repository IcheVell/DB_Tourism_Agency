import { http } from './http';
import type {
    CurrentUser,
    LoginRequest,
    LoginResponse,
    RegisterRequest,
    RegisterResponse,
} from '../types/auth';

function normalizeCurrentUser(raw: any): CurrentUser {
    return {
        id: Number(raw?.id ?? raw?.user_id),
        login: String(raw?.login ?? ''),
        email: String(raw?.email ?? ''),
        roles: Array.isArray(raw?.roles)
            ? raw.roles.map(String)
            : raw?.role
                ? [String(raw.role)]
                : [],
        permissions: Array.isArray(raw?.permissions)
            ? raw.permissions.map(String)
            : [],
    };
}

export async function login(request: LoginRequest): Promise<LoginResponse> {
    const { data } = await http.post('/auth/login', request);

    return {
        access_token: String(data.access_token ?? ''),
        token_type: String(data.token_type ?? 'Bearer'),
    };
}

export async function register(request: RegisterRequest): Promise<RegisterResponse> {
    const { data } = await http.post('/auth/register', request);

    return {
        id: Number(data.id),
        login: String(data.login),
        email: String(data.email),
        role: String(data.role),
        tourist_id: Number(data.tourist_id),
    };
}

export async function getMe(): Promise<CurrentUser> {
    const { data } = await http.get('/api/v1/me');
    return normalizeCurrentUser(data);
}