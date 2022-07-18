package main

import (
	"log"
	"net/http"

	"github.com/AbhinavAchha/golang-tutorial/router"
)

func main() {
	var r = router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
