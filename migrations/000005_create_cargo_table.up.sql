CREATE TABLE cargo_types (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE
);

CREATE TABLE flight_types (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE
);

CREATE TABLE flights (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    capacity INT NOT NULL,
    flight_date TIMESTAMP NOT NULL,
    flight_number INT NOT NULL UNIQUE,
    flight_type_id BIGINT,

    CONSTRAINT chk_flights_capacity
    CHECK (capacity >= 0),

    CONSTRAINT fk_flight_type_id
        FOREIGN KEY (flight_type_id)
        REFERENCES flight_types(id)
);

CREATE TABLE cargo_statements (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    status non_empty_varchar_255 NOT NULL,
    group_member_id BIGINT NOT NULL ,

    CONSTRAINT chk_cargo_statements_status
        CHECK (status IN ('draft', 'weighed', 'packed', 'ready_for_shipment', 'shipped', 'cancelled')),

    CONSTRAINT fk_cargo_statements_group_member_id
        FOREIGN KEY (group_member_id)
        REFERENCES group_members(id)
);

CREATE TABLE cargo_shipments (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    shipped_at TIMESTAMP,
    status non_empty_varchar_255 NOT NULL,
    flight_id BIGINT NOT NULL,
    cargo_statement_id BIGINT NOT NULL,

    CONSTRAINT chk_cargo_shipments_status
        CHECK (status IN ('pending', 'shipped', 'cancelled', 'delivered')),

    CONSTRAINT fk_cargo_shipments_flight_it
        FOREIGN KEY (flight_id)
        REFERENCES flights(id),

    CONSTRAINT fk_cargo_shipments_cargo_statement_id
        FOREIGN KEY (cargo_statement_id)
        REFERENCES cargo_statements(id)
);

CREATE TABLE cargo_items (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    weight_kg NUMERIC(10, 3) NOT NULL,
    volumetric_weight_kg NUMERIC(10, 3) NOT NULL,
    places_count INT NOT NULL,

    marking non_empty_varchar_255,
    packaged_at TIMESTAMPTZ,
    item_number non_empty_varchar_255 NOT NULL,

    cargo_type_id BIGINT NOT NULL,
    cargo_statement_id BIGINT NOT NULL,

    CONSTRAINT chk_cargo_items_weight_kg
        CHECK (weight_kg > 0),

    CONSTRAINT chk_cargo_items_volumetric_weight_kg
        CHECK (volumetric_weight_kg >= 0),

    CONSTRAINT chk_cargo_items_places_count
        CHECK (places_count > 0),

    CONSTRAINT uq_cargo_items_statement_item_number
        UNIQUE (cargo_statement_id, item_number),

    CONSTRAINT fk_cargo_items_cargo_type_id
        FOREIGN KEY (cargo_type_id)
        REFERENCES cargo_types(id),

    CONSTRAINT fk_cargo_items_cargo_statement_id
        FOREIGN KEY (cargo_statement_id)
        REFERENCES cargo_statements(id)
);