import { createBrowserRouter } from 'react-router-dom';
import { ProtectedRoute } from '../components/common/ProtectedRoute';
import { AppLayout } from '../components/layout/AppLayout';
import { LoginPage } from '../pages/auth/LoginPage';
import { RegisterPage } from '../pages/auth/RegisterPage';
import { CrudListPage } from '../pages/crud/CrudListPage';
import { DashboardPage } from '../pages/dashboard/DashboardPage';
import { ProfilePage } from '../pages/dashboard/ProfilePage';
import { NotFoundPage } from '../pages/errors/NotFoundPage';
import { MeListPage } from '../pages/me/MeListPage';
import { ReportsPage } from '../pages/reports/ReportsPage';
import { MyDocumentPage } from '../pages/me/MyDocumentPage';
import { ExcursionSchedulePage } from '../pages/me/ExcursionSchedulePage';

export const router = createBrowserRouter([
    {
        path: '/login',
        element: <LoginPage />,
    },
    {
        path: '/register',
        element: <RegisterPage />,
    },
    {
        element: <ProtectedRoute />,
        children: [
            {
                element: <AppLayout />,
                children: [
                    { index: true, element: <DashboardPage /> },
                    { path: 'profile', element: <ProfilePage /> },
                    { path: 'me/tours', element: <MeListPage type="tours" /> },
                    { path: 'me/visas', element: <MeListPage type="visas" /> },
                    { path: 'me/accommodations', element: <MeListPage type="accommodations" /> },
                    { path: 'me/excursions', element: <MeListPage type="excursions" /> },
                    { path: 'me/cargo', element: <MeListPage type="cargo" /> },
                    { path: 'crud/:resourceKey', element: <CrudListPage /> },
                    { path: 'reports', element: <ReportsPage /> },
                    { path: '*', element: <NotFoundPage /> },
                    { path: 'me/identity-document', element: <MyDocumentPage /> },
                    { path: 'excursion-schedules', element: <ExcursionSchedulePage /> },
                ],
            },
        ],
    },
]);