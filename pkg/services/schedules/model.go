package schedules

import (
	"time"

	"github.com/schedlinkk/server/pkg/services/users"
)

type ScheduleID string
type EventID string

type Schedule struct {
	ID         ScheduleID  `json:"id"`
	UID        users.UID   `json:"user_id"`
	Events     []Event     `json:"events"`
	SharedWith []users.UID `json:"shared_with"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type Event struct {
	ID          EventID    `json:"id"`
	ScheduleID  ScheduleID `json:"schedule_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartTime   time.Time  `json:"start_time"`
	EndTime     time.Time  `json:"end_time"`
}
