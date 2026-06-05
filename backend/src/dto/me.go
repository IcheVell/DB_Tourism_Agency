package dto

import "time"

type MeTourResponse struct {
	GroupMemberID       int64     `json:"group_member_id"`
	GroupID             int64     `json:"group_id"`
	GroupName           string    `json:"group_name"`
	ArrivalDate         time.Time `json:"arrival_date"`
	DepartureDate       time.Time `json:"departure_date"`
	CategoryID          int64     `json:"category_id"`
	CategoryName        string    `json:"category_name"`
	DesiredHotelID      *int64    `json:"desired_hotel_id,omitempty"`
	DesiredHotelName    *string   `json:"desired_hotel_name,omitempty"`
	DesiredHotelAddress *string   `json:"desired_hotel_address,omitempty"`
}

type MeTourListResponse struct {
	Items []MeTourResponse `json:"items"`
	Total int64            `json:"total"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
}

type MeVisaResponse struct {
	ID                 int64      `json:"id"`
	Number             *string    `json:"number,omitempty"`
	DestinationCountry string     `json:"destination_country"`
	Status             string     `json:"status"`
	CreatedAt          time.Time  `json:"created_at"`
	SubmittedAt        *time.Time `json:"submitted_at,omitempty"`
	DecisionAt         *time.Time `json:"decision_at,omitempty"`
	IssuedAt           *time.Time `json:"issued_at,omitempty"`
	ValidFrom          *time.Time `json:"valid_from,omitempty"`
	ValidUntil         *time.Time `json:"valid_until,omitempty"`
}

type MeVisaListResponse struct {
	Items []MeVisaResponse `json:"items"`
	Total int64            `json:"total"`
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
}

type MeAccommodationResponse struct {
	ID           int64      `json:"id"`
	Status       string     `json:"status"`
	CheckInAt    time.Time  `json:"check_in_at"`
	CheckOutAt   *time.Time `json:"check_out_at,omitempty"`
	HotelID      int64      `json:"hotel_id"`
	HotelName    string     `json:"hotel_name"`
	HotelAddress string     `json:"hotel_address"`
	RoomID       int64      `json:"room_id"`
	RoomNumber   int        `json:"room_number"`
	RoomCapacity int        `json:"room_capacity"`
	RoomPrice    float64    `json:"room_price"`
	GroupID      int64      `json:"group_id"`
	GroupName    string     `json:"group_name"`
}

type MeAccommodationListResponse struct {
	Items []MeAccommodationResponse `json:"items"`
	Total int64                     `json:"total"`
	Page  int                       `json:"page"`
	Limit int                       `json:"limit"`
}

type MeExcursionResponse struct {
	BookingID            int64     `json:"booking_id"`
	BookedAt             time.Time `json:"booked_at"`
	BookingStatus        string    `json:"booking_status"`
	TouristRating        *int      `json:"tourist_rating,omitempty"`
	ScheduleID           int64     `json:"schedule_id"`
	Price                float64   `json:"price"`
	StartTime            time.Time `json:"start_time"`
	EndTime              time.Time `json:"end_time"`
	ScheduleStatus       string    `json:"schedule_status"`
	ExcursionID          int64     `json:"excursion_id"`
	ExcursionName        string    `json:"excursion_name"`
	ExcursionDescription *string   `json:"excursion_description,omitempty"`
	AgencyID             int64     `json:"agency_id"`
	AgencyName           string    `json:"agency_name"`
}

type MeExcursionListResponse struct {
	Items []MeExcursionResponse `json:"items"`
	Total int64                 `json:"total"`
	Page  int                   `json:"page"`
	Limit int                   `json:"limit"`
}

type MeCargoResponse struct {
	CargoStatementID   int64      `json:"cargo_statement_id"`
	StatementStatus    string     `json:"statement_status"`
	CargoItemID        int64      `json:"cargo_item_id"`
	ItemNumber         string     `json:"item_number"`
	Marking            *string    `json:"marking,omitempty"`
	PlacesCount        int        `json:"places_count"`
	WeightKg           float64    `json:"weight_kg"`
	VolumetricWeightKg float64    `json:"volumetric_weight_kg"`
	PackagedAt         *time.Time `json:"packaged_at,omitempty"`
	CargoTypeID        int64      `json:"cargo_type_id"`
	CargoTypeName      string     `json:"cargo_type_name"`
	ShipmentID         *int64     `json:"shipment_id,omitempty"`
	ShipmentStatus     *string    `json:"shipment_status,omitempty"`
	ShippedAt          *time.Time `json:"shipped_at,omitempty"`
	FlightID           *int64     `json:"flight_id,omitempty"`
	FlightNumber       *int       `json:"flight_number,omitempty"`
	FlightDate         *time.Time `json:"flight_date,omitempty"`
}

type MeCargoListResponse struct {
	Items []MeCargoResponse `json:"items"`
	Total int64             `json:"total"`
	Page  int               `json:"page"`
	Limit int               `json:"limit"`
}

type MeIdentityDocumentResponse struct {
	ID             int64   `json:"id"`
	TouristID      int64   `json:"tourist_id"`
	DocumentType   string  `json:"document_type"`
	DocumentSeries string  `json:"document_series"`
	DocumentNumber string  `json:"document_number"`
	IssueDate      string  `json:"issue_date"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	IssuedBy       string  `json:"issued_by"`
	Citizenship    string  `json:"citizenship"`
}

type CreateMeIdentityDocumentRequest struct {
	DocumentType   string  `json:"document_type"`
	DocumentSeries string  `json:"document_series"`
	DocumentNumber string  `json:"document_number"`
	IssueDate      string  `json:"issue_date"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	IssuedBy       string  `json:"issued_by"`
	Citizenship    string  `json:"citizenship"`
}

type UpdateMeIdentityDocumentRequest struct {
	DocumentType   *string `json:"document_type,omitempty"`
	DocumentSeries *string `json:"document_series,omitempty"`
	DocumentNumber *string `json:"document_number,omitempty"`
	IssueDate      *string `json:"issue_date,omitempty"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	IssuedBy       *string `json:"issued_by,omitempty"`
	Citizenship    *string `json:"citizenship,omitempty"`
}

type CreateMeExcursionBookingRequest struct {
	ExcursionScheduleID int64  `json:"excursion_schedule_id"`
	GroupMemberID       *int64 `json:"group_member_id,omitempty"`
}

type MeExcursionBookingResponse struct {
	ID                  int64  `json:"id"`
	ExcursionScheduleID int64  `json:"excursion_schedule_id"`
	GroupMemberID       int64  `json:"group_member_id"`
	Status              string `json:"status"`
}
