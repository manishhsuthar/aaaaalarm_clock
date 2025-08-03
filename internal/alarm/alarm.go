package alarm

import "time"

type Alarm struct {
    Time     time.Time
    IsActive bool
}

func (a *Alarm) SetTime(t time.Time) {
    a.Time = t
}

func (a *Alarm) Activate() {
    a.IsActive = true
}

func (a *Alarm) Deactivate() {
    a.IsActive = false
}