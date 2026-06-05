# Tourist Agency Frontend

React + Vite + TypeScript клиент для backend `TouristAgencyApp`.

## Запуск

```bash
npm install
cp .env.example .env
npm run dev
```

По умолчанию backend ожидается на `http://localhost:8080`.

## Основная структура

- `src/api` — HTTP-клиент и методы backend API.
- `src/app` — роутер и auth context.
- `src/components` — переиспользуемые UI-блоки.
- `src/pages` — страницы приложения.
- `src/types` — TypeScript-типы.
- `src/config` — конфигурация меню, CRUD-ресурсов и reports.
