package credStorageErrors

import "fmt"

type UserAlreadyExistsError struct {
	User string
}

type InvalidUserError struct {
	User string
}

type InvalidTokenError struct {
}

func (e InvalidTokenError) Error() string {
	return "Invalid token"
}

func (e UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("cannot save user %s, user already exists", e.User)
}

func (e InvalidUserError) Error() string {
	return fmt.Sprintf("invalid user %s ", e.User)
}
