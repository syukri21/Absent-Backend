package socketIo

import (
	"fmt"
	"sync"

	socketio "github.com/googollee/go-socket.io"
)

var once sync.Once

// SocketIO ...
type SocketIO struct {
	Server *socketio.Server
}

// Run ...
func (s *SocketIO) Run() {

	s.Server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Join("teacher")

		rooms := s.Rooms()
		for _, room := range rooms {
			println(room)
		}
		return nil
	})

	s.Server.OnConnect("/absent", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		print(s.URL().RawQuery)
		s.Join("teacher")

		rooms := s.Rooms()
		for _, room := range rooms {
			println(room)
		}
		return nil
	})

	s.Server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	s.Server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		println(msg)
		return "recv " + msg
	})

	s.Server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	s.Server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	s.Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})
}

var instance *SocketIO

// GetSocketIO ....
func GetSocketIO() *SocketIO {

	once.Do(func() {

		server, err := socketio.NewServer(nil)

		if err != nil {
			panic(err)
		}

		instance = &SocketIO{Server: server}
	})

	return instance
}
