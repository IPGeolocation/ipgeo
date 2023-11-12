// cache/cache.go
package cache

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

    "github.com/IPGeolocation/ipgeo/models"
    "github.com/IPGeolocation/ipgeo/dbmanager"
	"go.etcd.io/bbolt"
)

const (
    CacheBucket = "Cache"
    ConfigBucket = "Config"
    CacheEnabledKey = "true"
)

func HandleCacheCommand(args []string) {
	    db := dbmanager.GetDB()

    if len(args) == 0 {
        fmt.Println("Usage: ipgeo cache [subcommand]")
        fmt.Println("Subcommands:")
        fmt.Println("  clear - Clears the cache.")
        fmt.Println("  count - Displays the number of entries in the cache.")
        fmt.Println("  enable - Enables caching.")
        fmt.Println("  disable - Disables caching.")
        fmt.Println("  status - Displays the current status of caching.")
        return
    } else if len(args) > 1 {
        fmt.Println("Invalid cache subcommand. Available subcommands: clear, count, enable, disable")
        return
    }

    switch args[0] {
    case "clear":
        err := ClearCache(db)
        if err != nil {
            if err == bbolt.ErrBucketNotFound {
                fmt.Println("Cache is already empty.")
                return
            } else{
                fmt.Println("Error clearing cache:", err)
            }
        } else {
            fmt.Println("Cache cleared successfully.")
        }

    case "count":
        count, err := CountCacheEntries(db)
        if err != nil {
            fmt.Println("Error counting cache entries:", err)
        } else {
            fmt.Printf("There are %d entries in the cache.\n", count)
        }

    case "enable", "disable":
        enable := strings.ToLower(args[0]) == "enable"
        err := SetCacheEnabled(enable)
        if err != nil {
            fmt.Printf("Error setting cache %s: %s\n", args[0], err)
        } else {
            fmt.Printf("Cache %s successfully.\n", args[0])
        }
    case "status":
    enabled, err := GetCacheStatus()
    if err != nil {
        fmt.Println("Error checking cache status:", err)
    } else {
        status := "disabled"
        if enabled {
            status = "enabled"
        }
        fmt.Printf("Cache is currently %s.\n", status)
    }

    default:
        fmt.Println("Invalid cache subcommand. Available subcommands: clear, count, enable, disable")
    }
}

func SaveToCache(key string, response []byte) error {
	    db := dbmanager.GetDB()
    return db.Update(func(tx *bbolt.Tx) error {
        b, _ := tx.CreateBucketIfNotExists([]byte(CacheBucket))
        cacheEntry := models.CacheEntry{Response: response, Timestamp: time.Now()}
        encoded, err := json.Marshal(cacheEntry)
        if err != nil {
            return err
        }
        return b.Put([]byte(key), encoded)
    })
}

func GetFromCache(key string) ([]byte, bool, error) {
		    db := dbmanager.GetDB()
    var cacheEntry models.CacheEntry
    found := false
    err := db.View(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte(CacheBucket))
        if b == nil {
            return nil
        }
        cachedData := b.Get([]byte(key))
        if cachedData == nil {
            return nil
        }
        err := json.Unmarshal(cachedData, &cacheEntry)
        if err != nil {
            return err
        }
        if time.Since(cacheEntry.Timestamp) > 10*time.Minute {
            return b.Delete([]byte(key))
        }
        found = true
        return nil
    })
    return cacheEntry.Response, found, err
}

func ClearCache(db *bbolt.DB) error {
    return db.Update(func(tx *bbolt.Tx) error {
        return tx.DeleteBucket([]byte(CacheBucket))
    })
}

func CountCacheEntries(db *bbolt.DB) (int, error) {
    count := 0
    err := db.View(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte(CacheBucket))
        if b == nil {
            return nil
        }
        return b.ForEach(func(k, v []byte) error {
            count++
            return nil
        })
    })
    return count, err
}

func SetCacheEnabled(enabled bool) error {
    db := dbmanager.GetDB()
    return db.Update(func(tx *bbolt.Tx) error {
        b, _ := tx.CreateBucketIfNotExists([]byte(ConfigBucket))
        val := "false"
        if enabled {
            val = "true"
        }
        return b.Put([]byte(CacheEnabledKey), []byte(val))
    })
}
func GetCacheStatus() (bool, error) {
    db := dbmanager.GetDB()
    enabled := true
    err := db.View(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte(ConfigBucket))
        if b == nil {
            return nil
        }
        val := b.Get([]byte(CacheEnabledKey))
        if val != nil && string(val) == "false" {
            enabled = false
        }
        return nil
    })
    return enabled, err
}
