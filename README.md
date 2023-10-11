# AuthSystem 🥷🏽🔐 (WIP)

<p align="center">
   <a href="http://makeapullrequest.com"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat" alt=""></a>
   <a href="https://github.com/goodnessuc/authsystem/actions"><img src="https://github.com/goodnessuc/authsystem/actions/workflows/test.yml/badge.svg" alt="Github Actions"></a>
   <a href="https://golang.org"><img src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg" alt="made-with-Go"></a>
   <a href="https://goreportcard.com/report/github.com/goodnessuc/authsystem"><img src="https://goreportcard.com/badge/github.com/goodnessuc/authsystem" alt="GoReportCard"></a>
   <a href="https://github.com/goodnessuc/authsystem"><img src="https://img.shields.io/github/go-mod/go-version/goodnessuc/authsystem.svg" alt="Go.mod version"></a>
   <a href="https://github.com/Goodnessuc/AuthSystem/blob/main/LICENSE"><img src="https://img.shields.io/github/license/goodnessuc/authsystem.svg" alt="LICENCE"></a>
</p>

A Complete Backend Service with User Authentication and Authorization 

## Description
This repository contains a backend service developed in GoLang that handles user authentication and authorization. 

It includes functionalities like user registration, login, password reset, social authentication (e.g., GitHub, Google, Twitter), and JWT authentication. 

The service is designed with security in mind, implementing measures to prevent common attacks such as SQL injection, CSRF attacks, and brute force attacks.


## Features
- [x] User Management
- [x] User Authentication
    - [x] User Registration: Allows new users to create an account.
    - [x] User Login: Allows existing users to log in.
- [x] Session Management
    - [x] Create Session: Creates a session when a user logs in.
    - [x] Destroy Session: Destroys the session when a user logs out.
- [x] Password Management
    - [x] Password Reset: Allows users to reset their password.
    - [x] Password Hashing: Hashes passwords before storing them in the database.
- [ ] Social Authentication: Allows users to log in using their social media accounts.
  - [ ] GitHub Authentication: Allows users to log in through GitHub
  - [ ] Google Authentication: Allows users to log in through Google
  - [ ] Personal Authentication: Implementing Personal OAuth in Go for the app
- [x] JWT Authentication: Uses JSON Web Tokens (JWT) for secure information transmission.
- [ ] Security Measures
    - [x] SQL Injection Prevention: Prevents SQL injection attacks.
    - [x] CSRF Prevention: Prevents Cross-Site Request Forgery (CSRF) attacks.
    - [ ] Slowloris Prevention: Prevents Slowloris Attacks


## Contributions
Contributions, issues, and feature requests are welcome!

## License
MIT

