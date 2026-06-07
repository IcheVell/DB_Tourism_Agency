CREATE DOMAIN non_empty_varchar_255 AS VARCHAR(255)
    CHECK (btrim(VALUE) <> '');

CREATE TABLE permissions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    code non_empty_varchar_255 NOT NULL UNIQUE,
    description VARCHAR(255)
);

CREATE TABLE roles (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE,
    description VARCHAR(255)
);


CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login non_empty_varchar_255 NOT NULL UNIQUE,
    password_hash non_empty_varchar_255 NOT NULL,
    email non_empty_varchar_255 NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_users_login
        CHECK (login ~ '^[a-zA-Z][a-zA-Z0-9_]{2,49}$'),

    CONSTRAINT chk_users_email
        CHECK (email ~* '^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$')
);

CREATE TABLE role_permissions (
    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    CONSTRAINT pk_role_permissions PRIMARY KEY(role_id, permission_id),

    CONSTRAINT fk_role_permissions_role_id
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_role_permissions_permission_id
        FOREIGN KEY (permission_id)
        REFERENCES permissions(id)
        ON DELETE CASCADE
);

CREATE TABLE user_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    CONSTRAINT pk_user_roles PRIMARY KEY(user_id, role_id),
    CONSTRAINT uq_user_roles_user_id UNIQUE (user_id),

    CONSTRAINT fk_user_roles_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_user_roles_role_id
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE
);


CREATE TABLE tourist_categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE
);

CREATE TABLE tourists (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name non_empty_varchar_255 NOT NULL,
    last_name non_empty_varchar_255 NOT NULL,
    middle_name VARCHAR(50),
    sex VARCHAR(10) NOT NULL,
    birth_date DATE NOT NULL,
    user_id BIGINT UNIQUE,
    desired_hotel_id BIGINT,

    CONSTRAINT chk_tourists_sex
        CHECK (sex IN ('male', 'female')),

    CONSTRAINT fk_tourists_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
);

CREATE TABLE identity_documents (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    document_type VARCHAR(20) NOT NULL,
    document_series non_empty_varchar_255 NOT NULL,
    document_number non_empty_varchar_255 NOT NULL,
    expiration_date DATE,
    issued_by non_empty_varchar_255 NOT NULL,
    issue_date DATE NOT NULL,
    citizenship non_empty_varchar_255 NOT NULL,
    tourist_id BIGINT NOT NULL,

    CONSTRAINT chk_identity_documents_document_type
        CHECK (document_type IN ('PASSPORT', 'BIRTH_CERTIFICATE', 'INTERNATIONAL_PASSPORT')),

    CONSTRAINT chk_identity_documents_dates
        CHECK (expiration_date IS NULL OR expiration_date > issue_date),

    CONSTRAINT uq_identity_documents_document
        UNIQUE (document_type, document_series, document_number, citizenship),

    CONSTRAINT fk_identity_documents_tourist_id
        FOREIGN KEY (tourist_id)
        REFERENCES tourists(id)
);

CREATE TABLE tourist_groups (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    arrival_date TIMESTAMPTZ NOT NULL,
    departure_date TIMESTAMPTZ NOT NULL,
    name non_empty_varchar_255 NOT NULL UNIQUE,

    CONSTRAINT chk_tourist_groups_arrival_less_departure
        CHECK (arrival_date < departure_date)
);

CREATE TABLE group_members (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tourist_group_id BIGINT NOT NULL,
    tourist_category_id BIGINT NOT NULL,
    tourist_id BIGINT NOT NULL,

    CONSTRAINT uq_group_members_group_tourist
        UNIQUE (tourist_group_id, tourist_id),

    CONSTRAINT fk_group_members_tourist_group_id
        FOREIGN KEY (tourist_group_id)
        REFERENCES tourist_groups(id),

    CONSTRAINT fk_group_members_category_id
        FOREIGN KEY (tourist_category_id)
        REFERENCES tourist_categories(id),

    CONSTRAINT fk_group_members_tourist_id
        FOREIGN KEY (tourist_id)
        REFERENCES tourists(id)
);

