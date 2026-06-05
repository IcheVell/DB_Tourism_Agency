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
	ErrVisaNotFound                  = errors.New("visa not found")
	ErrVisaInvalidNumber             = errors.New("visa number is required for issued status")
	ErrVisaInvalidDestinationCountry = errors.New("destination_country is required")
	ErrVisaInvalidStatus             = errors.New("visa status is invalid")
	ErrVisaInvalidTouristID          = errors.New("tourist_id must be positive")
	ErrVisaInvalidSubmittedAt        = errors.New("submitted_at must be RFC3339 datetime")
	ErrVisaInvalidDecisionAt         = errors.New("decision_at must be RFC3339 datetime")
	ErrVisaInvalidIssuedAt           = errors.New("issued_at must be RFC3339 datetime")
	ErrVisaInvalidValidFrom          = errors.New("valid_from must be YYYY-MM-DD date")
	ErrVisaInvalidValidUntil         = errors.New("valid_until must be YYYY-MM-DD date")
	ErrVisaInvalidTiming             = errors.New("visa timestamps are inconsistent")
	ErrVisaInvalidValidityPeriod     = errors.New("valid_from must be before valid_until")
	ErrVisaIssuedFieldsRequired      = errors.New("issued visa requires number, issued_at, valid_from and valid_until")
)

type VisaService struct {
	repo *repository.VisaRepository
}

func NewVisaService(repo *repository.VisaRepository) *VisaService {
	return &VisaService{
		repo: repo,
	}
}

func (s *VisaService) Create(req dto.CreateVisaRequest) (*dto.VisaResponse, error) {
	visa, err := buildVisaFromCreateRequest(req)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(visa); err != nil {
		return nil, err
	}

	return toVisaResponse(visa), nil
}

func (s *VisaService) GetByID(id int64) (*dto.VisaResponse, error) {
	visa, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if visa == nil {
		return nil, ErrVisaNotFound
	}

	return toVisaResponse(visa), nil
}

func (s *VisaService) List(page int, pageSize int) (*dto.VisaListResponse, error) {
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

	visas, total, err := s.repo.FindAll(pageSize, offset)
	if err != nil {
		return nil, err
	}

	items := make([]dto.VisaResponse, 0, len(visas))
	for _, visa := range visas {
		items = append(items, *toVisaResponse(&visa))
	}

	return &dto.VisaListResponse{
		Items:    items,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	}, nil
}

func (s *VisaService) Update(id int64, req dto.UpdateVisaRequest) (*dto.VisaResponse, error) {
	visa, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if visa == nil {
		return nil, ErrVisaNotFound
	}

	updatedVisa, err := buildVisaFromUpdateRequest(req, visa.CreatedAt)
	if err != nil {
		return nil, err
	}

	visa.Number = updatedVisa.Number
	visa.DestinationCountry = updatedVisa.DestinationCountry
	visa.Status = updatedVisa.Status
	visa.SubmittedAt = updatedVisa.SubmittedAt
	visa.DecisionAt = updatedVisa.DecisionAt
	visa.IssuedAt = updatedVisa.IssuedAt
	visa.ValidFrom = updatedVisa.ValidFrom
	visa.ValidUntil = updatedVisa.ValidUntil
	visa.TouristID = updatedVisa.TouristID

	if err := s.repo.Update(visa); err != nil {
		return nil, err
	}

	return toVisaResponse(visa), nil
}

func (s *VisaService) Delete(id int64) error {
	visa, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if visa == nil {
		return ErrVisaNotFound
	}

	return s.repo.Delete(visa)
}

func buildVisaFromCreateRequest(req dto.CreateVisaRequest) (*models.Visa, error) {
	createdAt := time.Now().UTC()

	return buildVisa(
		req.Number,
		req.DestinationCountry,
		req.Status,
		createdAt,
		req.SubmittedAt,
		req.DecisionAt,
		req.IssuedAt,
		req.ValidFrom,
		req.ValidUntil,
		req.TouristID,
	)
}

func buildVisaFromUpdateRequest(req dto.UpdateVisaRequest, createdAt time.Time) (*models.Visa, error) {
	return buildVisa(
		req.Number,
		req.DestinationCountry,
		req.Status,
		createdAt,
		req.SubmittedAt,
		req.DecisionAt,
		req.IssuedAt,
		req.ValidFrom,
		req.ValidUntil,
		req.TouristID,
	)
}

