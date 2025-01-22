package routes

import (
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
