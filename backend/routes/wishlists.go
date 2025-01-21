package routes

import (
	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
	"github.com/gofiber/fiber/v2"
)

type WishlistSerial struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description"`
	User        UserSerial `json:"user"`
}

func CreateResWishlist(wishlist models.Wishlist, user UserSerial) WishlistSerial {
	return WishlistSerial{ID: wishlist.ID, Name: wishlist.Name, Description: wishlist.Description, User: user}
}

func AllWishlists(c *fiber.Ctx) error {
	var wishlists []models.Wishlist
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
	query = query.Where("user_id = ?", tokenUserID)

	if err := query.Find(&wishlists).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error fetching wishlists", "message": err.Error()})
	}
	// need to include User
	return c.Status(200).JSON(wishlists)
}
