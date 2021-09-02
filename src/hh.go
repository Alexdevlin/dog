package main

import (
	"fmt"
)

/**/
func main() {
	a, _, _ := GetData(1)
	_, b, c := GetData(1)
	fmt.Println(a, b, c)

}

func GetData(a int) (int, string, int) {
	return a, "dasd", 500
}
