package handler

import (
	"net/http"
	"time"
)

// EventsForWeekHandler Try to find events for week
// GET /events_for_week
// Fields:
// 	start_day: format 2006-01-02
func (h *Handler) EventsForWeekHandler(req *http.Request) APIResponse {
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	events, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24*7))
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, events)
}
