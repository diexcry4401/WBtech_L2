package main

import (
	"fmt"
	"sort"
	"strings"
)

// normalizeWord нормализует слово, приводя его к нижнему регистру и сортируя его буквы.
// Это помогает идентифицировать анаграммы, приводя все слова к единому представлению.
func normalizeWord(word string) string {
	word = strings.ToLower(word) // Приведение слова к нижнему регистру
	runes := []rune(word)        // Преобразование строки в слайс рун для корректной сортировки Unicode символов
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j] // Сортировка рун в алфавитном порядке
	})
	return string(runes) // Преобразование отсортированных рун обратно в строку
}

// findAnagrams находит все множества анаграмм в заданном словаре.
func findAnagrams(words []string) *map[string][]string {
	anagrams := make(map[string][]string)    // Мапа для хранения групп анаграмм
	uniqueWords := make(map[string]struct{}) // Мапа для хранения уникальных слов

	// Группируем слова по их нормализованной форме.
	for _, word := range words {
		loweredWord := strings.ToLower(word) // Приведение слова к нижнему регистру
		if _, exists := uniqueWords[loweredWord]; exists {
			continue // Пропуск дублирующихся слов
		}
		uniqueWords[loweredWord] = struct{}{} // Добавление уникального слова в мапу

		normalized := normalizeWord(loweredWord)                         // Нормализация слова
		anagrams[normalized] = append(anagrams[normalized], loweredWord) // Добавление слова в группу анаграмм
	}

	// Формируем результат, исключая множества с одним элементом и сортируя слова в каждом множестве.
	result := make(map[string][]string)
	for _, group := range anagrams {
		if len(group) > 1 { // Исключение множеств с одним элементом
			sort.Strings(group)      // Сортировка слов в множестве по алфавиту
			result[group[0]] = group // Ключом является первое встретившееся слово в множестве
		}
	}

	return &result // Возвращаем ссылку на результирующую мапу
}

func main() {
	// Массив слов для поиска анаграмм
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "столик", "п", "п"}
	anagrams := findAnagrams(words)     // Вызов функции поиска анаграмм
	for key, group := range *anagrams { // Вывод результатов
		fmt.Printf("%s: %v\n", key, group)
	}
}
