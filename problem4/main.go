package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func FindUser(users []User, id int) (User, error) {
	for _, v := range users {
		if v.Id == id {
			return v, nil
		}
	}

	return User{}, fmt.Errorf("User with id %d not fouund", id)
}

func main() {
	userlist := []User{
		{
			Id:   1,
			Name: "brian",
		}, {
			Id:   2,
			Name: "kyaw",
		},
	}

	user, err := FindUser(userlist, 1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found %+v \n", user)
	}

	user2, err := FindUser(userlist, 5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found %+v \n", user2)
	}
}
