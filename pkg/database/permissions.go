package database

// Permission constants
const (
	CanCreateAccount = iota
	CanUpdateDetails
	CanViewUsers
	CanDeactivateUsers
	CanGetUsersByFullName
)

// RolePermissionsMap maps roles to their associated permissions.
var RolePermissionsMap = map[int][]int{
	RoleBasicUser: {
		CanCreateAccount,
	},
	RoleUser: {
		CanCreateAccount,
		CanUpdateDetails,
	},
	RoleAdmin: {
		CanCreateAccount,
		CanUpdateDetails,
		CanViewUsers,
		CanDeactivateUsers,
		CanGetUsersByFullName,
	},
}
