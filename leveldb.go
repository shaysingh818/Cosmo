package main

import (
	"fmt"
)



func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}


/* level db functions */
func (a *App) PutKey(key []byte, content []byte) bool {
    // csm
    a.mu.Lock()
    defer a.mu.Unlock()

    err := a.db.Put([]byte(key), content, nil)
    if err != nil {
        return false
    }
	fmt.Println("Wrote key value")
    fmt.Println("Stored file in db")
    return true
}


func (a *App) RetrieveKey(key []byte) []byte {
    //csm
    data, err := a.db.Get([]byte(key), nil)
    CheckError(err)
	return data
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


func (a *App) ViewKeys() []string {
	// retrieve all keys
	var keys []string
    iter := a.db.NewIterator(nil, nil)
    for iter.Next() {
        key := iter.Key()
		keys = append(keys, key)
    }

    iter.Release()
    err := iter.Error()
    if err != nil {
		fmt.Println("Something went wrong")
    }

    return keys
}

