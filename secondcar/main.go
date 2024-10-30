package main

import (
	"tce/init"

	//"os"
	"tce/CollectRoute"
	//"tce/sdkInit"
	//"tce/service"
	//"time"
)

func main() {
	r, db := Init.InitGinGrom()
	//setup := Init.InitFabric()
	r = CollectRoute.CollectRoute(r, db)
	if r.Run(":8000") != nil {
		return
	}
}
