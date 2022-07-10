package main

import (
	"fmt"
	"sync"
	"github.com/syndtr/goleveldb/leveldb"
)

/* level db functions */
func (a *App) PutKey(key []byte, content []byte) bool {
    // csm
    a.mu.Lock()
    defer a.mu.Unlock()

    err := a.db.Put([]byte(key), content, nil)
    if err != nil {
        return false
    }
    fmt.Println("Stored file in db")
    return true
}


func (a *App) RetrieveKey(key []byte) {
    //csm
    a.mu.Lock()
    defer a.mu.Unlock()

    data, err := a.db.Get([]byte(key), nil)
    CheckError(err)
    fmt.Println(string(data))
}


func (a *App) DeleteKey(key []byte) bool {
    //csm
    a.mu.Lock()
    defer a.mu.Unlock()

    err := a.db.Delete([]byte(key), nil)
    if err != nil {
        return false
    }
    return true
}


func (a *App) ViewKeys() bool {

    iter := a.db.NewIterator(nil, nil)
    for iter.Next() {
        key := iter.Key()
        value := iter.Value()
        fmt.Println("%v : %v", key, value)
    }

    iter.Release()
    err := iter.Error()
    if err != nil {
        return false
    }

    return true
}

