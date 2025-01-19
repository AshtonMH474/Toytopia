package seeders

func SeedAll() {
	SeedUsers()
	SeedToys()
}

func UndoSeeds() {
	UndoToys()
	UndoAllUsers()
}
