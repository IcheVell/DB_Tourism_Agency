package repository

import (
	"TouristAgencyApp/src/dto"

	"gorm.io/gorm"
)

type MeRepository struct {
	db *gorm.DB
}

func NewMeRepository(db *gorm.DB) *MeRepository {
	return &MeRepository{
		db: db,
	}
}

func (r *MeRepository) FindTouristIDByUserID(userID int64) (*int64, error) {
	var touristID int64

	err := r.db.
		Raw(`
			SELECT id
			FROM tourists
			WHERE user_id = ?
		`, userID).
		Scan(&touristID).
		Error

	if err != nil {
		return nil, err
	}

	if touristID == 0 {
		return nil, nil
	}

	return &touristID, nil
}

func (r *MeRepository) FindToursByUserID(userID int64, limit int, offset int) ([]dto.MeTourResponse, int64, error) {
	var items []dto.MeTourResponse
	var total int64

	if err := r.db.Raw(`
		SELECT COUNT(*)
		FROM group_members gm
		JOIN tourists t ON t.id = gm.tourist_id
		WHERE t.user_id = ?
	`, userID).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Raw(`
		SELECT
			gm.id AS group_member_id,
			tg.id AS group_id,
			tg.name AS group_name,
			tg.arrival_date,
			tg.departure_date,
			tc.id AS category_id,
			tc.name AS category_name,
			h.id AS desired_hotel_id,
			h.name AS desired_hotel_name,
			h.address AS desired_hotel_address
		FROM group_members gm
		JOIN tourists t ON t.id = gm.tourist_id
		JOIN tourist_groups tg ON tg.id = gm.tourist_group_id
		JOIN tourist_categories tc ON tc.id = gm.tourist_category_id
		LEFT JOIN hotels h ON h.id = gm.desired_hotel_id
		WHERE t.user_id = ?
		ORDER BY tg.arrival_date DESC, tg.id DESC
		LIMIT ? OFFSET ?
	`, userID, limit, offset).Scan(&items).Error

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *MeRepository) FindVisasByUserID(userID int64, limit int, offset int) ([]dto.MeVisaResponse, int64, error) {
	var items []dto.MeVisaResponse
	var total int64

	if err := r.db.Raw(`
		SELECT COUNT(*)
		FROM visas v
		JOIN tourists t ON t.id = v.tourist_id
		WHERE t.user_id = ?
	`, userID).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Raw(`
		SELECT
			v.id,
			v.number,
			v.destination_country,
			v.status,
			v.created_at,
			v.submitted_at,
			v.decision_at,
			v.issued_at,
			v.valid_from,
			v.valid_until
		FROM visas v
		JOIN tourists t ON t.id = v.tourist_id
		WHERE t.user_id = ?
		ORDER BY v.created_at DESC, v.id DESC
		LIMIT ? OFFSET ?
	`, userID, limit, offset).Scan(&items).Error

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *MeRepository) FindAccommodationsByUserID(userID int64, limit int, offset int) ([]dto.MeAccommodationResponse, int64, error) {
	var items []dto.MeAccommodationResponse
	var total int64

	if err := r.db.Raw(`
		SELECT COUNT(*)
		FROM accommodations a
		JOIN group_members gm ON gm.id = a.group_member_id
		JOIN tourists t ON t.id = gm.tourist_id
		WHERE t.user_id = ?
	`, userID).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Raw(`
		SELECT
			a.id,
			a.status,
			a.check_in_at,
			a.check_out_at,
			h.id AS hotel_id,
			h.name AS hotel_name,
			h.address AS hotel_address,
			hr.id AS room_id,
			hr.room_number,
			hr.capacity AS room_capacity,
			hr.price AS room_price,
			tg.id AS group_id,
			tg.name AS group_name
		FROM accommodations a
		JOIN hotel_rooms hr ON hr.id = a.hotel_room_id
		JOIN hotels h ON h.id = hr.hotel_id
		JOIN group_members gm ON gm.id = a.group_member_id
		JOIN tourist_groups tg ON tg.id = gm.tourist_group_id
		JOIN tourists t ON t.id = gm.tourist_id
		WHERE t.user_id = ?
		ORDER BY a.check_in_at DESC, a.id DESC
		LIMIT ? OFFSET ?
	`, userID, limit, offset).Scan(&items).Error

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *MeRepository) FindExcursionsByUserID(userID int64, limit int, offset int) ([]dto.MeExcursionResponse, int64, error) {
	var items []dto.MeExcursionResponse
	var total int64

	if err := r.db.Raw(`
		SELECT COUNT(*)
		FROM excursion_bookings eb
		JOIN group_members gm ON gm.id = eb.group_member_id
		JOIN tourists t ON t.id = gm.tourist_id
		WHERE t.user_id = ?
	`, userID).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Raw(`
		SELECT
			eb.id AS booking_id,
			eb.booked_at,
			eb.status AS booking_status,
			eb.tourist_rating,
			es.id AS schedule_id,
			es.price,
			es.start_time,
			es.end_time,
			es.status AS schedule_status,
			e.id AS excursion_id,
			e.name AS excursion_name,
			e.description AS excursion_description,
			ea.id AS agency_id,
			ea.name AS agency_name
		FROM excursion_bookings eb
		JOIN excursion_schedule es ON es.id = eb.excursion_schedule_id
		JOIN excursions e ON e.id = es.excursion_id
		JOIN excursion_agencies ea ON ea.id = es.excursion_agency_id
		JOIN group_members gm ON gm.id = eb.group_member_id
		JOIN tourists t ON t.id = gm.tourist_id
		WHERE t.user_id = ?
		ORDER BY es.start_time DESC, eb.id DESC
		LIMIT ? OFFSET ?
	`, userID, limit, offset).Scan(&items).Error

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *MeRepository) FindCargoByUserID(userID int64, limit int, offset int) ([]dto.MeCargoResponse, int64, error) {
	var items []dto.MeCargoResponse
	var total int64

	if err := r.db.Raw(`
		SELECT COUNT(*)
		FROM cargo_items ci
		JOIN cargo_statements cs ON cs.id = ci.cargo_statement_id
		JOIN group_members gm ON gm.id = cs.group_member_id
		JOIN tourists t ON t.id = gm.tourist_id
		WHERE t.user_id = ?
	`, userID).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Raw(`
		SELECT
			cs.id AS cargo_statement_id,
			cs.status AS statement_status,
			ci.id AS cargo_item_id,
			ci.item_number,
			ci.marking,
			ci.places_count,
			ci.weight_kg,
			ci.volumetric_weight_kg,
			ci.packaged_at,
			ct.id AS cargo_type_id,
			ct.name AS cargo_type_name,
			csh.id AS shipment_id,
			csh.status AS shipment_status,
			csh.shipped_at,
			f.id AS flight_id,
			f.flight_number,
			f.flight_date
		FROM cargo_items ci
		JOIN cargo_types ct ON ct.id = ci.cargo_type_id
		JOIN cargo_statements cs ON cs.id = ci.cargo_statement_id
		JOIN group_members gm ON gm.id = cs.group_member_id
		JOIN tourists t ON t.id = gm.tourist_id
		LEFT JOIN cargo_shipments csh ON csh.cargo_statement_id = cs.id
		LEFT JOIN flights f ON f.id = csh.flight_id
		WHERE t.user_id = ?
		ORDER BY cs.id DESC, ci.id ASC
		LIMIT ? OFFSET ?
	`, userID, limit, offset).Scan(&items).Error

	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

func (r *MeRepository) FindIdentityDocumentByUserID(userID int64) (*dto.MeIdentityDocumentResponse, error) {
	var document dto.MeIdentityDocumentResponse

	err := r.db.Raw(`
		SELECT
			idoc.id,
			idoc.tourist_id,
			idoc.document_type,
			idoc.document_series,
			idoc.document_number,
			TO_CHAR(idoc.issue_date, 'YYYY-MM-DD') AS issue_date,
			CASE
				WHEN idoc.expiration_date IS NULL THEN NULL
				ELSE TO_CHAR(idoc.expiration_date, 'YYYY-MM-DD')
			END AS expiration_date,
			idoc.issued_by,
			idoc.citizenship
		FROM identity_documents idoc
		JOIN tourists t ON t.id = idoc.tourist_id
		WHERE t.user_id = ?
		ORDER BY idoc.id DESC
		LIMIT 1
	`, userID).Scan(&document).Error

	if err != nil {
		return nil, err
	}

	if document.ID == 0 {
		return nil, nil
	}

	return &document, nil
}

func (r *MeRepository) HasIdentityDocument(userID int64) (bool, error) {
	var count int64

	err := r.db.Raw(`
		SELECT COUNT(*)
		FROM identity_documents idoc
		JOIN tourists t ON t.id = idoc.tourist_id
		WHERE t.user_id = ?
	`, userID).Scan(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *MeRepository) CreateIdentityDocument(
	touristID int64,
	documentType string,
	documentSeries string,
	documentNumber string,
	issueDate string,
	expirationDate *string,
	issuedBy string,
	citizenship string,
) (*dto.MeIdentityDocumentResponse, error) {
	var document dto.MeIdentityDocumentResponse

	err := r.db.Raw(`
		INSERT INTO identity_documents (
			tourist_id,
			document_type,
			document_series,
			document_number,
			issue_date,
			expiration_date,
			issued_by,
			citizenship
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING
			id,
			tourist_id,
			document_type,
			document_series,
			document_number,
			TO_CHAR(issue_date, 'YYYY-MM-DD') AS issue_date,
			CASE
				WHEN expiration_date IS NULL THEN NULL
				ELSE TO_CHAR(expiration_date, 'YYYY-MM-DD')
			END AS expiration_date,
			issued_by,
			citizenship
	`, touristID, documentType, documentSeries, documentNumber, issueDate, expirationDate, issuedBy, citizenship).Scan(&document).Error

	if err != nil {
		return nil, err
	}

	return &document, nil
}

func (r *MeRepository) UpdateIdentityDocument(
	documentID int64,
	documentType string,
	documentSeries string,
	documentNumber string,
	issueDate string,
	expirationDate *string,
	issuedBy string,
	citizenship string,
) (*dto.MeIdentityDocumentResponse, error) {
	var document dto.MeIdentityDocumentResponse

	err := r.db.Raw(`
		UPDATE identity_documents
		SET
			document_type = ?,
			document_series = ?,
			document_number = ?,
			issue_date = ?,
			expiration_date = ?,
			issued_by = ?,
			citizenship = ?
		WHERE id = ?
		RETURNING
			id,
			tourist_id,
			document_type,
			document_series,
			document_number,
			TO_CHAR(issue_date, 'YYYY-MM-DD') AS issue_date,
			CASE
				WHEN expiration_date IS NULL THEN NULL
				ELSE TO_CHAR(expiration_date, 'YYYY-MM-DD')
			END AS expiration_date,
			issued_by,
			citizenship
	`, documentType, documentSeries, documentNumber, issueDate, expirationDate, issuedBy, citizenship, documentID).Scan(&document).Error

	if err != nil {
		return nil, err
	}

	if document.ID == 0 {
		return nil, nil
	}

	return &document, nil
}

func (r *MeRepository) FindGroupMemberIDsByUserID(userID int64) ([]int64, error) {
	var ids []int64

	err := r.db.Raw(`
		SELECT gm.id
		FROM group_members gm
		JOIN tourists t ON t.id = gm.tourist_id
		WHERE t.user_id = ?
		ORDER BY gm.id ASC
	`, userID).Scan(&ids).Error

	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *MeRepository) GroupMemberBelongsToUser(userID int64, groupMemberID int64) (bool, error) {
	var count int64

	err := r.db.Raw(`
		SELECT COUNT(*)
		FROM group_members gm
		JOIN tourists t ON t.id = gm.tourist_id
		WHERE t.user_id = ?
		  AND gm.id = ?
	`, userID, groupMemberID).Scan(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *MeRepository) ExcursionScheduleExists(excursionScheduleID int64) (bool, error) {
	var count int64

	err := r.db.Raw(`
		SELECT COUNT(*)
		FROM excursion_schedule
		WHERE id = ?
		  AND status <> 'cancelled'
	`, excursionScheduleID).Scan(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *MeRepository) ExcursionBookingExists(excursionScheduleID int64, groupMemberID int64) (bool, error) {
	var count int64

	err := r.db.Raw(`
		SELECT COUNT(*)
		FROM excursion_bookings
		WHERE excursion_schedule_id = ?
		  AND group_member_id = ?
		  AND status <> 'cancelled'
	`, excursionScheduleID, groupMemberID).Scan(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *MeRepository) CreateExcursionBooking(
	excursionScheduleID int64,
	groupMemberID int64,
) (*dto.MeExcursionBookingResponse, error) {
	var booking dto.MeExcursionBookingResponse

	err := r.db.Raw(`
		INSERT INTO excursion_bookings (
			excursion_schedule_id,
			group_member_id,
			status,
			booked_at
		)
		VALUES (?, ?, 'booked', NOW())
		RETURNING
			id,
			excursion_schedule_id,
			group_member_id,
			status
	`, excursionScheduleID, groupMemberID).Scan(&booking).Error

	if err != nil {
		return nil, err
	}

	return &booking, nil
}
