package users

import (
	"context"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model `json:"-"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Age        uint   `json:"age" binding:"required,gte=18"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Phone      uint   `json:"phone" binding:"required, startswith=0, len=11"`
	IsActive   bool   `json:"is_active" binding:"required"`
}

type Service interface {
	CreateUser(context.Context, User) error
	GetUserByID(context.Context, uint) (User, error)
	GetByEmail(context.Context, string) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	GetUserByFullName(context.Context, string) (*User, error)
	UpdateUserByID(context.Context, User) error
	DeactivateUserByID(context.Context, uint) error
}

// StoreImpl  is the blueprint for the users logic
type StoreImpl struct {
	Store Service
}

// NewStoreImpl returns a new instance of the StoreImpl struct.
func NewStoreImpl(store Service) *StoreImpl {
	return &StoreImpl{
		Store: store,
	}
}

func (u *StoreImpl) CreateUser(ctx context.Context, user User) error {
	if err := u.Store.CreateUser(ctx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}
	return nil
}

func (u *StoreImpl) GetUserByID(ctx context.Context, id uint) (User, error) {
	user, err := u.Store.GetUserByID(ctx, id)
	if err != nil {
		log.Printf("Error fetching user with ID %v: %v", id, err)
		return user, err
	}
	return user, nil
}

func (u *StoreImpl) UpdateUserByID(ctx context.Context, user User) error {
	if err := u.Store.UpdateUserByID(ctx, user); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

func (u *StoreImpl) DeactivateUserByID(ctx context.Context, id uint) error {
	if err := u.Store.DeactivateUserByID(ctx, id); err != nil {
		log.Printf("Error deactivating user with ID %v: %v", id, err)
		return err
	}

	return nil
}

func (u *StoreImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := u.Store.GetByEmail(ctx, email)
	if err != nil {
		log.Printf("Error fetching user with email %v: %v", email, err)
		return nil, err
	}

	return user, nil
}

func (u *StoreImpl) GetByUsername(ctx context.Context, username string) (*User, error) {
	user, err := u.Store.GetByUsername(ctx, username)
	if err != nil {
		log.Printf("Error fetching user with username %v: %v", username, err)
		return nil, err
	}

	return user, nil
}

func (u *StoreImpl) GetUserByFullName(ctx context.Context, fullName string) (*User, error) {
	// Implement the logic to get a user by full name using the Store service
	user, err := u.Store.GetUserByFullName(ctx, fullName)
	if err != nil {
		log.Printf("Error fetching user by full name  %v: %v", fullName, err)
		return nil, err
	}
	return user, nil
}
