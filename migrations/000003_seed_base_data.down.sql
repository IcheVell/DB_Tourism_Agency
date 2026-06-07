DELETE FROM financial_operations
WHERE description IN (
    'Аэропорт встретил группу и счёт тоже',
    'Предоплата экскурсии Тайны кривой улочки'
);

DELETE FROM flights
WHERE flight_number IN (7001, 8801, 9901);

DELETE FROM tourist_groups
WHERE name IN (
    'Группа Майские Чемоданы',
    'Группа Загорелый Шоп-Тур'
);

DELETE FROM excursion_schedule
WHERE start_time IN (
    '2026-06-10 10:00:00+00'::timestamptz,
    '2026-06-11 11:00:00+00'::timestamptz,
    '2026-06-12 19:00:00+00'::timestamptz,
    '2026-06-13 09:30:00+00'::timestamptz,
    '2026-06-14 16:00:00+00'::timestamptz
);

DELETE FROM excursions
WHERE name IN (
    'Тайны кривой улочки',
    'Охота за лучшим магнитиком',
    'Ночной марш по сувенирным лавкам',
    'Легенда о потерянном чемодане',
    'Гастротур Без Шансов на Диету'
);

DELETE FROM excursion_agencies
WHERE name IN (
    'ООО Весёлый Гид',
    'Бюро Потерянный Турист',
    'Агентство Фото у Фонтана'
);

DELETE FROM hotel_rooms
WHERE hotel_id IN (
    SELECT id
    FROM hotels
    WHERE name IN (
        'Отель Сонный Чемодан',
        'Гостиница Уставший Пингвин',
        'Отель Виза Почти Готова'
    )
);

DELETE FROM hotels
WHERE name IN (
    'Отель Сонный Чемодан',
    'Гостиница Уставший Пингвин',
    'Отель Виза Почти Готова'
);

DELETE FROM financial_categories
WHERE name IN (
    'Оплата тура',
    'Оплата экскурсии',
    'Оплата перевозки груза',
    'Оплата страховки груза',
    'Прочий доход',
    'Гостиница',
    'Перевозки',
    'Экскурсии',
    'Визы',
    'Непредвиденные расходы',
    'Аэропорт: загрузка самолета',
    'Аэропорт: разгрузка самолета',
    'Аэропорт: взлет-посадка',
    'Аэропорт: диспетчерские услуги',
    'Хранение груза',
    'Упаковка груза',
    'Страховка груза',
    'Расходы представительства'
);

DELETE FROM flight_types
WHERE name IN ('passenger', 'cargo', 'cargo_passenger');

DELETE FROM cargo_types
WHERE name IN (
    'Одежда',
    'Электроника',
    'Бытовая техника',
    'Сувениры и магнитики',
    'Документы',
    'Чемоданные тайны',
    'Прочее'
);

DELETE FROM tourist_categories
WHERE name IN ('rest', 'shop', 'child');

DELETE FROM user_roles
WHERE user_id IN (
    SELECT id
    FROM users
    WHERE login IN ('manager', 'superadmin')
);

DELETE FROM users
WHERE login IN ('manager', 'superadmin');

DELETE FROM role_permissions
WHERE role_id IN (
    SELECT id
    FROM roles
    WHERE name IN ('tourist', 'manager', 'super_admin')
);

DELETE FROM permissions
WHERE code IN (
    'profile.read', 'profile.update',
    'own_tours.read', 'own_visas.read', 'own_accommodations.read',
    'own_excursions.read', 'own_excursions.create',
    'own_cargo.read', 'own_cargo.create',
    'users.read', 'users.create', 'users.update', 'users.delete',
    'roles.read', 'roles.create', 'roles.update', 'roles.delete',
    'permissions.read', 'permissions.create', 'permissions.update', 'permissions.delete',
    'user_roles.read', 'user_roles.create', 'user_roles.update', 'user_roles.delete',
    'role_permissions.read', 'role_permissions.create', 'role_permissions.update', 'role_permissions.delete',
    'acl.manage',
    'tourists.read', 'tourists.create', 'tourists.update', 'tourists.delete',
    'tourist_categories.read', 'tourist_categories.create', 'tourist_categories.update', 'tourist_categories.delete',
    'tourist_groups.read', 'tourist_groups.create', 'tourist_groups.update', 'tourist_groups.delete',
    'group_members.read', 'group_members.create', 'group_members.update', 'group_members.delete',
    'child_companions.read', 'child_companions.create', 'child_companions.update', 'child_companions.delete',
    'identity_documents.read', 'identity_documents.create', 'identity_documents.update', 'identity_documents.delete',
    'visas.read', 'visas.create', 'visas.update', 'visas.delete',
    'hotels.read', 'hotels.create', 'hotels.update', 'hotels.delete',
    'hotel_rooms.read', 'hotel_rooms.create', 'hotel_rooms.update', 'hotel_rooms.delete',
    'accommodations.read', 'accommodations.create', 'accommodations.update', 'accommodations.delete',
    'excursion_agencies.read', 'excursion_agencies.create', 'excursion_agencies.update', 'excursion_agencies.delete',
    'excursions.read', 'excursions.create', 'excursions.update', 'excursions.delete',
    'excursion_schedule.read', 'excursion_schedule.create', 'excursion_schedule.update', 'excursion_schedule.delete',
    'excursion_bookings.read', 'excursion_bookings.create', 'excursion_bookings.update', 'excursion_bookings.delete',
    'cargo_types.read', 'cargo_types.create', 'cargo_types.update', 'cargo_types.delete',
    'flight_types.read', 'flight_types.create', 'flight_types.update', 'flight_types.delete',
    'flights.read', 'flights.create', 'flights.update', 'flights.delete',
    'cargo_statements.read', 'cargo_statements.create', 'cargo_statements.update', 'cargo_statements.delete',
    'cargo_items.read', 'cargo_items.create', 'cargo_items.update', 'cargo_items.delete',
    'cargo_shipments.read', 'cargo_shipments.create', 'cargo_shipments.update', 'cargo_shipments.delete',
    'financial_categories.read', 'financial_categories.create', 'financial_categories.update', 'financial_categories.delete',
    'financial_operations.read', 'financial_operations.create', 'financial_operations.update', 'financial_operations.delete',
    'reports.customs_list.read', 'reports.accommodation.read', 'reports.tourists_count.read',
    'reports.tourist_info.read', 'reports.hotels.read', 'reports.excursions.read',
    'reports.flight_load.read', 'reports.cargo_turnover.read', 'reports.financial.read',
    'reports.profitability.read', 'reports.tourist_categories_percent.read'
);

DELETE FROM roles
WHERE name IN ('tourist', 'manager', 'super_admin');
