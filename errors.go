package main

import (
  "errors"
  "fmt"
)

func main() {

  message := "Hey"

  ///function to create a new error
  myError := errors.New("WRONG MESSAGE OUTPUT")

if message != "Joy Kiboi" {
  fmt.Println(myError)
}
  
}