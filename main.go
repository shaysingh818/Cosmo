package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
	"log"
	"net/http"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/gorilla/mux"
)



type App struct {
	dbpath string
	db *leveldb.DB
	mu sync.Mutex
	router *mux.Router
	port int
	timeout time.Duration
	peers []string
}


func (a *App) Run(addr string){
	fmt.Println("Running kvs server on port %d", a.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.router))
}


func main(){

	/* parse command line args */
	port := flag.Int("port", 3000, "Port for main server to listen on")
	pdb := flag.String("db", "server", "path to leveldb")
	flag.Parse()

	/* create app instance here */
	db, err := leveldb.OpenFile(*pdb, nil)
	if err != nil {
		panic(fmt.Sprintf("LevelDB open failed: %s", err))
	}
	defer db.Close()

	a := App{}
	a.db = db
	a.dbpath = *pdb
	a.timeout = 1
	a.peers = []string{"shay"}
	a.port = *port
	a.router = mux.NewRouter()
	a.initRoutes()

	a.Run("testing")

}
