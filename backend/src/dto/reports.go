package dto

import "time"

type CustomsTouristReportItem struct {
	GroupID        int64     `json:"group_id"`
	GroupName      string    `json:"group_name"`
	TouristID      int64     `json:"tourist_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	MiddleName     *string   `json:"middle_name,omitempty"`
	Sex            string    `json:"sex"`
	BirthDate      time.Time `json:"birth_date"`
	CategoryID     int64     `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	DocumentType   *string   `json:"document_type,omitempty"`
	DocumentSeries *string   `json:"document_series,omitempty"`
	DocumentNumber *string   `json:"document_number,omitempty"`
	Citizenship    *string   `json:"citizenship,omitempty"`
}

type AccommodationReportItem struct {
	GroupID      int64      `json:"group_id"`
	GroupName    string     `json:"group_name"`
	TouristID    int64      `json:"tourist_id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	MiddleName   *string    `json:"middle_name,omitempty"`
	CategoryID   int64      `json:"category_id"`
	CategoryName string     `json:"category_name"`
	HotelID      int64      `json:"hotel_id"`
	HotelName    string     `json:"hotel_name"`
	RoomID       int64      `json:"room_id"`
	RoomNumber   int        `json:"room_number"`
	CheckInAt    time.Time  `json:"check_in_at"`
	CheckOutAt   *time.Time `json:"check_out_at,omitempty"`
	Status       string     `json:"status"`
}

type TouristCountReportResponse struct {
	Count int64 `json:"count"`
}

type TouristTripReportItem struct {
	GroupID       int64     `json:"group_id"`
	GroupName     string    `json:"group_name"`
	ArrivalDate   time.Time `json:"arrival_date"`
	DepartureDate time.Time `json:"departure_date"`
	CategoryName  string    `json:"category_name"`
}

type TouristHotelReportItem struct {
	HotelID    int64      `json:"hotel_id"`
	HotelName  string     `json:"hotel_name"`
	RoomNumber int        `json:"room_number"`
	CheckInAt  time.Time  `json:"check_in_at"`
	CheckOutAt *time.Time `json:"check_out_at,omitempty"`
	Status     string     `json:"status"`
}

type TouristExcursionReportItem struct {
	ExcursionID   int64     `json:"excursion_id"`
	ExcursionName string    `json:"excursion_name"`
	AgencyID      int64     `json:"agency_id"`
	AgencyName    string    `json:"agency_name"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	Status        string    `json:"status"`
	TouristRating *int      `json:"tourist_rating,omitempty"`
}

type TouristCargoReportItem struct {
	CargoStatementID   int64   `json:"cargo_statement_id"`
	StatementStatus    string  `json:"statement_status"`
	CargoTypeName      string  `json:"cargo_type_name"`
	ItemNumber         string  `json:"item_number"`
	Marking            *string `json:"marking,omitempty"`
	PlacesCount        int     `json:"places_count"`
	WeightKg           float64 `json:"weight_kg"`
	VolumetricWeightKg float64 `json:"volumetric_weight_kg"`
}

type TouristInfoReportResponse struct {
	TouristID   int64                        `json:"tourist_id"`
	FirstName   string                       `json:"first_name"`
	LastName    string                       `json:"last_name"`
	MiddleName  *string                      `json:"middle_name,omitempty"`
	VisitsCount int64                        `json:"visits_count"`
	Trips       []TouristTripReportItem      `json:"trips"`
	Hotels      []TouristHotelReportItem     `json:"hotels"`
	Excursions  []TouristExcursionReportItem `json:"excursions"`
	Cargo       []TouristCargoReportItem     `json:"cargo"`
}

type HotelOccupancyReportItem struct {
	HotelID       int64  `json:"hotel_id"`
	HotelName     string `json:"hotel_name"`
	RoomsOccupied int64  `json:"rooms_occupied"`
	PeopleCount   int64  `json:"people_count"`
}

