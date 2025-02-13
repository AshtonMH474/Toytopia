package routes

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
)

var (
	bucket = "toytopiaimages"
	region = "us-east-2"
)

type ToyImageSerial struct {
	// not model Toy, see this as serialzer
	ID         uint            `json:"id"`
	PrimaryImg bool            `json:"primary_img"`
	ImgUrl     string          `json:"img_url"`
	Toy        ToySerialNoUser `json:"toy"`
}
type NoToy struct {
	ID         uint   `json:"id"`
	PrimaryImg bool   `json:"primary_img"`
	ImgUrl     string `json:"img_url"`
}

func CreateResImage(image models.ToyImage, toy ToySerialNoUser) ToyImageSerial {
	return ToyImageSerial{ID: image.ID, PrimaryImg: image.PrimaryImg, ImgUrl: image.ImgUrl, Toy: toy}
}
func CreateNoToyImage(image models.ToyImage) NoToy {
	return NoToy{ID: image.ID, PrimaryImg: image.PrimaryImg, ImgUrl: image.ImgUrl}
}
func FindImagesByToyId(id int, images *[]models.ToyImage) error {
	query := database.Database.Db.Model(&models.ToyImage{})
	query = query.Where("toy_id = ?", id)
	if err := query.Find(&images).Error; err != nil {
		return errors.New("images do not exist")
	}
	return nil

}

func initS3Session() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("SECRET_AWS_KEY"), ""),
	})
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}
	return s3.New(sess)

}

func CreateToyImage(c *fiber.Ctx) error {
	// checks for authorziation
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
	var user models.User

	// finds the user based off token id
	if err := findUser(int(tokenUserID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// checks to see if id is an int
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Please ensure the ID is an integer",
		})
	}
	// finds toy based off int
	var toy models.Toy
	if err := findToy(id, &toy); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// if toyid and userid are not same cant create img
	if toy.UserId != int(user.ID) {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unathorized",
		})
	}
	// parses body
	var image models.ToyImage
	if err := c.BodyParser(&image); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// toyId becomes toy id
	image.ToyId = int(id)

	// validates info
	validationErrors := make(map[string]string)
	if len(image.ImgUrl) < 1 {
		validationErrors["imgUrl"] = "Img url required"
	}

	if len(validationErrors) > 0 {
		return c.Status(400).JSON(fiber.Map{"errors": validationErrors})
	}

	// grabs image
	resp, err := http.Get(image.ImgUrl)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch image", "error": err.Error()})
	}
	defer resp.Body.Close()
	// Read the image content into memory
	imageData, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to read image content", "error": err.Error()})
	}

	// Create a buffer to store the image content
	imageBuffer := bytes.NewReader(imageData)

	// Generate a unique file name
	fileName := fmt.Sprintf("image_%d.jpg", time.Now().Unix())

	// Initialize AWS S3 client
	s3Client := initS3Session()

	// Upload the image to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   imageBuffer,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to upload image", "error": err.Error()})
	}

	// Construct the S3 URL for the uploaded image
	imageURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, fileName)
	// the image new url is this one
	image.ImgUrl = imageURL
	// creates image and returns it
	database.Database.Db.Create(&image)
	resToy := NoUserResToy(toy)
	resImage := CreateResImage(image, resToy)
	return c.Status(201).JSON(resImage)

}
