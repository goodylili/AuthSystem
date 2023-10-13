## **Types Documentation**

## **Overview**

This document outlines the distinct data types utilized within our system, providing a clear understanding of the data structure and the constraintegers imposed on various data fields. These data types form the backbone of our system, ensuring data consistency and structured integereraction across different functionalities. Whether you're a developer, an administrator, or an end-user, comprehending these data types will provide a clearer insight integero how data is handled within the system.


## **Index**

- [Overview](#overview)
- [User Type](#user-type)
- [Permissions](#permissions)
- [Roles and Permissions Mapping](#Roles-and-Permissions-Mapping)
    - [Role Constants](#role-constants)
    - [Role Permissions Map](#role-permissions-map)
---

### <a name="user-type"></a>1. User Type


The User type encapsulates the essential information related to an individual user within the system. This includes their personal details, contact information, and account status. It's crucial to ensure that each field adheres to the defined constraintegers to maintegerain data integeregrity and user privacy.

| Field      | Type   | Description                             | Constraintegers                           |
|------------|--------|-----------------------------------------|---------------------------------------|
| `username` | string | Username assigned to the user.          | Required                              |
| `email`    | string | Email address linked to the user.       | Required, Valid Email Format          |
| `password` | string | User's hashed password (write-only).    | Required, At least 8 characters       |
| `is_active`| bool   | Indicates if the user is active.        | Required, true or false               |
| `first_name`| string | First name of the user.                | Required                              |
| `last_name`| string | Last name of the user.                  | Required                              |
| `age`      | Integer  | Age of the user.                        | Required, Greater than or equal to 18 |
| `phone`    | Integer  | Phone number of the user.               | Required, Starts with 0, Length = 11  |
| `role_id`  | Integer  | Role ID associated with the user.       | Required, Less than or equal to 3     |

---

### <a name="permissions"></a>2. Permissions


Permissions are predefined constants that represent the different levels of access a user can have within the system. They are crucial for implementing role-based access control, ensuring that users have the appropriate level of access to perform various actions within the system.

| Constant           | Value | Description           |
|--------------------|-------|-----------------------|
| `CanCreateAccount` | 0     | Can create an account |
| `CanUpdateDetails` | 1     | Can update details    |
| `CanViewUsers`     | 2     | Can view users        |
| `CanDeactivateUsers`| 3    | Can deactivate users  |
| `CanGetUsersByFullName`| 4 | Can get users by full name |

---

### <a name="roles-and-permissions-mapping"></a>3. Roles and Permissions Mapping

Roles and Permissions Mapping is a system by which different roles are associated with specific permissions. This mapping is crucial for managing access control within the system, ensuring that each user has the right permissions based on their role. The following tables provide a detailed breakdown of the role constants and the permissions mapped to each role.

#### <a name="role-constants"></a>Role Constants


| Constant      | Value | Description   |
|---------------|-------|---------------|
| `RoleBasicUser`| 0    | Basic User    |
| `RoleUser`    | 1     | User          |
| `RoleAdmin`   | 2     | Admin         |

#### <a name="role-permissions-map"></a>Role Permissions Map

| Role         | Permissions |
|--------------|-------------|
| `RoleBasicUser` | `CanCreateAccount` |
| `RoleUser`   | `CanCreateAccount`, `CanUpdateDetails` |
| `RoleAdmin`  | `CanCreateAccount`, `CanUpdateDetails`, `CanViewUsers`, `CanDeactivateUsers`, `CanGetUsersByFullName` |

The `RolePermissionsMap` is a mapping from roles to their associated permissions, defining what actions each role can perform within the system.