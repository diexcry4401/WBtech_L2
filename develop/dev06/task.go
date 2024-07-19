package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pborman/getopt"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// splitText разбивает строку на колонки по указанному разделителю.
func splitText(text, delimiter string) []string {
	return strings.Split(text, delimiter)
}

// processLine обрабатывает одну строку в зависимости от флагов.
func processLine(line string, fields int, delimiter string, separated bool) string {
	// Проверяем наличие разделителя в строке.
	if separated && !strings.Contains(line, delimiter) {
		return "" // Если флаг -s установлен и разделителя нет, возвращаем пустую строку.
	}

	// Разбиваем строку на колонки.
	columns := splitText(line, delimiter)

	// Проверяем, что строка содержит достаточно колонок.
	if len(columns) < fields {
		return line // Если колонок меньше, чем запрашивается, возвращаем оригинальную строку.
	}

	// Возвращаем значение из указанной колонки.
	return columns[fields-1]
}

// main функция программы.
func main() {
	// Определяем флаги командной строки.
	fields := getopt.IntLong("fields", 'f', 0, "выбрать поля (колонки)")
	delimiter := getopt.StringLong("delimiter", 'd', "\t", "использовать другой разделитель")
	separated := getopt.BoolLong("separated", 's', "только строки с разделителем")

	// Парсим флаги командной строки.
	getopt.Parse()

	// Проверяем, что количество полей (колонок) указано корректно.
	if *fields <= 0 {
		log.Fatal("Число полей (колонок) должно быть больше 0.")
	}

	// Читаем строки из стандартного ввода.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		result := processLine(line, *fields, *delimiter, *separated)
		if result != "" {
			fmt.Println(result)
		}
	}

	// Проверяем на наличие ошибок при чтении стандартного ввода.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
