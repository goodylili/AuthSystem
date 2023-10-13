# **Users API Documentation**

## **Overview**

The Users API allows for the creation, retrieval, updating, and management of user accounts within your platform, aligning with a robust authentication and authorization system to ensure secure access to resources.

## **Index**

- **[Endpointegers](#endpointegers)**
    - [Create User](#create-user)
    - [Retrieve User by ID](#retrieve-user-by-id)
    - [Update User](#update-user)
    - [Set User Activity Status](#set-user-activity-status)
    - [Retrieve User by Email](#retrieve-user-by-email)
    - [Retrieve User by Username](#retrieve-user-by-username)
    - [Retrieve User by Full Name](#retrieve-user-by-full-name)
    - [Change Password](#change-password)
    - [Reset Password](#reset-password)
    - [Forgot Password](#forgot-password)

### **Base URL**: `/users/api/v1`

---

### **Models**

### **The User Object**

| Field       | Type    | Description                          | Restrictions               |
|-------------|---------|--------------------------------------|----------------------------|
| `username`  | string  | Username assigned to the user.       | Unique                     |
| `email`     | string  | Email address linked to the user.    | Unique, Valid Email Format |
| `password`  | string  | User's hashed password (write-only). | At least 8 characters      |
| `is_active` | boolean | Indicates if the user is active.     | true or false              |

---

## **Endpointegers**:

### <a name="create-user"></a>**1. Create User**

- **Endpointeger**: `/create`
- **HTTP Method**: `POST`
- **Description**: Registers a new user with the provided details.

**Request Body**:

```json
{
  "username": "string",
  "email": "string",
  "password": "string",
  "is_active": boolean
}
```

**Responses**:

- `201 Created`: Successfully created a user. Returns the user data.
- `400 Bad Request`: Invalid input or malformed request.
- `500 integerernal Server Error`: Unexpected server error.

---

### <a name="retrieve-user-by-id"></a>**2. Retrieve User by ID**

- **Endpointeger**: `/{id}`
- **HTTP Method**: `GET`
- **Description**: Fetches details of a specific user using their unique ID.

| Parameter | Type | Description                   | Required |
|-----------|------|-------------------------------|----------|
| id        | integer  | Unique identifier of the user | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the user data.
- `400 Bad Request`: Invalid ID format.
- `404 Not Found`: User with the provided ID doesn't exist.

---

### <a name="update-user"></a>**3. Update User**

- **Endpointeger**: `/{id}/update`
- **HTTP Method**: `PUT`
- **Description**: Modifies the details of an existing user.

| Parameter | Type | Description                   | Required |
|-----------|------|-------------------------------|----------|
| id        | integer  | Unique identifier of the user | Yes      |

**Request Body**:

```json
{
  "username": "string",
  "email": "string",
  "password": "string",
  "is_active": boolean
}
```

**Responses**:

- `200 OK`: Successfully updated the user data.
- `400 Bad Request`: Invalid input or malformed request.
- `404 Not Found`: User with the provided ID doesn't exist.
- `500 integerernal Server Error`: Unexpected server error.

---

### <a name="set-user-activity-status"></a>**4. Set User Activity Status**

- **Endpointeger**: `/{id}/status`
- **HTTP Method**: `PUT`
- **Description**: Updates the status of a user (activate/deactivate).

| Parameter | Type | Description                    | Required |
|-----------|------|--------------------------------|----------|
| id        | integer  | Unique identifier of the user. | Yes      |

**Request Body**:

```json
{
  "is_active": boolean
}
```

**Responses**:

- `200 OK`: Successfully updated the user's status.
- `400 Bad Request`: Invalid input or malformed request.
- `404 Not Found`: User with the provided ID doesn't exist.
- `500 integerernal Server Error`: Unexpected server error.

---

### <a name="retrieve-user-by-email"></a>**5. Retrieve User by Email**

- **Endpointeger**: `/email/{email}`
- **HTTP Method**: `GET`
- **Description**: Fetches user details based on their email address.

| Parameter | Type   | Description                | Required |
|-----------|--------|----------------------------|----------|
| email     | string | Email address of the user. | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the user data.
- `400 Bad Request`: Invalid email format.
- `404 Not Found`: User with the provided email doesn't exist.
- `500 integerernal Server Error`: Unexpected server error.

---

### <a name="retrieve-user-by-username"></a>**6. Retrieve User by Username**

- **Endpointeger**: `/username/{username}`
- **HTTP Method**: `GET`
- **Description**: Fetches user details based on their username.

| Parameter | Type   | Description           | Required |
|-----------|--------|-----------------------|----------|
| username  | string | Username of the user. | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the user data.
- `404 Not Found`: User with the provided username doesn't exist.
- `500 integerernal Server Error`: Unexpected server error.

---

### <a name="retrieve-user-by-full-name"></a>**7. Retrieve User by Full Name**

- **Endpointeger**: `/fullname/{fullName}`
- **HTTP Method**: `GET`
- **Description**: Fetches user details based on their full name.

| Parameter | Type   | Description             | Required |
|-----------|--------|-------------------------|----------|
| fullName  | string | Full name of the user.  | Yes      |

**Responses**:

- `200 OK`: Successfully fetched the user data.
- `400 Bad Request`: Invalid name format.
- `404 Not Found`: User with the provided full name doesn't exist.
- `500 integerernal Server Error`: Unexpected server error.

---

### <a name="change-password"></a>**8. Change Password**

- **Endpointeger**: `/password/change`
- **HTTP Method**: `PUT`
- **Description**: Changes the user's password.

**Request Body**:

```json
{
   "username": "string",
   "old_password": "string",
   "new_password": "string"
}
```

**Responses**:

- `200 OK`: Password successfully changed.
- `400 Bad Request`: Invalid input or malformed request.
- `500 integerernal Server Error`: Unexpected server error.

---

### <a name="reset-password"></a>**9. Reset Password**

- **Endpointeger**: `/password/reset`
- **HTTP Method**: `PUT`
- **Description**: Resets the user's password.

**Request Body**:

```json
{
   "username": "string",
   "email": "string",
   "password": "new_password"
}
```

**Responses**:

- `200 OK`:

Password successfully reset.
- `400 Bad Request`: Invalid input or malformed request.
- `500 integerernal Server Error`: Unexpected server error.

---

### <a name="forgot-password"></a>**10. Forgot Password**

- **Endpointeger**: `/password/forgot`
- **HTTP Method**: `POST`
- **Description**: Handles the forgot password process.

| Parameter | Type   | Description           | Required |
|-----------|--------|-----------------------|----------|
| email     | string | Email of the user.    | Yes      |

**Responses**:

- `200 OK`: Password reset instructions sent.
- `400 Bad Request`: Invalid email format.
- `500 integerernal Server Error`: Unexpected server error.

---