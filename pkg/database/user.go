package database

import (
	"Reusable-Auth-System/pkg/users"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Age       int64  `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Phone     int64  `gorm:"unique;not null"`
	IsActive  bool   `gorm:"not null"`
	RoleID    int64  `gorm:"not null"`
}

func (d *Database) CreateUser(ctx context.Context, user users.User) error {
	if !d.HasPermission(int64(user.ID), CanCreateAccount) {
		return errors.New("access denied: insufficient permissions")
	}

	d.Logger.WithFields(logrus.Fields{
		"function": "CreateUser",
		"user":     user,
	}).Info("Creating a new user")

	newUser := &User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Age:       user.Age,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
		RoleID:    RoleUser,
	}

	if err := d.Client.WithContext(ctx).Create(newUser).Error; err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "CreateUser",
			"error":    err,
		}).Error("Error creating user")
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (d *Database) GetUserByID(ctx context.Context, id int64) (users.User, error) {

	if !d.HasPermission(id, CanViewUsers) {
		return users.User{}, errors.New("access denied: insufficient permissions")
	}

	var user User
	err := d.Client.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "GetUserByID",
			"ID":       id,
			"error":    err,
		}).Error("Error fetching user by ID")
		return users.User{}, fmt.Errorf("error fetching user by ID %d: %w", id, err)
	}
	return users.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
	}, nil
}

func (d *Database) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	var user User

	err := d.Client.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "GetByEmail",
			"Email":    email,
			"error":    err,
		}).Error("Error fetching user by email")
		return nil, fmt.Errorf("error fetching user by email %s: %w", email, err)
	}

	if !d.HasPermission(int64(user.ID), CanViewUsers) {
		return &users.User{}, errors.New("access denied: insufficient permissions")
	}

	return &users.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
	}, nil
}

func (d *Database) GetByUsername(ctx context.Context, username string) (*users.User, error) {
	var user User
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "GetByUsername",
			"Username": username,
			"error":    err,
		}).Error("Error fetching user by username")
		return nil, fmt.Errorf("error fetching user by username %s: %w", username, err)
	}

	if !d.HasPermission(int64(user.ID), CanViewUsers) {
		return &users.User{}, errors.New("access denied: insufficient permissions")
	}

	return &users.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
	}, nil
}

func (d *Database) GetUserByFullName(ctx context.Context, fullName string) (*users.User, error) {
	var user User
	names := strings.Fields(fullName)
	if len(names) != 2 {
		err := errors.New("invalid full name format")
		d.Logger.WithFields(logrus.Fields{
			"function": "GetUserByFullName",
			"FullName": fullName,
			"error":    err,
		}).Error("Invalid full name format")
		return nil, err
	}

	firstName := names[0]
	lastName := names[1]

	err := d.Client.WithContext(ctx).Where("first_name = ? AND last_name = ?", firstName, lastName).First(&user).Error
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "GetUserByFullName",
			"FullName": fullName,
			"error":    err,
		}).Error("Error fetching user by full name")
		return nil, fmt.Errorf("error fetching user by full name %s: %w", fullName, err)
	}

	if !d.HasPermission(int64(user.ID), CanViewUsers) {
		return &users.User{}, errors.New("access denied: insufficient permissions")
	}

	return &users.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
	}, nil
}

func (d *Database) UpdateUserByID(ctx context.Context, user users.User, id int64) error {
	var existingUser User
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error; err != nil {
		d.Logger.Error("Error querying user", zap.Int64("ID", id), zap.Error(err))
		return fmt.Errorf("error querying user: %w", err)
	}
	updateColumns := make(map[string]interface{})

	if user.Username != "" {
		updateColumns["username"] = user.Username
	}
	if user.Email != "" {
		updateColumns["email"] = user.Email
	}
	if user.Phone != 0 {
		updateColumns["phone"] = user.Phone
	}
	if user.FirstName != "" {
		updateColumns["first_name"] = user.FirstName
	}
	if user.LastName != "" {
		updateColumns["last_name"] = user.LastName
	}
	if user.Age != 0 {
		updateColumns["age"] = user.Age
	}

	if len(updateColumns) == 0 {
		return nil // Nothing to update
	}

	if !d.HasPermission(int64(user.ID), CanUpdateDetails) {
		return errors.New("access denied: insufficient permissions")
	}

	if err := d.Client.WithContext(ctx).Model(&existingUser).Omit("RoleID", "IsActive", "Password").Updates(updateColumns).Error; err != nil {
		d.Logger.Error("Error updating user", zap.Int64("ID", id), zap.Error(err))
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}

func (d *Database) SetActivity(ctx context.Context, id int64, action bool) error {
	var user User
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "SetActivity",
			"ID":       id,
			"error":    err,
		}).Error("Error fetching user by ID for deactivation")
		return fmt.Errorf("error fetching user by ID %d for deactivation: %w", id, err)
	}
	user.IsActive = action
	if err := d.Client.WithContext(ctx).Save(&user).Error; err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "SetActivity",
			"ID":       id,
			"error":    err,
		}).Error("Error deactivating user")
		return fmt.Errorf("error deactivating user: %w", err)
	}

	if !d.HasPermission(int64(user.ID), CanUpdateDetails) {
		return errors.New("access denied: insufficient permissions")
	}

	return nil
}

func (d *Database) ResetPassword(ctx context.Context, newUser users.User) error {
	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "ResetPassword",
			"error":    err,
		}).Error("Error hashing password")
		return fmt.Errorf("error hashing password: %w", err)
	}

	result := d.Client.WithContext(ctx).Model(&User{}).
		Where("username = ? AND email = ? AND is_active = ?", newUser.Username, newUser.Email, newUser.IsActive).
		Updates(map[string]interface{}{"password": hashedPassword})

	if result.RowsAffected == 0 {
		return errors.New("no matching active user found with the provided username and email")
	}

	if result.Error != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "ResetPassword",
			"error":    result.Error,
		}).Error("Error updating password")
		return fmt.Errorf("error updating password: %w", result.Error)
	}

	return nil
}

func (d *Database) UpdateUserRoleID(ctx context.Context, id uint, newRoleID int64) error {
	var existingUser User
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&existingUser).Error; err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function": "UpdateUserRoleID",
			"ID":       id,
			"error":    err,
		}).Error("Error querying user")
		return fmt.Errorf("error querying user: %w", err)
	}

	existingUser.RoleID = newRoleID
	if err := d.Client.WithContext(ctx).Model(&existingUser).Updates(User{RoleID: newRoleID}).Error; err != nil {
		d.Logger.WithFields(logrus.Fields{
			"function":  "UpdateUserRoleID",
			"ID":        id,
			"newRoleID": newRoleID,
			"error":     err,
		}).Error("Error updating user role ID")
		return fmt.Errorf("error updating user role ID: %w", err)
	}

	return nil
}

func (d *Database) DeleteUserByID(ctx context.Context, id int64) error {
	result := d.Client.WithContext(ctx).Delete(&User{}, id)
	if result.Error != nil {
		d.Logger.WithFields(logrus.Fields{
			"id":    id,
			"error": result.Error,
		}).Error("Error deleting user")
		return result.Error
	}
	if result.RowsAffected == 0 {
		err := fmt.Errorf("no user found with id %v", id)
		d.Logger.WithFields(logrus.Fields{
			"id": id,
		}).Warn("No user found for deletion")
		return err
	}
	d.Logger.WithFields(logrus.Fields{
		"id": id,
	}).Info("User deleted successfully")
	return nil
}

func (d *Database) SignIn(ctx context.Context, username, password string) error {
	var user User
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return errors.New("invalid credentials")
		}
		d.Logger.WithFields(logrus.Fields{
			"function": "SignIn",
			"username": username,
			"error":    err,
		}).Error("Error fetching user by username")
		return fmt.Errorf("error fetching user by username %s: %w", username, err)
	}

	// Use ComparePassword function to compare the passwords
	if !ComparePassword(password, user.Password) {
		return errors.New("invalid credentials")
	}

	// If the passwords match, sign-in is successful, return nil (no error)
	return nil
}

// ChangePassword changes the user's password.
func (d *Database) ChangePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	var user User
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return fmt.Errorf("error fetching user by username: %w", err)
	}

	if !ComparePassword(oldPassword, user.Password) {
		return errors.New("invalid old password")
	}

	hashedNewPassword, err := HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("error hashing new password: %w", err)
	}

	result := d.Client.WithContext(ctx).Model(&User{}).
		Where("username = ?", username).
		Updates(map[string]interface{}{"password": hashedNewPassword})

	if result.Error != nil {
		return fmt.Errorf("error updating password: %w", result.Error)
	}

	return nil
}

// ForgotPassword handles the forgot password process.
func (d *Database) ForgotPassword(ctx context.Context, email string) error {
	_, err := d.GetByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("error fetching user by email: %w", err)
	}
	return nil
}
