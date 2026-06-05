package service

import "TouristAgencyApp/src/repository"

type ACLService struct {
	aclRepo *repository.ACLRepository
}

func NewACLService(aclRepo *repository.ACLRepository) *ACLService {
	return &ACLService{
		aclRepo: aclRepo,
	}
}

func (s *ACLService) HasPermission(userID int64, permissionCode string) (bool, error) {
	return s.aclRepo.HasPermission(userID, permissionCode)
}