CREATE TABLE visas (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    number non_empty_varchar_255,
    destination_country non_empty_varchar_255 NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    submitted_at TIMESTAMPTZ,
    decision_at TIMESTAMPTZ,
    issued_at TIMESTAMPTZ,
    valid_from DATE,
    valid_until DATE,
    tourist_id BIGINT NOT NULL,

    CONSTRAINT chk_visas_status
        CHECK (status IN ('draft', 'submitted', 'approved', 'rejected', 'issued', 'cancelled', 'expired')),

    CONSTRAINT chk_visas_timing
        CHECK ((decision_at IS NULL OR submitted_at IS NULL OR submitted_at <= decision_at) AND (issued_at IS NULL OR decision_at IS NULL OR decision_at <= issued_at)),

    CONSTRAINT chk_visas_valid
        CHECK (valid_from IS NULL OR valid_until IS NULL OR valid_from < valid_until),

    CONSTRAINT fk_visas_tourist_id
        FOREIGN KEY (tourist_id)
        REFERENCES tourists(id),

    CONSTRAINT chk_visas_issued_fields
        CHECK (status <> 'issued' OR (number IS NOT NULL AND issued_at IS NOT NULL AND valid_from IS NOT NULL AND valid_until IS NOT NULL))
);

CREATE TABLE child_companions (
    child_group_member_id BIGINT NOT NULL,
    adult_group_member_id BIGINT NOT NULL,

    CONSTRAINT pk_child_companions
        PRIMARY KEY (child_group_member_id, adult_group_member_id),

    CONSTRAINT fk_child_companions_child_id
        FOREIGN KEY (child_group_member_id)
        REFERENCES group_members(id),

    CONSTRAINT fk_child_companions_adult_id
        FOREIGN KEY (adult_group_member_id)
        REFERENCES group_members(id),

    CONSTRAINT chk_child_companions_not_same_member
        CHECK (child_group_member_id <> adult_group_member_id)
);


CREATE TABLE hotels (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    address non_empty_varchar_255 NOT NULL,
    name non_empty_varchar_255 NOT NULL,

    CONSTRAINT uq_hotels_address_name
        UNIQUE (address, name)
);

ALTER TABLE tourists
    ADD CONSTRAINT fk_tourists_desired_hotel_id
        FOREIGN KEY (desired_hotel_id)
        REFERENCES hotels(id);

CREATE TABLE hotel_rooms (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    room_number INT NOT NULL,
    capacity INT NOT NULL,
    price NUMERIC(12, 2) NOT NULL,
    hotel_id BIGINT NOT NULL,

    CONSTRAINT uq_hotel_rooms_room_number
        UNIQUE (hotel_id, room_number),

    CONSTRAINT chk_hotel_rooms_room_number
        CHECK (room_number > 0),

    CONSTRAINT chk_hotel_rooms_capacity
        CHECK (capacity > 0),

    CONSTRAINT chk_hotel_rooms_price
        CHECK (price > 0),

    CONSTRAINT fk_hotel_rooms_hotel_id
        FOREIGN KEY (hotel_id)
        REFERENCES hotels(id)
);

CREATE TABLE accommodations (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    status non_empty_varchar_255 NOT NULL,
    check_in_at TIMESTAMPTZ NOT NULL,
    check_out_at TIMESTAMPTZ,
    group_member_id BIGINT NOT NULL,
    hotel_room_id BIGINT NOT NULL,

    CONSTRAINT chk_accommodations_status
        CHECK (status IN ('reserved', 'checked_in', 'checked_out', 'cancelled')),

    CONSTRAINT chk_accommodations_timing
        CHECK (check_out_at IS NULL OR (check_out_at > check_in_at)),

    CONSTRAINT fk_accommodations_group_member_id
        FOREIGN KEY (group_member_id)
        REFERENCES group_members(id),

    CONSTRAINT fk_accommodations_hotel_room_id
        FOREIGN KEY (hotel_room_id)
        REFERENCES hotel_rooms(id)
);

ALTER TABLE group_members
    ADD COLUMN desired_hotel_id BIGINT;

ALTER TABLE group_members
    ADD CONSTRAINT fk_group_members_desired_hotel_id
        FOREIGN KEY (desired_hotel_id)
        REFERENCES hotels(id);


CREATE TABLE excursion_agencies (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE
);

CREATE TABLE excursions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL,
    description VARCHAR(500),

    CONSTRAINT uq_excursions_name
        UNIQUE (name)
);

CREATE TABLE excursion_schedule (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    price NUMERIC(12, 2) NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    capacity INT NOT NULL,
    status non_empty_varchar_255 NOT NULL,
    excursion_agency_id BIGINT NOT NULL,
    excursion_id BIGINT NOT NULL,

    CONSTRAINT chk_excursion_schedule_timing
        CHECK (start_time < end_time),

    CONSTRAINT chk_excursion_schedule_status
        CHECK (status IN ('planned', 'completed', 'cancelled')),

    CONSTRAINT chk_excursion_schedule_price
        CHECK (price >= 0),

    CONSTRAINT chk_excursion_schedule_capacity
        CHECK (capacity > 0),

    CONSTRAINT fk_excursion_schedule_excursion_agency_id
        FOREIGN KEY (excursion_agency_id)
        REFERENCES excursion_agencies(id),

    CONSTRAINT fk_excursion_schedule_excursion_id
        FOREIGN KEY (excursion_id)
        REFERENCES excursions(id)
);

