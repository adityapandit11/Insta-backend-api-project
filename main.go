package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const (
	databaseURL = "mongodb://127.0.0.1:27017/project/start?compressors=disabled&gssapiServiceName=mongodb"
	port        = 3000
)

func main() {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("could not open db connection: %v\n", err)
		return
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("could not ping to db: %v\n", err)
	}

	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit")
	s := &service.New(db, codec)
	h := handler.New(s)

	addr := fmt.Sprintf(":%d", port)
	log.Print("accepting connection on port %d\n", port)
	if err = http.ListenAndServe(addr, h); err != nil {
		log.Fatal("could not start server: %v\n", err)
	}
}
