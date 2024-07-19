package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	goPs "github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// cd изменяет текущую рабочую директорию.
func cd(request []string) {
	// Проверяем количество аргументов.
	switch len(request) {
	case 1:
		// Не указан аргумент для директории.
		fmt.Fprintln(os.Stderr, "Error: missing directory argument.")
	case 2:
		// Меняем текущую рабочую директорию.
		err := os.Chdir(request[1])
		if err != nil {
			// Ошибка при смене директории.
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		// Слишком много аргументов.
		fmt.Fprintln(os.Stderr, "Error: too many arguments.")
	}
}

// pwd выводит путь к текущей рабочей директории.
func pwd(request []string) {
	// Проверяем, что аргументов нет.
	if len(request) == 1 {
		path, err := os.Getwd() // Получаем текущую рабочую директорию.
		if err != nil {
			// Ошибка при получении директории.
			fmt.Fprintln(os.Stderr, err)
		} else {
			// Выводим текущую рабочую директорию.
			fmt.Println(path)
		}
	} else {
		// Если аргументы присутствуют, выводим ошибку.
		fmt.Fprintln(os.Stderr, "Error: too many arguments.")
	}
}

// echo выводит аргументы в STDOUT.
func echo(request []string) {
	// Начинаем с первого аргумента, так как первый элемент - команда.
	for i := 1; i < len(request); i++ {
		fmt.Printf("%s ", request[i])
	}
	fmt.Println() // Перенос строки после вывода всех аргументов.
}

// kill отправляет сигнал завершения процессу по PID.
func kill(request []string) {
	// Проверяем количество аргументов.
	if len(request) == 1 {
		fmt.Fprintln(os.Stderr, "Error: missing PID argument.")
		return
	}
	if len(request) > 2 {
		fmt.Fprintln(os.Stderr, "Error: too many arguments.")
		return
	}

	// Преобразуем PID из строки в целое число.
	pid, err := strconv.Atoi(request[1])
	if err != nil {
		// Ошибка при преобразовании PID.
		fmt.Fprintln(os.Stderr, "Error: invalid PID.")
		return
	}

	// Находим процесс по PID.
	process, err := os.FindProcess(pid)
	if err != nil {
		// Ошибка при нахождении процесса.
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// Отправляем сигнал завершения процессу.
	err = process.Kill()
	if err != nil {
		// Ошибка при попытке завершения процесса.
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

// ps выводит информацию о запущенных процессах.
func ps(request []string) {
	// Проверяем, что аргументов нет.
	if len(request) != 1 {
		fmt.Fprintln(os.Stderr, "Error: too many arguments.")
		return
	}

	// Получаем список всех запущенных процессов.
	sliceProc, _ := goPs.Processes()

	// Выводим информацию о каждом процессе.
	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v, Process ID: %v\n", proc.Executable(), proc.Pid())
	}
}

// executeCommand выполняет команду в оболочке и поддерживает конвейеры (pipes).
func executeCommand(command string) {
	// Разделяем команду на отдельные команды по пайпу.
	cmds := strings.Split(command, "|")
	var previousCmd *exec.Cmd

	for i, cmdStr := range cmds {
		// Очищаем пробелы и создаем команду.
		cmd := exec.Command(strings.TrimSpace(cmdStr))

		// Если это не первая команда, перенаправляем ввод из предыдущей команды.
		if i > 0 {
			cmd.Stdin, _ = previousCmd.StdoutPipe()
		}

		// Если это не последняя команда, перенаправляем вывод в стандартный вывод.
		if i < len(cmds)-1 {
			cmd.Stdout = os.Stdout
		} else {
			// Последняя команда также выводит в стандартный вывод.
			cmd.Stdout = os.Stdout
		}

		// Запускаем команду.
		if err := cmd.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting command %v: %v\n", cmdStr, err)
			return
		}
		previousCmd = cmd
	}

	// Ожидаем завершения последней команды.
	if previousCmd != nil {
		previousCmd.Wait()
	}
}

// main запускает интерактивный командный интерпретатор.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the custom shell. Type 'quit' to exit.")

	for {
		// Выводим приглашение для пользователя.
		fmt.Print("$ ")

		// Считываем ввод пользователя.
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		// Если пользователь вводит 'quit', выходим из цикла.
		if line == "quit" {
			break
		}

		// Проверяем наличие пайпов в команде.
		if strings.Contains(line, "|") {
			executeCommand(line)
		} else {
			// Разделяем ввод пользователя на команду и аргументы.
			request := strings.Fields(line)
			switch request[0] {
			case "cd":
				cd(request)
			case "pwd":
				pwd(request)
			case "echo":
				echo(request)
			case "kill":
				kill(request)
			case "ps":
				ps(request)
			default:
				// Неизвестная команда.
				fmt.Fprintln(os.Stderr, "Error: command not found.")
			}
		}
	}
}
