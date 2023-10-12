package users

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var log = logrus.New()

type User struct {
	gorm.Model `json:"-"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Age        int64  `json:"age" binding:"required,gte=18"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Phone      int64  `json:"phone" binding:"required, startswith=0, len=11"`
	IsActive   bool   `json:"is_active" binding:"required"`
	RoleID     int64  `json:"role_id" binding:"required, lte=3"`
}

type Service interface {
	CreateUser(context.Context, User) error
	GetUserByID(context.Context, int64) (User, error)
	GetByEmail(context.Context, string) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	GetUserByFullName(context.Context, string) (*User, error)
	UpdateUserByID(context.Context, User, int64) error
	SetActivity(context.Context, int64, bool) error
	UpdateUserRoleID(context.Context, uint, int64) error
	ResetPassword(context.Context, User) error
	DeleteUserByID(context.Context, int64) error
	Ping(ctx context.Context) error
	SignIn(context.Context, string, string) error
	ChangePassword(ctx context.Context, username string, oldPassword, newPassword string) error
	ForgotPassword(ctx context.Context, email string) error
}

// StoreImpl  is the blueprint for the users logic
type StoreImpl struct {
	Store Service
}

// NewService creates a new service
func NewService(store Service) StoreImpl {
	return StoreImpl{
		Store: store,
	}
}

func (u *StoreImpl) CreateUser(ctx context.Context, user User) error {
	if err := u.Store.CreateUser(ctx, user); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error creating user")
		return err
	}
	return nil
}

func (u *StoreImpl) GetUserByID(ctx context.Context, id int64) (User, error) {
	user, err := u.Store.GetUserByID(ctx, id)
	if err != nil {
		log.WithFields(logrus.Fields{
			"id":    id,
			"error": err,
		}).Error("Error fetching user by ID")
		return user, err
	}
	return user, nil
}

func (u *StoreImpl) UpdateUserByID(ctx context.Context, user User, id int64) error {
	if err := u.Store.UpdateUserByID(ctx, user, id); err != nil {
		log.WithFields(logrus.Fields{
			"user":  user,
			"error": err,
		}).Error("Error updating user")
		return err
	}

	return nil
}

func (u *StoreImpl) SetActivity(ctx context.Context, id int64, action bool) error {
	if err := u.Store.SetActivity(ctx, id, action); err != nil {
		log.WithFields(logrus.Fields{
			"id":    id,
			"error": err,
		}).Error("Error deactivating user by ID")
		return err
	}

	return nil
}

func (u *StoreImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := u.Store.GetByEmail(ctx, email)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error creating user")
		return nil, err
	}

	return user, nil
}

func (u *StoreImpl) GetByUsername(ctx context.Context, username string) (*User, error) {
	user, err := u.Store.GetByUsername(ctx, username)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error creating user")
		return nil, err
	}

	return user, nil
}

func (u *StoreImpl) GetUserByFullName(ctx context.Context, fullName string) (*User, error) {
	// Implement the logic to get a user by full name using the Store service
	user, err := u.Store.GetUserByFullName(ctx, fullName)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error creating user")
		return nil, err
	}
	return user, nil
}

func (u *StoreImpl) UpdateUserRoleID(ctx context.Context, id uint, newRoleID int64) error {
	if err := u.Store.UpdateUserRoleID(ctx, id, newRoleID); err != nil {
		log.WithFields(logrus.Fields{
			"id":        id,
			"newRoleID": newRoleID,
			"error":     err,
		}).Error("Error updating user role ID")
		return err
	}
	return nil
}

func (u *StoreImpl) ResetPassword(ctx context.Context, newUser User) error {
	if err := u.Store.ResetPassword(ctx, newUser); err != nil {
		log.WithFields(logrus.Fields{
			"user":  newUser,
			"error": err,
		}).Error("Error resetting password")
		return err
	}
	return nil
}

func (u *StoreImpl) DeleteUserByID(ctx context.Context, id int64) error {
	if err := u.Store.DeleteUserByID(ctx, id); err != nil {
		log.WithFields(logrus.Fields{
			"id":    id,
			"error": err,
		}).Error("Error deleting user by ID")
		return err
	}
	return nil

}

// Ping - pings the database to check if it is alive
func (u *StoreImpl) Ping(ctx context.Context) error {
	if err := u.Store.Ping(ctx); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error pinging database")
		return err
	}
	return nil
}

// SignIn - signs in a user
func (u *StoreImpl) SignIn(ctx context.Context, username, password string) error {
	if err := u.Store.SignIn(ctx, username, password); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error signing in user")
		return err
	}
	return nil
}

// ChangePassword changes the password
func (u *StoreImpl) ChangePassword(ctx context.Context, username, oldPassword, newPassword string) error {
	if err := u.Store.ChangePassword(ctx, username, oldPassword, newPassword); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error changing password")
		return err

	}
	return nil
}

// ForgotPassword - sends a password reset link to the user's email
func (u *StoreImpl) ForgotPassword(ctx context.Context, email string) error {
	if err := u.Store.ForgotPassword(ctx, email); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error sending password reset link")
		return err
	}
	return nil
}
