package main

import (
	"GoTask/work"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user2, err := user.Current()
	if nil == err {
		// fmt.Println(user2.HomeDir + "\\Documents\\eTaxSH3")
		fmt.Println(os.Chdir(user2.HomeDir + "\\Documents\\eTaxSH3"))
	}

	fmt.Println(os.Getenv("USERPROFILE"))

	ch := make(chan int)
	go work.AR3700BackDat()
	ch <- 1 // 写入chan
}
