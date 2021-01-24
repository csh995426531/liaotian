package websocket

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type wsMessage struct {
	MessageType int
	Data        []byte
}

type Connect struct {
	wsSocket  *websocket.Conn
	inChan    chan *wsMessage
	outChan   chan *wsMessage
	closeChan chan byte
	mutex     sync.Mutex
	isClosed  bool
}

func New(ctx *gin.Context) (conn *Connect, err error) {
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	wsSocket, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return nil, err
	}

	conn = &Connect{
		wsSocket:  wsSocket,
		inChan:    make(chan *wsMessage, 1000),
		outChan:   make(chan *wsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}

	// 读协程
	go conn.wsReadLoop()
	// 写协程
	go conn.wsWriteLoop()
	return
}

func (c *Connect) wsReadLoop() {
	for {
		// 从websocket读一个消息
		msgType, data, err := c.wsSocket.ReadMessage()
		if err != nil {
			goto ERR
		}
		if len(data) > 0 {
			req := &wsMessage{
				msgType,
				data,
			}
			// 放入请求队列
			select {
			case c.inChan <- req:
			case <-c.closeChan:
				goto CLOSED
			}
		}
	}
ERR:
	c.Close()
CLOSED:
}

func (c *Connect) wsWriteLoop() {
	for {
		select {
		case msg := <-c.outChan:
			// 写入websocket
			if err := c.wsSocket.WriteMessage(msg.MessageType, msg.Data); err != nil {
				goto ERR
			}
		case <-c.closeChan:
			goto CLOSED
		}
	}
ERR:
	c.Close()
CLOSED:
}

func (c *Connect) Write(messageType int, data []byte) error {
	select {
	case c.outChan <- &wsMessage{messageType, data}:
	case <-c.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (c *Connect) Read() (*wsMessage, error) {
	select {
	case msg := <-c.inChan:
		return msg, nil
	case <-c.closeChan:
	}
	return nil, errors.New("websocket closed")
}

func (c *Connect) Close() {
	c.wsSocket.Close()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.isClosed {
		c.isClosed = true
		close(c.closeChan)
	}
}
