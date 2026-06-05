import { NavLink, Outlet } from 'react-router-dom';
import { useAuth } from '../../app/AuthContext';
import { navigationItems, type NavigationItem } from '../../config/navigation';

export function AppLayout() {
    const { user, logout, hasPermission } = useAuth();

    const visibleItems = navigationItems.filter((item: NavigationItem) => {
        if (!item.permission) {
            return true;
        }

        return hasPermission(item.permission);
    });

    const groupedItems = visibleItems.reduce<Record<string, NavigationItem[]>>(
        (acc, item) => {
            const group = item.group ?? 'Общее';

            if (!acc[group]) {
                acc[group] = [];
            }

            acc[group].push(item);

            return acc;
        },
        {},
    );

    const roles = Array.isArray(user?.roles) ? user.roles : [];
    const rolesText = roles.length > 0 ? roles.join(', ') : 'без роли';

    return (
        <div className="app-shell">
            <aside className="sidebar">
                <div className="sidebar-title">tourist agency</div>

                <div className="sidebar-user">
                    <strong>{user?.login ?? '—'}</strong>
                    <span>{rolesText}</span>
                </div>

                <nav className="sidebar-nav">
                    {Object.entries(groupedItems).map(([group, items]) => (
                        <div className="sidebar-group" key={group}>
                            <div className="sidebar-group-title">{group}</div>

                            {items.map((item) => (
                                <NavLink
                                    key={item.to}
                                    to={item.to}
                                    end={item.to === '/'}
                                    className={({ isActive }) =>
                                        isActive ? 'sidebar-link active' : 'sidebar-link'
                                    }
                                >
                                    {item.label}
                                </NavLink>
                            ))}
                        </div>
                    ))}
                </nav>
            </aside>

            <div className="main-shell">
                <header className="topbar">
                    <button type="button" className="logout-button" onClick={logout}>
                        Выйти
                    </button>
                </header>

                <main className="main-content">
                    <Outlet />
                </main>
            </div>
        </div>
    );
}