ALTER TABLE group_members
    DROP CONSTRAINT fk_group_members_desired_hotel_id;

ALTER TABLE group_members
    DROP COLUMN desired_hotel_id;

DROP TABLE IF EXISTS accommodations;
DROP TABLE IF EXISTS hotel_rooms;
DROP TABLE IF EXISTS hotels;