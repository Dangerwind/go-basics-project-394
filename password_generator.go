package main
import (
	"fmt"
	"strings"
)

const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits    = "0123456789"
const special   = "!@#$%^&*"


func NextRandom(number int) int {
	return (16807 * number) % 2147483647
}


func GeneratePassword(length, seed int, useUppercase, useDigits, useSpecial bool) string {
	// Строчные буквы включены всегда
	alphabet := lowercase

	if useUppercase {
		alphabet += uppercase
	}

	if useDigits {
		alphabet += digits
	}

	if useSpecial {
		alphabet += special
	}

	alphabetLen := len(alphabet)
	result := ""
	current := seed 

	for i := 0; i < length; i++ {
		current = NextRandom(current)
		index := current % alphabetLen
		result = result + string(alphabet[index])
	}

	return result
}

func CheckPassword(password string) string {
	score := 0

	if len(password) >= 8 {
		score++
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			hasLower = true
		} else if c >= 'A' && c <= 'Z' {
			hasUpper = true
		} else if c >= '0' && c <= '9' {
			hasDigit = true
		} else if strings.ContainsRune(special, c) {
			hasSpecial = true
		}
	}

	if hasLower {
		score++
	}
	if hasUpper {
		score++
	}
	if hasDigit {
		score++
	}
	if hasSpecial {
		score++
	}

	var verdict string
	switch score {
	case 0, 1, 2:
		verdict = "Слабый"
	case 3:
		verdict = "Средний"
	case 4:
		verdict = "Надёжный"
	case 5:
		verdict = "Очень надёжный"
	}

	return fmt.Sprintf("%s пароль (оценка %d из 5)", verdict, score)
}