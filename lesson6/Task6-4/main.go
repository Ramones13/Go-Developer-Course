// Необходимо убрать повторяющийся код - поля Addss и Phone из структур:
// type user struct {
// ID int
// Name string
// Addss string
// Phone string
// }
// type employee struct {
// ID int
// Name string
// Addss string
// Phone string
// }
//
// После проведения рефакторинга строка fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone) должна выводить в консоль
// адрес и телефон пользователя и сотрудника соответственно

package main

import "fmt"

type contact struct {
	Addss string
	Phone string
}

type user struct {
	ID   int
	Name string
	contact
}

type employee struct {
	ID   int
	Name string
	contact
}

func main() {
	u := user{
		ID:   1,
		Name: "Пользователь 1",
		contact: contact{
			Addss: "ул. Косманавтов",
			Phone: "+7999xxxxxxx",
		},
	}

	e := employee{
		ID:   1,
		Name: "Сотрудник 2",
		contact: contact{
			Addss: "ул. Свободы",
			Phone: "+7666xxxxxxx",
		},
	}

	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
