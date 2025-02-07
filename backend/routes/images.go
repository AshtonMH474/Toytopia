package routes

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

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
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is missing",
		})
	}
}

func UploadImageFromURL(c *fiber.Ctx) error {
	// Get the image URL from the request body (JSON)
	var body struct {
		ImageURL string `json:"image_url"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Failed to parse request body", "error": err.Error()})
	}

	// Fetch the image from the URL
	resp, err := http.Get(body.ImageURL)
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

	// Return the image URL in the response
	return c.Status(201).JSON(fiber.Map{
		"message": "Image uploaded successfully",
		"url":     imageURL,
	})
}
