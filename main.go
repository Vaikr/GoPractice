package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("leveldb", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Put([]byte("user-1"), []byte("{\"username\":\"}"), nil)
	data, _ := db.Get([]byte("user-1"), nil)
	fmt.Println("data=", string(data))
}
