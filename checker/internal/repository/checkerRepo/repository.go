package checkerRepo

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"

type Repository interface {
	Create(checker *models.Checker) error
	Update(checker *models.Checker) error
}
