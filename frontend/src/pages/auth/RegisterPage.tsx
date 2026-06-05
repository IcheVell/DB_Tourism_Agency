import { FormEvent, useState } from 'react';
import { Link, Navigate, useNavigate } from 'react-router-dom';
import { getApiErrorMessage } from '../../api/http';
import { useAuth } from '../../app/AuthContext';

type SexValue = 'male' | 'female';

export function RegisterPage() {
    const [login, setLogin] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [passwordRepeat, setPasswordRepeat] = useState('');

    const [lastName, setLastName] = useState('');
    const [firstName, setFirstName] = useState('');
    const [middleName, setMiddleName] = useState('');
    const [sex, setSex] = useState<SexValue>('male');
    const [birthDate, setBirthDate] = useState('');

    const [error, setError] = useState<string | null>(null);
    const [submitting, setSubmitting] = useState(false);

    const auth = useAuth();
    const navigate = useNavigate();

    if (auth.authenticated) {
        return <Navigate to="/" replace />;
    }

    async function handleSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();

        setSubmitting(true);
        setError(null);

        if (password !== passwordRepeat) {
            setError('Пароли не совпадают');
            setSubmitting(false);
            return;
        }

        if (password.length < 8) {
            setError('Пароль должен быть не короче 8 символов');
            setSubmitting(false);
            return;
        }

        if (!lastName.trim()) {
            setError('Фамилия обязательна');
            setSubmitting(false);
            return;
        }

        if (!firstName.trim()) {
            setError('Имя обязательно');
            setSubmitting(false);
            return;
        }

        if (!birthDate) {
            setError('Дата рождения обязательна');
            setSubmitting(false);
            return;
        }

        try {
            await auth.register({
                login: login.trim(),
                email: email.trim(),
                password,

                last_name: lastName.trim(),
                first_name: firstName.trim(),
                middle_name: middleName.trim() || null,
                sex,
                birth_date: birthDate,
            });

            navigate('/login', {
                replace: true,
                state: {
                    message: 'Аккаунт туриста создан. Теперь войдите в систему.',
                },
            });
        } catch (caughtError) {
            setError(getApiErrorMessage(caughtError));
        } finally {
            setSubmitting(false);
        }
    }

    return (
        <div className="login-page">
            <form className="login-card register-card" onSubmit={handleSubmit}>
                <h1>Регистрация</h1>

                <div className="form-section-title">Данные аккаунта</div>

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
                    Email
                    <input
                        type="email"
                        value={email}
                        onChange={(event) => setEmail(event.target.value)}
                        autoComplete="email"
                        required
                    />
                </label>

                <label>
                    Пароль
                    <input
                        type="password"
                        value={password}
                        onChange={(event) => setPassword(event.target.value)}
                        autoComplete="new-password"
                        required
                    />
                </label>

                <label>
                    Повторите пароль
                    <input
                        type="password"
                        value={passwordRepeat}
                        onChange={(event) => setPasswordRepeat(event.target.value)}
                        autoComplete="new-password"
                        required
                    />
                </label>

                <div className="form-section-title">Данные туриста</div>

                <label>
                    Фамилия
                    <input
                        value={lastName}
                        onChange={(event) => setLastName(event.target.value)}
                        required
                    />
                </label>

                <label>
                    Имя
                    <input
                        value={firstName}
                        onChange={(event) => setFirstName(event.target.value)}
                        required
                    />
                </label>

                <label>
                    Отчество
                    <input
                        value={middleName}
                        onChange={(event) => setMiddleName(event.target.value)}
                    />
                </label>

                <label>
                    Пол
                    <select
                        value={sex}
                        onChange={(event) => setSex(event.target.value as SexValue)}
                        required
                    >
                        <option value="male">Мужской</option>
                        <option value="female">Женский</option>
                    </select>
                </label>

                <label>
                    Дата рождения
                    <input
                        type="date"
                        value={birthDate}
                        onChange={(event) => setBirthDate(event.target.value)}
                        required
                    />
                </label>

                {error && <div className="error-box">{error}</div>}

                <button type="submit" disabled={submitting}>
                    {submitting ? 'Регистрация...' : 'Зарегистрироваться'}
                </button>

                <div className="auth-switch">
                    Уже есть аккаунт? <Link to="/login">Войти</Link>
                </div>
            </form>
        </div>
    );
}