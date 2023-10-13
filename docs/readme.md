# Getting Started with OurPlatform API

In this guide, we'll walk you through the initial steps to get the **OurPlatform** API up and running on your local
environment.

## Table of Contents

1. [Health Check (Alive)](#health-check-alive)
2. [API Endpoints Overview](#api-endpoints-overview)
3. [Further Reading](#further-reading)

---

### <a name="health-check-ping"></a>**1. Health Check (Ping)**

Once you have the server running, you can check the health of the service by using the following endpoint:

- **Endpoint**: `/alive`
- **HTTP Method**: `GET`
- **Description**: Checks the health of the service.

**Responses**:

- `200 OK`: Service is healthy.
- `500 Internal Server Error`: Service is down or facing issues.

---

### <a name="database-health-check"></a>**2. Database Health Check**

To ensure your database connectivity is intact:

- **Endpoint**: `/ready`
- **HTTP Method**: `GET`
- **Description**: Checks if the database is alive and responding.

| Parameter | Type | Description | Required |
|-----------|------|-------------|----------|
| -         | -    | -           | -        |

**Responses**:

- `200 OK`: Database is alive and responding.
- `500 Internal Server Error`: Database connection issues or the service is down.
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

