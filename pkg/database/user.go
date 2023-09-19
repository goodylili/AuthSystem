package database

import (
	"authenticationsystem/pkg/users"
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strings"
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
	IsActive  bool   `gorm:"not null"`
}

func (d *Database) CreateUser(ctx context.Context, user users.User) error {
	newUser := &User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Age:       user.Age,
		Password:  user.Password,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
	}

	if err := d.Client.WithContext(ctx).Create(newUser).Error; err != nil {
		return err
	}

	return nil
}

// GetUserByID returns the users with a specified id
func (d *Database) GetUserByID(ctx context.Context, id uint) (users.User, error) {
	user := User{}
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return users.User{}, err
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
		return nil, err
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
		return nil, err
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
		return nil, errors.New("invalid full name format")
	}

	firstName := names[0]
	lastName := names[1]

	if err := d.Client.WithContext(ctx).Where("first_name = ? AND last_name = ?", firstName, lastName).First(&user).Error; err != nil {
		return nil, err
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

func (d *Database) UpdateUserByID(ctx context.Context, user users.User) error {
	var newUser User

	// Check if users exists
	if err := d.Client.WithContext(ctx).Where("id = ?", user.ID).First(&newUser).Error; err != nil {
		return err
	}

	newUser = User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
		Phone:     user.Phone,
		IsActive:  user.IsActive,
	}

	// if the users exists and passwords match, update the database with the users's new details
	if err := d.Client.WithContext(ctx).Save(&newUser).Error; err != nil {
		return err
	}
	return nil
}

// DeactivateUserByID sets the user's IsActive status to false based on the provided ID.
func (d *Database) DeactivateUserByID(ctx context.Context, id uint) error {
	user := User{}
	if err := d.Client.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		// Return an error if the users is not found
		return err
	}
	// Set the user's IsActive status to false
	user.IsActive = false
	if err := d.Client.WithContext(ctx).Save(&user).Error; err != nil {
		return err
	}

	return nil
}
