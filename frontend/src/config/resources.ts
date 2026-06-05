export interface ResourceConfig {
    key: string;
    label: string;
    endpoint: string;
    readPermission: string;
    createPermission?: string;
    updatePermission?: string;
    deletePermission?: string;
    createTemplate?: Record<string, unknown>;
}

export const resources: Record<string, ResourceConfig> = {
    hotels: {
        key: 'hotels',
        label: 'Гостиницы',
        endpoint: '/api/v1/hotels',
        readPermission: 'hotels.read',
        createPermission: 'hotels.create',
        updatePermission: 'hotels.update',
        deletePermission: 'hotels.delete',
        createTemplate: {
            name: '',
            address: '',
            phone: '',
        },
    },

    hotelRooms: {
        key: 'hotelRooms',
        label: 'Номера',
        endpoint: '/api/v1/hotel-rooms',
        readPermission: 'hotel_rooms.read',
        createPermission: 'hotel_rooms.create',
        updatePermission: 'hotel_rooms.update',
        deletePermission: 'hotel_rooms.delete',
        createTemplate: {
            hotel_id: 1,
            room_number: 101,
            capacity: 2,
            price: 100,
            status: 'available',
        },
    },

    excursionAgencies: {
        key: 'excursionAgencies',
        label: 'Агентства экскурсий',
        endpoint: '/api/v1/excursion-agencies',
        readPermission: 'excursion_agencies.read',
        createPermission: 'excursion_agencies.create',
        updatePermission: 'excursion_agencies.update',
        deletePermission: 'excursion_agencies.delete',
    },

    excursions: {
        key: 'excursions',
        label: 'Экскурсии',
        endpoint: '/api/v1/excursions',
        readPermission: 'excursions.read',
        createPermission: 'excursions.create',
        updatePermission: 'excursions.update',
        deletePermission: 'excursions.delete',
        createTemplate: {
            name: '',
            description: '',
            duration_minutes: 60,
        },
    },

    excursionSchedules: {
        key: 'excursionSchedules',
        label: 'Расписание экскурсий',
        endpoint: '/api/v1/excursion-schedules',
        readPermission: 'excursion_schedule.read',
        createPermission: 'excursion_schedule.create',
        updatePermission: 'excursion_schedule.update',
        deletePermission: 'excursion_schedule.delete',
        createTemplate: {
            excursion_id: 1,
            excursion_agency_id: 1,
            start_time: '2026-06-03T10:00:00Z',
            end_time: '2026-06-03T12:00:00Z',
            price: 100,
            status: 'planned',
        },
    },

    tourists: {
        key: 'tourists',
        label: 'Туристы',
        endpoint: '/api/v1/tourists',
        readPermission: 'tourists.read',
        createPermission: 'tourists.create',
        updatePermission: 'tourists.update',
        deletePermission: 'tourists.delete',
    },

    touristCategories: {
        key: 'touristCategories',
        label: 'Категории туристов',
        endpoint: '/api/v1/tourist-categories',
        readPermission: 'tourist_categories.read',
        createPermission: 'tourist_categories.create',
        updatePermission: 'tourist_categories.update',
        deletePermission: 'tourist_categories.delete',
    },

    touristGroups: {
        key: 'touristGroups',
        label: 'Туристические группы',
        endpoint: '/api/v1/tourist-groups',
        readPermission: 'tourist_groups.read',
        createPermission: 'tourist_groups.create',
        updatePermission: 'tourist_groups.update',
        deletePermission: 'tourist_groups.delete',
    },

    groupMembers: {
        key: 'groupMembers',
        label: 'Участники групп',
        endpoint: '/api/v1/group-members',
        readPermission: 'group_members.read',
        createPermission: 'group_members.create',
        updatePermission: 'group_members.update',
        deletePermission: 'group_members.delete',
    },

    identityDocuments: {
        key: 'identityDocuments',
        label: 'Документы',
        endpoint: '/api/v1/identity-documents',
        readPermission: 'identity_documents.read',
        createPermission: 'identity_documents.create',
        updatePermission: 'identity_documents.update',
        deletePermission: 'identity_documents.delete',
    },

    visas: {
        key: 'visas',
        label: 'Визы',
        endpoint: '/api/v1/visas',
        readPermission: 'visas.read',
        createPermission: 'visas.create',
        updatePermission: 'visas.update',
        deletePermission: 'visas.delete',
    },

    accommodations: {
        key: 'accommodations',
        label: 'Расселение',
        endpoint: '/api/v1/accommodations',
        readPermission: 'accommodations.read',
        createPermission: 'accommodations.create',
        updatePermission: 'accommodations.update',
        deletePermission: 'accommodations.delete',
    },

    flightTypes: {
        key: 'flightTypes',
        label: 'Типы рейсов',
        endpoint: '/api/v1/flight-types',
        readPermission: 'flight_types.read',
        createPermission: 'flight_types.create',
        updatePermission: 'flight_types.update',
        deletePermission: 'flight_types.delete',
    },

    flights: {
        key: 'flights',
        label: 'Рейсы',
        endpoint: '/api/v1/flights',
        readPermission: 'flights.read',
        createPermission: 'flights.create',
        updatePermission: 'flights.update',
        deletePermission: 'flights.delete',
        createTemplate: {
            flight_number: 1001,
            flight_date: '2026-06-03T12:00:00Z',
            capacity: 180,
            flight_type_id: 1,
        },
    },

    cargoTypes: {
        key: 'cargoTypes',
        label: 'Виды груза',
        endpoint: '/api/v1/cargo-types',
        readPermission: 'cargo_types.read',
        createPermission: 'cargo_types.create',
        updatePermission: 'cargo_types.update',
        deletePermission: 'cargo_types.delete',
    },

    cargoStatements: {
        key: 'cargoStatements',
        label: 'Грузовые ведомости',
        endpoint: '/api/v1/cargo-statements',
        readPermission: 'cargo_statements.read',
        createPermission: 'cargo_statements.create',
        updatePermission: 'cargo_statements.update',
        deletePermission: 'cargo_statements.delete',
    },

    cargoItems: {
        key: 'cargoItems',
        label: 'Грузовые места',
        endpoint: '/api/v1/cargo-items',
        readPermission: 'cargo_items.read',
        createPermission: 'cargo_items.create',
        updatePermission: 'cargo_items.update',
        deletePermission: 'cargo_items.delete',
    },

    cargoShipments: {
        key: 'cargoShipments',
        label: 'Отправки груза',
        endpoint: '/api/v1/cargo-shipments',
        readPermission: 'cargo_shipments.read',
        createPermission: 'cargo_shipments.create',
        updatePermission: 'cargo_shipments.update',
        deletePermission: 'cargo_shipments.delete',
    },

    financialCategories: {
        key: 'financialCategories',
        label: 'Финансовые категории',
        endpoint: '/api/v1/financial-categories',
        readPermission: 'financial_categories.read',
        createPermission: 'financial_categories.create',
        updatePermission: 'financial_categories.update',
        deletePermission: 'financial_categories.delete',
    },

    financialOperations: {
        key: 'financialOperations',
        label: 'Финансовые операции',
        endpoint: '/api/v1/financial-operations',
        readPermission: 'financial_operations.read',
        createPermission: 'financial_operations.create',
        updatePermission: 'financial_operations.update',
        deletePermission: 'financial_operations.delete',
    },

    users: {
        key: 'users',
        label: 'Пользователи',
        endpoint: '/api/v1/users',
        readPermission: 'users.read',
        createPermission: 'users.create',
        updatePermission: 'users.update',
        deletePermission: 'users.delete',
    },

    roles: {
        key: 'roles',
        label: 'Роли',
        endpoint: '/api/v1/roles',
        readPermission: 'roles.read',
        createPermission: 'roles.create',
        updatePermission: 'roles.update',
        deletePermission: 'roles.delete',
    },

    permissions: {
        key: 'permissions',
        label: 'Права',
        endpoint: '/api/v1/permissions',
        readPermission: 'permissions.read',
        createPermission: 'permissions.create',
        updatePermission: 'permissions.update',
        deletePermission: 'permissions.delete',
    },

    excursionBookings: {
        key: 'excursionBookings',
        label: 'Записи на экскурсии',
        endpoint: '/api/v1/excursion-bookings',
        readPermission: 'excursion_bookings.read',
        createPermission: 'excursion_bookings.create',
        updatePermission: 'excursion_bookings.update',
        deletePermission: 'excursion_bookings.delete',
    },
};