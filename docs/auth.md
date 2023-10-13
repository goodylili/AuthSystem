## **Authentication Documentation**

## **Overview**

The authentication system within our platform is designed to provide secure access control, ensuring that only
authorized users can access certain parts of the system.

The system implements multiple authentication strategies including JWT (JSON Web Tokens) for session management, and
OAuth 2.0 for social login through providers like Google and GitHub.

This document provides an overview of the authentication mechanisms and outlines the workflow involved in authenticating
users.

## **Table of Contents**

- [Overview](#overview)
- [JWT Authentication](#jwt-authentication)
- [OAuth 2.0 Authentication](#oauth-authentication)
    - [GitHub OAuth](#github-oauth)
    - [Google OAuth](#google-oauth)

---

### <a name="jwt-authentication"></a>**1. JWT Authentication**

JWT (JSON Web Tokens) is a compact, URL-safe means of representing claims to be transferred between two parties. In our
system, JWTs are used to authenticate requests to the API. A token is generated upon successful login and must be
included in the HTTP header for subsequent requests to the API.

#### Generating a JWT

- **Endpoint**: `/auth/token`
- **HTTP Method**: `POST`
- **Description**: Generates a JWT for authenticated sessions.

**Request Body**:

```json
{
  "username": "string",
  "password": "string"
}
```

**Responses**:

- `200 OK`: Successfully generated JWT. Returns the token.
- `400 Bad Request`: Invalid input or malformed request.
- `401 Unauthorized`: Invalid credentials.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="github-authentication"></a>**GitHub Authentication**

GitHub Authentication allows users to log in using their GitHub accounts, facilitating a seamless login process
especially for developers.

#### **Endpoints**

- **Login**
    - **Endpoint**: `/auth/github/login`
    - **HTTP Method**: `GET`
    - **Description**: Redirects the user to the GitHub login page.

- **Callback**
    - **Endpoint**: `/auth/github/callback`
    - **HTTP Method**: `GET`
    - **Description**: Handles the callback from GitHub to retrieve user information.

#### **Flow**

The process is initiated by redirecting the user to the GitHub login page. Upon successful login, GitHub redirects back
to our system with a temporary code via the callback endpoint. This code is then exchanged for a token which can be used
to fetch the user's profile information from GitHub.

#### **Responses**

- `200 OK`: Successfully retrieved user information.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---

### <a name="google-authentication"></a>**Google Authentication**

Google Authentication allows users to log in using their Google accounts, providing a familiar login interface and
reducing the registration friction.

#### **Endpoints**

- **Login**
    - **Endpoint**: `/auth/google/login`
    - **HTTP Method**: `GET`
    - **Description**: Redirects the user to the Google login page.

- **Callback**
    - **Endpoint**: `/auth/google/callback`
    - **HTTP Method**: `GET`
    - **Description**: Handles the callback from Google to retrieve user information.

#### **Flow**

Similar to GitHub Authentication, the process begins by redirecting the user to the Google login page. Post login,
Google redirects back to our system with a temporary code via the callback endpoint. This code is then exchanged for a
token, enabling the retrieval of the user's profile information from Google.

#### **Responses**

- `200 OK`: Successfully retrieved user information.
- `400 Bad Request`: Invalid input or malformed request.
- `500 Internal Server Error`: Unexpected server error.

---


