
# listing04

Что выведет программа? Объяснить вывод программы.

```go
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}
```

Ответ: 

```
Вывод:
0 1 2 3 4 5 6 7 8 9 fatal error: all goroutines are asleep

У нас две горутины, main и анонимная объявленная в main. 
Канал ch используется для передачи данных между горутинами. 
Одна горутина отправляет данные в канал, а main горутина читает из него. 
Канал ch никогда не закрывается, поэтому главная горутина будет ждать новых данных после отправки всех 10 значений, 
что приводит к дедлоку.
```
