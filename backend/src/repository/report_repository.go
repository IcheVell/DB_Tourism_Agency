package repository

import (
	"time"

	"TouristAgencyApp/src/dto"

	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{
		db: db,
	}
}

func (r *ReportRepository) FindCustomsTourists(categoryID int64) ([]dto.CustomsTouristReportItem, error) {
	var items []dto.CustomsTouristReportItem

	err := r.db.Raw(`
		SELECT
			tg.id AS group_id,
			tg.name AS group_name,
			t.id AS tourist_id,
			t.first_name,
			t.last_name,
			t.middle_name,
			t.sex,
			t.birth_date,
			tc.id AS category_id,
			tc.name AS category_name,
			idoc.document_type,
			idoc.document_series,
			idoc.document_number,
			idoc.citizenship
		FROM group_members gm
		JOIN tourist_groups tg ON tg.id = gm.tourist_group_id
		JOIN tourists t ON t.id = gm.tourist_id
		JOIN tourist_categories tc ON tc.id = gm.tourist_category_id
		LEFT JOIN identity_documents idoc ON idoc.tourist_id = t.id
		WHERE (? = 0 OR tc.id = ?)
		ORDER BY tg.id ASC, t.last_name ASC, t.first_name ASC
	`, categoryID, categoryID).Scan(&items).Error

	return items, err
}

func (r *ReportRepository) FindAccommodationList(hotelID int64, categoryID int64) ([]dto.AccommodationReportItem, error) {
	var items []dto.AccommodationReportItem

	err := r.db.Raw(`
		SELECT
			tg.id AS group_id,
			tg.name AS group_name,
			t.id AS tourist_id,
			t.first_name,
			t.last_name,
			t.middle_name,
			tc.id AS category_id,
			tc.name AS category_name,
			h.id AS hotel_id,
			h.name AS hotel_name,
			hr.id AS room_id,
			hr.room_number,
			a.check_in_at,
			a.check_out_at,
			a.status
		FROM accommodations a
		JOIN hotel_rooms hr ON hr.id = a.hotel_room_id
		JOIN hotels h ON h.id = hr.hotel_id
		JOIN group_members gm ON gm.id = a.group_member_id
		JOIN tourist_groups tg ON tg.id = gm.tourist_group_id
		JOIN tourists t ON t.id = gm.tourist_id
		JOIN tourist_categories tc ON tc.id = gm.tourist_category_id
		WHERE (? = 0 OR h.id = ?)
		  AND (? = 0 OR tc.id = ?)
		ORDER BY h.name ASC, hr.room_number ASC, t.last_name ASC
	`, hotelID, hotelID, categoryID, categoryID).Scan(&items).Error

	return items, err
}

func (r *ReportRepository) CountTourists(from *time.Time, to *time.Time, categoryID int64) (int64, error) {
	var count int64

	err := r.db.Raw(`
		SELECT COUNT(DISTINCT gm.tourist_id)
		FROM group_members gm
		JOIN tourist_groups tg ON tg.id = gm.tourist_group_id
		WHERE (CAST(? AS timestamptz) IS NULL OR tg.arrival_date >= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR tg.departure_date <= CAST(? AS timestamptz))
		  AND (? = 0 OR gm.tourist_category_id = ?)
	`, from, from, to, to, categoryID, categoryID).Scan(&count).Error

	return count, err
}

