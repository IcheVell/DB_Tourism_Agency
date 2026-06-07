export type ResourceFieldType = 'text' | 'number' | 'date' | 'datetime-local' | 'select';

export interface ResourceSelectOption {
  label: string;
  value: string | number;
}

export interface ResourceSelectConfig {
  endpoint?: string;
  valueField?: string;
  labelFields?: string[];
  staticOptions?: ResourceSelectOption[];
}

export interface ResourceFieldConfig {
  name: string;
  label: string;
  type: ResourceFieldType;
  required?: boolean;
  nullable?: boolean;
  select?: ResourceSelectConfig;
}

export interface ResourceConfig {
  key: string;
  label: string;
  endpoint: string;
  readPermission: string;
  createPermission?: string;
  updatePermission?: string;
  deletePermission?: string;
  createTemplate?: Record<string, unknown>;
  formFields?: ResourceFieldConfig[];
  helperText?: string;
}

const statusOptions = {
  accommodation: [
    { label: 'Забронировано', value: 'reserved' },
    { label: 'Заселён', value: 'checked_in' },
    { label: 'Выселен', value: 'checked_out' },
    { label: 'Отменено', value: 'cancelled' },
  ],
  excursionSchedule: [
    { label: 'Запланирована', value: 'planned' },
    { label: 'Завершена', value: 'completed' },
    { label: 'Отменена', value: 'cancelled' },
  ],
  excursionBooking: [
    { label: 'Записан', value: 'booked' },
    { label: 'Посетил', value: 'visited' },
    { label: 'Отменена', value: 'cancelled' },
  ],
  visa: [
    { label: 'Черновик', value: 'draft' },
    { label: 'Подана', value: 'submitted' },
    { label: 'Одобрена', value: 'approved' },
    { label: 'Отклонена', value: 'rejected' },
    { label: 'Выдана', value: 'issued' },
    { label: 'Отменена', value: 'cancelled' },
    { label: 'Истекла', value: 'expired' },
  ],
  cargoStatement: [
    { label: 'Черновик', value: 'draft' },
    { label: 'Взвешено', value: 'weighed' },
    { label: 'Упаковано', value: 'packed' },
    { label: 'Готово к отправке', value: 'ready_for_shipment' },
    { label: 'Отправлено', value: 'shipped' },
    { label: 'Отменено', value: 'cancelled' },
  ],
  cargoShipment: [
    { label: 'Ожидает отправки', value: 'pending' },
    { label: 'Отправлено', value: 'shipped' },
    { label: 'Отменено', value: 'cancelled' },
    { label: 'Доставлено', value: 'delivered' },
  ],
  financialOperationType: [
    { label: 'Доход', value: 'income' },
    { label: 'Расход', value: 'expense' },
  ],
  sex: [
    { label: 'Мужской', value: 'male' },
    { label: 'Женский', value: 'female' },
  ],
  documentType: [
    { label: 'Паспорт', value: 'PASSPORT' },
    { label: 'Заграничный паспорт', value: 'INTERNATIONAL_PASSPORT' },
    { label: 'Свидетельство о рождении', value: 'BIRTH_CERTIFICATE' },
  ],
};

const relation = (endpoint: string, labelFields: string[], valueField = 'id'): ResourceSelectConfig => ({
  endpoint,
  valueField,
  labelFields,
});

