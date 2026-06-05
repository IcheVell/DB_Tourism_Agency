export interface NavigationItem {
    label: string;
    to: string;
    group?: string;
    permission?: string;
}

export const navigationItems: NavigationItem[] = [
    {
        label: 'Главная',
        to: '/',
        group: 'Общее',
    },
    {
        label: 'Профиль',
        to: '/profile',
        group: 'Общее',
        permission: 'profile.read',
    },

    {
        label: 'Мои документы',
        to: '/me/identity-document',
        group: 'Турист',
        permission: 'profile.read',
    },
    {
        label: 'Мои поездки',
        to: '/me/tours',
        group: 'Турист',
        permission: 'own_tours.read',
    },
    {
        label: 'Мои визы',
        to: '/me/visas',
        group: 'Турист',
        permission: 'own_visas.read',
    },
    {
        label: 'Моё расселение',
        to: '/me/accommodations',
        group: 'Турист',
        permission: 'own_accommodations.read',
    },
    {
        label: 'Мои экскурсии',
        to: '/me/excursions',
        group: 'Турист',
        permission: 'own_excursions.read',
    },
    {
        label: 'Мой груз',
        to: '/me/cargo',
        group: 'Турист',
        permission: 'own_cargo.read',
    },

    {
        label: 'Гостиницы',
        to: '/crud/hotels',
        group: 'Гостиницы',
        permission: 'hotels.read',
    },
    {
        label: 'Номера',
        to: '/crud/hotelRooms',
        group: 'Гостиницы',
        permission: 'hotel_rooms.read',
    },
    {
        label: 'Расселение',
        to: '/crud/accommodations',
        group: 'Гостиницы',
        permission: 'accommodations.read',
    },

    {
        label: 'Агентства экскурсий',
        to: '/crud/excursionAgencies',
        group: 'Экскурсии',
        permission: 'excursion_agencies.read',
    },
    {
        label: 'Экскурсии',
        to: '/crud/excursions',
        group: 'Экскурсии',
        permission: 'excursions.read',
    },
    {
        label: 'Расписание экскурсий',
        to: '/excursion-schedules',
        group: 'Экскурсии',
        permission: 'excursion_schedule.read',
    },
    {
        label: 'Управление расписанием',
        to: '/crud/excursionSchedules',
        group: 'Экскурсии',
        permission: 'excursion_schedule.create',
    },
    {
        label: 'Записи на экскурсии',
        to: '/crud/excursionBookings',
        group: 'Экскурсии',
        permission: 'excursion_bookings.read',
    },

    {
        label: 'Туристы',
        to: '/crud/tourists',
        group: 'Туристы',
        permission: 'tourists.read',
    },
    {
        label: 'Категории туристов',
        to: '/crud/touristCategories',
        group: 'Туристы',
        permission: 'tourist_categories.read',
    },
    {
        label: 'Группы',
        to: '/crud/touristGroups',
        group: 'Туристы',
        permission: 'tourist_groups.read',
    },
    {
        label: 'Участники групп',
        to: '/crud/groupMembers',
        group: 'Туристы',
        permission: 'group_members.read',
    },
    {
        label: 'Сопровождение детей',
        to: '/crud/childCompanions',
        group: 'Туристы',
        permission: 'child_companions.read',
    },
    {
        label: 'Документы туристов',
        to: '/crud/identityDocuments',
        group: 'Туристы',
        permission: 'identity_documents.read',
    },
    {
        label: 'Визы',
        to: '/crud/visas',
        group: 'Туристы',
        permission: 'visas.read',
    },

    {
        label: 'Типы рейсов',
        to: '/crud/flightTypes',
        group: 'Рейсы',
        permission: 'flight_types.read',
    },
    {
        label: 'Рейсы',
        to: '/crud/flights',
        group: 'Рейсы',
        permission: 'flights.read',
    },

    {
        label: 'Виды груза',
        to: '/crud/cargoTypes',
        group: 'Груз',
        permission: 'cargo_types.read',
    },
    {
        label: 'Грузовые ведомости',
        to: '/crud/cargoStatements',
        group: 'Груз',
        permission: 'cargo_statements.read',
    },
    {
        label: 'Грузовые места',
        to: '/crud/cargoItems',
        group: 'Груз',
        permission: 'cargo_items.read',
    },
    {
        label: 'Отправки груза',
        to: '/crud/cargoShipments',
        group: 'Груз',
        permission: 'cargo_shipments.read',
    },

    {
        label: 'Финансовые категории',
        to: '/crud/financialCategories',
        group: 'Финансы',
        permission: 'financial_categories.read',
    },
    {
        label: 'Финансовые операции',
        to: '/crud/financialOperations',
        group: 'Финансы',
        permission: 'financial_operations.read',
    },

    {
        label: 'Пользователи',
        to: '/crud/users',
        group: 'ACL',
        permission: 'users.read',
    },
    {
        label: 'Роли',
        to: '/crud/roles',
        group: 'ACL',
        permission: 'roles.read',
    },
    {
        label: 'Права',
        to: '/crud/permissions',
        group: 'ACL',
        permission: 'permissions.read',
    },
    {
        label: 'Роли пользователей',
        to: '/crud/userRoles',
        group: 'ACL',
        permission: 'user_roles.read',
    },
    {
        label: 'Права ролей',
        to: '/crud/rolePermissions',
        group: 'ACL',
        permission: 'role_permissions.read',
    },

    {
        label: 'Отчёты',
        to: '/reports',
        group: 'Отчёты',
        permission: 'reports.tourists_count.read',
    },
];