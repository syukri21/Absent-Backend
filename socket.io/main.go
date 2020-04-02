package socketIo

import (
	"backend-qrcode/middleware"
	"fmt"
	"net/url"
	"strconv"
	"sync"

	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/websocket"

	"github.com/dgrijalva/jwt-go"
	socketio "github.com/googollee/go-socket.io"
)

var once sync.Once

// SocketIO ...
type SocketIO struct {
	Server *socketio.Server
}

// VerifyJWTReturn ...
type VerifyJWTReturn struct {
	UserID string
	RoleID string
}

// VerifyJWT ...
func VerifyJWT(tokenString string) (*VerifyJWTReturn, *error) {
	claims, err := middleware.VerifyToken(tokenString)
	if err != nil {
		return nil, &err
	}
	UserID := strconv.FormatFloat(claims.(jwt.MapClaims)["user_id"].(float64), 'g', 1, 64)
	RoleID := strconv.FormatFloat(claims.(jwt.MapClaims)["role_id"].(float64), 'g', 1, 64)

	return &VerifyJWTReturn{UserID, RoleID}, nil
}

// Run ...
func (s *SocketIO) Run() {

	s.Server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		queryString, _ := url.ParseQuery(s.URL().RawQuery)
		protected := queryString.Get("protected")
		token := queryString.Get("token")
		data, err := VerifyJWT(token)

		if err != nil {
			s.Close()
		}

		if protected == "teacher" && data.RoleID == "1" {
			room := queryString.Get("room")
			print(room)
			if len(room) != 0 {
				s.Join(room)
			}
		} else {
			s.Close()
		}

		return nil
	})

	s.Server.OnEvent("/", "absent", func(s socketio.Conn, msg map[string]interface{}) {
		fmt.Printf("as %v", msg)
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

		server, err := socketio.NewServer(&engineio.Options{
			Transports: []transport.Transport{
				websocket.Default,
			},
		})

		if err != nil {
			panic(err)
		}

		instance = &SocketIO{Server: server}
	})

	return instance
}
