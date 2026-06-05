export type PermissionCode = string;

export interface CurrentUser {
    id: number;
    login: string;
    email: string;
    roles: string[];
    permissions: PermissionCode[];
}

export interface LoginRequest {
    login: string;
    password: string;
}

export interface RegisterRequest {
    login: string;
    email: string;
    password: string;

    first_name: string;
    last_name: string;
    middle_name?: string | null;
    sex: 'male' | 'female';
    birth_date: string;
}

export interface LoginResponse {
    access_token: string;
    token_type: string;
}

export interface RegisterResponse {
    id: number;
    login: string;
    email: string;
    role: string;
    tourist_id: number;
}