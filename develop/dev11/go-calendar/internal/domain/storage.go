package domain

import (
	"sync"
	"time"
)

//StorageInterface Interface of events storage
type StorageInterface interface {
	GetAll() ([]*Event, error)
	GetByID(id string) (*Event, error)
	GetByPeriod(from, to time.Time) ([]*Event, error)
	Save(event *Event) error
	Remove(event *Event) error
}

//Storage Storage struct
type Storage struct {
	sync.Mutex
	events map[string]*Event
}

//NewStorage Create new storage
func NewStorage() *Storage {
	eventStorage := make(map[string]*Event)
	return &Storage{events: eventStorage}
}

//GetAll Return list of all events
func (storage *Storage) GetAll() ([]*Event, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	result := make([]*Event, 0, len(storage.events))
	for _, event := range storage.events {
		result = append(result, event)
	}

	return result, nil
}

//GetByID Return event by ID
func (storage *Storage) GetByID(id string) (*Event, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	if event, ok := storage.events[id]; ok {
		return event, nil
	}

	return nil, nil
}

//GetByPeriod Return list of events by period
func (storage *Storage) GetByPeriod(from, to time.Time) ([]*Event, error) {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	result := make([]*Event, 0)

	for _, event := range storage.events {
		if event.DateFrom.Before(to) && event.DateTo.After(from) {
			result = append(result, event)
		}
	}

	return result, nil
}

//Save Create or update event in storage
func (storage *Storage) Save(event *Event) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	storage.events[event.ID.String()] = event

	return nil
}

//Remove Remove event from storage
func (storage *Storage) Remove(event *Event) error {
	storage.Mutex.Lock()
	defer storage.Mutex.Unlock()
	delete(storage.events, event.ID.String())

	return nil
}
