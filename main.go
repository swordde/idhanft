package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type totdo struct{}

func main() {
	case2()

	//	fmt.Print("file printeed11111")
	for {

		fmt.Println("\n enter the process u want to pick!!!")
		i := 0
		fmt.Print(i)
		fmt.Scan(&i)

		switch i {
		case 1:

			case1()

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
}

func case1() {
	var name string
	fmt.Scan(&name)
	bav := bufio.NewReader(os.Stdin)
	discription, _ := bav.ReadString('\n')
	tim := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%T", tim)

	value := stringcombiner6000(name, tim, discription)
	filename := "activities.txt"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(value); err != nil {
		fmt.Println("Error write to file:", err)
		return
	}
	fmt.Println("Data appended successfully")
}

func case2() {
	file, err := os.Open("activities.txt") // For read access.
	//	fmt.Print("file printeed123123")
	if err != nil {
		//		fmt.Print("this printed here in the error for file")
		log.Fatal(err)
	}

	/*
	   	//	fmt.Println("file printed here before the data")
	   	data := make([]byte, 1000)

	   	//	fmt.Println("file printed here after the data")
	   	count, err := file.Read(data)
	   	//	fmt.Println("file printed here after the count")
	   	if err != nil {
	   		fmt.Print("the todo is empty!!!")
	   	}

	   $~$	fmt.Printf(" all the tasks are \n  %q \n ", data[:count])
	*/
	scanner := bufio.NewScanner(file)
	scanner.Err()
	for scanner.Scan() {

		line := scanner.Text()
		fmt.Println(line)
	}
}

func stringcombiner6000(name, tim, description string) (str string) {
	str = "\n name:" + name + " " + "desription:" + description + " " + "setdate:" + tim + "\n"

	fmt.Println(str)
	return str
}
