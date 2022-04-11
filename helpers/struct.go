package helpers

// Create a new Student struct
type Student struct {
	Name    string `fake:"{name}"`
	Address string `fake:"{street}, {city}, {country}"`
	Job     string `fake:"{jobtitle}"`
	Reason  string `fake:"{sentence:5}"`
}
