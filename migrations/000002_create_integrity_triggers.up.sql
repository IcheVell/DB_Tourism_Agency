CREATE OR REPLACE FUNCTION trg_set_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at := now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE FUNCTION trg_set_updated_at();

CREATE OR REPLACE FUNCTION trg_check_child_companion()
    RETURNS TRIGGER AS $$
DECLARE
    child_group_id BIGINT;
    adult_group_id BIGINT;
    child_category_name TEXT;
    adult_category_name TEXT;
BEGIN
    SELECT gm.tourist_group_id, tc.name
    INTO child_group_id, child_category_name
    FROM group_members gm
             JOIN tourist_categories tc ON tc.id = gm.tourist_category_id
    WHERE gm.id = NEW.child_group_member_id;

    IF child_group_id IS NULL THEN
        RAISE EXCEPTION 'Child group member % does not exist', NEW.child_group_member_id;
    END IF;

    SELECT gm.tourist_group_id, tc.name
    INTO adult_group_id, adult_category_name
    FROM group_members gm
             JOIN tourist_categories tc ON tc.id = gm.tourist_category_id
    WHERE gm.id = NEW.adult_group_member_id;

    IF adult_group_id IS NULL THEN
        RAISE EXCEPTION 'Adult group member % does not exist', NEW.adult_group_member_id;
    END IF;

    IF lower(child_category_name) NOT IN ('child', 'children', 'ребенок', 'ребёнок') THEN
        RAISE EXCEPTION 'Group member % is not a child', NEW.child_group_member_id;
    END IF;

    IF lower(adult_category_name) IN ('child', 'children', 'ребенок', 'ребёнок') THEN
        RAISE EXCEPTION 'Group member % cannot be a child companion because this member is also child', NEW.adult_group_member_id;
    END IF;

    IF child_group_id <> adult_group_id THEN
        RAISE EXCEPTION 'Child and adult companion must belong to the same tourist group';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_child_companion
    BEFORE INSERT OR UPDATE ON child_companions
    FOR EACH ROW
EXECUTE FUNCTION trg_check_child_companion();

CREATE OR REPLACE FUNCTION trg_check_accommodation_integrity()
    RETURNS TRIGGER AS $$
DECLARE
    new_period TSTZRANGE;
BEGIN
    IF NEW.status = 'checked_out' AND NEW.check_out_at IS NULL THEN
        RAISE EXCEPTION 'check_out_at must be set when accommodation status is checked_out';
    END IF;

    IF NEW.check_out_at IS NOT NULL AND NEW.check_out_at <= NEW.check_in_at THEN
        RAISE EXCEPTION 'check_out_at must be greater than check_in_at';
    END IF;

    IF NEW.status IN ('reserved', 'checked_in') THEN
        new_period := tstzrange(
                NEW.check_in_at,
                COALESCE(NEW.check_out_at, 'infinity'::timestamptz),
                '[)'
                      );

        IF EXISTS (
            SELECT 1
            FROM accommodations a
            WHERE a.hotel_room_id = NEW.hotel_room_id
              AND a.id <> COALESCE(NEW.id, -1)
              AND a.status IN ('reserved', 'checked_in')
              AND tstzrange(
                          a.check_in_at,
                          COALESCE(a.check_out_at, 'infinity'::timestamptz),
                          '[)'
                  ) && new_period
        ) THEN
            RAISE EXCEPTION 'Hotel room % is already occupied or reserved for the selected period', NEW.hotel_room_id;
        END IF;

        IF EXISTS (
            SELECT 1
            FROM accommodations a
            WHERE a.group_member_id = NEW.group_member_id
              AND a.id <> COALESCE(NEW.id, -1)
              AND a.status IN ('reserved', 'checked_in')
              AND tstzrange(
                          a.check_in_at,
                          COALESCE(a.check_out_at, 'infinity'::timestamptz),
                          '[)'
                  ) && new_period
        ) THEN
            RAISE EXCEPTION 'Group member % already has active accommodation for the selected period', NEW.group_member_id;
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_accommodation_integrity
    BEFORE INSERT OR UPDATE ON accommodations
    FOR EACH ROW
EXECUTE FUNCTION trg_check_accommodation_integrity();

CREATE OR REPLACE FUNCTION trg_check_excursion_booking_integrity()
    RETURNS TRIGGER AS $$
DECLARE
    schedule_capacity INT;
    schedule_status TEXT;
    active_bookings_count INT;
