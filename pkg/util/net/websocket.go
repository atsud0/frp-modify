package net

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	mywebsocket "github.com/atsud0/websocket"
)

var (
	ErrWebsocketListenerClosed = errors.New("websocket listener closed")
)

const (
	FrpWebsocketPath = "/fffadmin"
)

type WebsocketListener struct {
	ln       net.Listener
	acceptCh chan net.Conn

	server    *http.Server
	httpMutex *http.ServeMux
}

// NewWebsocketListener to handle websocket connections
// ln: tcp listener for websocket connections
func NewWebsocketListener(ln net.Listener) (wl *WebsocketListener) {
	wl = &WebsocketListener{
		acceptCh: make(chan net.Conn),
	}

	muxer := http.NewServeMux()
	muxer.Handle(FrpWebsocketPath, mywebsocket.Handler(func(c *mywebsocket.Conn) {
		notifyCh := make(chan struct{})
		conn := WrapCloseNotifyConn(c, func() {
			close(notifyCh)
		})
		wl.acceptCh <- conn
		<-notifyCh
	}))

	wl.server = &http.Server{
		Addr:    ln.Addr().String(),
		Handler: muxer,
	}

	go wl.server.Serve(ln)
	return
}

func ListenWebsocket(bindAddr string, bindPort int) (*WebsocketListener, error) {
	tcpLn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", bindAddr, bindPort))
	if err != nil {
		return nil, err
	}
	l := NewWebsocketListener(tcpLn)
	return l, nil
}

func (p *WebsocketListener) Accept() (net.Conn, error) {
	c, ok := <-p.acceptCh
	if !ok {
		return nil, ErrWebsocketListenerClosed
	}
	return c, nil
}

func (p *WebsocketListener) Close() error {
	return p.server.Close()
}

func (p *WebsocketListener) Addr() net.Addr {
	return p.ln.Addr()
}
