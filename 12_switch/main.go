package main

import "fmt"

func main() {
	fmt.Println("# Switch")
	day := dayOfWeek(1)
	fmt.Println(day)

	day2 := dayOfWeekTwo(7)
	fmt.Println(day2)
}

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number."
	}
}

func dayOfWeekTwo(number int) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "Monday"
	case number == 3:
		return "Tuesday"
	case number == 4:
		return "Wednesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Friday"
	case number == 7:
		return "Saturday"
	default:
		return "Invalid number."
	}
}
