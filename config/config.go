package config

import (
	"fmt"
	"os"
	"time"

	"github.com/IPGeolocation/ipgeo/dbmanager"
	"go.etcd.io/bbolt"
	"golang.org/x/term"
)

const ConfigBucket = "Config"

func HandleConfig(args []string) {
    if len(args) < 1 {
        fmt.Println("Usage: ipgeo config <login|logout>")
        fmt.Println("  login  - Log in and save API key for the session.")
        fmt.Println("  logout - Log out and clear the saved API key.")
        return
    } else if len(args) > 1 {
        fmt.Println("Invalid config command. Use 'login' or 'logout'. Or use 'ipgeo config' to see help.")
        return
    }
    switch args[0] {

    case "login":
        fmt.Println("🔑 Logging in to IP Geolocation API...")
        fmt.Println("ℹ️  You can find your API key at https://ipgeolocation.io/a/account")

        apiKey, err := GetAPIKey()

        if err == nil {
            fmt.Printf("⚠️ One API key already exists: %s\n", apiKey)
            fmt.Println("Do you want to overwrite it? (y/n)")
            var response string
            fmt.Scanln(&response)
            if response != "y" {
                fmt.Println("❌ Login cancelled. Already saved API key will be used.")
                return
            }
        }
        fmt.Println("💡 Please enter your API key to log in:")
        bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading API key: %s\n", err)
            return
        }
        apiKey = string(bytePassword)
        if apiKey == "" {
            fmt.Println("❌ API key cannot be empty.")
            return
        }
        login(apiKey)
    case "logout":
        logout()
    default:
        fmt.Println("Invalid config command. Use 'login' or 'logout'. Or use 'ipgeo config' to see help.")
    }
}

func login(apiKey string) {
    db := dbmanager.GetDB()

    updateErr := db.Update(func(tx *bbolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists([]byte(ConfigBucket))
        if err != nil {
            return err
        }
        if err = b.Put([]byte("APIKey"), []byte(apiKey)); err != nil {
            return err
        }
        return b.Put([]byte("LoginTime"), []byte(time.Now().Format(time.RFC3339)))
    })

    if updateErr != nil {
        fmt.Fprintf(os.Stderr, "Error saving API key: %s\n", updateErr)
        return
    }

    fmt.Println("\n✅ API Key has been saved successfully!")
}

func logout() {
    db := dbmanager.GetDB() 

    var loginTime time.Time
    var err error 
    err = db.View(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte(ConfigBucket))
        if b == nil {
            return fmt.Errorf("You need to log in first.")
        }
        loginTimeBytes := b.Get([]byte("LoginTime"))
        if loginTimeBytes == nil {
            return fmt.Errorf("You need to log in first.")
        }
        var parseErr error
        loginTime, parseErr = time.Parse(time.RFC3339, string(loginTimeBytes))
        return parseErr
    })

    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        return
    }

    logoutTime := time.Now()
    err = db.Update(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte(ConfigBucket))
        if b == nil {
            return fmt.Errorf("Configuration bucket not found.")
        }
        if err := b.Delete([]byte("APIKey")); err != nil {
            return err
        }
        return b.Delete([]byte("LoginTime"))
    })

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error logging out: %s\n", err)
        return
    }
    fmt.Println("🔑 Logging out...")
    fmt.Printf("🔒 API Key cleared. You've been logged out successfully.\n")
    fmt.Printf("You logged in at: %s and logged out at: %s. Total Session time: %s.\n", 
        loginTime.Format("Jan 2, 2006 03:04 PM"), 
        logoutTime.Format("Jan 2, 2006 03:04 PM"), 
        formatDuration(logoutTime.Sub(loginTime)))
    fmt.Println("Thank you for using IP Geolocation CLI! Please provide your valuable feedback at https://ipgeolocation.io/contact-us.html")
}

func formatDuration(d time.Duration) string {
    totalMinutes := int(d.Minutes())
    hours := totalMinutes / 60
    minutes := totalMinutes % 60
    return fmt.Sprintf("%02d:%02d hours", hours, minutes)
}


func GetAPIKey() (string, error) {
    var apiKey string
    db := dbmanager.GetDB() 

    err := db.View(func(tx *bbolt.Tx) error {
        b := tx.Bucket([]byte(ConfigBucket))
        if b == nil {
            return fmt.Errorf("You need to log in first. Please use the following command to login: ipgeo config login")
        }
        apiKeyBytes := b.Get([]byte("APIKey"))
        if apiKeyBytes == nil {
            return fmt.Errorf("API key not found, please log in again.")
        }
        apiKey = string(apiKeyBytes)
        return nil
    })

    return apiKey, err
}

