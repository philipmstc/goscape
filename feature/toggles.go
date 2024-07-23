package feature

func IsAnimated() bool {
	return false 
}

func Load() {
	// eventually, loads feature configurations from CLI or DB 
}

func IsNewGame() bool {
	return true
}

func DetailedRecipeStrings() bool { 
	return false
}