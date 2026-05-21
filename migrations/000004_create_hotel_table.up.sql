CREATE TABLE hotels (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    address non_empty_varchar_255 NOT NULL,
    name non_empty_varchar_255 NOT NULL,

    CONSTRAINT uq_hotels_address_name
        UNIQUE (address, name)
);

CREATE TABLE hotel_rooms (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    room_number INT NOT NULL,
    capacity INT NOT NULL,
    price NUMERIC(12, 2) NOT NULL,
    hotel_id BIGINT NOT NULL,

    CONSTRAINT uq_hotel_rooms_room_number
        UNIQUE (id, room_number),

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