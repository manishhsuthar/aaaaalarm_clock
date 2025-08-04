package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	mathRand "math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"alarm-clock-go/internal/alarm"
	"alarm-clock-go/internal/utils"
)

var currentAlarm alarm.Alarm
var quotes = []string{
	"The best way to predict the future is to create it.",
	"The only way to do great work is to love what you do.",
	"Success is not final, failure is not fatal: it is the courage to continue that counts.",
	"Believe you can and you're halfway there.",
	"The future belongs to those who believe in the beauty of their dreams.",
}
var secret string

func generateSecret() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func setAlarmHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Time string `json:"time"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := time.Parse("15:04", data.Time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	currentAlarm.SetTime(t)
	currentAlarm.Activate()

	secret, err = generateSecret()
	if err != nil {
		http.Error(w, "Failed to generate secret", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Secret string `json:"secret"`
	}{Secret: secret})
}

func deactivateAlarmHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Secret string `json:"secret"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.Secret == secret {
		currentAlarm.Deactivate()
		json.NewEncoder(w).Encode(struct {
			Success bool `json:"success"`
		}{Success: true})
	} else {
		json.NewEncoder(w).Encode(struct {
			Success bool `json:"success"`
		}{Success: false})
	}
}

func checkAlarmHandler(w http.ResponseWriter, r *http.Request) {
	trigger := false
	if currentAlarm.IsActive && time.Now().Hour() == currentAlarm.Time.Hour() && time.Now().Minute() == currentAlarm.Time.Minute() {
		trigger = true
	}

	json.NewEncoder(w).Encode(struct {
		Trigger bool `json:"trigger"`
	}{Trigger: trigger})
}

func getQuoteHandler(w http.ResponseWriter, r *http.Request) {
	mathRand.Seed(time.Now().UnixNano())
	quote := quotes[mathRand.Intn(len(quotes))]
	json.NewEncoder(w).Encode(struct {
		Quote string `json:"quote"`
	}{Quote: quote})
}

func deactivateAlarmURLHandler(w http.ResponseWriter, r *http.Request) {
	querySecret := r.URL.Query().Get("secret")
	if querySecret == secret {
		currentAlarm.Deactivate()
		fmt.Fprintf(w, "Alarm deactivated successfully!")
	} else {
		http.Error(w, "Invalid secret", http.StatusUnauthorized)
	}
}

func main() {
	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	http.Handle("/", http.FileServer(http.Dir(filepath.Join(exPath, "web"))))
	http.HandleFunc("/set-alarm", setAlarmHandler)
	http.HandleFunc("/deactivate-alarm", deactivateAlarmHandler)
	http.HandleFunc("/check-alarm", checkAlarmHandler)
	http.HandleFunc("/get-quote", getQuoteHandler)
	http.HandleFunc("/deactivate", deactivateAlarmURLHandler)

	localIP, err := utils.GetLocalIP()
	if err != nil {
		fmt.Printf("Could not determine local IP address: %v\n", err)
		localIP = "localhost"
	}

	fmt.Printf("Server starting on port %s...\n", *port)
	fmt.Printf("Access it at http://%s:%s\n", localIP, *port)
	if err := http.ListenAndServe("0.0.0.0:"+*port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
