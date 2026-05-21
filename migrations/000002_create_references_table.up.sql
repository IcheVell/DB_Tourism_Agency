CREATE TABLE tourist_categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE
);

CREATE TABLE flight_types (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE
);

CREATE TABLE cargo_types (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE
);

CREATE TABLE financial_categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE,
    operation_type VARCHAR(10) NOT NULL

    CONSTRAINT chk_financial_categories_operation_type
        CHECK (operation_type IN ('income', 'expense'))
);

