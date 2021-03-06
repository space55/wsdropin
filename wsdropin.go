package wsdropin

import (
	"net"
	"net/http"

	"golang.org/x/net/websocket"
)

// Dial gives you a net.Conn that talks to a WS destination.
// Addr should be like "ws://localhost:8080/"
func Dial(addr string) (net.Conn, error) {
	return websocket.Dial(addr, "", "http://localhost/")
}

type WSServer struct {
	net.Listener
	conns chan net.Conn
}

func Listen(laddr string) (net.Listener, error) {
	listener, err := net.Listen("tcp", laddr)
	if err != nil {
		return nil, err
	}
	wss := &WSServer{
		listener,
		make(chan net.Conn),
	}
	http.Handle("/", websocket.Handler(wss.wsHandler))
	go http.Serve(listener, nil)
	return wss, nil
}

func (w *WSServer) Accept() (net.Conn, error) {
	return <-w.conns, nil
}

func (w *WSServer) Close() error {
	return w.Close()
}

func (w *WSServer) Addr() net.Addr {
	// This is still legit enough? Maybe?
	return w.Addr()
}

func (w *WSServer) wsHandler(ws *websocket.Conn) {
	w.conns <- ws
}
