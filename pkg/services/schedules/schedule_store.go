package schedules

type ScheduleStore interface {
	NewSchedule(Schedule) error
	UpdateSchedule(ScheduleID, Schedule) error
	DeleteSchedule(ScheduleID) error
	NewEvent(Event) error
	UpdateEvent(EventID, Event) error
	DeleteEvent(ScheduleID, EventID) error
}
