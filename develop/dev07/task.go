package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// or объединяет несколько каналов и возвращает новый канал, который будет закрыт,
// как только один из переданных каналов закроется.
func or(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{}) // создаем новый канал, который будет возвращен
	for _, val := range channels {
		go func(ch <-chan interface{}) {
			<-ch       // блокируемся до получения из канала
			close(res) // закрываем результирующий канал, когда канал ch закрывается
		}(val)
	}
	return res // возвращаем результирующий канал
}

func main() {
	// sig возвращает канал, который будет закрыт после задержки.
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c) // гарантирует закрытие канала после завершения работы горутины
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	) // ждем, пока один из каналов не закроется

	fmt.Printf("done after %v", time.Since(start)) // выводим время, прошедшее с начала ожидания
}
