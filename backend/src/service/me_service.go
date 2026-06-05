package service

import (
	"errors"
	"strings"
	"time"

	"TouristAgencyApp/src/dto"
	"TouristAgencyApp/src/repository"
)

var (
	ErrInvalidInput             = errors.New("invalid input")
	ErrMeTouristNotLinked       = errors.New("tourist profile is not linked to user")
	ErrMeDocumentNotFound       = errors.New("identity document not found")
	ErrMeDocumentAlreadyExists  = errors.New("identity document already exists")
	ErrMeDocumentRequired       = errors.New("identity document required")
	ErrMeGroupMemberNotFound    = errors.New("group member not found")
	ErrMeMultipleGroupMembers   = errors.New("multiple group members found")
	ErrMeScheduleNotFound       = errors.New("excursion schedule not found")
	ErrMeExcursionAlreadyBooked = errors.New("excursion already booked")
)

type MeService struct {
	meRepository *repository.MeRepository
}

func NewMeService(meRepository *repository.MeRepository) *MeService {
	return &MeService{
		meRepository: meRepository,
	}
}

func (s *MeService) Tours(userID int64, page int, limit int) (*dto.MeTourListResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	if err := s.ensureTouristLinked(userID); err != nil {
		return nil, err
	}

	page, limit, offset := normalizeMePagination(page, limit)

	items, total, err := s.meRepository.FindToursByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.MeTourListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *MeService) Visas(userID int64, page int, limit int) (*dto.MeVisaListResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	if err := s.ensureTouristLinked(userID); err != nil {
		return nil, err
	}

	page, limit, offset := normalizeMePagination(page, limit)

	items, total, err := s.meRepository.FindVisasByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.MeVisaListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *MeService) Accommodations(userID int64, page int, limit int) (*dto.MeAccommodationListResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	if err := s.ensureTouristLinked(userID); err != nil {
		return nil, err
	}

	page, limit, offset := normalizeMePagination(page, limit)

	items, total, err := s.meRepository.FindAccommodationsByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.MeAccommodationListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *MeService) Excursions(userID int64, page int, limit int) (*dto.MeExcursionListResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	if err := s.ensureTouristLinked(userID); err != nil {
		return nil, err
	}

	page, limit, offset := normalizeMePagination(page, limit)

	items, total, err := s.meRepository.FindExcursionsByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.MeExcursionListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *MeService) Cargo(userID int64, page int, limit int) (*dto.MeCargoListResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	if err := s.ensureTouristLinked(userID); err != nil {
		return nil, err
	}

	page, limit, offset := normalizeMePagination(page, limit)

	items, total, err := s.meRepository.FindCargoByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return &dto.MeCargoListResponse{
		Items: items,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *MeService) ensureTouristLinked(userID int64) error {
	touristID, err := s.meRepository.FindTouristIDByUserID(userID)
	if err != nil {
		return err
	}

	if touristID == nil {
		return ErrMeTouristNotLinked
	}

	return nil
}

func normalizeMePagination(page int, limit int) (int, int, int) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 20
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	return page, limit, offset
}

func (s *MeService) IdentityDocument(userID int64) (*dto.MeIdentityDocumentResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	if err := s.ensureTouristLinked(userID); err != nil {
		return nil, err
	}

	document, err := s.meRepository.FindIdentityDocumentByUserID(userID)
	if err != nil {
		return nil, err
	}

	if document == nil {
		return nil, ErrMeDocumentNotFound
	}

	return document, nil
}

func (s *MeService) CreateIdentityDocument(
	userID int64,
	request dto.CreateMeIdentityDocumentRequest,
) (*dto.MeIdentityDocumentResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	touristID, err := s.meRepository.FindTouristIDByUserID(userID)
	if err != nil {
		return nil, err
	}

	if touristID == nil {
		return nil, ErrMeTouristNotLinked
	}

	existingDocument, err := s.meRepository.FindIdentityDocumentByUserID(userID)
	if err != nil {
		return nil, err
	}

	if existingDocument != nil {
		return nil, ErrMeDocumentAlreadyExists
	}

	documentType := normalizeMeDocumentType(request.DocumentType)
	documentSeries := strings.TrimSpace(request.DocumentSeries)
	documentNumber := strings.TrimSpace(request.DocumentNumber)
	issueDate := strings.TrimSpace(request.IssueDate)
	expirationDate := normalizeMeStringPointer(request.ExpirationDate)
	issuedBy := strings.TrimSpace(request.IssuedBy)
	citizenship := strings.TrimSpace(request.Citizenship)

	if !isValidMeDocumentType(documentType) ||
		documentSeries == "" ||
		documentNumber == "" ||
		issueDate == "" ||
		issuedBy == "" ||
		citizenship == "" {
		return nil, ErrInvalidInput
	}

	if err := validateMeDocumentDates(issueDate, expirationDate); err != nil {
		return nil, ErrInvalidInput
	}

	return s.meRepository.CreateIdentityDocument(
		*touristID,
		documentType,
		documentSeries,
		documentNumber,
		issueDate,
		expirationDate,
		issuedBy,
		citizenship,
	)
}

