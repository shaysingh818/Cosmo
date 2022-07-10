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


func (a *App) initApp(pdb string, setPort int){

	/* create level db instance */
	a.dbpath = pdb
	db, err := leveldb.OpenFile(a.dbpath, nil)
	if err != nil {
		panic(fmt.Sprintf("LevelDB open failed: %s", err))
	}
	defer db.Close()

	a.db = db
	a.timeout = 1
	a.peers = []string{"shay"}
	a.port = setPort
	a.router = mux.NewRouter()
	a.initRoutes()


}

func (a *App) Run(addr string){
	fmt.Println("Running kvs server on port %d", a.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.port), a.router))
}


func main(){

	/* parse command line args */
	port := flag.Int("port", 3000, "Port for main server to listen on")
	pdb := flag.String("db", "server", "path to leveldb")

	a := App{}
	a.initApp(*pdb, *port)
	a.Run("testing")

}
