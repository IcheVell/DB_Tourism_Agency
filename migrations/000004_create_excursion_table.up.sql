CREATE TABLE excursion_agencies (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE
);

CREATE TABLE excursions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL,
    description VARCHAR(500)
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