package checkerRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/flagGeneratorDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
)

type Repository interface {
	checkerRepo
	vulnServiceRequestRepo
	vulnServiceRepo
	flagGeneratorRepo
}

type checkerRepo interface {
	CreateChecker(checker *checkerDomain.Checker) error
	UpdateChecker(checker *checkerDomain.Checker) error
	GetChecker(UUID uuid.UUID) (*checkerDomain.Checker, error)
	ListAllRegisteredCheckers() ([]*checkerDomain.Checker, error)
}

type vulnServiceRequestRepo interface {
	CreateRequestToVulnService(requestUUID, vulnServiceUUID uuid.UUID) error
	GetRequestToVulnService(requestUUID uuid.UUID) (*vulnServiceDomain.RequestToVulnService, error)
}

type flagGeneratorRepo interface {
	CreateFlag(flag *flagGeneratorDomain.Flag) error
}

type vulnServiceRepo interface {
	GetActiveVulnServiceList(UUID uuid.UUID) ([]*vulnServiceDomain.VulnService, error)
}

type checkerImgRepo interface {
}
