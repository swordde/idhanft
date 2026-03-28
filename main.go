package main

import (
	"fmt"
	"log"
	"os"
)

type totdo struct{}

func main() {
	for true {

		fmt.Println("enter the process u want to pick!!!")
		i := 0
		fmt.Scan(&i)

		switch i {
		case 1:
			var name, discription string
			fmt.Scan(&name)
			fmt.Scan(&discription)
			case1(name, discription)

		case 2:
			case2()

		case 3:
			fmt.Println("case 2")

		case 4:
			fmt.Println("case 2")

		case 5:
			fmt.Println("case 2")
		}
	}
	fmt.Print("")
}

func case1(name, description string) {
	_, err := os.Open("activities.txt") // For read access.
	if err != nil {
		log.Fatal(err)

		fmt.Print("the errorr here")
	}
	fmt.Println("process is set", name, description)

	bytes := []byte(description)
	err = os.WriteFile("activities.txt", bytes, 0o666)
	if err != nil {
		fmt.Print("")
		log.Fatal(err)
	}
}

func case2() {
	file, err := os.Open("activities.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" all the tasks are \n  %q \n ", data[:count])
}
