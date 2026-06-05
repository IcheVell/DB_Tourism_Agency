import { useAuth } from '../../app/AuthContext';

export function ProfilePage() {
    const { user } = useAuth();

    const roles = Array.isArray(user?.roles) ? user.roles : [];
    const permissions = Array.isArray(user?.permissions) ? user.permissions : [];

    return (
        <section className="page-section">
            <h1>Профиль</h1>

            <div className="card">
                <p>
                    <strong>ID:</strong> {user?.id ?? '—'}
                </p>

                <p>
                    <strong>Логин:</strong> {user?.login ?? '—'}
                </p>

                <p>
                    <strong>Email:</strong> {user?.email ?? '—'}
                </p>

                <p>
                    <strong>Роль:</strong> {roles.length > 0 ? roles.join(', ') : 'не назначена'}
                </p>

                <p>
                    <strong>Количество прав:</strong> {permissions.length}
                </p>
            </div>

            <div className="card">
                <h2>Права</h2>

                {permissions.length === 0 ? (
                    <p>Нет назначенных прав</p>
                ) : (
                    <ul className="permissions-list">
                        {permissions.map((permission) => (
                            <li key={permission}>{permission}</li>
                        ))}
                    </ul>
                )}
            </div>
        </section>
    );
}