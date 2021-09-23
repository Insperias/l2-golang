package domain

import "time"

//CalendarInterface Interface of calendar
type CalendarInterface interface {
	GetEvents() ([]*Event, error)
	GetEventByID(id string) (*Event, error)
	AddEvent(title string, from, to *time.Time) (*Event, error)
	UpdateEvent(event *Event) error
	RemoveEvent(event *Event) error
}

//Calendar Calendar struct
type Calendar struct {
	Title   string
	storage StorageInterface
}

//NewCalendar Init new calendar
func NewCalendar(title string, s StorageInterface) *Calendar {
	return &Calendar{
		Title:   title,
		storage: s,
	}
}

//GetEvents Return list of all events
func (c *Calendar) GetEvents() ([]*Event, error) {
	return c.storage.GetAll()
}

//GetEventByID Return event by ID
func (c *Calendar) GetEventByID(id string) (*Event, error) {
	return c.storage.GetByID(id)
}

//CreateEvent Create new event and save it to storage
func (c *Calendar) CreateEvent(title string, from, to time.Time) (*Event, error) {
	event := NewEvent(title, from, to)
	if err := c.storage.Save(event); err != nil {
		return nil, err
	}

	return event, nil
}

//UpdateEvent Update existing event
func (c *Calendar) UpdateEvent(event *Event) error {
	return c.storage.Save(event)
}

//RemoveEvent Remove existing event
func (c *Calendar) RemoveEvent(event *Event) error {
	return c.storage.Remove(event)
}