export const resources: Record<string, ResourceConfig> = {
  hotels: {
    key: 'hotels',
    label: 'Гостиницы',
    endpoint: '/api/v1/hotels',
    readPermission: 'hotels.read',
    createPermission: 'hotels.create',
    updatePermission: 'hotels.update',
    deletePermission: 'hotels.delete',
    createTemplate: { name: '', address: '' },
    formFields: [
      { name: 'name', label: 'Название', type: 'text', required: true },
      { name: 'address', label: 'Адрес', type: 'text', required: true },
    ],
  },

  hotelRooms: {
    key: 'hotelRooms',
    label: 'Номера',
    endpoint: '/api/v1/hotel-rooms',
    readPermission: 'hotel_rooms.read',
    createPermission: 'hotel_rooms.create',
    updatePermission: 'hotel_rooms.update',
    deletePermission: 'hotel_rooms.delete',
    helperText: 'Номер привязывается к гостинице через поле «Гостиница».',
    createTemplate: { hotel_id: 1, room_number: 101, capacity: 2, price: 100 },
    formFields: [
      { name: 'hotel_id', label: 'Гостиница', type: 'select', required: true, select: relation('/api/v1/hotels', ['name', 'address']) },
      { name: 'room_number', label: 'Номер комнаты', type: 'number', required: true },
      { name: 'capacity', label: 'Вместимость', type: 'number', required: true },
      { name: 'price', label: 'Цена', type: 'number', required: true },
    ],
  },

  accommodations: {
    key: 'accommodations',
    label: 'Расселение',
    endpoint: '/api/v1/accommodations',
    readPermission: 'accommodations.read',
    createPermission: 'accommodations.create',
    updatePermission: 'accommodations.update',
    deletePermission: 'accommodations.delete',
    helperText: 'Расселение связывает участника группы с конкретным номером гостиницы.',
    createTemplate: { status: 'reserved', check_in_at: '2026-06-03T14:00', check_out_at: '2026-06-10T12:00', hotel_room_id: 1, group_member_id: 1 },
    formFields: [
      { name: 'group_member_id', label: 'Участник группы', type: 'select', required: true, select: relation('/api/v1/group-members', ['id', 'tourist_id', 'tourist_group_id']) },
      { name: 'hotel_room_id', label: 'Номер гостиницы', type: 'select', required: true, select: relation('/api/v1/hotel-rooms', ['id', 'room_number', 'hotel_id']) },
      { name: 'status', label: 'Статус', type: 'select', required: true, select: { staticOptions: statusOptions.accommodation } },
      { name: 'check_in_at', label: 'Дата/время заселения', type: 'datetime-local', required: true },
      { name: 'check_out_at', label: 'Дата/время выезда', type: 'datetime-local', nullable: true },
    ],
  },

  excursionAgencies: {
    key: 'excursionAgencies',
    label: 'Агентства экскурсий',
    endpoint: '/api/v1/excursion-agencies',
    readPermission: 'excursion_agencies.read',
    createPermission: 'excursion_agencies.create',
    updatePermission: 'excursion_agencies.update',
    deletePermission: 'excursion_agencies.delete',
    createTemplate: { name: '' },
    formFields: [{ name: 'name', label: 'Название агентства', type: 'text', required: true }],
  },

  excursions: {
    key: 'excursions',
    label: 'Экскурсии',
    endpoint: '/api/v1/excursions',
    readPermission: 'excursions.read',
    createPermission: 'excursions.create',
    updatePermission: 'excursions.update',
    deletePermission: 'excursions.delete',
    helperText: 'Сама экскурсия — справочник маршрутов. Связь с агентством задаётся в расписании экскурсии.',
    createTemplate: { name: '', description: '' },
    formFields: [
      { name: 'name', label: 'Название экскурсии', type: 'text', required: true },
      { name: 'description', label: 'Описание', type: 'text', nullable: true },
    ],
  },

  excursionSchedules: {
    key: 'excursionSchedules',
    label: 'Расписание экскурсий',
    endpoint: '/api/v1/excursion-schedules',
    readPermission: 'excursion_schedule.read',
    createPermission: 'excursion_schedule.create',
    updatePermission: 'excursion_schedule.update',
    deletePermission: 'excursion_schedule.delete',
    helperText: 'Здесь экскурсия связывается с агентством: выбери экскурсию и агентство в выпадающих списках.',
    createTemplate: { excursion_id: 1, excursion_agency_id: 1, start_time: '2026-06-03T10:00', end_time: '2026-06-03T12:00', capacity: 20, price: 100, status: 'planned' },
    formFields: [
      { name: 'excursion_id', label: 'Экскурсия', type: 'select', required: true, select: relation('/api/v1/excursions', ['name']) },
      { name: 'excursion_agency_id', label: 'Агентство', type: 'select', required: true, select: relation('/api/v1/excursion-agencies', ['name']) },
      { name: 'start_time', label: 'Начало', type: 'datetime-local', required: true },
      { name: 'end_time', label: 'Окончание', type: 'datetime-local', required: true },
      { name: 'capacity', label: 'Мест', type: 'number', required: true },
      { name: 'price', label: 'Цена', type: 'number', required: true },
      { name: 'status', label: 'Статус', type: 'select', required: true, select: { staticOptions: statusOptions.excursionSchedule } },
    ],
  },

  excursionBookings: {
    key: 'excursionBookings',
    label: 'Записи на экскурсии',
    endpoint: '/api/v1/excursion-bookings',
    readPermission: 'excursion_bookings.read',
    createPermission: 'excursion_bookings.create',
    updatePermission: 'excursion_bookings.update',
    deletePermission: 'excursion_bookings.delete',
    createTemplate: { booked_at: '2026-06-03T10:00', excursion_schedule_id: 1, group_member_id: 1, status: 'booked', tourist_rating: null },
    formFields: [
      { name: 'excursion_schedule_id', label: 'Расписание экскурсии', type: 'select', required: true, select: relation('/api/v1/excursion-schedules', ['id', 'excursion_id', 'excursion_agency_id', 'start_time']) },
      { name: 'group_member_id', label: 'Участник группы', type: 'select', required: true, select: relation('/api/v1/group-members', ['id', 'tourist_id', 'tourist_group_id']) },
      { name: 'booked_at', label: 'Дата записи', type: 'datetime-local', required: true },
      { name: 'status', label: 'Статус', type: 'select', required: true, select: { staticOptions: statusOptions.excursionBooking } },
      { name: 'tourist_rating', label: 'Оценка туриста', type: 'number', nullable: true },
    ],
  },

  tourists: {
    key: 'tourists',
    label: 'Туристы',
    endpoint: '/api/v1/tourists',
    readPermission: 'tourists.read',
    createPermission: 'tourists.create',
    updatePermission: 'tourists.update',
    deletePermission: 'tourists.delete',
    helperText: 'Желаемая гостиница — это пожелание туриста. Фактическое расселение создаётся отдельно в разделе «Расселение».',
    createTemplate: { first_name: '', last_name: '', middle_name: null, sex: 'male', birth_date: '2000-01-01', user_id: null, desired_hotel_id: null },
    formFields: [
      { name: 'last_name', label: 'Фамилия', type: 'text', required: true },
      { name: 'first_name', label: 'Имя', type: 'text', required: true },
      { name: 'middle_name', label: 'Отчество', type: 'text', nullable: true },
      { name: 'sex', label: 'Пол', type: 'select', required: true, select: { staticOptions: statusOptions.sex } },
      { name: 'birth_date', label: 'Дата рождения', type: 'date', required: true },
      { name: 'user_id', label: 'Аккаунт пользователя', type: 'select', nullable: true, select: relation('/api/v1/users', ['login', 'email']) },
      { name: 'desired_hotel_id', label: 'Желаемая гостиница', type: 'select', nullable: true, select: relation('/api/v1/hotels', ['name', 'address']) },
    ],
  },

  touristCategories: {
    key: 'touristCategories',
    label: 'Категории туристов',
    endpoint: '/api/v1/tourist-categories',
    readPermission: 'tourist_categories.read',
    createPermission: 'tourist_categories.create',
    updatePermission: 'tourist_categories.update',
    deletePermission: 'tourist_categories.delete',
    createTemplate: { name: 'rest' },
    formFields: [{ name: 'name', label: 'Код категории', type: 'text', required: true }],
  },

  touristGroups: {
    key: 'touristGroups',
    label: 'Туристические группы',
    endpoint: '/api/v1/tourist-groups',
    readPermission: 'tourist_groups.read',
    createPermission: 'tourist_groups.create',
    updatePermission: 'tourist_groups.update',
    deletePermission: 'tourist_groups.delete',
    createTemplate: { name: '', arrival_date: '2026-06-03T10:00', departure_date: '2026-06-10T10:00' },
    formFields: [
      { name: 'name', label: 'Название группы', type: 'text', required: true },
      { name: 'arrival_date', label: 'Дата/время прилёта', type: 'datetime-local', required: true },
      { name: 'departure_date', label: 'Дата/время вылета', type: 'datetime-local', required: true },
    ],
  },

  groupMembers: {
    key: 'groupMembers',
    label: 'Участники групп',
    endpoint: '/api/v1/group-members',
    readPermission: 'group_members.read',
    createPermission: 'group_members.create',
    updatePermission: 'group_members.update',
    deletePermission: 'group_members.delete',
    helperText: 'Менеджер добавляет туриста в группу. Если «Желаемая гостиница» не выбрана, backend автоматически возьмёт гостиницу из заявки туриста.',
    createTemplate: { tourist_group_id: 1, tourist_category_id: 1, tourist_id: 1, desired_hotel_id: null },
    formFields: [
      { name: 'tourist_group_id', label: 'Группа', type: 'select', required: true, select: relation('/api/v1/tourist-groups', ['name', 'arrival_date']) },
      { name: 'tourist_id', label: 'Турист', type: 'select', required: true, select: relation('/api/v1/tourists', ['last_name', 'first_name', 'birth_date']) },
      { name: 'tourist_category_id', label: 'Категория туриста', type: 'select', required: true, select: relation('/api/v1/tourist-categories', ['name']) },
      { name: 'desired_hotel_id', label: 'Желаемая гостиница', type: 'select', nullable: true, select: relation('/api/v1/hotels', ['name', 'address']) },
    ],
  },

  childCompanions: {
    key: 'childCompanions',
    label: 'Сопровождение детей',
    endpoint: '/api/v1/child-companions',
    readPermission: 'child_companions.read',
    createPermission: 'child_companions.create',
    updatePermission: 'child_companions.update',
    deletePermission: 'child_companions.delete',
    createTemplate: { adult_group_member_id: 1, child_group_member_id: 2 },
    formFields: [
      { name: 'adult_group_member_id', label: 'Взрослый сопровождающий', type: 'select', required: true, select: relation('/api/v1/group-members', ['id', 'tourist_id', 'tourist_category_id']) },
      { name: 'child_group_member_id', label: 'Ребёнок', type: 'select', required: true, select: relation('/api/v1/group-members', ['id', 'tourist_id', 'tourist_category_id']) },
    ],
  },

  identityDocuments: {
    key: 'identityDocuments',
    label: 'Документы туристов',
    endpoint: '/api/v1/identity-documents',
    readPermission: 'identity_documents.read',
    createPermission: 'identity_documents.create',
    updatePermission: 'identity_documents.update',
    deletePermission: 'identity_documents.delete',
    createTemplate: { tourist_id: 1, document_type: 'PASSPORT', document_series: '', document_number: '', issue_date: '2020-01-01', expiration_date: null, issued_by: '', citizenship: 'Россия' },
    formFields: [
      { name: 'tourist_id', label: 'Турист', type: 'select', required: true, select: relation('/api/v1/tourists', ['last_name', 'first_name', 'birth_date']) },
      { name: 'document_type', label: 'Тип документа', type: 'select', required: true, select: { staticOptions: statusOptions.documentType } },
      { name: 'document_series', label: 'Серия', type: 'text', required: true },
      { name: 'document_number', label: 'Номер', type: 'text', required: true },
      { name: 'issue_date', label: 'Дата выдачи', type: 'date', required: true },
      { name: 'expiration_date', label: 'Действителен до', type: 'date', nullable: true },
      { name: 'issued_by', label: 'Кем выдан', type: 'text', required: true },
      { name: 'citizenship', label: 'Гражданство', type: 'text', required: true },
    ],
  },

  visas: {
    key: 'visas',
    label: 'Визы',
    endpoint: '/api/v1/visas',
    readPermission: 'visas.read',
    createPermission: 'visas.create',
    updatePermission: 'visas.update',
    deletePermission: 'visas.delete',
    createTemplate: {
      number: null,
      destination_country: '',
      status: 'draft',
      submitted_at: null,
      decision_at: null,
      issued_at: null,
      valid_from: null,
      valid_until: null,
      tourist_id: 1,
    },
    formFields: [
      { name: 'tourist_id', label: 'Турист', type: 'select', required: true, select: relation('/api/v1/tourists', ['last_name', 'first_name', 'birth_date']) },
      { name: 'destination_country', label: 'Страна назначения', type: 'text', required: true },
      { name: 'number', label: 'Номер визы', type: 'text', nullable: true },
      { name: 'status', label: 'Статус', type: 'select', required: true, select: { staticOptions: statusOptions.visa } },
      { name: 'submitted_at', label: 'Дата подачи', type: 'datetime-local', nullable: true },
      { name: 'decision_at', label: 'Дата решения', type: 'datetime-local', nullable: true },
      { name: 'issued_at', label: 'Дата выдачи', type: 'datetime-local', nullable: true },
      { name: 'valid_from', label: 'Действует с', type: 'date', nullable: true },
      { name: 'valid_until', label: 'Действует до', type: 'date', nullable: true },
    ],
  },

  flightTypes: {
    key: 'flightTypes',
    label: 'Типы рейсов',
    endpoint: '/api/v1/flight-types',
    readPermission: 'flight_types.read',
    createPermission: 'flight_types.create',
    updatePermission: 'flight_types.update',
    deletePermission: 'flight_types.delete',
    createTemplate: { name: 'cargo_passenger' },
    formFields: [{ name: 'name', label: 'Название типа рейса', type: 'text', required: true }],
  },

  flights: {
    key: 'flights',
    label: 'Рейсы',
    endpoint: '/api/v1/flights',
    readPermission: 'flights.read',
    createPermission: 'flights.create',
    updatePermission: 'flights.update',
    deletePermission: 'flights.delete',
    createTemplate: { flight_number: 1001, flight_date: '2026-06-03', capacity: 180, flight_type_id: 1 },
    formFields: [
      { name: 'flight_number', label: 'Номер рейса', type: 'number', required: true },
      { name: 'flight_date', label: 'Дата рейса', type: 'date', required: true },
      { name: 'capacity', label: 'Вместимость', type: 'number', required: true },
      { name: 'flight_type_id', label: 'Тип рейса', type: 'select', required: true, select: relation('/api/v1/flight-types', ['name']) },
    ],
  },

  cargoTypes: {
    key: 'cargoTypes',
    label: 'Виды груза',
    endpoint: '/api/v1/cargo-types',
    readPermission: 'cargo_types.read',
    createPermission: 'cargo_types.create',
    updatePermission: 'cargo_types.update',
    deletePermission: 'cargo_types.delete',
    createTemplate: { name: 'Прочее' },
    formFields: [{ name: 'name', label: 'Название вида груза', type: 'text', required: true }],
  },

  cargoStatements: {
    key: 'cargoStatements',
    label: 'Грузовые ведомости',
    endpoint: '/api/v1/cargo-statements',
    readPermission: 'cargo_statements.read',
    createPermission: 'cargo_statements.create',
    updatePermission: 'cargo_statements.update',
    deletePermission: 'cargo_statements.delete',
    createTemplate: { status: 'draft', group_member_id: 1 },
    formFields: [
      { name: 'group_member_id', label: 'Участник группы', type: 'select', required: true, select: relation('/api/v1/group-members', ['id', 'tourist_id', 'tourist_group_id']) },
      { name: 'status', label: 'Статус ведомости', type: 'select', required: true, select: { staticOptions: statusOptions.cargoStatement } },
    ],
  },

  cargoItems: {
    key: 'cargoItems',
    label: 'Грузовые места',
    endpoint: '/api/v1/cargo-items',
    readPermission: 'cargo_items.read',
    createPermission: 'cargo_items.create',
    updatePermission: 'cargo_items.update',
    deletePermission: 'cargo_items.delete',
    createTemplate: { item_number: '', weight_kg: 1, volumetric_weight_kg: 0, places_count: 1, marking: null, packaged_at: null, cargo_type_id: 1, cargo_statement_id: 1 },
    formFields: [
      { name: 'cargo_statement_id', label: 'Грузовая ведомость', type: 'select', required: true, select: relation('/api/v1/cargo-statements', ['id', 'status', 'group_member_id']) },
      { name: 'cargo_type_id', label: 'Вид груза', type: 'select', required: true, select: relation('/api/v1/cargo-types', ['name']) },
      { name: 'item_number', label: 'Номер места', type: 'text', required: true },
      { name: 'weight_kg', label: 'Вес, кг', type: 'number', required: true },
      { name: 'volumetric_weight_kg', label: 'Объёмный вес, кг', type: 'number', required: true },
      { name: 'places_count', label: 'Количество мест', type: 'number', required: true },
      { name: 'marking', label: 'Маркировка', type: 'text', nullable: true },
      { name: 'packaged_at', label: 'Дата/время упаковки', type: 'datetime-local', nullable: true },
    ],
  },

  cargoShipments: {
    key: 'cargoShipments',
    label: 'Отправки груза',
    endpoint: '/api/v1/cargo-shipments',
    readPermission: 'cargo_shipments.read',
    createPermission: 'cargo_shipments.create',
    updatePermission: 'cargo_shipments.update',
    deletePermission: 'cargo_shipments.delete',
    helperText: 'Здесь грузовая ведомость связывается с рейсом: выбери рейс и ведомость.',
    createTemplate: { shipped_at: null, status: 'pending', flight_id: 1, cargo_statement_id: 1 },
    formFields: [
      { name: 'flight_id', label: 'Рейс', type: 'select', required: true, select: relation('/api/v1/flights', ['flight_number', 'flight_date', 'capacity']) },
      { name: 'cargo_statement_id', label: 'Грузовая ведомость', type: 'select', required: true, select: relation('/api/v1/cargo-statements', ['id', 'status', 'group_member_id']) },
      { name: 'status', label: 'Статус отправки', type: 'select', required: true, select: { staticOptions: statusOptions.cargoShipment } },
      { name: 'shipped_at', label: 'Дата/время отправки', type: 'datetime-local', nullable: true },
    ],
  },

  financialCategories: {
    key: 'financialCategories',
    label: 'Финансовые категории',
    endpoint: '/api/v1/financial-categories',
    readPermission: 'financial_categories.read',
    createPermission: 'financial_categories.create',
    updatePermission: 'financial_categories.update',
    deletePermission: 'financial_categories.delete',
    createTemplate: { name: '', operation_type: 'expense' },
    formFields: [
      { name: 'name', label: 'Название категории', type: 'text', required: true },
      { name: 'operation_type', label: 'Тип операции', type: 'select', required: true, select: { staticOptions: statusOptions.financialOperationType } },
    ],
  },

  financialOperations: {
    key: 'financialOperations',
    label: 'Финансовые операции',
    endpoint: '/api/v1/financial-operations',
    readPermission: 'financial_operations.read',
    createPermission: 'financial_operations.create',
    updatePermission: 'financial_operations.update',
    deletePermission: 'financial_operations.delete',
    createTemplate: { amount: 100, operation_at: '2026-06-03T12:00', description: '', financial_category_id: 1, flight_id: null, visa_id: null, excursion_schedule_id: null, excursion_booking_id: null, cargo_shipment_id: null, cargo_statement_id: null, accommodation_id: null },
    formFields: [
      { name: 'financial_category_id', label: 'Финансовая категория', type: 'select', required: true, select: relation('/api/v1/financial-categories', ['name', 'operation_type']) },
      { name: 'amount', label: 'Сумма', type: 'number', required: true },
      { name: 'operation_at', label: 'Дата/время операции', type: 'datetime-local', required: true },
      { name: 'description', label: 'Описание', type: 'text', nullable: true },
      { name: 'flight_id', label: 'Связанный рейс', type: 'select', nullable: true, select: relation('/api/v1/flights', ['flight_number', 'flight_date']) },
      { name: 'visa_id', label: 'Связанная виза', type: 'select', nullable: true, select: relation('/api/v1/visas', ['id', 'destination_country', 'status']) },
      { name: 'excursion_schedule_id', label: 'Связанное расписание экскурсии', type: 'select', nullable: true, select: relation('/api/v1/excursion-schedules', ['id', 'start_time', 'excursion_id']) },
      { name: 'excursion_booking_id', label: 'Связанная запись на экскурсию', type: 'select', nullable: true, select: relation('/api/v1/excursion-bookings', ['id', 'status', 'group_member_id']) },
      { name: 'cargo_shipment_id', label: 'Связанная отправка груза', type: 'select', nullable: true, select: relation('/api/v1/cargo-shipments', ['id', 'status', 'flight_id']) },
      { name: 'cargo_statement_id', label: 'Связанная грузовая ведомость', type: 'select', nullable: true, select: relation('/api/v1/cargo-statements', ['id', 'status', 'group_member_id']) },
      { name: 'accommodation_id', label: 'Связанное расселение', type: 'select', nullable: true, select: relation('/api/v1/accommodations', ['id', 'status', 'group_member_id']) },
    ],
  },

  users: {
    key: 'users',
    label: 'Пользователи',
    endpoint: '/api/v1/users',
    readPermission: 'users.read',
    createPermission: 'users.create',
    updatePermission: 'users.update',
    deletePermission: 'users.delete',
    createTemplate: { login: '', email: '', password: 'password123', role_id: 1 },
    formFields: [
      { name: 'login', label: 'Логин', type: 'text', required: true },
      { name: 'email', label: 'Email', type: 'text', required: true },
      { name: 'password', label: 'Пароль', type: 'text', required: true },
      { name: 'role_id', label: 'Роль', type: 'select', required: true, select: relation('/api/v1/roles', ['name', 'description']) },
    ],
  },

  roles: {
    key: 'roles',
    label: 'Роли',
    endpoint: '/api/v1/roles',
    readPermission: 'roles.read',
    createPermission: 'roles.create',
    updatePermission: 'roles.update',
    deletePermission: 'roles.delete',
    createTemplate: { name: '', description: '' },
    formFields: [
      { name: 'name', label: 'Название роли', type: 'text', required: true },
      { name: 'description', label: 'Описание', type: 'text', nullable: true },
    ],
  },

  permissions: {
    key: 'permissions',
    label: 'Права',
    endpoint: '/api/v1/permissions',
    readPermission: 'permissions.read',
    createPermission: 'permissions.create',
    updatePermission: 'permissions.update',
    deletePermission: 'permissions.delete',
    createTemplate: { code: '', description: '' },
    formFields: [
      { name: 'code', label: 'Код права', type: 'text', required: true },
      { name: 'description', label: 'Описание', type: 'text', nullable: true },
    ],
  },

  userRoles: {
    key: 'userRoles',
    label: 'Роли пользователей',
    endpoint: '/api/v1/user-roles',
    readPermission: 'user_roles.read',
    createPermission: 'user_roles.create',
    updatePermission: 'user_roles.update',
    deletePermission: 'user_roles.delete',
    helperText: 'Обычная роль admin удалена. Назначай пользователям tourist, manager или super_admin.',
    createTemplate: { user_id: 1, role_id: 1 },
    formFields: [
      { name: 'user_id', label: 'Пользователь', type: 'select', required: true, select: relation('/api/v1/users', ['login', 'email']) },
      { name: 'role_id', label: 'Роль', type: 'select', required: true, select: relation('/api/v1/roles', ['name', 'description']) },
    ],
  },

  rolePermissions: {
    key: 'rolePermissions',
    label: 'Права ролей',
    endpoint: '/api/v1/role-permissions',
    readPermission: 'role_permissions.read',
    createPermission: 'role_permissions.create',
    updatePermission: 'role_permissions.update',
    deletePermission: 'role_permissions.delete',
    createTemplate: { role_id: 1, permission_id: 1 },
    formFields: [
      { name: 'role_id', label: 'Роль', type: 'select', required: true, select: relation('/api/v1/roles', ['name', 'description']) },
      { name: 'permission_id', label: 'Право', type: 'select', required: true, select: relation('/api/v1/permissions', ['code', 'description']) },
    ],
  },
};
