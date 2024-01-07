package helper

import "fmt"

func PanicIfError(err error) {
  fmt.Println(err)
  if err != nil {
    panic(err)
  }
}
