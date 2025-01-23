package seeders

func SeedAll() {
	SeedUsers()
	SeedToys()
	SeedWishlists()
}

func UndoSeeds() {
	UndoWishlistsToys()
	UndoAllWishLists()
	UndoToys()
	UndoAllUsers()
}
