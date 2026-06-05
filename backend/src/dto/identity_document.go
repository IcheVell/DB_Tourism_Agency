package dto

type CreateIdentityDocumentRequest struct {
	DocumentType   string  `json:"document_type"`
	DocumentSeries string  `json:"document_series"`
	DocumentNumber string  `json:"document_number"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	IssuedBy       string  `json:"issued_by"`
	IssueDate      string  `json:"issue_date"`
	Citizenship    string  `json:"citizenship"`
	TouristID      int64   `json:"tourist_id"`
}

type UpdateIdentityDocumentRequest struct {
	DocumentType   string  `json:"document_type"`
	DocumentSeries string  `json:"document_series"`
	DocumentNumber string  `json:"document_number"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	IssuedBy       string  `json:"issued_by"`
	IssueDate      string  `json:"issue_date"`
	Citizenship    string  `json:"citizenship"`
	TouristID      int64   `json:"tourist_id"`
}

type IdentityDocumentResponse struct {
	ID             int64   `json:"id"`
	DocumentType   string  `json:"document_type"`
	DocumentSeries string  `json:"document_series"`
	DocumentNumber string  `json:"document_number"`
	ExpirationDate *string `json:"expiration_date,omitempty"`
	IssuedBy       string  `json:"issued_by"`
	IssueDate      string  `json:"issue_date"`
	Citizenship    string  `json:"citizenship"`
	TouristID      int64   `json:"tourist_id"`
}

type IdentityDocumentListResponse struct {
	Items    []IdentityDocumentResponse `json:"items"`
	Page     int                        `json:"page"`
	PageSize int                        `json:"page_size"`
	Total    int64                      `json:"total"`
}
