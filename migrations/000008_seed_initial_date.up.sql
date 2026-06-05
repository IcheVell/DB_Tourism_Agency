DELETE FROM role_permissions
WHERE permission_id IN (
    SELECT id
    FROM permissions
    WHERE code = 'own_excursions.create'
);

DELETE FROM permissions
WHERE code = 'own_excursions.create';

INSERT INTO roles (name, description)
VALUES
    ('tourist', 'Обычный пользователь-турист'),
    ('manager', 'Менеджер представительства туристической фирмы'),
    ('admin', 'Администратор системы'),
    ('super_admin', 'Суперадминистратор системы')
ON CONFLICT (name) DO NOTHING;


INSERT INTO permissions (code, description)
VALUES
    ('profile.read', 'Просмотр собственного профиля'),
    ('profile.update', 'Изменение собственного профиля'),

    ('own_tours.read', 'Просмотр собственных поездок'),
    ('own_visas.read', 'Просмотр собственных виз'),
    ('own_accommodations.read', 'Просмотр собственного расселения'),
    ('own_excursions.read', 'Просмотр собственных экскурсий'),
    ('own_cargo.read', 'Просмотр собственного груза'),

    ('users.read', 'Просмотр пользователей'),
    ('users.create', 'Создание пользователей'),
    ('users.update', 'Изменение пользователей'),
    ('users.delete', 'Удаление пользователей'),

    ('roles.read', 'Просмотр ролей'),
    ('roles.create', 'Создание ролей'),
    ('roles.update', 'Изменение ролей'),
    ('roles.delete', 'Удаление ролей'),

    ('permissions.read', 'Просмотр прав доступа'),
    ('permissions.create', 'Создание прав доступа'),
    ('permissions.update', 'Изменение прав доступа'),
    ('permissions.delete', 'Удаление прав доступа'),

    ('user_roles.read', 'Просмотр связей пользователей и ролей'),
    ('user_roles.create', 'Назначение роли пользователю'),
    ('user_roles.update', 'Изменение роли пользователя'),
    ('user_roles.delete', 'Удаление роли пользователя'),

    ('role_permissions.read', 'Просмотр связей ролей и прав'),
    ('role_permissions.create', 'Назначение права роли'),
    ('role_permissions.update', 'Изменение права роли'),
    ('role_permissions.delete', 'Удаление права роли'),

    ('acl.manage', 'Управление ролями и правами доступа'),

    ('tourists.read', 'Просмотр туристов'),
    ('tourists.create', 'Создание туристов'),
    ('tourists.update', 'Изменение туристов'),
    ('tourists.delete', 'Удаление туристов'),

    ('tourist_categories.read', 'Просмотр категорий туристов'),
    ('tourist_categories.create', 'Создание категорий туристов'),
    ('tourist_categories.update', 'Изменение категорий туристов'),
    ('tourist_categories.delete', 'Удаление категорий туристов'),

    ('tourist_groups.read', 'Просмотр туристических групп'),
    ('tourist_groups.create', 'Создание туристических групп'),
    ('tourist_groups.update', 'Изменение туристических групп'),
    ('tourist_groups.delete', 'Удаление туристических групп'),

    ('group_members.read', 'Просмотр участников групп'),
    ('group_members.create', 'Добавление участников групп'),
    ('group_members.update', 'Изменение участников групп'),
    ('group_members.delete', 'Удаление участников групп'),

    ('child_companions.read', 'Просмотр сопровождающих детей'),
    ('child_companions.create', 'Создание связи сопровождения ребенка'),
    ('child_companions.update', 'Изменение связи сопровождения ребенка'),
    ('child_companions.delete', 'Удаление связи сопровождения ребенка'),

    ('identity_documents.read', 'Просмотр документов'),
    ('identity_documents.create', 'Создание документов'),
    ('identity_documents.update', 'Изменение документов'),
    ('identity_documents.delete', 'Удаление документов'),

    ('visas.read', 'Просмотр виз'),
    ('visas.create', 'Создание виз'),
    ('visas.update', 'Изменение виз'),
    ('visas.delete', 'Удаление виз'),

    ('hotels.read', 'Просмотр гостиниц'),
    ('hotels.create', 'Создание гостиниц'),
    ('hotels.update', 'Изменение гостиниц'),
    ('hotels.delete', 'Удаление гостиниц'),

    ('hotel_rooms.read', 'Просмотр номеров гостиниц'),
    ('hotel_rooms.create', 'Создание номеров гостиниц'),
    ('hotel_rooms.update', 'Изменение номеров гостиниц'),
    ('hotel_rooms.delete', 'Удаление номеров гостиниц'),

    ('accommodations.read', 'Просмотр расселения'),
    ('accommodations.create', 'Создание расселения'),
    ('accommodations.update', 'Изменение расселения'),
    ('accommodations.delete', 'Удаление расселения'),

    ('excursion_agencies.read', 'Просмотр экскурсионных агентств'),
    ('excursion_agencies.create', 'Создание экскурсионных агентств'),
    ('excursion_agencies.update', 'Изменение экскурсионных агентств'),
    ('excursion_agencies.delete', 'Удаление экскурсионных агентств'),

    ('excursions.read', 'Просмотр экскурсий'),
    ('excursions.create', 'Создание экскурсий'),
    ('excursions.update', 'Изменение экскурсий'),
    ('excursions.delete', 'Удаление экскурсий'),

    ('excursion_schedule.read', 'Просмотр расписания экскурсий'),
    ('excursion_schedule.create', 'Создание расписания экскурсий'),
    ('excursion_schedule.update', 'Изменение расписания экскурсий'),
    ('excursion_schedule.delete', 'Удаление расписания экскурсий'),

    ('excursion_bookings.read', 'Просмотр записей на экскурсии'),
    ('excursion_bookings.create', 'Создание записей на экскурсии'),
    ('excursion_bookings.update', 'Изменение записей на экскурсии'),
    ('excursion_bookings.delete', 'Удаление записей на экскурсии'),

    ('cargo_types.read', 'Просмотр видов груза'),
    ('cargo_types.create', 'Создание видов груза'),
    ('cargo_types.update', 'Изменение видов груза'),
    ('cargo_types.delete', 'Удаление видов груза'),

    ('flight_types.read', 'Просмотр типов рейсов'),
    ('flight_types.create', 'Создание типов рейсов'),
    ('flight_types.update', 'Изменение типов рейсов'),
    ('flight_types.delete', 'Удаление типов рейсов'),

    ('flights.read', 'Просмотр рейсов'),
    ('flights.create', 'Создание рейсов'),
    ('flights.update', 'Изменение рейсов'),
    ('flights.delete', 'Удаление рейсов'),

    ('cargo_statements.read', 'Просмотр грузовых ведомостей'),
    ('cargo_statements.create', 'Создание грузовых ведомостей'),
    ('cargo_statements.update', 'Изменение грузовых ведомостей'),
    ('cargo_statements.delete', 'Удаление грузовых ведомостей'),

    ('cargo_items.read', 'Просмотр грузовых мест'),
    ('cargo_items.create', 'Создание грузовых мест'),
    ('cargo_items.update', 'Изменение грузовых мест'),
    ('cargo_items.delete', 'Удаление грузовых мест'),

    ('cargo_shipments.read', 'Просмотр отправок груза'),
    ('cargo_shipments.create', 'Создание отправок груза'),
    ('cargo_shipments.update', 'Изменение отправок груза'),
    ('cargo_shipments.delete', 'Удаление отправок груза'),

    ('financial_categories.read', 'Просмотр финансовых категорий'),
    ('financial_categories.create', 'Создание финансовых категорий'),
    ('financial_categories.update', 'Изменение финансовых категорий'),
    ('financial_categories.delete', 'Удаление финансовых категорий'),

    ('financial_operations.read', 'Просмотр финансовых операций'),
    ('financial_operations.create', 'Создание финансовых операций'),
    ('financial_operations.update', 'Изменение финансовых операций'),
    ('financial_operations.delete', 'Удаление финансовых операций'),

    ('reports.customs_list.read', 'Просмотр таможенного списка'),
    ('reports.accommodation.read', 'Просмотр отчета по расселению'),
    ('reports.tourists_count.read', 'Просмотр количества туристов за период'),
    ('reports.tourist_info.read', 'Просмотр сведений о туристе'),
    ('reports.hotels.read', 'Просмотр гостиничных отчетов'),
    ('reports.excursions.read', 'Просмотр экскурсионных отчетов'),
    ('reports.flight_load.read', 'Просмотр загрузки рейса'),
    ('reports.cargo_turnover.read', 'Просмотр грузооборота склада'),
    ('reports.financial.read', 'Просмотр финансового отчета'),
    ('reports.profitability.read', 'Просмотр рентабельности'),
    ('reports.tourist_categories_percent.read', 'Просмотр процента категорий туристов')
