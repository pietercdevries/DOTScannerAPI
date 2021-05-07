package models

import "time"

type Scan struct {
	Id int64 `json:"id"`
	ScanTypeId int `json:"scan_type_id"`
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	ExpiresOn string `json:"expires_on"`
	CreatedByUserId int `json:"created_by_user_id"`
}

type ScanType struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	IsPaidVersion int `json:"is_paid_version"`
	DefaultExpirationTime time.Time `json:"default_expiration_time"`
	CreatedByUserId int `json:"created_by_user_id"`
}

type User struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	LoginAttempts string `json:"login_attempts"`
	ResetPasswordCode string `json:"reset_password_code"`
	PaidMember int `json:"paid_member"`
	NotificationOn string `json:"notificationOn"`
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	TokenExpireDate time.Time `json:"token_expire_date"`
	CreatedByUserId int `json:"created_by_user_id"`
}