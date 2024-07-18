package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/pborman/getopt"
)

// getFile читает строки из указанного файла и возвращает их в виде слайса строк.
// Если uniqueRequired установлено в true, то удаляет повторяющиеся строки.
func getFile(path string, uniqueRequired bool) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	set := make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	data := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if uniqueRequired {
			if _, ok := set[line]; !ok {
				set[line] = struct{}{}
			} else {
				continue
			}
		}
		data = append(data, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

// toSort сортирует строки. Если n установлено в true, то сортирует по числовому значению.
func toSort(data []string, n bool) []string {
	if n {
		sort.Slice(data, func(i, j int) bool {
			vi, err1 := strconv.Atoi(data[i])
			vj, err2 := strconv.Atoi(data[j])
			if err1 != nil || err2 != nil {
				// Если строки не могут быть преобразованы в числа, сортируем как строки.
				return data[i] < data[j]
			}
			return vi < vj
		})
	} else {
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
	}
	return data
}

func main() {
	// Парсинг аргументов командной строки.
	filename := getopt.String('f', "", "файл")
	n := getopt.Bool('n', "сортировка по числовому значению")
	r := getopt.Bool('r', "сортировка в обратном порядке")
	u := getopt.Bool('u', "не выводить повторяющиеся строки")
	getopt.Parse()

	// Получение строк из файла.
	file, err := getFile(*filename, *u)
	if err != nil {
		panic(err)
	}

	// Сортировка строк.
	file = toSort(file, *n)

	// Вывод отсортированных строк.
	if *r {
		for i := len(file) - 1; i >= 0; i-- {
			fmt.Println(file[i])
		}
	} else {
		for _, value := range file {
			fmt.Println(value)
		}
	}
}