ON CONFLICT (code) DO NOTHING;


DELETE FROM role_permissions
WHERE role_id IN (
    SELECT id
    FROM roles
    WHERE name IN ('tourist', 'manager', 'admin', 'super_admin')
);


INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
         JOIN permissions p ON p.code IN (
                                          'profile.read',
                                          'profile.update',

                                          'own_tours.read',
                                          'own_visas.read',
                                          'own_accommodations.read',
                                          'own_excursions.read',
                                          'own_cargo.read',

                                          'hotels.read',
                                          'hotel_rooms.read',

                                          'excursion_agencies.read',
                                          'excursions.read',
                                          'excursion_schedule.read'
    )
WHERE r.name = 'tourist'
ON CONFLICT (role_id, permission_id) DO NOTHING;


INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
         JOIN permissions p ON p.code IN (
                                          'profile.read',
                                          'profile.update',

                                          'tourists.read',
                                          'tourists.create',
                                          'tourists.update',

                                          'tourist_categories.read',

                                          'tourist_groups.read',
                                          'tourist_groups.create',
                                          'tourist_groups.update',

                                          'group_members.read',
                                          'group_members.create',
                                          'group_members.update',

                                          'child_companions.read',
                                          'child_companions.create',
                                          'child_companions.update',

                                          'identity_documents.read',
                                          'identity_documents.create',
                                          'identity_documents.update',

                                          'visas.read',
                                          'visas.create',
                                          'visas.update',

                                          'hotels.read',
                                          'hotel_rooms.read',

                                          'accommodations.read',
                                          'accommodations.create',
                                          'accommodations.update',

                                          'excursion_agencies.read',

                                          'excursions.read',
                                          'excursions.create',
                                          'excursions.update',

                                          'excursion_schedule.read',
                                          'excursion_schedule.create',
                                          'excursion_schedule.update',

                                          'excursion_bookings.read',
                                          'excursion_bookings.create',
                                          'excursion_bookings.update',

                                          'cargo_types.read',

                                          'flight_types.read',

                                          'flights.read',
                                          'flights.create',
                                          'flights.update',

                                          'cargo_statements.read',
                                          'cargo_statements.create',
                                          'cargo_statements.update',

                                          'cargo_items.read',
                                          'cargo_items.create',
                                          'cargo_items.update',

                                          'cargo_shipments.read',
                                          'cargo_shipments.create',
                                          'cargo_shipments.update',

                                          'financial_categories.read',

                                          'financial_operations.read',
                                          'financial_operations.create',
                                          'financial_operations.update',

                                          'reports.customs_list.read',
                                          'reports.accommodation.read',
                                          'reports.tourists_count.read',
                                          'reports.tourist_info.read',
                                          'reports.hotels.read',
                                          'reports.excursions.read',
                                          'reports.flight_load.read',
                                          'reports.cargo_turnover.read',
                                          'reports.financial.read',
                                          'reports.profitability.read',
                                          'reports.tourist_categories_percent.read'
    )
