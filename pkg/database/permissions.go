package database

// Role represents a role in the system.
type Role struct {
	RoleID          int    // Primary Key: Role ID
	RoleName        string // Unique: Role Name
	RoleDescription string // Role Description
}

// Permission represents a permission in the system.
type Permission struct {
	PermissionID   int    // Primary Key: Permission ID
	PermissionName string // Unique: Permission Name
	// Additional permission-related fields go here
}

// UserToRole represents the relationship between users and roles.
type UserToRole struct {
	UserID int // User ID
	RoleID int // Role ID
	// Primary Key: UserID + RoleID combined
}

// RoleToPermission represents the relationship between roles and permissions.
type RoleToPermission struct {
	RoleID       int // Role ID
	PermissionID int // Permission ID
	// Primary Key: RoleID + PermissionID combined
}
