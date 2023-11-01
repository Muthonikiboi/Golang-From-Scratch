package main
import "fmt"

func countDown(number int) {

  if number > 0 {
    fmt.Println(number)

    countDown(number - 1)
  } else {
    fmt.Println("Countdown Stops")
  }

 }

func main() {
  countDown(3)
}