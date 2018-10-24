package main

import (
	"fmt"
)

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

func(u user) notify () {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

func(u *user) changeEmail(email string) {
	u.email = email
}

func main() {


	bill := user{"bill","bill@email.com",10,true}

	fmt.Println(bill)

	lisa := &user{
		name:       "lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()


	fred := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level:"super",
	}

	fmt.Println(fred)

	type Duration int64

	//var dur Duration
	//dur = int64(1000)






}
