package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Функція для обфускації символів
func obfuscate(char byte) byte {
	return char ^ 0x5F
}

// Функція для деобфускації символів
func deobfuscate(char byte) byte {
	return char ^ 0x5F
}

// Функція для отримання правильного пароля
func getCorrectPassword() string {
	originalPassword := "SecurePass123"
	obfuscated := make([]byte, len(originalPassword))
	for i, c := range originalPassword {
		obfuscated[i] = obfuscate(byte(c))
	}
	return string(obfuscated)
}

// Функція для перевірки пароля
func checkPassword(input string) bool {
	obfuscatedPassword := getCorrectPassword()
	deobfuscated := make([]byte, len(obfuscatedPassword))
	for i := 0; i < len(obfuscatedPassword); i++ {
		deobfuscated[i] = deobfuscate(obfuscatedPassword[i])
	}
	return input == string(deobfuscated)
}

func main() {
	fmt.Println("Вітаємо у GoCrackMaster!")
	fmt.Print("Будь ласка, введіть пароль активації: ")

	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Помилка читання вводу:", err)
		return
	}

	// Видалення символів нового рядка
	userInput = userInput[:len(userInput)-1]

	// Невелика затримка для ускладнення аналізу
	time.Sleep(500 * time.Millisecond)

	if checkPassword(userInput) {
		fmt.Println("Вітаємо! Ви розбили GoCrackMaster.")
	} else {
		fmt.Println("Пароль невірний. Спробуйте ще раз.")
	}
}
