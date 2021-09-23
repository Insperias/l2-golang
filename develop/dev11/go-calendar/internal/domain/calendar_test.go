package domain

import (
	"testing"
	"time"
)

type testEventData struct {
	Title string
	From  time.Time
	To    time.Time
}

func TestCreateCalendar(t *testing.T) {
	calendarTitle := "Calendar"
	calendar := NewCalendar(calendarTitle, NewStorage())

	if calendar == nil {
		t.Fatalf("Failed to inititate calendar")
	}

	if calendar.Title != calendarTitle {
		t.Fatalf("Invalid calendar title: expect [%s], got [%s]", calendarTitle, calendar.Title)
	}
}

func TestCreateEvent(t *testing.T) {
	testData := &testEventData{
		"Event #1",
		time.Now().Add(12 * time.Hour),
		time.Now().Add(13 * time.Hour),
	}

	calendar := NewCalendar("Calendar", NewStorage())
	event, err := calendar.CreateEvent(testData.Title, testData.From, testData.To)
	if err != nil {
		t.Fatalf("Failed to create event: %s", err)
	}

	if event.ID.String() == "" {
		t.Fatalf("Invalid event ID: expect not empty")
	}

	if event.Title != testData.Title {
		t.Fatalf("Invalid event title: expect [%s], got [%s]", testData.Title, event.Title)
	}

	if event.DateFrom.Unix() != testData.From.Unix() {
		t.Fatalf("Invalid event date from: expect [%s], got [%s]", testData.From, event.DateFrom)
	}

	if event.DateTo.Unix() != testData.To.Unix() {
		t.Fatalf("Invalid event date to: expect [%s], got [%s]", testData.To, event.DateTo)
	}
}

func TestSaveEvent(t *testing.T) {
	testData := &testEventData{
		"Event #1",
		time.Now().Add(12 * time.Hour),
		time.Now().Add(13 * time.Hour),
	}

	calendar := NewCalendar("Calendar", NewStorage())
	event, err := calendar.CreateEvent(testData.Title, testData.From, testData.To)
	if err != nil {
		t.Fatalf("Failed to create event: %s", err)
	}

	testEventTitle := "Custom Event"
	event.Title = testEventTitle
	if err := calendar.UpdateEvent(event); err != nil {
		t.Fatalf("failed to update event: %s", err)
	}

	foundEvent, err := calendar.GetEventByID(event.ID.String())
	if err != nil {
		t.Fatalf("Failed to load event by ID [%s]: %s", event.ID.String(), err)
	}

	if foundEvent.Title != testEventTitle {
		t.Fatalf("Invalid event title: expect [%s], got [%s]", testEventTitle, foundEvent.Title)
	}
}

func TestRemoveEvent(t *testing.T) {
	eventTime := time.Now()
	calendar := NewCalendar("Calendar", NewStorage())
	event1, _ := calendar.CreateEvent("Event #1", eventTime, eventTime)
	_, _ = calendar.CreateEvent("Event #2", eventTime, eventTime)
	_, _ = calendar.CreateEvent("Event #3", eventTime, eventTime)

	if err := calendar.RemoveEvent(event1); err != nil {
		t.Fatalf("Faield to remove event: %s", err)
	}

	events, err := calendar.GetEvents()
	if err != nil {
		t.Fatalf("Failed to load all events: %s", err)
	}

	if totalEvents := len(events); totalEvents != 2 {
		t.Fatalf("Invlid quantity of events: expect 2, got %d", totalEvents)
	}

}