func (r *ReportRepository) FindTouristInfo(touristID int64) (*dto.TouristInfoReportResponse, error) {
	var result dto.TouristInfoReportResponse

	var baseInfo struct {
		TouristID  int64
		FirstName  string
		LastName   string
		MiddleName *string
	}

	if err := r.db.Raw(`
		SELECT
			id AS tourist_id,
			first_name,
			last_name,
			middle_name
		FROM tourists
		WHERE id = ?
	`, touristID).Scan(&baseInfo).Error; err != nil {
		return nil, err
	}

	if baseInfo.TouristID == 0 {
		return nil, nil
	}

	result.TouristID = baseInfo.TouristID
	result.FirstName = baseInfo.FirstName
	result.LastName = baseInfo.LastName
	result.MiddleName = baseInfo.MiddleName

	if err := r.db.Raw(`
		SELECT COUNT(DISTINCT tourist_group_id)
		FROM group_members
		WHERE tourist_id = ?
	`, touristID).Scan(&result.VisitsCount).Error; err != nil {
		return nil, err
	}

	if err := r.db.Raw(`
		SELECT
			tg.id AS group_id,
			tg.name AS group_name,
			tg.arrival_date,
			tg.departure_date,
			tc.name AS category_name
		FROM group_members gm
		JOIN tourist_groups tg ON tg.id = gm.tourist_group_id
		JOIN tourist_categories tc ON tc.id = gm.tourist_category_id
		WHERE gm.tourist_id = ?
		ORDER BY tg.arrival_date DESC
	`, touristID).Scan(&result.Trips).Error; err != nil {
		return nil, err
	}

	if err := r.db.Raw(`
		SELECT
			h.id AS hotel_id,
			h.name AS hotel_name,
			hr.room_number,
			a.check_in_at,
			a.check_out_at,
			a.status
		FROM accommodations a
		JOIN hotel_rooms hr ON hr.id = a.hotel_room_id
		JOIN hotels h ON h.id = hr.hotel_id
		JOIN group_members gm ON gm.id = a.group_member_id
		WHERE gm.tourist_id = ?
		ORDER BY a.check_in_at DESC
	`, touristID).Scan(&result.Hotels).Error; err != nil {
		return nil, err
	}

	if err := r.db.Raw(`
		SELECT
			e.id AS excursion_id,
			e.name AS excursion_name,
			ea.id AS agency_id,
			ea.name AS agency_name,
			es.start_time,
			es.end_time,
			eb.status,
			eb.tourist_rating
		FROM excursion_bookings eb
		JOIN excursion_schedule es ON es.id = eb.excursion_schedule_id
		JOIN excursions e ON e.id = es.excursion_id
		JOIN excursion_agencies ea ON ea.id = es.excursion_agency_id
		JOIN group_members gm ON gm.id = eb.group_member_id
		WHERE gm.tourist_id = ?
		ORDER BY es.start_time DESC
	`, touristID).Scan(&result.Excursions).Error; err != nil {
		return nil, err
	}

	if err := r.db.Raw(`
		SELECT
			cs.id AS cargo_statement_id,
			cs.status AS statement_status,
			ct.name AS cargo_type_name,
			ci.item_number,
			ci.marking,
			ci.places_count,
			ci.weight_kg,
			ci.volumetric_weight_kg
		FROM cargo_statements cs
		JOIN cargo_items ci ON ci.cargo_statement_id = cs.id
		JOIN cargo_types ct ON ct.id = ci.cargo_type_id
		JOIN group_members gm ON gm.id = cs.group_member_id
		WHERE gm.tourist_id = ?
		ORDER BY cs.id ASC, ci.item_number ASC
	`, touristID).Scan(&result.Cargo).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ReportRepository) FindHotelOccupancy(from *time.Time, to *time.Time) ([]dto.HotelOccupancyReportItem, error) {
	var items []dto.HotelOccupancyReportItem

	err := r.db.Raw(`
		SELECT
			h.id AS hotel_id,
			h.name AS hotel_name,
			COUNT(DISTINCT hr.id) AS rooms_occupied,
			COUNT(DISTINCT a.group_member_id) AS people_count
		FROM hotels h
		JOIN hotel_rooms hr ON hr.hotel_id = h.id
		JOIN accommodations a ON a.hotel_room_id = hr.id
		WHERE (CAST(? AS timestamptz) IS NULL OR a.check_in_at <= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR a.check_out_at IS NULL OR a.check_out_at >= CAST(? AS timestamptz))
		  AND a.status <> 'cancelled'
		GROUP BY h.id, h.name
		ORDER BY h.name ASC
	`, to, to, from, from).Scan(&items).Error

	return items, err
}

func (r *ReportRepository) CountExcursionTourists(from *time.Time, to *time.Time) (int64, error) {
	var count int64

	err := r.db.Raw(`
		SELECT COUNT(DISTINCT eb.group_member_id)
		FROM excursion_bookings eb
		JOIN excursion_schedule es ON es.id = eb.excursion_schedule_id
		WHERE eb.group_member_id IS NOT NULL
		  AND eb.status <> 'cancelled'
		  AND (CAST(? AS timestamptz) IS NULL OR es.start_time >= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR es.start_time <= CAST(? AS timestamptz))
	`, from, from, to, to).Scan(&count).Error

	return count, err
}

