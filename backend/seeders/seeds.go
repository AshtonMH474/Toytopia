package seeders

func SeedAll() {
	SeedUsers()
	SeedToys()
	SeedWishlists()
}

func UndoSeeds() {
	UndoWishlists()
	UndoToys()
	UndoAllUsers()
}
