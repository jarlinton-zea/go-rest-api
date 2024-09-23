package main

import "fmt"

// Instantiate and startup the go app
func Run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {
	fmt.Println("Our First Go RESt API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
