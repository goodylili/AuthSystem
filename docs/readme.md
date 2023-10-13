# Getting Started with OurPlatform API

In this guide, we'll walk you through the initial steps to get the **OurPlatform** API up and running on your local
environment.

## Table of Contents

1. [Health Check (Alive)](#health-check-alive)
2. [User Registration](#user-registration)
3. [Authentication](#authentication)
4. [API Endpoints Overview](#api-endpoints-overview)
5. [Further Reading](#further-reading)

---

### <a name="health-check-alive"></a>**1. Health Check (Alive)**

Once you have the server running, you can check its health using the following endpoint:

- **Endpoint**: `/alive`
- **HTTP Method**: `GET`
- **Description**: Checks if the server is alive and responding.

**Responses**:

- `200 OK`: Server is alive.

---

### <a name="user-registration"></a>**2. User Registration**

Begin by creating a new user account:

- **Endpoint**: `/api/v1/users/`
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

- `201 Created`: Successfully created a user.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="api-endpoints-overview"></a>**4. API Endpoints Overview**

For a deep dive into each category of endpoints, refer to the detailed documentation:

- [Users](./users.md)
- [Authentication](./auth.md)
- [Types](./types.md)

---

### <a name="further-reading"></a>**5. Further Reading**

As you delve into **OurPlatform** API, we recommend familiarizing yourself with the main [README.md](../README.md) for a
holistic understanding of the system, its architecture, and functionalities. For any issues or contributions, please
follow the guidelines specified in the main README.

