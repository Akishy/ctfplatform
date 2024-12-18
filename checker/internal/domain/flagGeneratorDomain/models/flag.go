package models

import "github.com/google/uuid"

type FlagStatus uint // да но забыл как
const (
	FLAG_PUSHED_TO_CHECKER      FlagStatus = iota // доставлен в чекер
	FLAG_PUSHED_TO_VULN_SERVICE                   // чекер успешно положил в уязв. сервис
)

type Flag struct {
	Flag   uuid.UUID
	Status FlagStatus
}
