export type ReportParamType = 'text' | 'number' | 'date' | 'select';

export interface ReportSelectOption {
  label: string;
  value: string | number;
}

export interface ReportSelectConfig {
  endpoint?: string;
  valueField?: string;
  labelFields?: string[];
  staticOptions?: ReportSelectOption[];
}

export interface ReportParamConfig {
  name: string;
  label: string;
  type: ReportParamType;
  required?: boolean;
  select?: ReportSelectConfig;
}

export interface ReportConfig {
  key: string;
  title: string;
  path: string;
  permission: string;
  params: ReportParamConfig[];
}

const relation = (
  endpoint: string,
  labelFields: string[],
  valueField = 'id',
): ReportSelectConfig => ({
  endpoint,
  valueField,
  labelFields,
});

const touristCategorySelect = relation('/api/v1/tourist-categories', ['name']);
const hotelSelect = relation('/api/v1/hotels', ['name', 'address']);
const touristSelect = relation('/api/v1/tourists', ['last_name', 'first_name', 'middle_name', 'id']);
const flightSelect = relation('/api/v1/flights', ['flight_number', 'flight_date']);
const groupSelect = relation('/api/v1/tourist-groups', ['name', 'arrival_date', 'departure_date']);

export const reportConfigs: ReportConfig[] = [
  {
    key: 'customs-tourists',
    title: 'Список туристов для таможни',
    path: '/api/v1/reports/customs-tourists',
    permission: 'reports.customs_list.read',
    params: [
      {
        name: 'category_id',
        label: 'Категория туриста',
        type: 'select',
        select: touristCategorySelect,
      },
    ],
  },
  {
    key: 'accommodation-list',
    title: 'Список на расселение',
    path: '/api/v1/reports/accommodation-list',
    permission: 'reports.accommodation.read',
    params: [
      {
        name: 'hotel_id',
        label: 'Гостиница',
        type: 'select',
        select: hotelSelect,
      },
      {
        name: 'category_id',
        label: 'Категория туриста',
        type: 'select',
        select: touristCategorySelect,
      },
    ],
  },
  {
    key: 'tourist-count',
    title: 'Количество туристов за период',
    path: '/api/v1/reports/tourist-count',
    permission: 'reports.tourists_count.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
      {
        name: 'category_id',
        label: 'Категория туриста',
        type: 'select',
        select: touristCategorySelect,
      },
    ],
  },
  {
    key: 'tourist-info',
    title: 'Сведения о туристе',
    path: '/api/v1/reports/tourist-info',
    permission: 'reports.tourist_info.read',
    params: [
      {
        name: 'tourist_id',
        label: 'Турист',
        type: 'select',
        required: true,
        select: touristSelect,
      },
    ],
  },
  {
    key: 'hotel-occupancy',
    title: 'Занятость гостиниц',
    path: '/api/v1/reports/hotel-occupancy',
    permission: 'reports.hotels.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
    ],
  },
  {
    key: 'excursion-tourist-count',
    title: 'Количество туристов на экскурсиях',
    path: '/api/v1/reports/excursion-tourist-count',
    permission: 'reports.excursions.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
    ],
  },
  {
    key: 'excursion-analytics',
    title: 'Популярные экскурсии и агентства',
    path: '/api/v1/reports/excursion-analytics',
    permission: 'reports.excursions.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
    ],
  },
  {
    key: 'flight-load',
    title: 'Загрузка рейса',
    path: '/api/v1/reports/flight-load',
    permission: 'reports.flight_load.read',
    params: [
      {
        name: 'flight_id',
        label: 'Рейс',
        type: 'select',
        required: true,
        select: flightSelect,
      },
    ],
  },
  {
    key: 'warehouse-turnover',
    title: 'Грузооборот склада',
    path: '/api/v1/reports/warehouse-turnover',
    permission: 'reports.cargo_turnover.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
    ],
  },
  {
    key: 'group-financial-report',
    title: 'Финансовый отчёт по группе',
    path: '/api/v1/reports/group-financial-report',
    permission: 'reports.financial.read',
    params: [
      {
        name: 'group_id',
        label: 'Туристическая группа',
        type: 'select',
        required: true,
        select: groupSelect,
      },
      {
        name: 'category_id',
        label: 'Категория туриста',
        type: 'select',
        select: touristCategorySelect,
      },
    ],
  },
  {
    key: 'income-expense',
    title: 'Доходы и расходы за период',
    path: '/api/v1/reports/income-expense',
    permission: 'reports.financial.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
    ],
  },
  {
    key: 'cargo-type-share',
    title: 'Доля видов груза',
    path: '/api/v1/reports/cargo-type-share',
    permission: 'reports.cargo_turnover.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
    ],
  },
  {
    key: 'profitability',
    title: 'Рентабельность представительства',
    path: '/api/v1/reports/profitability',
    permission: 'reports.profitability.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
    ],
  },
  {
    key: 'tourist-category-ratio',
    title: 'Процент отдыхающих и shop-туристов',
    path: '/api/v1/reports/tourist-category-ratio',
    permission: 'reports.tourist_categories_percent.read',
    params: [
      { name: 'from', label: 'Дата с', type: 'date' },
      { name: 'to', label: 'Дата по', type: 'date' },
      {
        name: 'rest_category_id',
        label: 'Категория отдыхающих',
        type: 'select',
        required: true,
        select: touristCategorySelect,
      },
      {
        name: 'shop_category_id',
        label: 'Категория shop-туристов',
        type: 'select',
        required: true,
        select: touristCategorySelect,
      },
    ],
  },
  {
    key: 'flight-tourists',
    title: 'Туристы указанного рейса',
    path: '/api/v1/reports/flight-tourists',
    permission: 'reports.flight_load.read',
    params: [
      {
        name: 'flight_id',
        label: 'Рейс',
        type: 'select',
        required: true,
        select: flightSelect,
      },
    ],
  },
];
