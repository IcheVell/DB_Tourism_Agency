package service

import (
	"errors"
	"strings"
	"time"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/models"
	"TouristAgencyApp/src/repository"
)

var (
	ErrIdentityDocumentNotFound              = errors.New("identity document not found")
	ErrIdentityDocumentInvalidType           = errors.New("identity document type is invalid")
	ErrIdentityDocumentInvalidSeries         = errors.New("identity document series is required")
	ErrIdentityDocumentInvalidNumber         = errors.New("identity document number is required")
	ErrIdentityDocumentInvalidIssuedBy       = errors.New("identity document issued_by is required")
	ErrIdentityDocumentInvalidIssueDate      = errors.New("identity document issue_date must be in YYYY-MM-DD format")
	ErrIdentityDocumentInvalidExpirationDate = errors.New("identity document expiration_date must be in YYYY-MM-DD format and greater than issue_date")
	ErrIdentityDocumentInvalidCitizenship    = errors.New("identity document citizenship is required")
	ErrIdentityDocumentInvalidTouristID      = errors.New("tourist_id must be positive")
)

type IdentityDocumentService struct {
	repo *repository.IdentityDocumentRepository
}

func NewIdentityDocumentService(repo *repository.IdentityDocumentRepository) *IdentityDocumentService {
	return &IdentityDocumentService{
		repo: repo,
	}
}

func (s *IdentityDocumentService) Create(req dto.CreateIdentityDocumentRequest) (*dto.IdentityDocumentResponse, error) {
	documentType := strings.TrimSpace(req.DocumentType)
	documentSeries := strings.TrimSpace(req.DocumentSeries)
	documentNumber := strings.TrimSpace(req.DocumentNumber)
	issuedBy := strings.TrimSpace(req.IssuedBy)
	citizenship := strings.TrimSpace(req.Citizenship)

	if !isValidDocumentType(documentType) {
		return nil, ErrIdentityDocumentInvalidType
	}

	if documentSeries == "" {
		return nil, ErrIdentityDocumentInvalidSeries
	}

	if documentNumber == "" {
		return nil, ErrIdentityDocumentInvalidNumber
	}

	if issuedBy == "" {
		return nil, ErrIdentityDocumentInvalidIssuedBy
	}

	if citizenship == "" {
		return nil, ErrIdentityDocumentInvalidCitizenship
	}

	if req.TouristID <= 0 {
		return nil, ErrIdentityDocumentInvalidTouristID
	}

	issueDate, err := parseDateYYYYMMDD(req.IssueDate)
	if err != nil {
		return nil, ErrIdentityDocumentInvalidIssueDate
	}

	expirationDate, err := parseOptionalDateYYYYMMDD(req.ExpirationDate)
	if err != nil {
		return nil, ErrIdentityDocumentInvalidExpirationDate
	}

	if expirationDate != nil && !expirationDate.After(issueDate) {
		return nil, ErrIdentityDocumentInvalidExpirationDate
	}

	document := &models.IdentityDocument{
		DocumentType:   documentType,
		DocumentSeries: documentSeries,
		DocumentNumber: documentNumber,
		ExpirationDate: expirationDate,
		IssuedBy:       issuedBy,
		IssueDate:      issueDate,
		Citizenship:    citizenship,
		TouristID:      req.TouristID,
	}

	if err := s.repo.Create(document); err != nil {
		return nil, err
	}

	return toIdentityDocumentResponse(document), nil
}

func (s *IdentityDocumentService) GetByID(id int64) (*dto.IdentityDocumentResponse, error) {
	document, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if document == nil {
		return nil, ErrIdentityDocumentNotFound
	}

	return toIdentityDocumentResponse(document), nil
}

func (s *IdentityDocumentService) List(page int, pageSize int) (*dto.IdentityDocumentListResponse, error) {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 20
	}

	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize

	documents, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.IdentityDocumentResponse, 0, len(documents))
	for _, document := range documents {
		items = append(items, *toIdentityDocumentResponse(&document))
	}

	return &dto.IdentityDocumentListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *IdentityDocumentService) Update(id int64, req dto.UpdateIdentityDocumentRequest) (*dto.IdentityDocumentResponse, error) {
	document, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if document == nil {
		return nil, ErrIdentityDocumentNotFound
	}

	documentType := strings.TrimSpace(req.DocumentType)
	documentSeries := strings.TrimSpace(req.DocumentSeries)
	documentNumber := strings.TrimSpace(req.DocumentNumber)
	issuedBy := strings.TrimSpace(req.IssuedBy)
	citizenship := strings.TrimSpace(req.Citizenship)

	if !isValidDocumentType(documentType) {
		return nil, ErrIdentityDocumentInvalidType
	}

	if documentSeries == "" {
		return nil, ErrIdentityDocumentInvalidSeries
	}

	if documentNumber == "" {
		return nil, ErrIdentityDocumentInvalidNumber
	}

	if issuedBy == "" {
		return nil, ErrIdentityDocumentInvalidIssuedBy
	}

	if citizenship == "" {
		return nil, ErrIdentityDocumentInvalidCitizenship
	}

	if req.TouristID <= 0 {
		return nil, ErrIdentityDocumentInvalidTouristID
	}

	issueDate, err := parseDateYYYYMMDD(req.IssueDate)
	if err != nil {
		return nil, ErrIdentityDocumentInvalidIssueDate
	}

	expirationDate, err := parseOptionalDateYYYYMMDD(req.ExpirationDate)
	if err != nil {
		return nil, ErrIdentityDocumentInvalidExpirationDate
	}

	if expirationDate != nil && !expirationDate.After(issueDate) {
		return nil, ErrIdentityDocumentInvalidExpirationDate
	}

	document.DocumentType = documentType
	document.DocumentSeries = documentSeries
	document.DocumentNumber = documentNumber
	document.ExpirationDate = expirationDate
	document.IssuedBy = issuedBy
	document.IssueDate = issueDate
	document.Citizenship = citizenship
	document.TouristID = req.TouristID

	if err := s.repo.Update(document); err != nil {
		return nil, err
	}

	return toIdentityDocumentResponse(document), nil
}

func (s *IdentityDocumentService) Delete(id int64) error {
	document, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if document == nil {
		return ErrIdentityDocumentNotFound
	}

	return s.repo.Delete(document)
}

func toIdentityDocumentResponse(document *models.IdentityDocument) *dto.IdentityDocumentResponse {
	var expirationDate *string
	if document.ExpirationDate != nil {
		formatted := document.ExpirationDate.Format("2006-01-02")
		expirationDate = &formatted
	}

	return &dto.IdentityDocumentResponse{
		ID:             document.ID,
		DocumentType:   document.DocumentType,
		DocumentSeries: document.DocumentSeries,
		DocumentNumber: document.DocumentNumber,
		ExpirationDate: expirationDate,
		IssuedBy:       document.IssuedBy,
		IssueDate:      document.IssueDate.Format("2006-01-02"),
		Citizenship:    document.Citizenship,
		TouristID:      document.TouristID,
	}
}

func isValidDocumentType(documentType string) bool {
	switch documentType {
	case "PASSPORT", "BIRTH_CERTIFICATE", "INTERNATIONAL_PASSPORT":
		return true
	default:
		return false
	}
}

func parseOptionalDateYYYYMMDD(value *string) (*time.Time, error) {
	if value == nil {
		return nil, nil
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil, nil
	}

	parsed, err := time.Parse("2006-01-02", trimmed)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}