type ExcursionTouristCountReportResponse struct {
	Count int64 `json:"count"`
}

type PopularExcursionReportItem struct {
	ExcursionID   int64  `json:"excursion_id"`
	ExcursionName string `json:"excursion_name"`
	BookingsCount int64  `json:"bookings_count"`
	VisitedCount  int64  `json:"visited_count"`
}

type AgencyQualityReportItem struct {
	AgencyID      int64   `json:"agency_id"`
	AgencyName    string  `json:"agency_name"`
	AverageRating float64 `json:"average_rating"`
	RatingsCount  int64   `json:"ratings_count"`
}

type ExcursionAnalyticsReportResponse struct {
	PopularExcursions []PopularExcursionReportItem `json:"popular_excursions"`
	QualityAgencies   []AgencyQualityReportItem    `json:"quality_agencies"`
}

type FlightLoadReportResponse struct {
	FlightID           int64     `json:"flight_id"`
	FlightNumber       int       `json:"flight_number"`
	FlightDate         time.Time `json:"flight_date"`
	Capacity           int       `json:"capacity"`
	PlacesCount        int64     `json:"places_count"`
	WeightKg           float64   `json:"weight_kg"`
	VolumetricWeightKg float64   `json:"volumetric_weight_kg"`
}

type WarehouseTurnoverReportResponse struct {
	PlacesCount                int64   `json:"places_count"`
	WeightKg                   float64 `json:"weight_kg"`
	FlightsCount               int64   `json:"flights_count"`
	CargoFlightsCount          int64   `json:"cargo_flights_count"`
	CargoPassengerFlightsCount int64   `json:"cargo_passenger_flights_count"`
}

type FinancialReportItem struct {
	CategoryID    int64   `json:"category_id"`
	CategoryName  string  `json:"category_name"`
	OperationType string  `json:"operation_type"`
	Amount        float64 `json:"amount"`
}

type IncomeExpenseReportItem struct {
	CategoryID    int64   `json:"category_id"`
	CategoryName  string  `json:"category_name"`
	OperationType string  `json:"operation_type"`
	Amount        float64 `json:"amount"`
}

type CargoTypeShareReportItem struct {
	CargoTypeID   int64   `json:"cargo_type_id"`
	CargoTypeName string  `json:"cargo_type_name"`
	PlacesCount   int64   `json:"places_count"`
	WeightKg      float64 `json:"weight_kg"`
	SharePercent  float64 `json:"share_percent"`
}

type ProfitabilityReportResponse struct {
	Income               float64 `json:"income"`
	Expense              float64 `json:"expense"`
	Profit               float64 `json:"profit"`
	ProfitabilityPercent float64 `json:"profitability_percent"`
}

type TouristCategoryRatioReportResponse struct {
	RestTouristsCount int64   `json:"rest_tourists_count"`
	ShopTouristsCount int64   `json:"shop_tourists_count"`
	Total             int64   `json:"total"`
	RestPercent       float64 `json:"rest_percent"`
	ShopPercent       float64 `json:"shop_percent"`
}

type FlightTouristReportItem struct {
	FlightID         int64     `json:"flight_id"`
	FlightNumber     int       `json:"flight_number"`
	FlightDate       time.Time `json:"flight_date"`
	GroupID          int64     `json:"group_id"`
	GroupName        string    `json:"group_name"`
	TouristID        int64     `json:"tourist_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	MiddleName       *string   `json:"middle_name,omitempty"`
	HotelName        *string   `json:"hotel_name,omitempty"`
	RoomNumber       *int      `json:"room_number,omitempty"`
	CargoStatementID int64     `json:"cargo_statement_id"`
	ItemNumber       string    `json:"item_number"`
	Marking          *string   `json:"marking,omitempty"`
	PlacesCount      int       `json:"places_count"`
	WeightKg         float64   `json:"weight_kg"`
}
