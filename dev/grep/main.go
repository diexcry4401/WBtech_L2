package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/pborman/getopt"
)

// openFile открывает файл и считывает его содержимое в слайс строк.
func openFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}

// getExpression компилирует регулярное выражение, добавляя опцию игнорирования регистра при необходимости.
func getExpression(pattern string, ignore bool) (*regexp.Regexp, error) {
	ignorePrefix := ""
	if ignore {
		ignorePrefix = "(?i)"
	}
	compiledExpession, err := regexp.Compile(ignorePrefix + pattern)
	if err != nil {
		return nil, err
	}
	return compiledExpession, nil
}

// getNumberOfIntersections считает количество строк, совпадающих с заданным регулярным выражением.
func getNumberOfIntersections(file []string, expression *regexp.Regexp) int {
	result := 0
	for _, str := range file {
		match := expression.Match([]byte(str))
		if match {
			result++
		}
	}
	return result
}

// reg обрабатывает строки файла, печатая строки в соответствии с заданными параметрами.
func reg(file []string, expression *regexp.Regexp, after, before int, number, invert bool) {
	for i, str := range file {
		match := expression.Match([]byte(str))
		if invert && !match {
			echo(file, i, after, before, number)
		} else if !invert && match {
			echo(file, i, after, before, number)
		}
	}
}

// echo выводит строки файла с учетом контекста (до и после совпадения) и номера строки.
func echo(file []string, i, after, before int, number bool) {
	startPoint := 0
	endPoint := len(file)
	if i-after > 0 {
		startPoint = i - after
	}
	if i+before < len(file) {
		endPoint = i + before
	}
	if endPoint != len(file) {
		endPoint++
	}
	fmt.Println("------------------------")
	for line := startPoint; line < endPoint; line++ {
		if number {
			fmt.Printf("%d: ", line+1)
		}
		fmt.Printf("%s\n", file[line])
	}
	fmt.Println("------------------------")
}

func main() {
	// Определение и парсинг флагов командной строки.
	pattern := getopt.String('e', "", "паттерн")
	path := getopt.String('f', "", "файл")
	after := getopt.IntLong("after", 'A', 0, "вывод N строк после совпадения")
	before := getopt.IntLong("before", 'B', 0, "вывод N строк до совпадения")
	inTheMiddle := getopt.IntLong("context", 'C', 0, "вывод N строк в районе совпадения")
	count := getopt.Bool('c', "вывести количество строк с совпадением")
	ignore := getopt.Bool('i', "игнорировать различия регистра")
	invert := getopt.Bool('v', "инвертировать вывод")
	number := getopt.Bool('n', "напечатать номер строки")

	getopt.Parse()

	// Чтение файла.
	file, err := openFile(*path)
	if err != nil {
		panic(err)
	}

	// Компиляция регулярного выражения.
	expression, err := getExpression(*pattern, *ignore)
	if err != nil {
		panic(err)
	}

	// Обработка флага подсчета совпадений.
	if *count {
		result := getNumberOfIntersections(file, expression)
		if *invert {
			result = len(file) - result
		}
		fmt.Println(result)
	} else {
		// Обработка флага контекста.
		if *after == 0 && *before == 0 && *inTheMiddle != 0 {
			*after = *inTheMiddle / 2
			*before = *inTheMiddle / 2
		}
		// Выполнение поиска с учетом параметров.
		reg(file, expression, *after, *before, *number, *invert)
	}
}
