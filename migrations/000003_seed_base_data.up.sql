INSERT INTO roles (name, description)
VALUES
    ('tourist', 'Обычный пользователь-турист'),
    ('manager', 'Менеджер представительства туристической фирмы'),
    ('super_admin', 'Суперадминистратор системы')
ON CONFLICT (name) DO UPDATE
SET description = EXCLUDED.description;

INSERT INTO permissions (code, description)
VALUES
    ('profile.read', 'Просмотр собственного профиля'),
    ('profile.update', 'Изменение собственного профиля'),
    ('own_tours.read', 'Просмотр собственных поездок'),
    ('own_visas.read', 'Просмотр собственных виз'),
    ('own_accommodations.read', 'Просмотр собственного расселения'),
    ('own_excursions.read', 'Просмотр собственных экскурсий'),
    ('own_excursions.create', 'Запись на экскурсию от имени текущего туриста'),
    ('own_cargo.read', 'Просмотр собственного груза'),
    ('own_cargo.create', 'Добавление собственного груза'),
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
ON CONFLICT (code) DO UPDATE
SET description = EXCLUDED.description;

DELETE FROM role_permissions
WHERE role_id IN (
    SELECT id
    FROM roles
    WHERE name IN ('tourist', 'manager', 'super_admin')
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
        'own_excursions.create',
        'own_cargo.read',
        'own_cargo.create',
        'hotels.read',
        'hotel_rooms.read',
        'excursion_agencies.read',
        'excursions.read',
        'excursion_schedule.read',
        'cargo_types.read'
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
        'hotels.create',
        'hotels.update',
        'hotel_rooms.read',
        'hotel_rooms.create',
        'hotel_rooms.update',
        'accommodations.read',
        'accommodations.create',
        'accommodations.update',
        'excursion_agencies.read',
        'excursion_agencies.create',
        'excursion_agencies.update',
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
    WHERE login IN ('manager', 'superadmin')
);

INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u
JOIN roles r ON
    (u.login = 'manager' AND r.name = 'manager')
    OR (u.login = 'superadmin' AND r.name = 'super_admin')
WHERE u.login IN ('manager', 'superadmin')
ON CONFLICT DO NOTHING;


INSERT INTO tourist_categories (name)
VALUES
    ('rest'),
    ('shop'),
    ('child')
ON CONFLICT (name) DO NOTHING;

INSERT INTO cargo_types (name)
VALUES
    ('Одежда'),
    ('Электроника'),
    ('Бытовая техника'),
    ('Сувениры и магнитики'),
    ('Документы'),
    ('Чемоданные тайны'),
    ('Прочее')
ON CONFLICT (name) DO NOTHING;

INSERT INTO flight_types (name)
VALUES
    ('passenger'),
    ('cargo'),
    ('cargo_passenger')
ON CONFLICT (name) DO NOTHING;

INSERT INTO financial_categories (name, operation_type)
VALUES
    ('Оплата тура', 'income'),
    ('Оплата экскурсии', 'income'),
    ('Оплата перевозки груза', 'income'),
    ('Оплата страховки груза', 'income'),
    ('Прочий доход', 'income'),
    ('Гостиница', 'expense'),
    ('Перевозки', 'expense'),
    ('Экскурсии', 'expense'),
    ('Визы', 'expense'),
    ('Непредвиденные расходы', 'expense'),
    ('Аэропорт: загрузка самолета', 'expense'),
    ('Аэропорт: разгрузка самолета', 'expense'),
    ('Аэропорт: взлет-посадка', 'expense'),
    ('Аэропорт: диспетчерские услуги', 'expense'),
    ('Хранение груза', 'expense'),
    ('Упаковка груза', 'expense'),
    ('Страховка груза', 'expense'),
    ('Расходы представительства', 'expense')
ON CONFLICT (name) DO UPDATE
SET operation_type = EXCLUDED.operation_type;


INSERT INTO hotels (name, address)
VALUES
    ('Отель Сонный Чемодан', 'Улица Забытых Паспортов, 7'),
    ('Гостиница Уставший Пингвин', 'Площадь Таможенного Досмотра, 3'),
    ('Отель Виза Почти Готова', 'Проспект Бесконечной Регистрации, 14')
ON CONFLICT (address, name) DO NOTHING;

INSERT INTO hotel_rooms (hotel_id, room_number, capacity, price)
SELECT h.id, room_data.room_number, room_data.capacity, room_data.price
FROM hotels h
JOIN (VALUES
    ('Отель Сонный Чемодан', 101, 2, 85.00),
    ('Отель Сонный Чемодан', 102, 3, 120.00),
    ('Гостиница Уставший Пингвин', 201, 2, 95.00),
    ('Гостиница Уставший Пингвин', 202, 4, 150.00),
    ('Отель Виза Почти Готова', 301, 1, 70.00),
    ('Отель Виза Почти Готова', 302, 2, 110.00)
) AS room_data(hotel_name, room_number, capacity, price)
    ON room_data.hotel_name = h.name
ON CONFLICT (hotel_id, room_number) DO NOTHING;


INSERT INTO excursion_agencies (name)
VALUES
    ('ООО Весёлый Гид'),
    ('Бюро Потерянный Турист'),
    ('Агентство Фото у Фонтана')
