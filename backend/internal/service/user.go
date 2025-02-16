package service

import (
	"fmt"

	"library-service/internal/constant"
	"library-service/internal/model"
	"library-service/internal/port"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo port.UserRepository
	jwt      port.JWT
}

func NewUserService(userRepo port.UserRepository, jwt port.JWT) *UserService {
	return &UserService{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

// CreateUser creates a new user in the system.
// It first checks if the username already exists in the system.
// If the username does not exist, it will hash the password and create a new user.
// If the username already exists, it will return an error.
func (u *UserService) CreateUser(user *model.User) error {
	if user.Role != constant.MemberRole {
		return fmt.Errorf("invalid role")
	}

	// Check if the username already exists in the system
	usernameExist, err := u.userRepo.HasUsername(user.Username)
	if err != nil {
		return err
	}

	if usernameExist {
		return fmt.Errorf("username already exist")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the hashed password to the user struct
	user.Password = string(hashedPassword)

	// Create the new user
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

// CreateLibrarian creates a new librarian in the system.
// It first checks if the username already exists in the system.
// If the username does not exist, it will hash the password and create a new librarian.
// If the username already exists, it will return an error.
func (u *UserService) CreateLibrarian(user *model.User) error {
	if user.Role != constant.Librarian {
		return fmt.Errorf("invalid role")
	}

	// Check if the username already exists in the system
	usernameExist, err := u.userRepo.HasUsername(user.Username)
	if err != nil {
		return err
	}

	if usernameExist {
		return fmt.Errorf("username already exist")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Set the hashed password to the user struct
	user.Password = string(hashedPassword)

	// Create the new user
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

// Authentication performs the authentication process for a user.
// It first checks if the username and password matches the one in the database.
// If the username and password matches, it will generate a JWT token for the user.
// If the username and password does not match, or if there is an error while
// retrieving the user from the database, it will return an error.
func (u *UserService) Authentication(username, password string) (string, string, error) {
	// Retrieve the user from the database based on the username
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", "", err
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			return "", "", fmt.Errorf("invalid credential")
		}
		return "", "", err
	}

	// Generate a JWT token for the user
	token := u.jwt.Generate(user.Username, user.Role)

	return token, user.Role, nil
}

func (u *UserService) GetAllMember() ([]*model.User, error) {
	users, err := u.userRepo.GetAllMember()
	if err != nil {
		return nil, err
	}

	return users, nil
}
