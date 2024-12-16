package models

type User struct {
	Id       int    `db:"user_id"`
	Username string `db:"username"`
}

type Team struct {
	TeamId     string `db:"team_id"`
	Name       string `db:"name"`
	InviteLink string `db:"invite_link"`
	Members    []User
}
