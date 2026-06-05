DROP TRIGGER IF EXISTS check_cargo_item_integrity ON cargo_items;
DROP FUNCTION IF EXISTS trg_check_cargo_item_integrity();

DROP TRIGGER IF EXISTS check_visa_integrity ON visas;
DROP FUNCTION IF EXISTS trg_check_visa_integrity();

DROP TRIGGER IF EXISTS check_cargo_shipment_integrity ON cargo_shipments;
DROP FUNCTION IF EXISTS trg_check_cargo_shipment_integrity();

DROP TRIGGER IF EXISTS check_excursion_booking_integrity ON excursion_bookings;
DROP FUNCTION IF EXISTS trg_check_excursion_booking_integrity();

DROP TRIGGER IF EXISTS check_accommodation_integrity ON accommodations;
DROP FUNCTION IF EXISTS trg_check_accommodation_integrity();

DROP TRIGGER IF EXISTS check_child_companion ON child_companions;
DROP FUNCTION IF EXISTS trg_check_child_companion();

DROP TRIGGER IF EXISTS set_users_updated_at ON users;
DROP FUNCTION IF EXISTS trg_set_updated_at();