WHERE r.name = 'manager'
ON CONFLICT (role_id, permission_id) DO NOTHING;


INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
         JOIN permissions p ON p.code IN (
                                          'profile.read',
                                          'profile.update',

                                          'users.read',

                                          'roles.read',
                                          'permissions.read',
                                          'user_roles.read',
                                          'role_permissions.read',

                                          'tourists.read',
                                          'tourists.create',
                                          'tourists.update',
                                          'tourists.delete',

                                          'tourist_categories.read',
                                          'tourist_categories.create',
                                          'tourist_categories.update',
                                          'tourist_categories.delete',

                                          'tourist_groups.read',
                                          'tourist_groups.create',
                                          'tourist_groups.update',
                                          'tourist_groups.delete',

                                          'group_members.read',
                                          'group_members.create',
                                          'group_members.update',
                                          'group_members.delete',

                                          'child_companions.read',
                                          'child_companions.create',
                                          'child_companions.update',
                                          'child_companions.delete',

                                          'identity_documents.read',
                                          'identity_documents.create',
                                          'identity_documents.update',
                                          'identity_documents.delete',

                                          'visas.read',
                                          'visas.create',
                                          'visas.update',
                                          'visas.delete',

                                          'hotels.read',
                                          'hotels.create',
                                          'hotels.update',
                                          'hotels.delete',

                                          'hotel_rooms.read',
                                          'hotel_rooms.create',
                                          'hotel_rooms.update',
                                          'hotel_rooms.delete',

                                          'accommodations.read',
                                          'accommodations.create',
                                          'accommodations.update',
                                          'accommodations.delete',

                                          'excursion_agencies.read',
                                          'excursion_agencies.create',
                                          'excursion_agencies.update',
                                          'excursion_agencies.delete',

                                          'excursions.read',
                                          'excursions.create',
                                          'excursions.update',
                                          'excursions.delete',

                                          'excursion_schedule.read',
                                          'excursion_schedule.create',
                                          'excursion_schedule.update',
                                          'excursion_schedule.delete',

                                          'excursion_bookings.read',
                                          'excursion_bookings.create',
                                          'excursion_bookings.update',
                                          'excursion_bookings.delete',

                                          'cargo_types.read',
                                          'cargo_types.create',
                                          'cargo_types.update',
                                          'cargo_types.delete',

                                          'flight_types.read',
                                          'flight_types.create',
                                          'flight_types.update',
                                          'flight_types.delete',

                                          'flights.read',
                                          'flights.create',
                                          'flights.update',
                                          'flights.delete',

                                          'cargo_statements.read',
                                          'cargo_statements.create',
                                          'cargo_statements.update',
                                          'cargo_statements.delete',

                                          'cargo_items.read',
                                          'cargo_items.create',
                                          'cargo_items.update',
                                          'cargo_items.delete',

                                          'cargo_shipments.read',
                                          'cargo_shipments.create',
                                          'cargo_shipments.update',
                                          'cargo_shipments.delete',

                                          'financial_categories.read',
                                          'financial_categories.create',
                                          'financial_categories.update',
                                          'financial_categories.delete',

                                          'financial_operations.read',
                                          'financial_operations.create',
                                          'financial_operations.update',
                                          'financial_operations.delete',

                                          'reports.customs_list.read',
                                          'reports.accommodation.read',
                                          'reports.tourists_count.read',
                                          'reports.tourist_info.read',
                                          'reports.hotels.read',
                                          'reports.excursions.read',
                                          'reports.flight_load.read',
                                          'reports.cargo_turnover.read',
                                          'reports.financial.read',
                                          'reports.profitability.read',
                                          'reports.tourist_categories_percent.read'
    )
