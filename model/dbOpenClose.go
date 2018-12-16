
package model

import(
		"github.com/boltdb/bolt"
		"runtime"
		"path"
		"log"
		"fmt"
	)

var db * bolt.DB
var open bool

func Open() error {
    var err error
    _, filename, _, _ := runtime.Caller(0)  // get full path of this file
	dbfile := path.Join(path.Dir(filename), "my.db")
	fmt.Println(dbfile)
    //config := &bolt.Options{Timeout: 1 * time.Second}
    db, err = bolt.Open(dbfile, 0600, nil)
    if err != nil {
        log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error{
		tx.CreateBucketIfNotExists([]byte("testBucket"))
		return nil
	})
    open = true
    return nil
}

func Close() {
    open = false
    db.Close()
}