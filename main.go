package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type totdo struct{}

func main() {
	for {

		fmt.Println("\n enter the process u want to pick!!!")
		i := 0
		fmt.Print(i)
		fmt.Scan(&i)

		switch i {
		case 1:
			var name, discription string
			fmt.Scan(&name)
			fmt.Scan(&discription)
			tim := time.Now().Format("2006-01-02 15:04:05")
			fmt.Printf("%T", tim)

			case1(name, discription, tim)

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

func case1(name, description, tim string) {
	value := stringcombiner6000(name, description, tim)
	filename := "activities.txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(value); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data appended successfully")
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

func stringcombiner6000(name, descrition, tim string) (str string) {
	str = "\n name:" + name + " " + "desription" + descrition + "setdate:" + tim + "\n"

	fmt.Println(str)
	return str
}
