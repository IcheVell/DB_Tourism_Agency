import { FormEvent, useState } from 'react';
import { Link, Navigate, useLocation, useNavigate } from 'react-router-dom';
import { getApiErrorMessage } from '../../api/http';
import { useAuth } from '../../app/AuthContext';

export function LoginPage() {
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState<string | null>(null);
    const [submitting, setSubmitting] = useState(false);

    const auth = useAuth();
    const navigate = useNavigate();
    const location = useLocation();

    const state = location.state as { from?: { pathname?: string }; message?: string } | null;
    const from = state?.from?.pathname ?? '/';

    if (auth.authenticated) {
        return <Navigate to="/" replace />;
    }

    async function handleSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();

        setSubmitting(true);
        setError(null);

        try {
            await auth.login({ login, password });
            navigate(from, { replace: true });
        } catch (caughtError) {
            setError(getApiErrorMessage(caughtError));
        } finally {
            setSubmitting(false);
        }
    }

    return (
        <div className="login-page">
            <form className="login-card" onSubmit={handleSubmit}>
                <h1>Вход</h1>

                {state?.message && <div className="success-box">{state.message}</div>}

                <label>
                    Логин
                    <input
                        value={login}
                        onChange={(event) => setLogin(event.target.value)}
                        autoComplete="username"
                        required
                    />
                </label>

                <label>
                    Пароль
                    <input
                        type="password"
                        value={password}
                        onChange={(event) => setPassword(event.target.value)}
                        autoComplete="current-password"
                        required
                    />
                </label>

                {error && <div className="error-box">{error}</div>}

                <button type="submit" disabled={submitting}>
                    {submitting ? 'Вход...' : 'Войти'}
                </button>

                <div className="auth-switch">
                    Нет аккаунта? <Link to="/register">Зарегистрироваться</Link>
                </div>
            </form>
        </div>
    );
}