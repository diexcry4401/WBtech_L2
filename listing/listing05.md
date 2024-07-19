
# listing05

Что выведет программа? Объяснить вывод программы.

```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```

Ответ:

```
Вывод:
error

В Go, интерфейсы могут содержать значение и тип. Даже если значение интерфейса nil, тип может быть установлен.
В случае функции test, она возвращает nil типа *customError. 
Когда это значение присваивается переменной интерфейса err, интерфейс содержит nil значение, но с типом *customError.
Проверка if err != nil проверяет не только значение, но и тип интерфейса. 
Так как тип не является nil (это *customError), условие err != nil выводит true, и программа выводит "error".
```
