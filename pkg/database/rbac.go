package database

type Permission struct {
	CanDelete        Role
	CanCreateAccount Role
}

type Role struct {
	ID          uint
	Name        string
	Permissions []Permission
}
