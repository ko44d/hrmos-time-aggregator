package main

import (
	"github.com/ko44d/hrmos-time-aggregator/pkg/di"
)

func main() {

	di := di.NewDI()
	eoc := di.EmployeeOvertimeController()
	err := eoc.Aggregate()
	if err != nil {
		panic(err)
	}

}
