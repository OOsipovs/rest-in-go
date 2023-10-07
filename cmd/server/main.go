package main

import "fmt"

//Rsponsible for our go app instantiation and start
func Run() error {
	fmt.Println("starting up our application")
	return nil
}

func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
