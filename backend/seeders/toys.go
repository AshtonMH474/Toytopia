package seeders

import (
	"log"
	"time"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
)

func SeedToys() {
	toys := []models.Toy{
		{ReleaseDate: time.Date(1984, time.January, 18, 0, 0, 0, 0, time.UTC), Price: 1000, ProductType: "Optimus Prime g1 Orginial", Theme: "Transformers", Count: 2, Available: true, UserId: 1, Company: "Hasbro", Rating: 5.0},
		{ReleaseDate: time.Date(1995, time.December, 15, 0, 0, 0, 0, time.UTC), Price: 250.75, ProductType: "Buzz Lightyear Action Figure", Theme: "Toy Story", Count: 5, Available: true, UserId: 1, Company: "Disney", Rating: 4.5},
		{ReleaseDate: time.Date(2001, time.November, 3, 0, 0, 0, 0, time.UTC), Price: 500, ProductType: "Millennium Falcon", Theme: "Star Wars", Count: 1, Available: true, UserId: 1, Company: "Disney", Rating: 5.0},
		{ReleaseDate: time.Date(1989, time.May, 22, 0, 0, 0, 0, time.UTC), Price: 75, ProductType: "Michelangelo Action Figure", Theme: "Teenage Mutant Ninja Turtles", Count: 12, Available: true, UserId: 1, Company: "Playmates Toys", Rating: 3.5},
		{ReleaseDate: time.Date(2010, time.March, 15, 0, 0, 0, 0, time.UTC), Price: 20.99, ProductType: "Lightning McQueen Diecast", Theme: "Cars", Count: 30, Available: true, UserId: 1, Company: "Disney", Rating: 3.0},
		{ReleaseDate: time.Date(1978, time.October, 17, 0, 0, 0, 0, time.UTC), Price: 1000, ProductType: "Luke Skywalker Original Action Figure", Theme: "Star Wars", Count: 3, Available: true, UserId: 3, Company: "Disney", Rating: 4.5},
		{ReleaseDate: time.Date(1998, time.June, 10, 0, 0, 0, 0, time.UTC), Price: 150, ProductType: "Red Ranger Figurine", Theme: "Power Rangers", Count: 8, Available: true, UserId: 2, Company: "Hasbro", Rating: 4.5},
		{ReleaseDate: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC), Price: 35.5, ProductType: "Baby Yoda Plush", Theme: "The Mandalorian", Count: 50, Available: true, UserId: 2, Company: "Disney", Rating: 2.3},
		{ReleaseDate: time.Date(1980, time.July, 5, 0, 0, 0, 0, time.UTC), Price: 500, ProductType: "Barbie Dream House", Theme: "Barbie", Count: 4, Available: true, UserId: 1, Company: "Mattel", Rating: 4.8},
		{ReleaseDate: time.Date(2003, time.March, 10, 0, 0, 0, 0, time.UTC), Price: 30.99, ProductType: "Barbie Fashionista Doll", Theme: "Barbie", Count: 20, Available: true, UserId: 1, Company: "Mattel", Rating: 4.6},
		{ReleaseDate: time.Date(2007, time.September, 25, 0, 0, 0, 0, time.UTC), Price: 100, ProductType: "LEGO Star Wars X-Wing Fighter", Theme: "Star Wars", Count: 15, Available: true, UserId: 2, Company: "LEGO", Rating: 5.0},
		{ReleaseDate: time.Date(1994, time.May, 1, 0, 0, 0, 0, time.UTC), Price: 200, ProductType: "LEGO Castle Blacksmith Shop", Theme: "Castle", Count: 10, Available: true, UserId: 3, Company: "LEGO", Rating: 4.7},
		{ReleaseDate: time.Date(2015, time.October, 15, 0, 0, 0, 0, time.UTC), Price: 50, ProductType: "LEGO Friends Heartlake City", Theme: "Friends", Count: 25, Available: true, UserId: 2, Company: "LEGO", Rating: 4.3},
		{ReleaseDate: time.Date(1991, time.August, 10, 0, 0, 0, 0, time.UTC), Price: 80, ProductType: "LEGO Pirates Ship", Theme: "Pirates", Count: 10, Available: true, UserId: 1, Company: "LEGO", Rating: 4.8},
		{ReleaseDate: time.Date(2018, time.March, 12, 0, 0, 0, 0, time.UTC), Price: 250, ProductType: "LEGO Star Wars Millennium Falcon", Theme: "Star Wars", Count: 5, Available: true, UserId: 3, Company: "LEGO", Rating: 4.9},
		{ReleaseDate: time.Date(2012, time.November, 22, 0, 0, 0, 0, time.UTC), Price: 15, ProductType: "Barbie Fashion Accessories", Theme: "Barbie", Count: 30, Available: true, UserId: 1, Company: "Mattel", Rating: 4.1},
		{ReleaseDate: time.Date(1999, time.July, 14, 0, 0, 0, 0, time.UTC), Price: 300, ProductType: "LEGO Ninjago Dragon", Theme: "Ninjago", Count: 8, Available: true, UserId: 2, Company: "LEGO", Rating: 4.7},
		{ReleaseDate: time.Date(2010, time.May, 18, 0, 0, 0, 0, time.UTC), Price: 75, ProductType: "LEGO City Police Car", Theme: "City", Count: 50, Available: true, UserId: 1, Company: "LEGO", Rating: 4.4},
		{ReleaseDate: time.Date(2017, time.February, 28, 0, 0, 0, 0, time.UTC), Price: 10.99, ProductType: "Barbie Camper Van", Theme: "Barbie", Count: 40, Available: true, UserId: 1, Company: "Mattel", Rating: 4.2},
		{ReleaseDate: time.Date(2005, time.April, 9, 0, 0, 0, 0, time.UTC), Price: 15, ProductType: "LEGO Creator 3-in-1 Car", Theme: "Creator", Count: 12, Available: true, UserId: 3, Company: "LEGO", Rating: 4.6},
		{ReleaseDate: time.Date(1990, time.December, 1, 0, 0, 0, 0, time.UTC), Price: 25, ProductType: "Barbie Princess Doll", Theme: "Barbie", Count: 22, Available: true, UserId: 2, Company: "Mattel", Rating: 4.5},
		{ReleaseDate: time.Date(1996, time.May, 17, 0, 0, 0, 0, time.UTC), Price: 150, ProductType: "LEGO Space Shuttle", Theme: "Space", Count: 6, Available: true, UserId: 2, Company: "LEGO", Rating: 4.8},
		{ReleaseDate: time.Date(2011, time.April, 3, 0, 0, 0, 0, time.UTC), Price: 250, ProductType: "LEGO Technic Race Car", Theme: "Technic", Count: 7, Available: true, UserId: 3, Company: "LEGO", Rating: 4.9},
		{ReleaseDate: time.Date(2013, time.June, 20, 0, 0, 0, 0, time.UTC), Price: 65, ProductType: "LEGO Minecraft Set", Theme: "Minecraft", Count: 15, Available: true, UserId: 1, Company: "LEGO", Rating: 4.5},
		{ReleaseDate: time.Date(1997, time.August, 11, 0, 0, 0, 0, time.UTC), Price: 40, ProductType: "LEGO Star Wars Tie Fighter", Theme: "Star Wars", Count: 20, Available: true, UserId: 2, Company: "LEGO", Rating: 4.6},
		{ReleaseDate: time.Date(2016, time.March, 22, 0, 0, 0, 0, time.UTC), Price: 55, ProductType: "LEGO Friends Heartlake Spa", Theme: "Friends", Count: 30, Available: true, UserId: 3, Company: "LEGO", Rating: 4.4},
		{ReleaseDate: time.Date(2021, time.August, 10, 0, 0, 0, 0, time.UTC), Price: 10, ProductType: "Barbie Mini Doll", Theme: "Barbie", Count: 60, Available: true, UserId: 2, Company: "Mattel", Rating: 3.8},
		{ReleaseDate: time.Date(2014, time.July, 10, 0, 0, 0, 0, time.UTC), Price: 20, ProductType: "LEGO Super Mario Starter Course", Theme: "Super Mario", Count: 18, Available: true, UserId: 1, Company: "LEGO", Rating: 4.9},
		{ReleaseDate: time.Date(2000, time.October, 23, 0, 0, 0, 0, time.UTC), Price: 80, ProductType: "Barbie Wedding Set", Theme: "Barbie", Count: 35, Available: true, UserId: 3, Company: "Mattel", Rating: 4.7},
		{ReleaseDate: time.Date(2019, time.November, 10, 0, 0, 0, 0, time.UTC), Price: 40, ProductType: "LEGO Harry Potter Diagon Alley", Theme: "Harry Potter", Count: 10, Available: true, UserId: 1, Company: "LEGO", Rating: 5.0},
		{ReleaseDate: time.Date(2006, time.March, 5, 0, 0, 0, 0, time.UTC), Price: 90, ProductType: "LEGO Creator House", Theme: "Creator", Count: 9, Available: true, UserId: 2, Company: "LEGO", Rating: 4.3},
		{ReleaseDate: time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC), Price: 15, ProductType: "LEGO City Race Car", Theme: "City", Count: 30, Available: true, UserId: 3, Company: "LEGO", Rating: 4.6},
		{ReleaseDate: time.Date(1998, time.May, 17, 0, 0, 0, 0, time.UTC), Price: 120, ProductType: "LEGO Castle Dragon", Theme: "Castle", Count: 15, Available: true, UserId: 1, Company: "LEGO", Rating: 4.9},
		{ReleaseDate: time.Date(1988, time.October, 12, 0, 0, 0, 0, time.UTC), Price: 20, ProductType: "Leonardo Action Figure", Theme: "Teenage Mutant Ninja Turtles", Count: 25, Available: true, UserId: 1, Company: "Playmates Toys", Rating: 4.5},
		{ReleaseDate: time.Date(1989, time.March, 5, 0, 0, 0, 0, time.UTC), Price: 25, ProductType: "Donatello Action Figure", Theme: "Teenage Mutant Ninja Turtles", Count: 18, Available: true, UserId: 2, Company: "Playmates Toys", Rating: 4.7},
		{ReleaseDate: time.Date(1990, time.June, 21, 0, 0, 0, 0, time.UTC), Price: 30, ProductType: "Shredder Action Figure", Theme: "Teenage Mutant Ninja Turtles", Count: 15, Available: true, UserId: 3, Company: "Playmates Toys", Rating: 4.6},
		{ReleaseDate: time.Date(1991, time.April, 10, 0, 0, 0, 0, time.UTC), Price: 35, ProductType: "Splinter Action Figure", Theme: "Teenage Mutant Ninja Turtles", Count: 20, Available: true, UserId: 1, Company: "Playmates Toys", Rating: 4.8},
		{ReleaseDate: time.Date(1992, time.November, 7, 0, 0, 0, 0, time.UTC), Price: 40, ProductType: "Raphael Action Figure", Theme: "Teenage Mutant Ninja Turtles", Count: 22, Available: true, UserId: 2, Company: "Playmates Toys", Rating: 4.7},
		{ReleaseDate: time.Date(1995, time.March, 15, 0, 0, 0, 0, time.UTC), Price: 60, ProductType: "Bebop and Rocksteady Set", Theme: "Teenage Mutant Ninja Turtles", Count: 10, Available: true, UserId: 3, Company: "Playmates Toys", Rating: 4.5},
		{ReleaseDate: time.Date(1996, time.August, 23, 0, 0, 0, 0, time.UTC), Price: 50, ProductType: "Turtle Van Playset", Theme: "Teenage Mutant Ninja Turtles", Count: 8, Available: true, UserId: 1, Company: "Playmates Toys", Rating: 4.9},
		{ReleaseDate: time.Date(1997, time.December, 9, 0, 0, 0, 0, time.UTC), Price: 70, ProductType: "TMNT Sewer Playset", Theme: "Teenage Mutant Ninja Turtles", Count: 12, Available: true, UserId: 2, Company: "Playmates Toys", Rating: 4.6},
		{ReleaseDate: time.Date(1999, time.July, 4, 0, 0, 0, 0, time.UTC), Price: 18, ProductType: "Teenage Mutant Ninja Turtles Pizza Thrower", Theme: "Teenage Mutant Ninja Turtles", Count: 15, Available: true, UserId: 3, Company: "Playmates Toys", Rating: 4.3},
		{ReleaseDate: time.Date(2001, time.February, 20, 0, 0, 0, 0, time.UTC), Price: 40, ProductType: "TMNT Secret Sewer Lair", Theme: "Teenage Mutant Ninja Turtles", Count: 25, Available: true, UserId: 1, Company: "Playmates Toys", Rating: 4.7},
		{ReleaseDate: time.Date(2003, time.July, 15, 0, 0, 0, 0, time.UTC), Price: 45, ProductType: "TMNT Battle Shells Set", Theme: "Teenage Mutant Ninja Turtles", Count: 18, Available: true, UserId: 2, Company: "Playmates Toys", Rating: 4.6},
		{ReleaseDate: time.Date(2005, time.May, 10, 0, 0, 0, 0, time.UTC), Price: 90, ProductType: "TMNT Turtle Blimp", Theme: "Teenage Mutant Ninja Turtles", Count: 8, Available: true, UserId: 3, Company: "Playmates Toys", Rating: 4.8},
		{ReleaseDate: time.Date(2007, time.August, 23, 0, 0, 0, 0, time.UTC), Price: 25, ProductType: "TMNT Foot Soldier", Theme: "Teenage Mutant Ninja Turtles", Count: 20, Available: true, UserId: 1, Company: "Playmates Toys", Rating: 4.4},
		{ReleaseDate: time.Date(2011, time.March, 30, 0, 0, 0, 0, time.UTC), Price: 55, ProductType: "TMNT Raph's Shellraiser", Theme: "Teenage Mutant Ninja Turtles", Count: 10, Available: true, UserId: 2, Company: "Playmates Toys", Rating: 4.7},
		{ReleaseDate: time.Date(2013, time.June, 5, 0, 0, 0, 0, time.UTC), Price: 35, ProductType: "TMNT The Technodrome", Theme: "Teenage Mutant Ninja Turtles", Count: 18, Available: true, UserId: 3, Company: "Playmates Toys", Rating: 4.6},
		{ReleaseDate: time.Date(2015, time.August, 19, 0, 0, 0, 0, time.UTC), Price: 60, ProductType: "TMNT Battle Shells Leonardo", Theme: "Teenage Mutant Ninja Turtles", Count: 12, Available: true, UserId: 1, Company: "Playmates Toys", Rating: 4.8},
		{ReleaseDate: time.Date(2017, time.November, 11, 0, 0, 0, 0, time.UTC), Price: 45, ProductType: "TMNT Ultimate Shellraiser", Theme: "Teenage Mutant Ninja Turtles", Count: 9, Available: true, UserId: 2, Company: "Playmates Toys", Rating: 4.7},
		{ReleaseDate: time.Date(2019, time.May, 13, 0, 0, 0, 0, time.UTC), Price: 50, ProductType: "TMNT Sewer Playset", Theme: "Teenage Mutant Ninja Turtles", Count: 8, Available: true, UserId: 3, Company: "Playmates Toys", Rating: 4.9},
		{ReleaseDate: time.Date(2020, time.January, 8, 0, 0, 0, 0, time.UTC), Price: 30, ProductType: "TMNT Bebop and Rocksteady", Theme: "Teenage Mutant Ninja Turtles", Count: 15, Available: true, UserId: 1, Company: "Playmates Toys", Rating: 4.6},
		{ReleaseDate: time.Date(1995, time.September, 20, 0, 0, 0, 0, time.UTC), Price: 150, ProductType: "Power Rangers Red Ranger Action Figure", Theme: "Power Rangers", Count: 10, Available: true, UserId: 2, Company: "Hasbro", Rating: 4.5},
		{ReleaseDate: time.Date(2002, time.November, 3, 0, 0, 0, 0, time.UTC), Price: 40, ProductType: "GI Joe Snake Eyes Action Figure", Theme: "GI Joe", Count: 15, Available: true, UserId: 1, Company: "Hasbro", Rating: 4.7},
		{ReleaseDate: time.Date(2006, time.March, 18, 0, 0, 0, 0, time.UTC), Price: 60, ProductType: "Nerf N-Strike Maverick", Theme: "Nerf", Count: 20, Available: true, UserId: 2, Company: "Hasbro", Rating: 4.6},
		{ReleaseDate: time.Date(2010, time.June, 21, 0, 0, 0, 0, time.UTC), Price: 25, ProductType: "Transformers Bumblebee", Theme: "Transformers", Count: 5, Available: true, UserId: 1, Company: "Hasbro", Rating: 4.9},
		{ReleaseDate: time.Date(2014, time.April, 10, 0, 0, 0, 0, time.UTC), Price: 70, ProductType: "GI Joe Cobra Commander", Theme: "GI Joe", Count: 8, Available: true, UserId: 3, Company: "Hasbro", Rating: 4.5},
		{ReleaseDate: time.Date(2015, time.July, 30, 0, 0, 0, 0, time.UTC), Price: 15, ProductType: "Monopoly Game", Theme: "Monopoly", Count: 30, Available: true, UserId: 2, Company: "Hasbro", Rating: 4.2},
		{ReleaseDate: time.Date(2017, time.November, 5, 0, 0, 0, 0, time.UTC), Price: 100, ProductType: "Nerf Rival Nemesis", Theme: "Nerf", Count: 10, Available: true, UserId: 1, Company: "Hasbro", Rating: 4.7},
		{ReleaseDate: time.Date(2018, time.October, 17, 0, 0, 0, 0, time.UTC), Price: 25, ProductType: "Transformers Power of the Primes Optimus Prime", Theme: "Transformers", Count: 12, Available: true, UserId: 2, Company: "Hasbro", Rating: 4.8},
		{ReleaseDate: time.Date(2019, time.March, 21, 0, 0, 0, 0, time.UTC), Price: 45, ProductType: "Nerf Rival Nemesis", Theme: "Nerf", Count: 8, Available: true, UserId: 3, Company: "Hasbro", Rating: 4.7},
		{ReleaseDate: time.Date(2020, time.June, 1, 0, 0, 0, 0, time.UTC), Price: 80, ProductType: "Monopoly: The Office Edition", Theme: "Monopoly", Count: 50, Available: true, UserId: 1, Company: "Hasbro", Rating: 4.4},
		{ReleaseDate: time.Date(2021, time.January, 10, 0, 0, 0, 0, time.UTC), Price: 22, ProductType: "GI Joe Classified Snake Eyes", Theme: "GI Joe", Count: 25, Available: true, UserId: 2, Company: "Hasbro", Rating: 4.6},
	}

	for _, toy := range toys {
		if err := database.Database.Db.Create(&toy).Error; err != nil {
			log.Printf("Failed to seed Toy: %v\n", err)
		}
	}
}

func UndoToys() {
	// Delete all records in the toys table
	if err := database.Database.Db.Exec("DELETE FROM toys").Error; err != nil {
		log.Printf("Failed to delete all toys: %v\n", err)
	} else {
		log.Println("Successfully deleted all toys from the table.")
	}

	// Reset the auto-increment ID for the products table
	switch database.Database.Db.Dialector.Name() {
	case "sqlite":
		// SQLite-specific reset
		if err := database.Database.Db.Exec("DELETE FROM sqlite_sequence WHERE name = 'toys'").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (SQLite): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toys table (SQLite).")
		}
	case "postgres":
		// PostgreSQL-specific reset
		sequenceName := "toys_id_seq" // Default naming convention in PostgreSQL
		if err := database.Database.Db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (PostgreSQL): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toys table (PostgreSQL).")
		}
	default:
		log.Println("Unsupported database type. Auto-increment ID reset skipped.")
	}
}
