package main

import (
	"github.com/ko44d/hrmos-time-aggregator/pkg/di"
	"log"
)

func main() {

	di := di.NewDI()
	eoc := di.EmployeeOvertimeController()
	tpc := di.TopPageController()
	r := SetupRouter(tpc, eoc)

	log.Println("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
