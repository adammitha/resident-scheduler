package main

import (
	"fmt"

	"github.com/adammitha/resident-scheduler/scheduler"
)

func main() {
	fmt.Println(scheduler.Residents("residents.csv"))
}