func (s *MeService) UpdateIdentityDocument(
	userID int64,
	request dto.UpdateMeIdentityDocumentRequest,
) (*dto.MeIdentityDocumentResponse, error) {
	if userID <= 0 {
		return nil, ErrInvalidInput
	}

	currentDocument, err := s.meRepository.FindIdentityDocumentByUserID(userID)
	if err != nil {
		return nil, err
	}

	if currentDocument == nil {
		return nil, ErrMeDocumentNotFound
	}

	documentType := currentDocument.DocumentType
	documentSeries := currentDocument.DocumentSeries
	documentNumber := currentDocument.DocumentNumber
	issueDate := currentDocument.IssueDate
	expirationDate := currentDocument.ExpirationDate
	issuedBy := currentDocument.IssuedBy
	citizenship := currentDocument.Citizenship

	if request.DocumentType != nil {
		documentType = normalizeMeDocumentType(*request.DocumentType)
	}

	if request.DocumentSeries != nil {
		documentSeries = strings.TrimSpace(*request.DocumentSeries)
	}

	if request.DocumentNumber != nil {
		documentNumber = strings.TrimSpace(*request.DocumentNumber)
	}

	if request.IssueDate != nil {
		issueDate = strings.TrimSpace(*request.IssueDate)
	}

	if request.ExpirationDate != nil {
		expirationDate = normalizeMeStringPointer(request.ExpirationDate)
	}

	if request.IssuedBy != nil {
		issuedBy = strings.TrimSpace(*request.IssuedBy)
	}

	if request.Citizenship != nil {
		citizenship = strings.TrimSpace(*request.Citizenship)
	}

	if !isValidMeDocumentType(documentType) ||
		documentSeries == "" ||
		documentNumber == "" ||
		issueDate == "" ||
		issuedBy == "" ||
		citizenship == "" {
		return nil, ErrInvalidInput
	}

	if err := validateMeDocumentDates(issueDate, expirationDate); err != nil {
		return nil, ErrInvalidInput
	}

	updatedDocument, err := s.meRepository.UpdateIdentityDocument(
		currentDocument.ID,
		documentType,
		documentSeries,
		documentNumber,
		issueDate,
		expirationDate,
		issuedBy,
		citizenship,
	)

	if err != nil {
		return nil, err
	}

	if updatedDocument == nil {
		return nil, ErrMeDocumentNotFound
	}

	return updatedDocument, nil
}

func (s *MeService) CreateExcursionBooking(
	userID int64,
	request dto.CreateMeExcursionBookingRequest,
) (*dto.MeExcursionBookingResponse, error) {
	if userID <= 0 || request.ExcursionScheduleID <= 0 {
		return nil, ErrInvalidInput
	}

	if err := s.ensureTouristLinked(userID); err != nil {
		return nil, err
	}

	hasDocument, err := s.meRepository.HasIdentityDocument(userID)
	if err != nil {
		return nil, err
	}

	if !hasDocument {
		return nil, ErrMeDocumentRequired
	}

	scheduleExists, err := s.meRepository.ExcursionScheduleExists(request.ExcursionScheduleID)
	if err != nil {
		return nil, err
	}

	if !scheduleExists {
		return nil, ErrMeScheduleNotFound
	}

	groupMemberID, err := s.resolveGroupMemberID(userID, request.GroupMemberID)
	if err != nil {
		return nil, err
	}

	alreadyBooked, err := s.meRepository.ExcursionBookingExists(
		request.ExcursionScheduleID,
		groupMemberID,
	)

	if err != nil {
		return nil, err
	}

	if alreadyBooked {
		return nil, ErrMeExcursionAlreadyBooked
	}

	return s.meRepository.CreateExcursionBooking(
		request.ExcursionScheduleID,
		groupMemberID,
	)
}

func (s *MeService) resolveGroupMemberID(userID int64, requestedGroupMemberID *int64) (int64, error) {
	if requestedGroupMemberID != nil {
		if *requestedGroupMemberID <= 0 {
			return 0, ErrInvalidInput
		}

		exists, err := s.meRepository.GroupMemberBelongsToUser(userID, *requestedGroupMemberID)
		if err != nil {
			return 0, err
		}

		if !exists {
			return 0, ErrMeGroupMemberNotFound
		}

		return *requestedGroupMemberID, nil
	}

	groupMemberIDs, err := s.meRepository.FindGroupMemberIDsByUserID(userID)
	if err != nil {
		return 0, err
	}

	if len(groupMemberIDs) == 0 {
		return 0, ErrMeGroupMemberNotFound
	}

	if len(groupMemberIDs) > 1 {
		return 0, ErrMeMultipleGroupMembers
	}

	return groupMemberIDs[0], nil
}

func normalizeMeStringPointer(value *string) *string {
	if value == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		return nil
	}

	return &trimmed
}

func normalizeMeDocumentType(value string) string {
	return strings.ToUpper(strings.TrimSpace(value))
}

func isValidMeDocumentType(value string) bool {
	switch value {
	case "PASSPORT", "BIRTH_CERTIFICATE", "INTERNATIONAL_PASSPORT":
		return true
	default:
		return false
	}
}

func validateMeDocumentDates(issueDate string, expirationDate *string) error {
	parsedIssueDate, err := time.Parse("2006-01-02", issueDate)
	if err != nil {
		return err
	}

	if expirationDate == nil {
		return nil
	}

	parsedExpirationDate, err := time.Parse("2006-01-02", *expirationDate)
	if err != nil {
		return err
	}

	if !parsedExpirationDate.After(parsedIssueDate) {
		return errors.New("expiration date must be after issue date")
	}

	return nil
}
