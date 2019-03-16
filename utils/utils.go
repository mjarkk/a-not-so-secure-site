package utils

import (
	"crypto/rand"
	"math/big"
	mathRand "math/rand"
	"strings"
)

// RandomString generates a purly random string with the lenght of n
func RandomString(length int) (string, error) {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(2147483647))
	if err != nil {
		return "", err
	}
	r := mathRand.New(mathRand.NewSource(randomNumber.Int64()))
	possibleLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	l := int64(len(possibleLetters))

	toReturn := ""
	for i := 0; i < length; i++ {
		toReturn = toReturn + string(possibleLetters[r.Int63n(l)])
	}

	return toReturn, nil
}

// GetDomain returns the domain of a input string
// Example:
//    domain := other.GetDomain("https://test.somedomain.com:8080/idk", true)
//    fmt.Println(domain) // -> test.somedomain.com
func GetDomain(input string, removePort bool) string {
	workingOn := input
	if strings.Contains(input, "http://") || strings.Contains(input, "https://") {
		workingOn = strings.Split(input, "//")[1]
	}

	if removePort {
		workingOn = strings.Split(workingOn, ":")[0]
	}
	workingOn = strings.Split(workingOn, "/")[0]
	workingOn = strings.Split(workingOn, "#")[0]
	workingOn = strings.Split(workingOn, "?")[0]

	return workingOn
}
