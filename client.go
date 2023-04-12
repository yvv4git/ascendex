package ascendex

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	mx   sync.Mutex
	conn Connection
	url  string
}

func NewClient(urlCreator URLCreator) *Client {
	return &Client{
		url: urlCreator.Create(),
	}
}

func (c *Client) Connection() error {
	dialer := websocket.Dialer{
		Subprotocols: []string{"json"},
	}

	wssConn, _, err := dialer.Dial(c.url, nil)
	if err != nil {
		return fmt.Errorf("connection could not be established: %#v\n", err)
	}
	c.mx.Lock()
	defer c.mx.Unlock()

	c.conn = wssConn

	return nil
}

func (c *Client) Disconnect() {
	c.mx.Lock()
	defer c.mx.Unlock()
	if c.conn == nil {
		fmt.Printf("failed to close connection\n")
		return
	}

	err := c.conn.Close()
	if err != nil {
		fmt.Printf("failed to close connection: %v\n", err)
		return
	}
	c.conn = nil
}

func (c *Client) SubscribeToChannel(symbol string) error {
	if c.conn == nil {
		return ErrConnectionClosed
	}

	preparedSymbol, err := SymbolPrepare(symbol)
	if err != nil {
		return fmt.Errorf("error on prepare symbol: %w", err)
	}

	if err = c.conn.WriteJSON(map[string]string{"op": "sub", "ch": "bbo:" + preparedSymbol}); err != nil {
		c.Disconnect()

		return fmt.Errorf("error on write json data to connection: %w", err)
	}

	return nil
}

func (c *Client) ReadMessagesFromChannel(ch chan<- BestOrderBook) {
	if c.conn == nil {
		fmt.Println("failed read message on nil conn")

		return
	}

	for {
		var m Message

		if err := c.conn.ReadJSON(&m); err != nil {
			fmt.Printf("failed read message from channel: %v", err)
		}

		var bbo BestOrderBook
		if m.M == "bbo" {
			if err := ParseOrderBookData(&bbo, m.Data); err == nil {
				ch <- bbo
			}
		}
	}
}

func (c *Client) WriteMessagesToChannel() {
	var msg = map[string]string{"op": "ping"}

	if c.conn != nil {
		if err := c.conn.WriteJSON(msg); err != nil {
			fmt.Printf("error on write json to conn: %v \n", err)
		}
	}
}
