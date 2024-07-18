package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"
)

// StringUnpack функция, которая распаковывает строку.
func StringUnpack(str string) (string, error) {
	s := []rune(str)

	// Проверяем, что строка не пустая и не начинается с цифры.
	if str != "" && unicode.IsDigit(s[0]) {
		return "", fmt.Errorf("invalid string: cannot start with digit: %s", string(s[0]))
	}

	var (
		newStr []rune
		numStr string
		char   rune
	)

	for i, r := range s {
		// Если символ цифра, добавляем его к numStr.
		if unicode.IsDigit(r) {
			if numStr == "" { // Проверяем, чтобы символ не перезаписался на число.
				char = s[i-1] // Сохраняем предыдущий символ.
			}
			numStr += string(s[i]) // Добавляем цифру к numStr.
			if i == len(str)-1 {   // Если это последний символ строки, распаковываем символы.
				symbolsUnpack(&newStr, &numStr, char)
			}
		} else {
			// Распаковываем накопленные символы и добавляем текущий символ в результат.
			symbolsUnpack(&newStr, &numStr, char)
			newStr = append(newStr, r)
		}
	}

	return string(newStr), nil
}

// symbolsUnpack функция, которая распаковывает символ.
func symbolsUnpack(newStr *[]rune, numStr *string, char rune) {
	if *numStr != "" {
		num, err := strconv.Atoi(*numStr) // Преобразуем строку в число.
		if err != nil {
			fmt.Println("error:", err)
			newStr = nil
			return
		}

		// Добавляем char num-1 раз, так как один раз он уже был добавлен.
		for i := 1; i < num; i++ {
			*newStr = append(*newStr, char)
		}

		*numStr = "" // Обнуляем numStr после распаковки.
	}
}

func main() {
	result, err := StringUnpack("abcd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result) // "aaaabccddddde"
}