WHERE r.name = 'admin'
ON CONFLICT (role_id, permission_id) DO NOTHING;


INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r
         CROSS JOIN permissions p
WHERE r.name = 'super_admin'
ON CONFLICT (role_id, permission_id) DO NOTHING;

INSERT INTO users (
    login,
    email,
    password_hash
)
VALUES
    (
        'manager',
        'manager@example.com',
        '$2a$10$DygOP/hHQiFZ90.InJZYgeXOmswb/D8KijQ1tpvHvsp3nwooawP7W'
    ),
    (
        'admin',
        'admin@example.com',
        '$2a$10$8omefsCRzTzBIiStkR5bIOphm2f7K6Tt3ozheelwWSL4h8UqfLgsC'
    ),
    (
        'superadmin',
        'superadmin@example.com',
        '$2a$10$HcoC2X72pDO9cgPHIG5Ijet72x3D9lFLzF4kch3nctVAfXcGDHnba'
    )
ON CONFLICT (login) DO UPDATE
    SET
        email = EXCLUDED.email,
        password_hash = EXCLUDED.password_hash;


DELETE FROM user_roles
WHERE user_id IN (
    SELECT id
    FROM users
    WHERE login IN ('manager', 'admin', 'superadmin')
);


INSERT INTO user_roles (
    user_id,
    role_id
)
SELECT u.id, r.id
FROM users u
         JOIN roles r ON
    (u.login = 'manager' AND r.name = 'manager')
        OR (u.login = 'admin' AND r.name = 'admin')
        OR (u.login = 'superadmin' AND r.name = 'super_admin')
WHERE u.login IN ('manager', 'admin', 'superadmin')
ON CONFLICT DO NOTHING;