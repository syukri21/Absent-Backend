package socket

import (
	socketio "github.com/googollee/go-socket.io"
)

// Socket ....
type Socket struct {
	Server *socketio.Server
}

// Serve ...
func (s Socket) Serve() {
	s.Server.Serve()
}

// Close ...
func (s Socket) Close() error {
	return s.Server.Close()
}

// Listen ...
func (s Socket) Listen() {
	s.OnAbsentCreated()
	s.OnAbsentDeleted()
	s.OnConnect()
}

var instance *Socket

// NewSocket ...
func NewSocket() (*Socket, error) {

	var err error
	var server *socketio.Server

	if instance == nil {
		server, err = socketio.NewServer(nil)
		instance = &Socket{
			Server: server,
		}
	}

	return instance, err
}
