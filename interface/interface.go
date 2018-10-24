package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify(){
	fmt.Printf("sending user emial to %s<%s>\n", u.name, u.email)
}


func sendNotification(n notifier){
	n.notify()
}

type admin struct {
	user
	level string
}

func (a *admin) notify() {
	fmt.Printf("sending admin emial to %s<%s>\n", a.name, a.email)
}




func main() {

	ad := admin {
		user: user {
			name: "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	ad.user.notify()
	ad.notify()
	sendNotification(&ad)

	//bill := user{"bill", "bill@email.com"}
	//
	//sendNotification(&bill)
	//
	//lisa := admin{"Lisa", "lisa@email.com"}
	//
	//sendNotification(&lisa)




}