func (r *ReportRepository) FindExcursionAnalytics(from *time.Time, to *time.Time) (*dto.ExcursionAnalyticsReportResponse, error) {
	var result dto.ExcursionAnalyticsReportResponse

	if err := r.db.Raw(`
		SELECT
			e.id AS excursion_id,
			e.name AS excursion_name,
			COUNT(eb.id) AS bookings_count,
			COUNT(CASE WHEN eb.status = 'visited' THEN 1 END) AS visited_count
		FROM excursions e
		JOIN excursion_schedule es ON es.excursion_id = e.id
		LEFT JOIN excursion_bookings eb ON eb.excursion_schedule_id = es.id
		WHERE (CAST(? AS timestamptz) IS NULL OR es.start_time >= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR es.start_time <= CAST(? AS timestamptz))
		GROUP BY e.id, e.name
		ORDER BY bookings_count DESC, visited_count DESC, e.name ASC
	`, from, from, to, to).Scan(&result.PopularExcursions).Error; err != nil {
		return nil, err
	}

	if err := r.db.Raw(`
		SELECT
			ea.id AS agency_id,
			ea.name AS agency_name,
			COALESCE(AVG(eb.tourist_rating), 0) AS average_rating,
			COUNT(eb.tourist_rating) AS ratings_count
		FROM excursion_agencies ea
		JOIN excursion_schedule es ON es.excursion_agency_id = ea.id
		LEFT JOIN excursion_bookings eb ON eb.excursion_schedule_id = es.id
		WHERE (CAST(? AS timestamptz) IS NULL OR es.start_time >= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR es.start_time <= CAST(? AS timestamptz))
		GROUP BY ea.id, ea.name
		ORDER BY average_rating DESC, ratings_count DESC, ea.name ASC
	`, from, from, to, to).Scan(&result.QualityAgencies).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *ReportRepository) FindFlightLoad(flightID int64) (*dto.FlightLoadReportResponse, error) {
	var result dto.FlightLoadReportResponse

	err := r.db.Raw(`
		SELECT
			f.id AS flight_id,
			f.flight_number,
			f.flight_date,
			f.capacity,
			COALESCE(SUM(ci.places_count), 0) AS places_count,
			COALESCE(SUM(ci.weight_kg), 0) AS weight_kg,
			COALESCE(SUM(ci.volumetric_weight_kg), 0) AS volumetric_weight_kg
		FROM flights f
		LEFT JOIN cargo_shipments csh ON csh.flight_id = f.id AND csh.status <> 'cancelled'
		LEFT JOIN cargo_items ci ON ci.cargo_statement_id = csh.cargo_statement_id
		WHERE f.id = ?
		GROUP BY f.id, f.flight_number, f.flight_date, f.capacity
	`, flightID).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	if result.FlightID == 0 {
		return nil, nil
	}

	return &result, nil
}

func (r *ReportRepository) FindWarehouseTurnover(from *time.Time, to *time.Time) (*dto.WarehouseTurnoverReportResponse, error) {
	var result dto.WarehouseTurnoverReportResponse

	err := r.db.Raw(`
		SELECT
			COALESCE(SUM(ci.places_count), 0) AS places_count,
			COALESCE(SUM(ci.weight_kg), 0) AS weight_kg,
			COUNT(DISTINCT f.id) AS flights_count,
			COUNT(DISTINCT CASE
				WHEN lower(ft.name) LIKE '%cargo%' OR lower(ft.name) LIKE '%груз%'
				THEN f.id
			END) AS cargo_flights_count,
			COUNT(DISTINCT CASE
				WHEN lower(ft.name) LIKE '%passenger%' OR lower(ft.name) LIKE '%пасс%'
				THEN f.id
			END) AS cargo_passenger_flights_count
		FROM cargo_shipments csh
		JOIN flights f ON f.id = csh.flight_id
		LEFT JOIN flight_types ft ON ft.id = f.flight_type_id
		JOIN cargo_items ci ON ci.cargo_statement_id = csh.cargo_statement_id
		WHERE csh.status <> 'cancelled'
		  AND (CAST(? AS timestamptz) IS NULL OR csh.shipped_at >= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR csh.shipped_at <= CAST(? AS timestamptz))
	`, from, from, to, to).Scan(&result).Error

	return &result, err
}

