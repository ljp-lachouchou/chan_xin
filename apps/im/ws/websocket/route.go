package websocket

type HandlerFunc func(server *Server, conn *Connection, msg *Message)
type Route struct {
	Method  string
	Handler HandlerFunc
}
