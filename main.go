package main

import "fmt"

func main() {
	// Проверка функции из условий задачи
	fmt.Println(CheckPassword("abc"))         // Слабый пароль (оценка 1 из 5)
	fmt.Println(CheckPassword("abcdefgh"))    // Слабый пароль (оценка 2 из 5)
	fmt.Println(CheckPassword("abcdef1234"))  // Средний пароль (оценка 3 из 5)
	fmt.Println(CheckPassword("Abcdef1234"))  // Надёжный пароль (оценка 4 из 5)
	fmt.Println(CheckPassword("Abcdef123!")) 
}