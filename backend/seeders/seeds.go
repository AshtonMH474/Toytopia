package seeders

func SeedAll() {
	SeedUsers()
	SeedToys()
	SeedWishlists()
	SeedReviews()
	SeedToyImages()
}

func UndoSeeds() {
	UndoToyImages()
	UndoReviews()
	UndoWishlistsToys()
	UndoAllWishLists()
	UndoToys()
	UndoAllUsers()
}
