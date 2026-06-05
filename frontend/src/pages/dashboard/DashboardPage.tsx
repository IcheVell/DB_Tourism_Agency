import { useAuth } from '../../app/AuthContext';

export function DashboardPage() {
    const { user } = useAuth();

    const rolesText = user?.roles && user.roles.length > 0
        ? user.roles.join(', ')
        : 'не назначена';

    const permissionsCount = user?.permissions?.length ?? 0;

    return (
        <section className="page-section">
            <h1>Главная</h1>

            <div className="dashboard-grid">
                <div className="card">
                    <h2>Пользователь</h2>

                    <p>
                        <strong>Логин:</strong> {user?.login ?? '—'}
                    </p>

                    <p>
                        <strong>Email:</strong> {user?.email ?? '—'}
                    </p>

                    <p>
                        <strong>Роль:</strong> {rolesText}
                    </p>
                </div>

                <div className="card">
                    <h2>Права</h2>

                    <p>
                        <strong>Назначено прав:</strong> {permissionsCount}
                    </p>
                </div>
            </div>
        </section>
    );
}