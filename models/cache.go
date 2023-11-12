// models/cache.go
package models

import "time"

type CacheEntry struct {
    Response  []byte    `json:"response"`
    Timestamp time.Time `json:"timestamp"`
}