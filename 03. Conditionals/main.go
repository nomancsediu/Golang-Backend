package main

import "fmt"

func main() {
    score := 75

    // if / else if / else
    if score >= 80{
		fmt.Println("A+")
	}else if score>=75{
		fmt.Println("A")
	}else if score>=70{
		println("A-")
	}else if score>=40{
		println("D")
	}else
	{
		fmt.Println("F")
	}

	//switch - cleaner for many conditions
	day := "Monday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend!")
	case "Monday":
		fmt.Println("Bad day :(")
	default:
		fmt.Println("Weekday")
	}

}