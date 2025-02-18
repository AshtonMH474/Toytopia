package routes

import (
	"errors"
	"os"
	"strings"
	"time"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
	"github.com/gofiber/fiber/v2"
)

type ToySerial struct {
	ID          uint       `json:"id"`
	ReleaseDate time.Time  `json:"release_date"`
	Price       float64    `json:"price"`
	ProductType string     `json:"product_type"`
	Theme       string     `json:"theme"`
	Count       int        `json:"count"`
	Available   bool       `json:"available"`
	Rating      float64    `json:"rating"`
	User        UserSerial `json:"user"`
	Images      []NoToy    `json:"images"`
	Company     string     `json:"company"`
}
type ToySerialNoUser struct {
	// not model Toy, see this as serialzer
	ID          uint      `json:"id"`
	ReleaseDate time.Time `json:"release_date"`
	Price       float64   `json:"price"`
	ProductType string    `json:"product_type"`
	Theme       string    `json:"theme"`
	Count       int       `json:"count"`
	Available   bool      `json:"available"`
	Rating      float64   `json:"rating"`
	Company     string    `json:"company"`
}

func CreateResToyImages(toy models.Toy, user UserSerial, images []NoToy) ToySerial {
	return ToySerial{ID: toy.ID, ReleaseDate: toy.ReleaseDate, Price: toy.Price, ProductType: toy.ProductType, Theme: toy.Theme, Count: toy.Count, Available: toy.Available, Company: toy.Company, User: user, Images: images, Rating: toy.Rating}
}

func NoUserResToy(toy models.Toy) ToySerialNoUser {
	return ToySerialNoUser{ID: toy.ID, Rating: toy.Rating, ReleaseDate: toy.ReleaseDate, Price: toy.Price, ProductType: toy.ProductType, Theme: toy.Theme, Count: toy.Count, Available: toy.Available, Company: toy.Company}
}

func SearchToys(c *fiber.Ctx) error {
	var toys []models.Toy
	query := database.Database.Db.Model(&models.Toy{})

	// retriveing query params
	productType := c.Query("product_type")
	theme := c.Query("theme")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")
	companyParam := c.Query("company")

	// Apply filters based on the parameters
	enviroment := os.Getenv("NODE_ENV")
	isProduction := enviroment == "production"

	if isProduction {
		if theme != "" {
			query = query.Where("theme ILIKE ?", "%"+theme+"%")

		}
		if productType != "" {
			query = query.Where("product_type ILIKE ?", "%"+productType+"%")
		}
	} else {
		if productType != "" {
			query = query.Where("product_type LIKE ?", "%"+productType+"%")
		}
		if theme != "" {
			query = query.Where("theme LIKE ?", "%"+theme+"%")
		}
	}

	if minPrice != "" {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice != "" {
		query = query.Where("price <= ?", maxPrice)
	}
	if companyParam != "" {
		companies := strings.Split(companyParam, ",")

		// Trim whitespace from each company name
		for i, company := range companies {
			companies[i] = strings.TrimSpace(company)
		}

		// Apply the filter with the "IN" SQL operator for multiple companies
		query = query.Where("company IN (?)", companies)
	}

	// Filter for availability (true or false)
	available := c.Query("available")
	if available != "" {
		var availableBool bool
		if available == "true" {
			availableBool = true
		} else if available == "false" {
			availableBool = false
		} else {
			// Handle invalid input for available
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid value for 'available'. Use 'true' or 'false'.",
			})
		}
		query = query.Where("available = ?", availableBool)
	}

	if err := query.Find(&toys).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error fetching toys", "message": err.Error()})
	}
	var newToys []ToySerial
	for _, toy := range toys {
		var resImages []NoToy
		var images []models.ToyImage
		if err := FindImagesByToyId(int(toy.ID), &images); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		for _, image := range images {
			resImage := CreateNoToyImage(image)
			resImages = append(resImages, resImage)
		}
		var user models.User
		if err := findUser(int(toy.UserId), &user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		resUser := CreateResUser(user)
		resToy := CreateResToyImages(toy, resUser, resImages)
		newToys = append(newToys, resToy)
	}

	// Return the results as a JSON response
	return c.Status(200).JSON(newToys)
}
func CreateToy(c *fiber.Ctx) error {
	// seeing if token
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}
	// checking for the user data in token
	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}

	// id of user in token
	tokenUserID := uint(userData["id"].(float64))
	var user models.User

	// finds the user based off token id
	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var toy models.Toy
	if err := c.BodyParser(&toy); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// validates info from req
	validationErrors := make(map[string]string)
	if len(toy.ProductType) < 1 {
		validationErrors["productType"] = "Product Type needed"
	}
	if toy.ReleaseDate.IsZero() {
		validationErrors["releaseDate"] = "Release Date needed"
	}
	if len(toy.Theme) < 1 {
		validationErrors["theme"] = "Theme needed"
	}

	if len(validationErrors) > 0 {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
	}

	toy.UserId = int(user.ID)

	database.Database.Db.Create(&toy)
	var images []NoToy
	resUser := CreateResUser(user)
	resToy := CreateResToyImages(toy, resUser, images)
	return c.Status(201).JSON(resToy)
}

