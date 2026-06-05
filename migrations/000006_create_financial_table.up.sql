CREATE TABLE financial_categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE,
    operation_type VARCHAR(10) NOT NULL,

    CONSTRAINT chk_financial_categories_operation_type
        CHECK (operation_type IN ('income', 'expense'))
);

CREATE TABLE financial_operations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    amount NUMERIC(12, 2) NOT NULL,
    operation_at TIMESTAMPTZ NOT NULL,
    description non_empty_varchar_255,
    financial_category_id BIGINT NOT NULL,
    flight_id BIGINT,
    visa_id BIGINT,
    excursion_schedule_id BIGINT,
    excursion_booking_id BIGINT,
    cargo_shipment_id BIGINT,
    cargo_statement_id BIGINT,
    accommodation_id BIGINT,

    CONSTRAINT chk_financial_operations_amount
        CHECK (amount > 0),

    CONSTRAINT chk_financial_operations_source
        CHECK (
            flight_id IS NOT NULL
            OR visa_id IS NOT NULL
            OR excursion_schedule_id IS NOT NULL
            OR excursion_booking_id IS NOT NULL
            OR cargo_shipment_id IS NOT NULL
            OR cargo_statement_id IS NOT NULL
            OR accommodation_id IS NOT NULL
        ),

    CONSTRAINT fk_financial_operations_financial_category_id
        FOREIGN KEY (financial_category_id)
            REFERENCES financial_categories(id),

    CONSTRAINT fk_financial_operations_flight_id
        FOREIGN KEY (flight_id)
        REFERENCES flights(id),

    CONSTRAINT fk_financial_operations_visa_id
        FOREIGN KEY (visa_id)
            REFERENCES visas(id),

    CONSTRAINT fk_financial_operations_excursion_schedule_id
        FOREIGN KEY (excursion_schedule_id)
            REFERENCES excursion_schedule(id),

    CONSTRAINT fk_financial_operations_excursion_booking_id
        FOREIGN KEY (excursion_booking_id)
            REFERENCES excursion_bookings(id),

    CONSTRAINT fk_financial_operations_cargo_shipment_id
        FOREIGN KEY (cargo_shipment_id)
            REFERENCES cargo_shipments(id),

    CONSTRAINT fk_financial_operations_cargo_statement_id
        FOREIGN KEY (cargo_statement_id)
            REFERENCES cargo_statements(id),

    CONSTRAINT fk_financial_operations_accommadation_id
        FOREIGN KEY (accommodation_id)
            REFERENCES accommodations(id)
);