package handler

import (
	"net/http"
	"time"
)

type eventsForDateRequest struct {
	StartDay time.Time
}

// EventsForDayHandler Try to find events for day
// GET /events_for_day
// Fields:
// 	start_day: format 2006-01-02
func (h *Handler) EventsForDayHandler(req *http.Request) APIResponse {
	data := &eventsForDateRequest{}
	if err := data.parse(req); err != nil {
		return h.Error(http.StatusBadRequest, err)
	}

	events, err := h.Storage.GetByPeriod(data.StartDay, data.StartDay.Add(time.Hour*24))
	if err != nil {
		return h.Error(http.StatusInternalServerError, err)
	}

	return h.JSON(http.StatusOK, events)
}

func (data *eventsForDateRequest) parse(req *http.Request) error {
	var err error
	if err := req.ParseForm(); err != nil {
		return err
	}

	if data.StartDay, err = time.Parse("2006-01-02", req.FormValue("start_day")); err != nil {
		return err
	}

	return nil
}
