package checkerRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/flagGeneratorDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
)

type Repository interface {
	CreateChecker(checker *checkerDomain.Checker) error
	UpdateChecker(checker *checkerDomain.Checker) error
	GetChecker(UUID uuid.UUID) (*checkerDomain.Checker, error)
	GetVulnServiceList(UUID uuid.UUID) ([]*vulnServiceDomain.VulnService, error)
	CreateRequestToVulnService(requestUUID, vulnServiceUUID uuid.UUID) error
	GetRequestToVulnService(requestUUID uuid.UUID) (*vulnServiceDomain.RequestToVulnService, error)
	CreateFlag(flag *flagGeneratorDomain.Flag) error
}
