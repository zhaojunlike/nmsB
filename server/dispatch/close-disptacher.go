package dispatch

import (
	"log"

	"../http/websocket"
	"./context"
	"./handler/response"
)

func dispatchWebSocketClose(ctx *context.DispatchContext, webSocket *websocket.WebSocket) error {
	log.Printf("websocket [ %s ] to remote [ %s ] terminated\n", webSocket.Id, webSocket.RemoteAddr())
	// get the connection manager
	connectionManager := ctx.ConnectionManager
	// unregister the closed websocket
	connectionManager.Unregister(webSocket)

	// notify about changed server
	err := response.PushServerStatus(ctx)
	return err
}
