package handler

import (
	"net/http"
	"time"
)

// EventsForMonthHandler Try to find events for month
// GET /events_for_month
// Fields:
// 	start_day: format 2006-01-02
func (h *Handler) EventsForMonthsHandler(req *http.Request) APIResponse {
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}
	events, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24*30))
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, events)
}
