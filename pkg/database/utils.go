package database

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePassword compares the password with the hashed password
func ComparePassword(password, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}

// HasPermission checks if the user has a permission
func (d *Database) HasPermission(userID int64, permission int) bool {
	var user User
	err := d.Client.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return false
	}
	for _, perm := range RolePermissionsMap[int(user.RoleID)] {
		if perm == permission {
			return true
		}
	}
	return false
}
