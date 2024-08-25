package network

import (
	. "chat_server"
	"github.com/gorilla/websocket"
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: MessageBufferSize, CheckOrigin: func(r *http.Request) bool {return true}}


type Room struct {
	Forward chan *message // 수신되는 메세지 보관값
	//들어오는 메세지를 다른 클라이언트에게 전송

	Join chan *Client // Socket이 연결되는 경우에 작동
	Leave chan *Client // Socket이 끊어지는 경우에 작동

	Clients map[*Client]bool // 현재 방에 있는 Client 정보 저장
}

type Message struct{
	Name 	string
	Message string
	Time 	int64
}

type Client struct {
	Send chan *message
	Room *Room
	Name string
	Socket *websocket.Conn
}

func NewRoom() *Room {
	return &Room{
		Forward: make{chan *message},
		Join: 	 make{chan *Client},
		Leave:	 make{chan *Client},
		Clients: make{map[*Client]bool},
	}
}




func (r *Room) SocketServe(c *gin.Context) { //gin Context를 연결해줘야 gin 프레임워크 사용시 api 연결이 가능함
	// upgrader.CheckOrigin = func(r *http.Request) bool {
	// 	return true
	// }
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
}