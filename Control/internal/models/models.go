package models

type User struct {
	Id       int    `db:"user_id" json:"id"`
	Username string `db:"username" json:"username"`
}

type Team struct {
	TeamId     string `db:"team_id" json:"team_id"`
	Name       string `db:"name" json:"name"`
	InviteLink string `db:"invite_link" json:"invite_link"`
	Members    []User
}
