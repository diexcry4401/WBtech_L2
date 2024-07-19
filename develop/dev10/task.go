package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"dev10/server"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type Client struct {
	Conn net.Conn // Соединение с сервером
}

// Dial устанавливает TCP-соединение с сервером по указанному адресу и таймауту.
func (c *Client) Dial(address string, timeout time.Duration) error {
	// Пытаемся установить соединение с сервером с указанным таймаутом.
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err // Возвращаем ошибку, если соединение не удалось установить.
	}
	c.Conn = conn // Сохраняем соединение в поле Conn клиента.
	return nil
}

// DoEcho отправляет сообщение на сервер и получает ответ.
func (c *Client) DoEcho(msg string) (string, error) {
	// Отправляем сообщение на сервер.
	if _, err := c.Conn.Write([]byte(msg)); err != nil {
		return "", err // Возвращаем ошибку, если не удалось отправить сообщение.
	}

	// Буфер для чтения ответа от сервера.
	buf := make([]byte, 1024)
	_, err := c.Conn.Read(buf) // Читаем ответ от сервера.
	if err != nil {
		if err != io.EOF {
			return "", err // Возвращаем ошибку, если чтение не удалось и это не EOF.
		}
	}

	fmt.Println()           // Печатаем пустую строку для разделения сообщений.
	return string(buf), nil // Возвращаем ответ от сервера.
}

// NewClient создает и возвращает новый экземпляр клиента.
func NewClient() *Client {
	return &Client{}
}

func main() {
	// Запускаем тестовый TCP сервер в отдельной горутине.
	go server.StartTCP()
	time.Sleep(time.Second) // Даем серверу время на запуск.

	// Обрабатываем аргументы командной строки.
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout") // Таймаут подключения.
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		// Если количество аргументов меньше двух, выводим сообщение о корректном использовании.
		fmt.Println("Usage: go-telnet [--timeout=<duration>] host port")
		os.Exit(1)
	}
	host := args[0] // Хост для подключения.
	port := args[1] // Порт для подключения.

	client := NewClient() // Создаем нового клиента.
	if err := client.Dial(fmt.Sprintf("%s:%s", host, port), *timeout); err != nil {
		// Если не удалось установить соединение, выводим сообщение об ошибке и завершаем программу.
		fmt.Printf("error dialing to a server[%s]: %s\n", port, err)
		os.Exit(1)
	}
	defer client.Conn.Close() // Обеспечиваем закрытие соединения при завершении работы функции main.

	// Создаем канал для сигналов о завершении работы.
	exit := make(chan struct{})
	go func() {
		defer client.Conn.Close()           // Обеспечиваем закрытие соединения при завершении работы горутины.
		reader := bufio.NewReader(os.Stdin) // Создаем читатель для STDIN.
		for {
			text, err := reader.ReadString('\n') // Читаем строку из STDIN.
			if err != nil {
				if err == io.EOF {
					// Если достигнут конец файла (Ctrl+D), выводим сообщение и отправляем сигнал о завершении.
					fmt.Println("Ctrl+D pressed. Closing connection...")
					exit <- struct{}{}
				} else {
					// Если произошла ошибка при чтении, выводим сообщение об ошибке.
					fmt.Println("Error reading from STDIN:", err)
				}
				break
			}
			// Отправляем сообщение на сервер и получаем ответ.
			resp, err := client.DoEcho(text)
			if err != nil {
				// Если произошла ошибка при отправке сообщения или чтении ответа, выводим сообщение об ошибке и завершаем работу.
				fmt.Printf("error [DoMessage]: %s\n", err)
				exit <- struct{}{}
			}
			fmt.Println(resp) // Выводим ответ от сервера.
		}
	}()

	// Ожидаем сигнала о завершении работы.
	<-exit
}
