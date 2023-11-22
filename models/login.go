package models

import "errors"

type LoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (lq *LoginRequest) Validate() error {
	if lq.Username == "" {
		return errors.New("username is required")
	}
	if lq.Password == "" {
		return errors.New("unauthorized access")
	}
	return nil
}

type ResponseLogin struct {
	UserId int64  `json:"UserId"`
	Token  string `json:"Token"`
}