CREATE TABLE excursion_bookings (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    booked_at TIMESTAMPTZ NOT NULL,
    tourist_rating INT,
    status non_empty_varchar_255 NOT NULL,
    excursion_schedule_id BIGINT NOT NULL,
    group_member_id BIGINT,

    CONSTRAINT chk_excursion_bookings_status
        CHECK (status IN ('booked', 'visited', 'cancelled')),

    CONSTRAINT chk_excursion_bookings_tourist_rating
        CHECK (tourist_rating > 0 AND tourist_rating <= 5),

    CONSTRAINT fk_excursion_booking_excursion_schedule_id
        FOREIGN KEY (excursion_schedule_id)
        REFERENCES excursion_schedule(id),

    CONSTRAINT fk_excursion_bookings_group_member_id
        FOREIGN KEY (group_member_id)
        REFERENCES group_members(id),

    CONSTRAINT uq_excursion_bookings_group_member_id_excursion_schedule_id
        UNIQUE (group_member_id, excursion_schedule_id)
);


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

    CONSTRAINT fk_cargo_shipments_flight_id
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

    CONSTRAINT fk_financial_operations_accommodation_id
        FOREIGN KEY (accommodation_id)
            REFERENCES accommodations(id)
);


CREATE INDEX idx_tourists_user_id ON tourists(user_id);
CREATE INDEX idx_tourists_desired_hotel_id ON tourists(desired_hotel_id);
CREATE INDEX idx_identity_documents_tourist_id ON identity_documents(tourist_id);
CREATE INDEX idx_tourist_groups_dates ON tourist_groups(arrival_date, departure_date);
CREATE INDEX idx_group_members_group_id ON group_members(tourist_group_id);
CREATE INDEX idx_group_members_tourist_id ON group_members(tourist_id);
CREATE INDEX idx_group_members_category_id ON group_members(tourist_category_id);
CREATE INDEX idx_group_members_desired_hotel_id ON group_members(desired_hotel_id);
CREATE INDEX idx_visas_tourist_id ON visas(tourist_id);
CREATE INDEX idx_visas_status ON visas(status);
CREATE INDEX idx_hotel_rooms_hotel_id ON hotel_rooms(hotel_id);
CREATE INDEX idx_accommodations_group_member_id ON accommodations(group_member_id);
CREATE INDEX idx_accommodations_hotel_room_id ON accommodations(hotel_room_id);
CREATE INDEX idx_accommodations_period ON accommodations(check_in_at, check_out_at);
CREATE INDEX idx_excursion_schedule_excursion_id ON excursion_schedule(excursion_id);
CREATE INDEX idx_excursion_schedule_agency_id ON excursion_schedule(excursion_agency_id);
CREATE INDEX idx_excursion_schedule_time ON excursion_schedule(start_time, end_time);
CREATE INDEX idx_excursion_bookings_schedule_id ON excursion_bookings(excursion_schedule_id);
CREATE INDEX idx_excursion_bookings_group_member_id ON excursion_bookings(group_member_id);
CREATE INDEX idx_flights_flight_type_id ON flights(flight_type_id);
CREATE INDEX idx_flights_flight_date ON flights(flight_date);
CREATE INDEX idx_cargo_statements_group_member_id ON cargo_statements(group_member_id);
CREATE INDEX idx_cargo_shipments_flight_id ON cargo_shipments(flight_id);
CREATE INDEX idx_cargo_shipments_statement_id ON cargo_shipments(cargo_statement_id);
CREATE INDEX idx_cargo_items_statement_id ON cargo_items(cargo_statement_id);
CREATE INDEX idx_cargo_items_type_id ON cargo_items(cargo_type_id);
CREATE INDEX idx_financial_operations_category_id ON financial_operations(financial_category_id);
CREATE INDEX idx_financial_operations_operation_at ON financial_operations(operation_at);
CREATE INDEX idx_financial_operations_flight_id ON financial_operations(flight_id);
CREATE INDEX idx_financial_operations_visa_id ON financial_operations(visa_id);
CREATE INDEX idx_financial_operations_excursion_schedule_id ON financial_operations(excursion_schedule_id);
CREATE INDEX idx_financial_operations_excursion_booking_id ON financial_operations(excursion_booking_id);
CREATE INDEX idx_financial_operations_cargo_shipment_id ON financial_operations(cargo_shipment_id);
CREATE INDEX idx_financial_operations_cargo_statement_id ON financial_operations(cargo_statement_id);
CREATE INDEX idx_financial_operations_accommodation_id ON financial_operations(accommodation_id);
