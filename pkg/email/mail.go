package email

// SendWelcomeMessage sends a welcome message to a new user
func SendWelcomeMessage(username, url, recipient string) error {
	message := Message{
		Header:       "My App Africa",
		UsersName:    username,
		Introduction: "Hey" + username + ", Welcome to My App Africa",
		Content:      "Thank you for signing up recipient My App Africa. Please click the link below recipient activate your account.",
		URL:          url,
		Action:       "Activate Account",
	}
	body, err := LoadEmail(message)
	if err != nil {
		return err
	}

	err = SendEmail(recipient, "Welcome to My App Africa", body)
	if err != nil {
		return err
	}

	return nil
}

// SendPassWordResetMessage sends a password reset message to a user
func SendPassWordResetMessage(username, url, recipient string) error {
	message := Message{
		Header:       "My App Africa",
		UsersName:    username,
		Introduction: "Hey" + username,
		Content:      "You have requested to reset your password. Please click the link below to reset your password.",
		URL:          url,
		Action:       "Reset Password",
	}
	body, err := LoadEmail(message)
	if err != nil {
		return err
	}

	err = SendEmail(recipient, "Welcome to My App Africa", body)
	if err != nil {
		return err
	}
	return nil
}
