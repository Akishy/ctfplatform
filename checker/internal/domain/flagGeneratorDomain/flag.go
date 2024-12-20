package flagGeneratorDomain

import "github.com/google/uuid"

type FlagStatus uint

const (
	FLAG_PUSHED_TO_CHECKER      FlagStatus = iota // доставлен в чекер
	FLAG_PUSHED_TO_VULN_SERVICE                   // чекер успешно положил в уязв. сервис
)

type Flag struct {
	UUID   uuid.UUID // id флага
	Flag   uuid.UUID // содержимое флага
	Status FlagStatus
}
