package routes

import (
	"errors"
	"log"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
	"github.com/gofiber/fiber/v2"
)

type WishlistSerial struct {
	ID          uint              `json:"id"`
	Name        string            `json:"name" gorm:"not null"`
	Description string            `json:"description"`
	User        UserSerial        `json:"user"`
	Toys        []ToySerialNoUser `json:"toys"`
}

func CreateResWishlist(wishlist models.Wishlist, user UserSerial, toys []ToySerialNoUser) WishlistSerial {
	return WishlistSerial{ID: wishlist.ID, Name: wishlist.Name, Description: wishlist.Description, User: user, Toys: toys}
}

func AllWishlists(c *fiber.Ctx) error {
	var wishlists []models.Wishlist
	var resWishlists []WishlistSerial
	query := database.Database.Db.Model(&models.Wishlist{})
	// seeing if token
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}

	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}

	// id of user in token
	tokenUserID := uint(userData["id"].(float64))
	// checks wishlist by userId
	query = query.Where("user_id = ?", tokenUserID)

	var user models.User
	// finds the user based off token id
	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Fetch the wishlists for the user
	if err := query.Preload("Toys").Find(&wishlists).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Error fetching wishlists",
			"message": err.Error(),
		})
	}

	resUser := CreateResUser(user)
	// creates res for json
	for _, wishlist := range wishlists {
		var toys []ToySerialNoUser
		for _, toy := range wishlist.Toys {
			resToy := NoUserResToy(toy)
			toys = append(toys, resToy)
		}

		resWishlist := CreateResWishlist(wishlist, resUser, toys)

		resWishlists = append(resWishlists, resWishlist)
	}

	return c.Status(200).JSON(resWishlists)
}

func GetWishlist(c *fiber.Ctx) error {
	// seeing if token
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}

	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}
	tokenUserID := uint(userData["id"].(float64))
	// pulls wishlist Id
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("please ensure Id is int")
	}

	var wishlist models.Wishlist
	// finds wishlist
	if err := findWishlist(id, &wishlist); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	// if user id is not same as token
	if wishlist.UserId != int(tokenUserID) {
		return c.Status(403).JSON(fiber.Map{
			"error": "You are not authorized to update this wishlist",
		})
	}
	var user models.User
	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// creates res toys user and wishlist and returns in json
	var toys []ToySerialNoUser
	for _, toy := range wishlist.Toys {
		resToy := NoUserResToy(toy)
		toys = append(toys, resToy)
	}
	resUser := CreateResUser(user)
	resWishlist := CreateResWishlist(wishlist, resUser, toys)
	return c.Status(200).JSON(resWishlist)
}

func CreateWishlist(c *fiber.Ctx) error {
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

	// parses body of req into wishlists
	var wishlist models.Wishlist
	if err := c.BodyParser(&wishlist); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validates req body
	validationErrors := make(map[string]string)
	if len(wishlist.Name) < 1 {
		validationErrors["Name"] = "Name needed"
	}
	if len(validationErrors) > 0 {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
	}

	wishlist.UserId = int(user.ID)

	// creates wishlists
	database.Database.Db.Create(&wishlist)
	resUser := CreateResUser(user)
	// empty toys array for res func bcz new wishlists means no toys
	var toys []ToySerialNoUser
	resWishlist := CreateResWishlist(wishlist, resUser, toys)
	return c.Status(201).JSON(resWishlist)
}

func AddToy(c *fiber.Ctx) error {
	// seeing if token
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}

	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}
	tokenUserID := uint(userData["id"].(float64))
	var user models.User
	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("please ensure Id is int")
	}

	var wishlist models.Wishlist
	// finds wishlist
	if err := findWishlist(id, &wishlist); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	// if user id is not same as token
	if wishlist.UserId != int(user.ID) {
		return c.Status(403).JSON(fiber.Map{
			"error": "You are not authorized to update this wishlist",
		})
	}

	var toys_in_wishlists []models.ToysInWishlist
	if err := c.BodyParser(&toys_in_wishlists); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	for _, toy := range toys_in_wishlists {
		toy.WishlistId = int(wishlist.ID)
		if err := database.Database.Db.Create(&toy).Error; err != nil {
			return c.Status(403).JSON(fiber.Map{
				"error": "Failed To Add Toy",
			})
		}
	}

	// finds wishlist
	if err := findWishlist(id, &wishlist); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var toys []ToySerialNoUser
	for _, toy := range wishlist.Toys {
		resToy := NoUserResToy(toy)
		toys = append(toys, resToy)
	}
	resUser := CreateResUser(user)
	resWishlist := CreateResWishlist(wishlist, resUser, toys)
	return c.Status(201).JSON(resWishlist)
}

func RemoveToy(c *fiber.Ctx) error {
	// seeing if token
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}

	userData, ok := extractUserDataFromToken(c)
	if !ok || userData == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}
	tokenUserID := uint(userData["id"].(float64))
	var user models.User
	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	log.Printf("test1")
	wishlistId, err := c.ParamsInt("wishlistId")
	toyId, err := c.ParamsInt("toyId")

	if err != nil {
		return c.Status(400).JSON("please ensure Id is int")
	}

	var wishlist models.Wishlist
	// finds wishlist
	if err := findWishlist(wishlistId, &wishlist); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	log.Printf("test2")

	// if user id is not same as token
	if wishlist.UserId != int(user.ID) {
		return c.Status(403).JSON(fiber.Map{
			"error": "You are not authorized to update this wishlist",
		})
	}
	log.Printf("test3")

	if err := database.Database.Db.Where("wishlist_id = ? AND toy_id = ?", wishlistId, toyId).Delete(&models.ToysInWishlist{}).Error; err != nil {
		log.Printf("test4")
		return c.Status(400).JSON(fiber.Map{
			"error": "couldn't delete toy from wishlist",
		})
	}
	log.Printf("test5")
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully deleted",
	})
}

func UpdateWishlist(c *fiber.Ctx) error {
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
	var wishlist models.Wishlist
	if err := findWishlist(id, &wishlist); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// if token id and wishlist userId r not same
	if wishlist.UserId != int(tokenUserID) || wishlist.UserId == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Unathorzied"})
	}

	type UpdatedWishlist struct {
		Name        string `json:"name" gorm:"not null"`
		Description string `json:"description"`
	}

	var data UpdatedWishlist
	if err := c.BodyParser(&data); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	if len(data.Name) > 0 {
		wishlist.Name = data.Name
	}
	if len(data.Description) > 0 {
		wishlist.Description = data.Description
	}

	database.Database.Db.Save(&wishlist)
	var toys []ToySerialNoUser
	for _, toy := range wishlist.Toys {
		resToy := NoUserResToy(toy)
		toys = append(toys, resToy)
	}
	resUser := CreateResUser(user)
	resWishlist := CreateResWishlist(wishlist, resUser, toys)
	return c.Status(200).JSON(resWishlist)
}

func DeleteWishlist(c *fiber.Ctx) error {
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
	var wishlist models.Wishlist
	if err := findWishlist(id, &wishlist); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// if token id and wishlist userId r not same
	if wishlist.UserId != int(tokenUserID) || wishlist.UserId == 0 {
		return c.Status(401).JSON(fiber.Map{"message": "Unathorzied"})
	}

	if err := database.Database.Db.Delete(&wishlist).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON(resMessage("Successfully Deleted"))
}

func findWishlist(id int, wishlist *models.Wishlist) error {
	database.Database.Db.Preload("Toys").Find(&wishlist, "id = ?", id)
	if wishlist.ID == 0 {
		return errors.New("wishlist does not exist")
	}
	return nil
}
