package routes

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SafeUser struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

type UserSerial struct {
	// not model user, see this as serialzer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

func CreateResUser(user models.User) UserSerial {
	return UserSerial{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Username: user.Username}
}

func LoginHandler(c *fiber.Ctx) error {

	var user models.User

	// Parse the incoming JSON request body into the user struct
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	fmt.Println(user.Password, user.Email)

	if len(user.Password) < 8 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Password must be 8 characters or more",
		})
	}

	if user.Email == "" || !strings.Contains(user.Email, "@") {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invaild Email",
		})
	}

	// Find user by credential (username or email)
	foundUser, err := findUserByCredential(user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}
		// Handle other database errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	// Compare password with the hashed password stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		// If password does not match
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Set token cookie
	safeUser := SafeUser{
		ID:        foundUser.ID,
		Email:     foundUser.Email,
		Username:  foundUser.Username,
		FirstName: foundUser.FirstName,
		LastName:  foundUser.LastName,
	}

	_, err = SetTokenCookie(c, safeUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to set token cookie"})
	}

	// Return success response
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"user": safeUser,
	})

}

func SignupHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Parse request body
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
		}

		// Input validation
		validationErrors := validateSignupInput(user)
		if len(validationErrors) > 0 {
			return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"errors": validationErrors})
		}

		// Check if user with email or username already exists
		var existingUser models.User
		if err := database.Database.Db.Where("email = ? OR username = ?", user.Email, user.Username).First(&existingUser).Error; err == nil {
			errors := make(map[string]string)
			if strings.EqualFold(existingUser.Email, user.Email) {
				errors["email"] = "User with that email already exists"
			}
			if strings.EqualFold(existingUser.Username, user.Username) {
				errors["username"] = "User with that username already exists"
			}
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"message": "User already exists",
				"errors":  errors,
			})
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
		}

		// Create the user
		user.Password = string(hashedPassword)
		if err := database.Database.Db.Create(&user).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create user"})
		}

		// Set token cookie
		safeUser := SafeUser{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		_, err = SetTokenCookie(c, safeUser)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to set token cookie"})
		}

		// Return success response
		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"user": safeUser,
		})
	}
}

func Logout(c *fiber.Ctx) error {
	// Clear the cookie
	c.ClearCookie("token")

	// Return a JSON response
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func GetUser(c *fiber.Ctx) error {

	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.JSON(fiber.Map{"user": nil})
	}

	if ok && userData != nil {
		// Safe user data
		var user models.User
		id := int(userData["id"].(float64))
		if err := findUser(id, &user); err != nil {
			return c.Status(400).JSON(err.Error())
		}

		safeUser := SafeUser{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		return c.JSON(fiber.Map{
			"user": safeUser,
		})
	}

	// If no user found, return null user
	return c.JSON(fiber.Map{
		"user": nil,
	})
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	// Extract the token from the Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}

	// Extract user data from the token
	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}

	// Get the user ID from the token (you could also extract other data like username if needed)
	tokenUserID := uint(userData["id"].(float64))

	// Parse the user ID from the URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please ensure the ID is an integer",
		})
	}

	// Check if the token's user is the same as the one being updated
	if id != int(tokenUserID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You are not authorized to update this user",
		})
	}

	var user models.User

	if err != nil {
		return c.Status(400).JSON("please ensure Id is int")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// new struct for new Updated Data
	type UpdatedUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	// checks if theres data
	var data UpdatedUser
	if err := c.BodyParser(&data); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	// checks if firstName and lastName are there
	if len(data.FirstName) > 0 {
		user.FirstName = data.FirstName
	}
	if len(data.LastName) > 0 {
		user.LastName = data.LastName
	}

	// saves user data
	database.Database.Db.Save(&user)

	resUser := CreateResUser(user)
	return c.Status(200).JSON(resUser)
}

func DeleteUser(c *fiber.Ctx) error {
	// Extract the token from the Authorization header
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}

	// Extract user data from the token
	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}

	// Get the user ID from the token (you could also extract other data like username if needed)
	tokenUserID := uint(userData["id"].(float64))

	// Parse the user ID from the URL parameter
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please ensure the ID is an integer",
		})
	}

	// Check if the token's user is the same as the one being updated
	if id != int(tokenUserID) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "You are not authorized to delete this user",
		})
	}

	var user models.User

	if err != nil {
		return c.Status(400).JSON("please ensure Id is int")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(resMessage("Successfully Deleted"))
}

type ResMessage struct {
	Message string
}

func resMessage(str string) ResMessage {
	return ResMessage{Message: str}
}

func validateSignupInput(user models.User) map[string]string {
	errors := make(map[string]string)

	if user.Email == "" || !strings.Contains(user.Email, "@") {
		errors["email"] = "Invalid email"
	}
	if len(user.Username) < 4 {
		errors["username"] = "Username is required and must be at least 4 characters"
	}
	if strings.Contains(user.Username, "@") {
		errors["username"] = "Username cannot be an email"
	}
	if user.FirstName == "" {
		errors["firstName"] = "First Name is required"
	}
	if user.LastName == "" {
		errors["lastName"] = "Last Name is required"
	}
	if len(user.Password) < 8 { // Assuming `models.User` has a `Password` field
		errors["password"] = "Password must be 8 characters or more"
	}

	return errors
}

func SetTokenCookie(ctx *fiber.Ctx, user SafeUser) (string, error) {
	var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Secret key from environment variable
	jwtExpiresInStr := os.Getenv("JWT_EXPIRES_IN")  // e.g., "604800" for 1 week
	jwtExpiresIn, err := strconv.Atoi(jwtExpiresInStr)
	if err != nil {
		jwtExpiresIn = 604800 // Default to 1 week (in seconds) if parsing fails
	}
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": SafeUser{
			ID:        user.ID,
			Email:     user.Email,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		"exp": time.Now().Add(time.Duration(jwtExpiresIn) * time.Second).Unix(),
	})

	// Sign the token
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Determine if it's a production environment
	isProduction := os.Getenv("NODE_ENV") == "production"

	// Set the token as a cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    signedToken,
		MaxAge:   jwtExpiresIn,
		HTTPOnly: true,
		Secure:   isProduction,
		SameSite: func() string {
			if isProduction {
				return "Lax"
			}
			return "None"
		}(),
	})

	return signedToken, nil
}

func findUserByCredential(cred string) (*models.User, error) {
	var user models.User
	// Query the database to find a user by username or email
	err := database.Database.Db.Where("username = ? OR email = ?", cred, cred).First(&user).Error
	if err != nil {
		// If an error occurs, return nil for user and the error
		return nil, err
	}
	// If no error, return the user and nil for the error
	return &user, nil
}

// helper function to extract data from user
func extractUserDataFromToken(c *fiber.Ctx) (map[string]interface{}, bool) {
	// Get the JWT secret from the environment variable
	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	tokenString := c.Cookies("token")

	// Check if the token is provided
	if tokenString == "" {
		return nil, false
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		// Return the secret for verification
		return jwtSecret, nil
	})

	// Handle errors with the token
	if err != nil || !token.Valid {
		return nil, false
	}

	// Extract the claims (user data) from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, false
	}

	// Get the user data from claims
	userData, ok := claims["data"].(map[string]interface{})
	if !ok || userData == nil {
		return nil, false
	}

	// Return the user data and success
	return userData, true
}
