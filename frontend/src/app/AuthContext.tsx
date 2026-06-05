import {
    createContext,
    useCallback,
    useContext,
    useEffect,
    useMemo,
    useState,
} from 'react';
import {
    getMe,
    login as loginRequest,
    register as registerRequest,
} from '../api/authApi';
import {
    clearAccessToken,
    getAccessToken,
    setAccessToken,
} from '../api/tokenStorage';
import type {
    CurrentUser,
    LoginRequest,
    RegisterRequest,
    RegisterResponse,
} from '../types/auth';

interface AuthContextValue {
    user: CurrentUser | null;
    loading: boolean;
    authenticated: boolean;
    login: (request: LoginRequest) => Promise<void>;
    register: (request: RegisterRequest) => Promise<RegisterResponse>;
    logout: () => void;
    refreshMe: () => Promise<void>;
    hasPermission: (permission: string) => boolean;
}

const AuthContext = createContext<AuthContextValue | null>(null);

export function AuthProvider({ children }: { children: React.ReactNode }) {
    const [user, setUser] = useState<CurrentUser | null>(null);
    const [loading, setLoading] = useState(true);

    const refreshMe = useCallback(async () => {
        const currentUser = await getMe();
        setUser(currentUser);
    }, []);

    useEffect(() => {
        async function bootstrap() {
            const token = getAccessToken();

            if (!token) {
                setLoading(false);
                return;
            }

            try {
                await refreshMe();
            } catch {
                clearAccessToken();
                setUser(null);
            } finally {
                setLoading(false);
            }
        }

        void bootstrap();
    }, [refreshMe]);

    const login = useCallback(async (request: LoginRequest) => {
        const response = await loginRequest(request);

        if (!response.access_token) {
            throw new Error('Сервер не вернул access_token');
        }

        setAccessToken(response.access_token);

        const currentUser = await getMe();
        setUser(currentUser);
    }, []);

    const register = useCallback(async (request: RegisterRequest) => {
        const response = await registerRequest(request);
        return response;
    }, []);

    const logout = useCallback(() => {
        clearAccessToken();
        setUser(null);
    }, []);

    const hasPermission = useCallback(
        (permission: string) => user?.permissions.includes(permission) ?? false,
        [user],
    );

    const value = useMemo<AuthContextValue>(
        () => ({
            user,
            loading,
            authenticated: Boolean(user),
            login,
            register,
            logout,
            refreshMe,
            hasPermission,
        }),
        [hasPermission, loading, login, logout, refreshMe, register, user],
    );

    return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuth(): AuthContextValue {
    const value = useContext(AuthContext);

    if (!value) {
        throw new Error('useAuth must be used inside AuthProvider');
    }

    return value;
}