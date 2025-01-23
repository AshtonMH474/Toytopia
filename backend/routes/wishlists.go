package routes

import (
	"errors"

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

func findWishlist(id int, wishlist *models.Wishlist) error {
	database.Database.Db.Preload("Toys").Find(&wishlist, "id = ?", id)
	if wishlist.ID == 0 {
		return errors.New("wishlist does not exist")
	}
	return nil
}
