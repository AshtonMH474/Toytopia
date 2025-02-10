package seeders

import (
	"log"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
)

func SeedToyImages() {
	images := []models.ToyImage{
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/Baby-Yoda-Plush-Stuffed-Toy-Plushie-Cute-Animal-Pillow-Grogu-Baby-Yoda-Star-Wars-Kids-Doll-Gift-10-inch-26-cm_b0514180-f418-49bf-aaae-05fd87088b9f.c00883bd8f234eb2d97ea171a678024f.avif", PrimaryImg: true, ToyId: 8},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/redranger.jpg", PrimaryImg: true, ToyId: 7},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/luke.jpg", PrimaryImg: true, ToyId: 6},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/Radiator-Springs-Lightning-McQueen-Diecast-Car-Disney-Cars_33243daa-d660-4423-a1a2-5cdce7d1cf4d.5079ef90c615abdf4886a7500f76e767.avif", PrimaryImg: true, ToyId: 5},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/BST-AXN_TMNT_IDW_Michelangelo-Battle-Ready_Package_1.webp", PrimaryImg: true, ToyId: 4},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/mellianflacon.jpg", PrimaryImg: true, ToyId: 3},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/buzz.jpg", PrimaryImg: true, ToyId: 2},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/optimus.png", PrimaryImg: true, ToyId: 1},
		{ImgUrl: "https://cdn11.bigcommerce.com/s-cy4lua1xoh/images/stencil/1280x1280/products/33464/427451/82b6d5d8-8de5-5398-b619-ec94fd9ae919__05565.1719337509.jpg?c=1?imbypass=on", PrimaryImg: true, ToyId: 9},
		{ImgUrl: "https://static.wikia.nocookie.net/barbie/images/3/3f/Fashionistas_Sassy_Doll_R9882_%2801%29.jpg/revision/latest?cb=20210509105558", PrimaryImg: true, ToyId: 10},
		{ImgUrl: "https://i.ytimg.com/vi/GmOfwtQzXcY/maxresdefault.jpg", PrimaryImg: true, ToyId: 11},
		{ImgUrl: "https://i.ebayimg.com/images/g/CbIAAOSwk9lifrye/s-l640.jpg", PrimaryImg: true, ToyId: 12},
		{ImgUrl: "https://m.media-amazon.com/images/I/91NWGxqXBML.jpg", PrimaryImg: true, ToyId: 13},
		{ImgUrl: "https://static.wikia.nocookie.net/lego/images/2/22/6271_Imperial_Flagship.jpg/revision/latest/thumbnail/width/360/height/360?cb=20090817092702", PrimaryImg: true, ToyId: 14},
		{ImgUrl: "https://www.lego.com/cdn/cs/set/assets/blt6c7f809caec392ac/75212_alt1.jpg", PrimaryImg: true, ToyId: 15},
		{ImgUrl: "https://m.media-amazon.com/images/I/91Ed0nNU6yL.jpg", PrimaryImg: true, ToyId: 16},
		{ImgUrl: "https://m.media-amazon.com/images/I/71XKTVkQbkL.jpg", PrimaryImg: true, ToyId: 17},
		{ImgUrl: "https://m.media-amazon.com/images/I/71zUtgiuNWL.jpg", PrimaryImg: true, ToyId: 18},
		{ImgUrl: "https://m.media-amazon.com/images/I/812LsRuocJL.jpg", PrimaryImg: true, ToyId: 19},
		{ImgUrl: "https://i.ebayimg.com/00/s/MTYwMFgxNjAw/z/UFQAAOSwi0xZ6bDs/$_57.JPG?set_id=8800005007", PrimaryImg: true, ToyId: 20},
		{ImgUrl: "https://media-photos.depop.com/b1/43595479/2010975236_661bda186dc747bc9a7e2eb25371e6a1/P0.jpg", PrimaryImg: true, ToyId: 21},
		{ImgUrl: "https://m.media-amazon.com/images/I/91AhKl85vEL.jpg", PrimaryImg: true, ToyId: 22},
		{ImgUrl: "https://m.media-amazon.com/images/I/81NdUM-decL._AC_UF894,1000_QL80_.jpg", PrimaryImg: true, ToyId: 23},
		{ImgUrl: "https://m.media-amazon.com/images/I/614cywpje5L.jpg", PrimaryImg: true, ToyId: 24},
		{ImgUrl: "https://m.media-amazon.com/images/I/81JsaPhdmfL.jpg", PrimaryImg: true, ToyId: 25},
		{ImgUrl: "https://m.media-amazon.com/images/I/614FNqe9S9L._AC_UF894,1000_QL80_.jpg", PrimaryImg: true, ToyId: 26},
		{ImgUrl: "https://m.media-amazon.com/images/I/71jMC0Bu7JL.jpg", PrimaryImg: true, ToyId: 27},
		{ImgUrl: "https://assets.nintendo.com/image/upload/c_fill,w_1200/q_auto:best/f_auto/dpr_2.0/ncom/en_US/products/merchandise/toys%20and%20games/lego/legor-super-mario-adventures-with-mario-starter-course/113448-lego-super-mario-adventures-with-mario-starter-course-package-front-1200x675", PrimaryImg: true, ToyId: 28},
		{ImgUrl: "https://barbielistholland.wordpress.com/wp-content/uploads/2013/11/2000-barbie-doll-ken-skipper-kelly-wedding-day-set-i-can-be-bride-goom.jpg?w=848", PrimaryImg: true, ToyId: 29},
		{ImgUrl: "https://m.media-amazon.com/images/I/71MSrSjLhaL._AC_UF894,1000_QL80_.jpg", PrimaryImg: true, ToyId: 30},
		{ImgUrl: "https://m.media-amazon.com/images/I/71zPHnDbZmL.jpg", PrimaryImg: true, ToyId: 31},
		{ImgUrl: "https://i5.walmartimages.com/asr/f901b097-6613-4f03-b03c-32cd71f84183.050580d419e3801680a56ee8cd8a94e5.jpeg?odnHeight=612&odnWidth=612&odnBg=FFFFFF", PrimaryImg: true, ToyId: 32},
		{ImgUrl: "https://m.media-amazon.com/images/I/81oVcmRzb6L._AC_UF894,1000_QL80_.jpg", PrimaryImg: true, ToyId: 33},
		{ImgUrl: "https://i.ebayimg.com/images/g/0jcAAOSw-CxjDWcK/s-l1200.jpg", PrimaryImg: true, ToyId: 34},
		{ImgUrl: "https://i5.walmartimages.com/asr/a73b734e-af75-4eaf-9cc1-d036c1341734.5e9d56e2746b4fba10f23ae1c3a71077.jpeg?odnHeight=768&odnWidth=768&odnBg=FFFFFF", PrimaryImg: true, ToyId: 35},
		{ImgUrl: "https://m.media-amazon.com/images/I/71hxoHkuEFL._AC_UF894,1000_QL80_.jpg", PrimaryImg: true, ToyId: 36},
		{ImgUrl: "https://i.ebayimg.com/images/g/CeQAAOSwMT1isJFL/s-l400.jpg", PrimaryImg: true, ToyId: 37},
		{ImgUrl: "https://i.ebayimg.com/images/g/jjgAAOSw6Ghlw75q/s-l1200.jpg", PrimaryImg: true, ToyId: 38},
		{ImgUrl: "https://i.ebayimg.com/images/g/z4oAAOSwHvBi1YjZ/s-l1200.jpg", PrimaryImg: true, ToyId: 39},
		{ImgUrl: "https://m.media-amazon.com/images/I/61gKm2DOa1L.jpg", PrimaryImg: true, ToyId: 40},
		{ImgUrl: "https://www.transformerland.com/image/reference_images/51496.jpg", PrimaryImg: true, ToyId: 41},
		{ImgUrl: "https://images.bigbadtoystore.com/images/p/full/2015/12/PLM11387_i.jpg", PrimaryImg: true, ToyId: 42},
		{ImgUrl: "https://m.media-amazon.com/images/I/81sKUP9YJCL.jpg", PrimaryImg: true, ToyId: 43},
		{ImgUrl: "https://i.redd.it/et2os0xt7nub1.jpg", PrimaryImg: true, ToyId: 44},
		{ImgUrl: "https://media.entertainmentearth.com/assets/images/a00ee5636b314488a24d4ef5c25161caxl.jpg", PrimaryImg: true, ToyId: 45},
		{ImgUrl: "https://i.ebayimg.com/images/g/hsIAAOSwa-5mPDAX/s-l1200.jpg", PrimaryImg: true, ToyId: 46},
		{ImgUrl: "https://i.ebayimg.com/images/g/R~kAAOSwzx9gYC1U/s-l400.jpg", PrimaryImg: true, ToyId: 47},
		{ImgUrl: "https://m.media-amazon.com/images/I/81p4TuYmixL.jpg", PrimaryImg: true, ToyId: 48},
		{ImgUrl: "https://www.jnltradinginc.com/cdn/shop/products/JNL_83f0ee48-726c-42ba-8306-9402a3764652_grande.jpg?v=1456098332", PrimaryImg: true, ToyId: 49},
		{ImgUrl: "https://images-na.ssl-images-amazon.com/images/I/71hz9zTNAtL._AC_SL1257_.jpg", PrimaryImg: true, ToyId: 50},
		{ImgUrl: "https://i5.walmartimages.com/seo/Teenage-Mutant-Ninja-Turtles-Mutant-Mayhem-Sewer-Lair-Multicolor-Action-Figure-Playset-by-Playmates_3fd0ed6f-f618-4842-a3a9-a9eb74e40a55.e58b8f001838ac05a349b428081ea090.jpeg", PrimaryImg: true, ToyId: 51},
		{ImgUrl: "https://collecticontoys.com/cdn/shop/products/neca-tmnt-bebop-rocksteady-2pack-second-run-action-figure-toys_1200x1200.jpg?v=1604386757", PrimaryImg: true, ToyId: 52},
		{ImgUrl: "https://m.media-amazon.com/images/I/41fe3cvauAL.jpg", PrimaryImg: true, ToyId: 53},
		{ImgUrl: "https://m.media-amazon.com/images/I/61P194QC+vL.jpg", PrimaryImg: true, ToyId: 54},
		{ImgUrl: "https://m.media-amazon.com/images/I/51A88NG-6DL.jpg", PrimaryImg: true, ToyId: 55},
		{ImgUrl: "https://news.tfw2005.com/wp-content/uploads/sites/10/2010/02/Battle-Ops-Bumblebee-robot_1266123996.jpg", PrimaryImg: true, ToyId: 56},
		{ImgUrl: "https://m.media-amazon.com/images/I/71icLNGZbcL.jpg", PrimaryImg: true, ToyId: 57},
		{ImgUrl: "https://m.media-amazon.com/images/I/91T6FHqQuNL.jpg", PrimaryImg: true, ToyId: 58},
		{ImgUrl: "https://m.media-amazon.com/images/I/719oOPOzt2L.jpg", PrimaryImg: true, ToyId: 59},
		{ImgUrl: "https://m.media-amazon.com/images/I/81DMS9-vSOL.jpg", PrimaryImg: true, ToyId: 60},
		{ImgUrl: "https://i.ebayimg.com/images/g/PecAAOSwJ3BhoYJN/s-l1200.jpg", PrimaryImg: true, ToyId: 61},
		{ImgUrl: "https://m.media-amazon.com/images/I/91DOO+MlcHL._AC_UF894,1000_QL80_.jpg", PrimaryImg: true, ToyId: 62},
		{ImgUrl: "https://tcsrockets.com/wp-content/uploads/2021/10/SnakeEyesOrigins.jpeg", PrimaryImg: true, ToyId: 63},
	}
	for _, image := range images {
		if err := database.Database.Db.Create(&image).Error; err != nil {
			log.Printf("Failed to seed image: %v\n", err)
		}
	}
}

func UndoToyImages() {
	if err := database.Database.Db.Exec("DELETE FROM toy_images").Error; err != nil {
		log.Printf("Failed to delete all toy_images: %v\n", err)
	} else {
		log.Println("Successfully deleted all toy_images from the table.")
	}

	// Reset the auto-increment ID for the products table
	switch database.Database.Db.Dialector.Name() {
	case "sqlite":
		// SQLite-specific reset
		if err := database.Database.Db.Exec("DELETE FROM sqlite_sequence WHERE name = 'toy_images'").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (SQLite): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toy_images table (SQLite).")
		}
	case "postgres":
		// PostgreSQL-specific reset
		sequenceName := "toy_images_id_seq" // Default naming convention in PostgreSQL
		if err := database.Database.Db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (PostgreSQL): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toy_images table (PostgreSQL).")
		}
	default:
		log.Println("Unsupported database type. Auto-increment ID reset skipped.")
	}
}
