package main

const lowercase = "abcdefghijklmnopqrstuvwxyz"


func GeneratePassword(length int) string {
	alphabet := lowercase
	alphabetLen := len(alphabet) 
	result := ""

	for i := 0; i < length; i++ {
		index := i % alphabetLen 
		result = result + string(alphabet[index]) 
	}

	return result
}