ON CONFLICT (name) DO NOTHING;

INSERT INTO excursions (name, description)
VALUES
    ('Тайны кривой улочки', 'Пешеходная экскурсия по местам, где навигатор сдаётся первым.'),
    ('Охота за лучшим магнитиком', 'Маршрут по сувенирным лавкам с критическим сравнением магнитиков.'),
    ('Ночной марш по сувенирным лавкам', 'Вечерняя экскурсия для тех, кто днём обещал ничего не покупать.'),
    ('Легенда о потерянном чемодане', 'Городская история о багаже, который видел больше стран, чем турист.'),
    ('Гастротур Без Шансов на Диету', 'Экскурсия по местной кухне с высоким риском полюбить десерты.')
ON CONFLICT DO NOTHING;

INSERT INTO excursion_schedule (
    excursion_id,
    excursion_agency_id,
    start_time,
    end_time,
    capacity,
    price,
    status
)
SELECT e.id, a.id, schedule_data.start_time::timestamptz, schedule_data.end_time::timestamptz, schedule_data.capacity, schedule_data.price, 'planned'
FROM (VALUES
    ('Тайны кривой улочки', 'ООО Весёлый Гид', '2026-06-10 10:00:00+00', '2026-06-10 12:00:00+00', 25, 45.00),
    ('Охота за лучшим магнитиком', 'Бюро Потерянный Турист', '2026-06-11 11:00:00+00', '2026-06-11 13:30:00+00', 18, 35.00),
    ('Ночной марш по сувенирным лавкам', 'Агентство Фото у Фонтана', '2026-06-12 19:00:00+00', '2026-06-12 21:00:00+00', 20, 40.00),
    ('Легенда о потерянном чемодане', 'ООО Весёлый Гид', '2026-06-13 09:30:00+00', '2026-06-13 12:30:00+00', 30, 50.00),
    ('Гастротур Без Шансов на Диету', 'Бюро Потерянный Турист', '2026-06-14 16:00:00+00', '2026-06-14 19:00:00+00', 15, 65.00)
) AS schedule_data(excursion_name, agency_name, start_time, end_time, capacity, price)
JOIN excursions e ON e.name = schedule_data.excursion_name
JOIN excursion_agencies a ON a.name = schedule_data.agency_name
WHERE NOT EXISTS (
    SELECT 1
    FROM excursion_schedule existing_schedule
    WHERE existing_schedule.excursion_id = e.id
      AND existing_schedule.excursion_agency_id = a.id
      AND existing_schedule.start_time = schedule_data.start_time::timestamptz
);


INSERT INTO tourist_groups (name, arrival_date, departure_date)
VALUES
    ('Группа Майские Чемоданы', '2026-06-09 08:00:00+00', '2026-06-19 20:00:00+00'),
    ('Группа Загорелый Шоп-Тур', '2026-07-05 09:30:00+00', '2026-07-15 22:00:00+00')
ON CONFLICT (name) DO UPDATE
SET
    arrival_date = EXCLUDED.arrival_date,
    departure_date = EXCLUDED.departure_date;

INSERT INTO flights (flight_number, flight_date, capacity, flight_type_id)
SELECT flight_data.flight_number, flight_data.flight_date::timestamp, flight_data.capacity, ft.id
FROM (VALUES
    (7001, '2026-06-09 08:00:00', 180, 'passenger'),
    (8801, '2026-06-19 21:00:00', 12000, 'cargo'),
    (9901, '2026-07-05 09:30:00', 220, 'cargo_passenger')
) AS flight_data(flight_number, flight_date, capacity, flight_type_name)
JOIN flight_types ft ON ft.name = flight_data.flight_type_name
ON CONFLICT (flight_number) DO UPDATE
SET
    flight_date = EXCLUDED.flight_date,
    capacity = EXCLUDED.capacity,
    flight_type_id = EXCLUDED.flight_type_id;


INSERT INTO financial_operations (amount, operation_at, description, financial_category_id, flight_id)
SELECT 850.00, '2026-06-09 07:00:00+00'::timestamptz, 'Аэропорт встретил группу и счёт тоже', fc.id, f.id
FROM financial_categories fc
JOIN flights f ON f.flight_number = 7001
WHERE fc.name = 'Аэропорт: загрузка самолета'
  AND NOT EXISTS (
      SELECT 1
      FROM financial_operations fo
      WHERE fo.description = 'Аэропорт встретил группу и счёт тоже'
  );

INSERT INTO financial_operations (amount, operation_at, description, financial_category_id, excursion_schedule_id)
SELECT 675.00, '2026-06-10 09:00:00+00'::timestamptz, 'Предоплата экскурсии Тайны кривой улочки', fc.id, es.id
FROM financial_categories fc
JOIN excursions e ON e.name = 'Тайны кривой улочки'
JOIN excursion_schedule es ON es.excursion_id = e.id
WHERE fc.name = 'Экскурсии'
  AND NOT EXISTS (
      SELECT 1
      FROM financial_operations fo
      WHERE fo.description = 'Предоплата экскурсии Тайны кривой улочки'
  )
LIMIT 1;
