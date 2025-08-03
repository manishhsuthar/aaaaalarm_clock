package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "alarm-clock-go/internal/alarm"
    "time"
)

var currentAlarm alarm.Alarm

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
    fmt.Fprintf(w, "Alarm set for %s", t.Format("15:04"))
}

func deactivateAlarmHandler(w http.ResponseWriter, r *http.Request) {
    currentAlarm.Deactivate()
    fmt.Fprint(w, "Alarm deactivated")
}

func main() {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)

    http.Handle("/", http.FileServer(http.Dir(filepath.Join(exPath, "web"))))
    http.HandleFunc("/set-alarm", setAlarmHandler)
    http.HandleFunc("/deactivate-alarm", deactivateAlarmHandler)

    fmt.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
