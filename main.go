package main

import (
	"fmt"
	"github.com/ngenohkevin/poc-ark/db"
)

func main() {
	fmt.Println("This is a proof of concept for ark realtors")
	db.ConnectDb()
}
