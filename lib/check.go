package lib

// Check handles errors quickly
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
