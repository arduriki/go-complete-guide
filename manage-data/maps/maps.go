package maps

import "fmt"

func main() {
	// Create a map - a dict in other languages.
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	// Access to a value key.
	fmt.Println(websites["Amazon Web Services"])
	// Add a new key-value pair.
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	// Delete a key-value pair.
	delete(websites, "Google")
	fmt.Println(websites)
}