BEGIN
    SELECT es.capacity, es.status
    INTO schedule_capacity, schedule_status
    FROM excursion_schedule es
    WHERE es.id = NEW.excursion_schedule_id
        FOR UPDATE;

    IF schedule_capacity IS NULL THEN
        RAISE EXCEPTION 'Excursion schedule % does not exist', NEW.excursion_schedule_id;
    END IF;

    IF NEW.status = 'booked' AND schedule_status <> 'planned' THEN
        RAISE EXCEPTION 'Cannot book excursion schedule % because its status is %', NEW.excursion_schedule_id, schedule_status;
    END IF;

    IF NEW.status = 'visited' AND NEW.tourist_rating IS NULL THEN
        RAISE EXCEPTION 'tourist_rating must be set when booking status is visited';
    END IF;

    IF NEW.status <> 'visited' AND NEW.tourist_rating IS NOT NULL THEN
        RAISE EXCEPTION 'tourist_rating can be set only when booking status is visited';
    END IF;

    IF NEW.tourist_rating IS NOT NULL AND (NEW.tourist_rating < 1 OR NEW.tourist_rating > 5) THEN
        RAISE EXCEPTION 'tourist_rating must be between 1 and 5';
    END IF;

    IF NEW.status IN ('booked', 'visited') THEN
        SELECT COUNT(*)
        INTO active_bookings_count
        FROM excursion_bookings eb
        WHERE eb.excursion_schedule_id = NEW.excursion_schedule_id
          AND eb.id <> COALESCE(NEW.id, -1)
          AND eb.status IN ('booked', 'visited');

        IF active_bookings_count + 1 > schedule_capacity THEN
            RAISE EXCEPTION 'Excursion schedule % capacity exceeded', NEW.excursion_schedule_id;
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_excursion_booking_integrity
    BEFORE INSERT OR UPDATE ON excursion_bookings
    FOR EACH ROW
EXECUTE FUNCTION trg_check_excursion_booking_integrity();

CREATE OR REPLACE FUNCTION trg_check_cargo_shipment_integrity()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.tourist_rating IS NOT NULL
        AND NEW.status <> 'visited' THEN
        RAISE EXCEPTION 'tourist_rating can be set only when booking status is visited';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_cargo_shipment_integrity
    BEFORE INSERT OR UPDATE ON cargo_shipments
    FOR EACH ROW
EXECUTE FUNCTION trg_check_cargo_shipment_integrity();

CREATE OR REPLACE FUNCTION trg_check_visa_integrity()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'submitted' AND NEW.submitted_at IS NULL THEN
        RAISE EXCEPTION 'submitted_at must be set when visa status is submitted';
    END IF;

    IF NEW.status = 'rejected'
        AND (NEW.submitted_at IS NULL OR NEW.decision_at IS NULL) THEN
        RAISE EXCEPTION 'submitted_at and decision_at must be set when visa status is rejected';
    END IF;

    IF NEW.status = 'issued'
        AND (
           NEW.number IS NULL
               OR NEW.submitted_at IS NULL
               OR NEW.decision_at IS NULL
               OR NEW.issued_at IS NULL
               OR NEW.valid_from IS NULL
               OR NEW.valid_until IS NULL
           ) THEN
        RAISE EXCEPTION 'number, submitted_at, decision_at, issued_at, valid_from and valid_until must be set when visa status is issued';
    END IF;

    IF NEW.status = 'expired'
        AND (NEW.valid_until IS NULL OR NEW.valid_until >= current_date) THEN
        RAISE EXCEPTION 'visa can have expired status only when valid_until is less than current date';
    END IF;

    IF NEW.valid_from IS NOT NULL
        AND NEW.valid_until IS NOT NULL
        AND NEW.valid_from >= NEW.valid_until THEN
        RAISE EXCEPTION 'valid_from must be less than valid_until';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_visa_integrity
    BEFORE INSERT OR UPDATE ON visas
    FOR EACH ROW
EXECUTE FUNCTION trg_check_visa_integrity();

CREATE OR REPLACE FUNCTION trg_check_cargo_item_integrity()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.weight_kg <= 0 THEN
        RAISE EXCEPTION 'cargo item weight_kg must be greater than zero';
    END IF;

    IF NEW.volumetric_weight_kg < 0 THEN
        RAISE EXCEPTION 'cargo item volumetric_weight_kg must be greater than or equal to zero';
    END IF;

    IF NEW.places_count <= 0 THEN
        RAISE EXCEPTION 'cargo item places_count must be greater than zero';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_cargo_item_integrity
    BEFORE INSERT OR UPDATE ON cargo_items
    FOR EACH ROW
EXECUTE FUNCTION trg_check_cargo_item_integrity();