func (r *ReportRepository) FindGroupFinancialReport(groupID int64, categoryID int64) ([]dto.FinancialReportItem, error) {
	var items []dto.FinancialReportItem

	err := r.db.Raw(`
		SELECT
			fc.id AS category_id,
			fc.name AS category_name,
			fc.operation_type,
			COALESCE(SUM(fo.amount), 0) AS amount
		FROM financial_operations fo
		JOIN financial_categories fc ON fc.id = fo.financial_category_id
		LEFT JOIN accommodations a ON a.id = fo.accommodation_id
		LEFT JOIN excursion_bookings eb ON eb.id = fo.excursion_booking_id
		LEFT JOIN visas v ON v.id = fo.visa_id
		LEFT JOIN cargo_statements cs1 ON cs1.id = fo.cargo_statement_id
		LEFT JOIN cargo_shipments csh ON csh.id = fo.cargo_shipment_id
		LEFT JOIN cargo_statements cs2 ON cs2.id = csh.cargo_statement_id
		LEFT JOIN group_members gm ON
			gm.id = a.group_member_id
			OR gm.id = eb.group_member_id
			OR gm.id = cs1.group_member_id
			OR gm.id = cs2.group_member_id
			OR gm.tourist_id = v.tourist_id
		WHERE (? = 0 OR gm.tourist_group_id = ?)
		  AND (? = 0 OR gm.tourist_category_id = ?)
		GROUP BY fc.id, fc.name, fc.operation_type
		ORDER BY fc.operation_type ASC, fc.name ASC
	`, groupID, groupID, categoryID, categoryID).Scan(&items).Error

	return items, err
}

func (r *ReportRepository) FindIncomeExpense(from *time.Time, to *time.Time) ([]dto.IncomeExpenseReportItem, error) {
	var items []dto.IncomeExpenseReportItem

	err := r.db.Raw(`
		SELECT
			fc.id AS category_id,
			fc.name AS category_name,
			fc.operation_type,
			COALESCE(SUM(fo.amount), 0) AS amount
		FROM financial_operations fo
		JOIN financial_categories fc ON fc.id = fo.financial_category_id
		WHERE (CAST(? AS timestamptz) IS NULL OR fo.operation_at >= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR fo.operation_at <= CAST(? AS timestamptz))
		GROUP BY fc.id, fc.name, fc.operation_type
		ORDER BY fc.operation_type ASC, fc.name ASC
	`, from, from, to, to).Scan(&items).Error

	return items, err
}

func (r *ReportRepository) FindCargoTypeShare(from *time.Time, to *time.Time) ([]dto.CargoTypeShareReportItem, error) {
	var items []dto.CargoTypeShareReportItem

	err := r.db.Raw(`
		WITH cargo_by_type AS (
			SELECT
				ct.id AS cargo_type_id,
				ct.name AS cargo_type_name,
				COALESCE(SUM(ci.places_count), 0) AS places_count,
				COALESCE(SUM(ci.weight_kg), 0) AS weight_kg
			FROM cargo_types ct
			JOIN cargo_items ci ON ci.cargo_type_id = ct.id
			JOIN cargo_statements cs ON cs.id = ci.cargo_statement_id
			LEFT JOIN cargo_shipments csh ON csh.cargo_statement_id = cs.id
			WHERE (CAST(? AS timestamptz) IS NULL OR csh.shipped_at >= CAST(? AS timestamptz) OR ci.packaged_at >= CAST(? AS timestamptz))
			  AND (CAST(? AS timestamptz) IS NULL OR csh.shipped_at <= CAST(? AS timestamptz) OR ci.packaged_at <= CAST(? AS timestamptz))
			GROUP BY ct.id, ct.name
		),
		total AS (
			SELECT COALESCE(SUM(weight_kg), 0) AS total_weight
			FROM cargo_by_type
		)
		SELECT
			cargo_by_type.cargo_type_id,
			cargo_by_type.cargo_type_name,
			cargo_by_type.places_count,
			cargo_by_type.weight_kg,
			CASE
				WHEN total.total_weight = 0 THEN 0
				ELSE ROUND((cargo_by_type.weight_kg / total.total_weight) * 100, 2)
			END AS share_percent
		FROM cargo_by_type, total
		ORDER BY share_percent DESC, cargo_type_name ASC
	`, from, from, from, to, to, to).Scan(&items).Error

	return items, err
}

