package handlers

import (
	"io"
)

type SSEService interface {
	MarshalTo(io.Writer) error
	GetEvents()
	// GetAllPlayers() ([]*services.Player, error)
	// GetPlayerById(id int) (services.Player, error)
	// UpdatePlayer(player services.Player) error
}

func NewSSEHandler(ss SSEService) *SSEHandler {
	return &SSEHandler{ss}
}

type SSEHandler struct {
	SSEService SSEService
}

// func (sh *SSEHandler) ShowEvents(c echo.Context) error {
// 	log.Printf("SSE client connected, ip: %v", c.RealIP())
//
// 	w := c.Response()
// 	w.Header().Set("Content-Type", "text/event-stream")
// 	w.Header().Set("Cache-Control", "no-cache")
// 	w.Header().Set("Connection", "keep-alive")
//
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-c.Request().Context().Done():
// 			log.Printf("SSE client disconnected, ip: %v", c.RealIP())
// 			return nil
// 		case <-ticker.C:
// 			event := services.Event{
// 				Data: []byte("time: " + time.Now().Format(time.RFC3339Nano)),
// 			}
// 			// if err := sh.SSEService.MarshalTo(w); err != nil {
// 			if err := event.Event.MarshalTo(w); err != nil {
// 				return err
// 			}
// 			w.Flush()
// 		}
// 	}
// }
