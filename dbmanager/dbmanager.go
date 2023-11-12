// dbmanager/dbmanager.go
package dbmanager

import (
	"time"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB

const DbPath = "ipgeolocation.db"

func OpenDB() error {
    var err error
    db, err = bbolt.Open(DbPath, 0600, &bbolt.Options{Timeout: 1 * time.Second})
    return err
}

func CloseDB() error {
    if db != nil {
        return db.Close()
    }
    return nil
}

func GetDB() *bbolt.DB {
    return db
}