func (r *ReportRepository) FindProfitability(from *time.Time, to *time.Time) (*dto.ProfitabilityReportResponse, error) {
	var result dto.ProfitabilityReportResponse

	err := r.db.Raw(`
		SELECT
			COALESCE(SUM(CASE WHEN fc.operation_type = 'income' THEN fo.amount ELSE 0 END), 0) AS income,
			COALESCE(SUM(CASE WHEN fc.operation_type = 'expense' THEN fo.amount ELSE 0 END), 0) AS expense
		FROM financial_operations fo
		JOIN financial_categories fc ON fc.id = fo.financial_category_id
		WHERE (CAST(? AS timestamptz) IS NULL OR fo.operation_at >= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR fo.operation_at <= CAST(? AS timestamptz))
	`, from, from, to, to).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	result.Profit = result.Income - result.Expense

	if result.Expense > 0 {
		result.ProfitabilityPercent = (result.Profit / result.Expense) * 100
	}

	return &result, nil
}

func (r *ReportRepository) FindTouristCategoryRatio(
	from *time.Time,
	to *time.Time,
	restCategoryID int64,
	shopCategoryID int64,
) (*dto.TouristCategoryRatioReportResponse, error) {
	var result dto.TouristCategoryRatioReportResponse

	err := r.db.Raw(`
		SELECT
			COUNT(DISTINCT CASE WHEN gm.tourist_category_id = ? THEN gm.tourist_id END) AS rest_tourists_count,
			COUNT(DISTINCT CASE WHEN gm.tourist_category_id = ? THEN gm.tourist_id END) AS shop_tourists_count,
			COUNT(DISTINCT CASE WHEN gm.tourist_category_id IN (?, ?) THEN gm.tourist_id END) AS total
		FROM group_members gm
		JOIN tourist_groups tg ON tg.id = gm.tourist_group_id
		WHERE (CAST(? AS timestamptz) IS NULL OR tg.arrival_date >= CAST(? AS timestamptz))
		  AND (CAST(? AS timestamptz) IS NULL OR tg.arrival_date <= CAST(? AS timestamptz))
	`, restCategoryID, shopCategoryID, restCategoryID, shopCategoryID, from, from, to, to).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	if result.Total > 0 {
		result.RestPercent = float64(result.RestTouristsCount) / float64(result.Total) * 100
		result.ShopPercent = float64(result.ShopTouristsCount) / float64(result.Total) * 100
	}

	return &result, nil
}

func (r *ReportRepository) FindFlightTourists(flightID int64) ([]dto.FlightTouristReportItem, error) {
	var items []dto.FlightTouristReportItem

	err := r.db.Raw(`
		SELECT
			f.id AS flight_id,
			f.flight_number,
			f.flight_date,
			tg.id AS group_id,
			tg.name AS group_name,
			t.id AS tourist_id,
			t.first_name,
			t.last_name,
			t.middle_name,
			h.name AS hotel_name,
			hr.room_number,
			cs.id AS cargo_statement_id,
			ci.item_number,
			ci.marking,
			ci.places_count,
			ci.weight_kg
		FROM flights f
		JOIN cargo_shipments csh ON csh.flight_id = f.id
		JOIN cargo_statements cs ON cs.id = csh.cargo_statement_id
		JOIN cargo_items ci ON ci.cargo_statement_id = cs.id
		JOIN group_members gm ON gm.id = cs.group_member_id
		JOIN tourist_groups tg ON tg.id = gm.tourist_group_id
		JOIN tourists t ON t.id = gm.tourist_id
		LEFT JOIN accommodations a ON a.group_member_id = gm.id AND a.status <> 'cancelled'
		LEFT JOIN hotel_rooms hr ON hr.id = a.hotel_room_id
		LEFT JOIN hotels h ON h.id = hr.hotel_id
		WHERE f.id = ?
		ORDER BY tg.name ASC, t.last_name ASC, t.first_name ASC, ci.item_number ASC
	`, flightID).Scan(&items).Error

	return items, err
}
