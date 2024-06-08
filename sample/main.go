package main

import (
	"fmt"
	"log"
	"sample/core"
	"sample/db"
	"sample/utils"
)

func main() {
	// _ = pagi.Paginate[interface{}]()
	// paging

	param := utils.ProvideFiberHttpServiceParams()
	fiberSrv := utils.InitializeHTTPService(param)

	database, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Failed to start Database:", err)
	}
	db.DB = database

	repo := core.NewRepository(db.DB)
	srv := core.NewServices(repo)
	handler := core.NewHandlers(srv)

	fiberSrv.Get("/", handler.HandleGetTodoes)

	portString := fmt.Sprintf(":%v", param.Port)
	err = fiberSrv.Listen(portString)

	if err != nil {
		log.Fatal("Failed to start golang Fiber server:", err)
	}

}
