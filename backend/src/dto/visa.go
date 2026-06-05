package dto

type CreateVisaRequest struct {
	Number             *string `json:"number,omitempty"`
	DestinationCountry string  `json:"destination_country"`
	Status             string  `json:"status"`

	SubmittedAt *string `json:"submitted_at,omitempty"`
	DecisionAt  *string `json:"decision_at,omitempty"`
	IssuedAt    *string `json:"issued_at,omitempty"`

	ValidFrom  *string `json:"valid_from,omitempty"`
	ValidUntil *string `json:"valid_until,omitempty"`

	TouristID int64 `json:"tourist_id"`
}

type UpdateVisaRequest struct {
	Number             *string `json:"number,omitempty"`
	DestinationCountry string  `json:"destination_country"`
	Status             string  `json:"status"`

	SubmittedAt *string `json:"submitted_at,omitempty"`
	DecisionAt  *string `json:"decision_at,omitempty"`
	IssuedAt    *string `json:"issued_at,omitempty"`

	ValidFrom  *string `json:"valid_from,omitempty"`
	ValidUntil *string `json:"valid_until,omitempty"`

	TouristID int64 `json:"tourist_id"`
}

type VisaResponse struct {
	ID                 int64   `json:"id"`
	Number             *string `json:"number,omitempty"`
	DestinationCountry string  `json:"destination_country"`
	Status             string  `json:"status"`

	CreatedAt   string  `json:"created_at"`
	SubmittedAt *string `json:"submitted_at,omitempty"`
	DecisionAt  *string `json:"decision_at,omitempty"`
	IssuedAt    *string `json:"issued_at,omitempty"`

	ValidFrom  *string `json:"valid_from,omitempty"`
	ValidUntil *string `json:"valid_until,omitempty"`

	TouristID int64 `json:"tourist_id"`
}

type VisaListResponse struct {
	Items    []VisaResponse `json:"items"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
	Total    int64          `json:"total"`
}
