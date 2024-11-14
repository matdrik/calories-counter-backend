package ws

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"time"
)

func (c *Client) readMessages(ctx context.Context) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close(websocket.StatusNormalClosure, "closing")
	}()
	for {
		var msg interface{}
		err := wsjson.Read(ctx, c.conn, &msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		c.hub.broadcast <- msg
	}
}

func (c *Client) writeMessages(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-c.send:
			if !ok {
				return
			}
			err := wsjson.Write(ctx, c.conn, msg)
			if err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}
}

// WebSocket handler adapted for Gin
func HandleWebSocket(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Print("HandleWebSocket")

		conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
			OriginPatterns: []string{"*"},
		})
		if err != nil {
			log.Println("WebSocket Accept error:", err)
			return
		}

		client := &Client{
			conn: conn,
			send: make(chan interface{}),
			hub:  hub,
		}
		hub.register <- client

		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Minute)
		defer cancel()

		go client.writeMessages(ctx)
		client.readMessages(ctx)
	}
}
