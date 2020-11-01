package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {

	userblacklist := func(name string) bool {
		return name == "root"
	}
	registerUser("root", userblacklist)

}
