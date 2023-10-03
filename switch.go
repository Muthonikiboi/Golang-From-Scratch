package main
import "fmt"

func main(){
	//Switch case example 1
	day := 5

	switch day{
		case 1:
			fmt.Println("Sunday")

		case 2:
			fmt.Println("Monday")

		case 3:
			fmt.Println("Tuesday")

		case 4:
			fmt.Println("Wednesday")

		case 5:
			fmt.Println("Thursday")

		case 6:
			fmt.Println("Friday")

		case 7:
			fmt.Println("Saturday")

		default:
			fmt.Println("Invadid Day")
	}

	//Example 2 Switch statement:multiple
	dayIs:="Friday"

	switch dayIs{
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	case "Saturday","Sunday":
		fmt.Println("Weekend")

	default:
		fmt.Println("Invalid day")
	}

	//with no argument
	func
}