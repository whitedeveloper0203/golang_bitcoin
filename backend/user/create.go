package user

import (
	"github.com/jinzhu/gorm"
	"github.com/tamizhvendan/gomidway/postgres"
)

func Create(db *gorm.DB, user *User) (uint, error) {
	err := db.Create(user).Error
	if err != nil {
		if postgres.IsUniqueConstraintError(err, UniqueConstraintAuth0UUID) {
			return 0, &Auth0UUIDDuplicateError{Auth0UUID: user.Auth0UUID}
		}
		if postgres.IsUniqueConstraintError(err, UniqueConstraintEmail) {
			return 0, &EmailDuplicateError{Email: user.Email}
		}
		return 0, err
	}
	return user.ID, nil
}
