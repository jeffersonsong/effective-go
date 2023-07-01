package main

import "fmt"

type Object struct {
	owner string
}

func (obj *Object) Owner() string {
	return obj.owner
}

func (obj *Object) SetOwner(user string) {
	obj.owner = user
}

func (obj *Object) String() string {
	return fmt.Sprintf("owner %s", obj.owner)
}

func main() {
	user := "jeff"
	obj := Object{"james"}
	fmt.Printf("obj: %v\n", obj)

	owner := obj.Owner()
	fmt.Printf("owner: %v\n", owner)
	if owner != user {
		obj.SetOwner(user)
	}
	fmt.Printf("owner: %v\n", owner)
	fmt.Printf("obj: %v\n", obj)
}
