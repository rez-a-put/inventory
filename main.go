package main

import (
	dbs "inventory/database"
	r "inventory/router"
)

func main() {
	dbs.ConnectMysql()
	r.InitRoutes()
}
