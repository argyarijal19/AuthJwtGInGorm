package models

type UserSimgoa struct {
	Username    string
	Plasma      string
	Description string
	Password    string
}

type LoginUser struct {
	Username string
	Password string
}

type NewAccessToken struct {
	AccsessToken string
}
