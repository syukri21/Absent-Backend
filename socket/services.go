package socket

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

// OnConnect ...
func (s Socket) OnConnect() {
	s.Server.OnConnect("/socket", func(sio socketio.Conn) error {
		sio.SetContext("")
		fmt.Println("connected:", sio.ID())
		return nil
	})
}

// OnAbsentCreated ...
func (s Socket) OnAbsentCreated() {
	s.Server.OnEvent("/absent", "ABSENT_CREATED", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
}

// OnAbsentDeleted ...
func (s Socket) OnAbsentDeleted() {
	s.Server.OnEvent("/absent", "ABSENT_DELETED", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
}
