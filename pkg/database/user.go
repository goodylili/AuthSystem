package database

import (
	"authenticationsystem/pkg/user"
	"context"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Age       uint   `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Phone     uint   `gorm:"unique;not null"`
}

func (d *Database) CreateUser(ctx context.Context, user *user.User) error {
	dbUser := &User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Age:       user.Age,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
	}

	if err := d.Client.WithContext(ctx).Create(dbUser).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID returns the user with a specified id
func (d *Database) GetUserByID(ctx context.Context, id int64) (*user.User, error) {
	var dbUser User
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&dbUser).Error; err != nil {
		return nil, err
	}
	return &user.User{
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Username:  dbUser.Username,
		Age:       dbUser.Age,
		Password:  dbUser.Password,
		Email:     dbUser.Email,
		Phone:     dbUser.Phone,
	}, nil
}

func (d *Database) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	var dbUser User
	err := d.Client.WithContext(ctx).Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
	return &user.User{
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Username:  dbUser.Username,
		Age:       dbUser.Age,
		Password:  dbUser.Password,
		Email:     dbUser.Email,
		Phone:     dbUser.Phone,
	}, nil
}

func (d *Database) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	var dbUser User
	err := d.Client.WithContext(ctx).Where("username = ?", username).First(&dbUser).Error
	if err != nil {
		return nil, err
	}
	return &user.User{
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Username:  dbUser.Username,
		Age:       dbUser.Age,
		Password:  dbUser.Password,
		Email:     dbUser.Email,
		Phone:     dbUser.Phone,
	}, nil
}

func (d *Database) UpdateUser(ctx context.Context, user *user.User) error {
	var dbUser User

	// Check if user exists
	if err := d.Client.WithContext(ctx).Where("id = ?", user.ID).First(&dbUser).Error; err != nil {
		return err
	}

	// Check if the passwords match using the comparePasswords function
	if user.Password != dbUser.Password {
		return errors.New("password does not match")
	}

	dbUser.Username = user.Username
	dbUser.Email = user.Email

	// If the user exists and passwords match, update the database with the user's new details
	if err := d.Client.WithContext(ctx).Save(&dbUser).Error; err != nil {
		return err
	}
	return nil
}