func UpdateToy(c *fiber.Ctx) error {
	// seeing if token
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}

	// seeing if user data in token
	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}

	// id of user from token
	tokenUserID := uint(userData["id"].(float64))
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please ensure the ID is an integer",
		})
	}
	var toy models.Toy
	var user models.User
	// finds toy
	if err := findToy(id, &toy); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// if token id and toy userId r not same
	if toy.UserId != int(tokenUserID) || toy.UserId == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Unathorzied"})
	}

	type UpdatedToy struct {
		ReleaseDate time.Time `json:"release_date"`
		Price       float64   `json:"price"`
		ProductType string    `json:"product_type"`
		Theme       string    `json:"theme"`
		Count       *int      `json:"count"`
		Available   bool      `json:"available"`
	}

	var data UpdatedToy
	if err := c.BodyParser(&data); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	// sees what needs to be updated
	if len(data.ProductType) > 0 {
		toy.ProductType = data.ProductType
	}
	if len(data.Theme) > 0 {
		toy.Theme = data.Theme
	}
	if data.Price != 0 {
		toy.Price = data.Price
	}
	if !data.ReleaseDate.IsZero() {
		toy.ReleaseDate = data.ReleaseDate
	}
	if data.Count != nil && *data.Count != toy.Count {
		toy.Count = *data.Count
	}

	// if count is above 0 its avaivale true
	if toy.Count <= 0 {
		toy.Available = false
	} else {
		toy.Available = true
	}

	database.Database.Db.Save(&toy)
	var resImages []NoToy
	var images []models.ToyImage
	if err := FindImagesByToyId(int(toy.ID), &images); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	for _, image := range images {
		resImage := CreateNoToyImage(image)
		resImages = append(resImages, resImage)
	}

	resUser := CreateResUser(user)
	resToy := CreateResToyImages(toy, resUser, resImages)

	return c.Status(200).JSON(resToy)

}

func DeleteToy(c *fiber.Ctx) error {
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

	// Get the user ID from the token
	tokenUserID := uint(userData["id"].(float64))

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please ensure the ID is an integer",
		})
	}

	var toy models.Toy
	var user models.User
	if err := findToy(id, &toy); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if toy.UserId != int(tokenUserID) || toy.UserId == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Unathorzied"})
	}
	var images []models.ToyImage
	if err := FindImagesByToyId(int(toy.ID), &images); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	for _, image := range images {
		if err := database.Database.Db.Delete(&image).Error; err != nil {
			return c.Status(404).JSON(err.Error())
		}
	}

	if err := database.Database.Db.Delete(&toy).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(resMessage("Successfully Deleted"))
}

func findToy(id int, toy *models.Toy) error {
	database.Database.Db.Find(&toy, "id = ?", id)
	if toy.ID == 0 {
		return errors.New("toy does not exist")
	}
	return nil
}
