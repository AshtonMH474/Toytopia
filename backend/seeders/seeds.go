package seeders

func SeedAll() {
	SeedUsers()
	SeedToys()
	SeedWishlists()
	SeedReviews()
}

func UndoSeeds() {
	UndoReviews()
	UndoWishlistsToys()
	UndoAllWishLists()
	UndoToys()
	UndoAllUsers()
}
