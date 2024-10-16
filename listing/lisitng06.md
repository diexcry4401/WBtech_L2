# listing06

Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передаче их в качестве аргументов функции.

```go
package main
 
import (
  "fmt"
)
 
func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}
 
func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}
```

Ответ:

```
Вывод:
[3 2 3]

Слайс в Go состоит из трех элементов:

Указатель на массив (array pointer): указывает на первый элемент массива, который хранит элементы слайса.
Длина (length): количество элементов в слайсе.
Вместимость (capacity): количество элементов, которые могут быть помещены в массив без выделения новой памяти.
Когда мы передаем слайс в функцию, передается копия этих трех элементов. 
Хотя сама структура передается по значению, она содержит указатель на базовый массив, 
поэтому изменения элементов слайса внутри функции отражаются на исходном слайсе, но до изменения его длины или емкости.
```

```
i[0] = "3"
Это изменение затрагивает оригинальный слайс s, поскольку слайс передается по ссылке. Теперь s выглядит так: ["3", "2", "3"].

i = append(i, "4")
Следующие изменения (после первого append) больше не влияют на оригинальный слайс, так как append вызывает выделение нового массива. Поэтому конечный вывод программы — ["3", "2", "3"].
```