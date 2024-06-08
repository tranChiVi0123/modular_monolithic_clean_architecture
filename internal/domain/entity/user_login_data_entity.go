package entity

import "time"

type UsersLoginData struct {
	UserAccountID                string `gorm:"primaryKey;autoIncrement:true"`
	LoginName                    string `gorm:"unique"`
	PasswordHash                 string
	PasswordSalt                 string
	HasingAlgorithmID            int
	EmailAddress                 string `gorm:"unique"`
	EmailAddressVerificationCode string
	EmailAddressVerificationTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ValidationEmailToken         string
	ValidationEmailTime          time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	EmailVadationStatusID        int       // TODO: change the column name to EmailValidationStatusID
	PasswordRecoveryToken        string
	PasswordRecoveryTime         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt                    time.Time
	UpdatedAt                    time.Time
}
