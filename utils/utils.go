package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

// Common first names for random generation
var firstNames = []string{
	"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph",
	"Thomas", "Charles", "Mary", "Patricia", "Jennifer", "Linda", "Elizabeth",
	"Barbara", "Susan", "Jessica", "Sarah", "Karen", "Igor", "Anna", "Daniel",
	"Emma", "Sophia", "Oliver", "Noah", "Ava", "Isabella", "Mia",
}

// Common last names for random generation
var lastNames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis",
	"Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson",
	"Thomas", "Taylor", "Moore", "Jackson", "Martin", "Lee", "Thompson", "White",
}

// Email domains for random generation
var emailDomains = []string{
	"gmail.com", "yahoo.com", "hotmail.com", "outlook.com", "icloud.com",
	"mail.com", "protonmail.com", "aol.com",
}

func GenerateRandomName() string {
	firstName := firstNames[rand.Intn(len(firstNames))]
	lastName := lastNames[rand.Intn(len(lastNames))]
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func GenerateRandomPhone() string {
	// Generate Israeli-like phone number (050/052/054/055/058)
	prefixes := []string{"050", "052", "054", "055", "058"}
	prefix := prefixes[rand.Intn(len(prefixes))]

	// Generate 7 random digits
	numbers := make([]byte, 7)
	for i := 0; i < 7; i++ {
		numbers[i] = byte(rand.Intn(10)) + '0'
	}

	return fmt.Sprintf("%s%s", prefix, string(numbers))
}

func GenerateRandomEmail() string {
	// Use random parts to create email
	name := strings.ToLower(firstNames[rand.Intn(len(firstNames))])
	number := rand.Intn(9999)
	domain := emailDomains[rand.Intn(len(emailDomains))]

	return fmt.Sprintf("%s%d@%s", name, number, domain)
}
