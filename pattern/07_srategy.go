package pattern

import "fmt"

// Strategy - интерфейс, определяющий метод для сортировки.
type Strategy interface {
	Sort([]int)
}

// BubbleSort - конкретная стратегия для пузырьковой сортировки.
type BubbleSort struct{}

func (b *BubbleSort) Sort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("Sorted using Bubble Sort:", arr)
}

// InsertionSort - конкретная стратегия для сортировки вставками.
type InsertionSort struct{}

func (i *InsertionSort) Sort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
	fmt.Println("Sorted using Insertion Sort:", arr)
}

// Context - контекст, использующий стратегию сортировки.
type Context struct {
	strategy Strategy
}

// SetStrategy - метод для установки конкретной стратегии.
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// ExecuteStrategy - метод для выполнения сортировки с использованием установленной стратегии.
func (c *Context) ExecuteStrategy(arr []int) {
	c.strategy.Sort(arr)
}

// func main() {
// 	// Исходный массив данных.
// 	data := []int{5, 2, 9, 1, 5, 6}

// 	// Создаем контекст.
// 	context := &Context{}

// 	// Устанавливаем стратегию пузырьковой сортировки и выполняем сортировку.
// 	context.SetStrategy(&BubbleSort{})
// 	context.ExecuteStrategy(data)

// 	// Восстанавливаем исходный массив данных.
// 	data = []int{5, 2, 9, 1, 5, 6}

// 	// Устанавливаем стратегию сортировки вставками и выполняем сортировку.
// 	context.SetStrategy(&InsertionSort{})
// 	context.ExecuteStrategy(data)
// }
