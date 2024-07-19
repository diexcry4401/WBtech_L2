package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/pborman/getopt"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// getURLandFilename извлекает URL и имя файла из аргументов командной строки.
// Возвращает URL и имя файла, которые можно использовать для загрузки и сохранения данных.
func getURLandFilename() (string, string) {
	// Получаем значение URL из аргументов командной строки.
	urlPath := getopt.StringLong("url", 'u', "", "URL to download")
	getopt.Parse() // Парсим аргументы командной строки.

	// Проверяем, является ли URL корректным.
	_, err := url.Parse(*urlPath)
	if err != nil {
		log.Fatal("Invalid URL:", err) // Завершаем выполнение программы, если URL некорректен.
	}

	// Разделяем URL на части, чтобы извлечь имя файла из последней части пути.
	splitedURL := strings.Split(*urlPath, "/")
	filename := splitedURL[len(splitedURL)-1]
	return *urlPath, filename // Возвращаем URL и имя файла.
}

// createFile создает новый файл для записи данных и возвращает указатель на него.
// Если файл не удается создать, программа завершится с ошибкой.
func createFile(filename string) *os.File {
	fmt.Printf("Creating file: %s\n", filename) // Выводим имя создаваемого файла.

	// Создаем файл для записи. Если файл уже существует, он будет перезаписан.
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Error creating file:", err) // Завершаем выполнение программы в случае ошибки.
	}
	return file // Возвращаем дескриптор созданного файла.
}

// getData загружает данные по указанному URL и сохраняет их в файл.
// Возвращает размер загруженных данных в байтах.
func getData(urlPath string, client *http.Client, file *os.File) int64 {
	// Выполняем HTTP GET запрос для получения данных по указанному URL.
	resp, err := client.Get(urlPath)
	if err != nil {
		log.Fatal("Error fetching URL:", err) // Завершаем выполнение программы, если запрос не удался.
	}
	defer resp.Body.Close() // Закрываем тело ответа после завершения функции.

	// Копируем данные из ответа HTTP запроса в созданный файл.
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal("Error saving data to file:", err) // Завершаем выполнение программы в случае ошибки записи.
	}
	return size // Возвращаем размер загруженных данных.
}

func main() {
	// Получаем URL и имя файла для сохранения данных.
	urlPath, filename := getURLandFilename()

	// Создаем файл для записи данных с полученным именем.
	file := createFile(filename)

	// Создаем HTTP клиент с функцией обработки перенаправлений.
	// Пропускаем перенаправления, чтобы скачать конечный файл по исходному URL.
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path // Устанавливаем только путь, игнорируя другие части URL.
			return nil
		},
	}

	// Загружаем данные по указанному URL и сохраняем их в созданный файл.
	size := getData(urlPath, &client, file)

	// Выводим сообщение о завершении загрузки, указывая URL и размер загруженного файла.
	fmt.Printf("Downloaded file %s with size %d bytes\n", urlPath, size)
}
