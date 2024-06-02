package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}
type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	return d.voice
}
func main() {
	var d *Duck
	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}
func Sing(b Bird) (string, error) {
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}

/*
Паника возникает потому что в d не равно nil. интерфейcы содержат тип помимо значения.
то есть в данном примере в d содержится тип *main.Duck со значением nil

Решения:

1. уйти от ссылочных типов. тогда в значении будет значение по умолчанию.
вариант не очень хороший. паники не возникнет, но и птица петь не будет если её не заполнить

2. Проверять в на nil в функции Sing для интерфейса.
плохой вариант - избыточные проверки.
потому что сам объект не должен быть ответственным за то что в него передали не верные значения

3. Положить в переменную какое-то значение.
Писать код аккуратно и всегда следить чтобы интерфейсы были инициализированы

*/
