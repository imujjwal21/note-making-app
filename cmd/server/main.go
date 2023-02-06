package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imujjwal21/note-making-app/internal/httptransport"
	"github.com/imujjwal21/note-making-app/notes"
	"github.com/imujjwal21/note-making-app/sqldb"
)

func main() {

	db := sqldb.ConnectDB()
	var port int

	flag.IntVar(&port, "port", 0, "Address to bind the socket on.")

	flag.Parse()

	server := &http.Server{Handler: httptransport.NewHandler(notes.NewInMemoryStorage(db))} // server is type of &http.Server

	go func() {

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port)) // It is only for listen like http.ListenAndServe(":8080", r)

		if err != nil {
			log.Panicf("cannot create tpc listener: %v", err)
		}

		log.Printf("      starting http server on %q", lis.Addr())
		if err := server.Serve(lis); err != nil { // It is only for Serve like http.ListenAndServe(":8080", r)
			log.Panicf("cannot start http server: %v", err)
		}

		// err := server.Serve(lis)
		// if err != nil {
		// 	log.Panicf("cannot start http server: %v", err)
		// }
	}()

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	log.Printf("Got exit signal %q. Bye", <-sig)
}