func buildVisa(
	number *string,
	destinationCountry string,
	status string,
	createdAt time.Time,
	submittedAtRaw *string,
	decisionAtRaw *string,
	issuedAtRaw *string,
	validFromRaw *string,
	validUntilRaw *string,
	touristID int64,
) (*models.Visa, error) {
	normalizedNumber := normalizeOptionalString(number)
	normalizedDestinationCountry := strings.TrimSpace(destinationCountry)
	normalizedStatus := strings.TrimSpace(status)

	if normalizedDestinationCountry == "" {
		return nil, ErrVisaInvalidDestinationCountry
	}

	if !isValidVisaStatus(normalizedStatus) {
		return nil, ErrVisaInvalidStatus
	}

	if touristID <= 0 {
		return nil, ErrVisaInvalidTouristID
	}

	submittedAt, err := parseOptionalDateTimeRFC3339(submittedAtRaw)
	if err != nil {
		return nil, ErrVisaInvalidSubmittedAt
	}

	decisionAt, err := parseOptionalDateTimeRFC3339(decisionAtRaw)
	if err != nil {
		return nil, ErrVisaInvalidDecisionAt
	}

	issuedAt, err := parseOptionalDateTimeRFC3339(issuedAtRaw)
	if err != nil {
		return nil, ErrVisaInvalidIssuedAt
	}

	validFrom, err := parseOptionalDateYYYYMMDD(validFromRaw)
	if err != nil {
		return nil, ErrVisaInvalidValidFrom
	}

	validUntil, err := parseOptionalDateYYYYMMDD(validUntilRaw)
	if err != nil {
		return nil, ErrVisaInvalidValidUntil
	}

	if submittedAt != nil && createdAt.After(*submittedAt) {
		return nil, ErrVisaInvalidTiming
	}

	if submittedAt != nil && decisionAt != nil && submittedAt.After(*decisionAt) {
		return nil, ErrVisaInvalidTiming
	}

	if decisionAt != nil && issuedAt != nil && decisionAt.After(*issuedAt) {
		return nil, ErrVisaInvalidTiming
	}

	if validFrom != nil && validUntil != nil && !validFrom.Before(*validUntil) {
		return nil, ErrVisaInvalidValidityPeriod
	}

	if normalizedStatus == "issued" {
		if normalizedNumber == nil || issuedAt == nil || validFrom == nil || validUntil == nil {
			return nil, ErrVisaIssuedFieldsRequired
		}
	}

	return &models.Visa{
		Number:             normalizedNumber,
		DestinationCountry: normalizedDestinationCountry,
		Status:             normalizedStatus,
		CreatedAt:          createdAt,
		SubmittedAt:        submittedAt,
		DecisionAt:         decisionAt,
		IssuedAt:           issuedAt,
		ValidFrom:          validFrom,
		ValidUntil:         validUntil,
		TouristID:          touristID,
	}, nil
}

func toVisaResponse(visa *models.Visa) *dto.VisaResponse {
	return &dto.VisaResponse{
		ID:                 visa.ID,
		Number:             visa.Number,
		DestinationCountry: visa.DestinationCountry,
		Status:             visa.Status,
		CreatedAt:          visa.CreatedAt.Format(time.RFC3339),
		SubmittedAt:        formatOptionalDateTimeRFC3339(visa.SubmittedAt),
		DecisionAt:         formatOptionalDateTimeRFC3339(visa.DecisionAt),
		IssuedAt:           formatOptionalDateTimeRFC3339(visa.IssuedAt),
		ValidFrom:          formatOptionalDateYYYYMMDD(visa.ValidFrom),
		ValidUntil:         formatOptionalDateYYYYMMDD(visa.ValidUntil),
		TouristID:          visa.TouristID,
	}
}

func isValidVisaStatus(status string) bool {
	switch status {
	case "draft", "submitted", "rejected", "issued", "cancelled", "expired":
		return true
	default:
		return false
	}
}

func parseOptionalDateTimeRFC3339(value *string) (*time.Time, error) {
	if value == nil {
		return nil, nil
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil, nil
	}

	parsed, err := time.Parse(time.RFC3339, trimmed)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}

func formatOptionalDateTimeRFC3339(value *time.Time) *string {
	if value == nil {
		return nil
	}

	formatted := value.Format(time.RFC3339)
	return &formatted
}

func formatOptionalDateYYYYMMDD(value *time.Time) *string {
	if value == nil {
		return nil
	}

	formatted := value.Format("2006-01-02")
	return &formatted
}
