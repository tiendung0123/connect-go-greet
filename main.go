package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/rs/cors"
	"github.com/tiendung0123/greet-proto/gen/go/proto"
)

type GreetServer struct{}

func (s *GreetServer) Greet(ctx context.Context, req *connect.Request[proto.GreetRequest]) (*connect.Response[proto.GreetResponse], error) {
	return connect.NewResponse(&proto.GreetResponse{
		Greeting: fmt.Sprintf("こんにちは、%sさん！", req.Msg.Name),
	}), nil
}

func main() {
	mux := http.NewServeMux()
	greeter := &GreetServer{}
	path, handler := proto.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	corsHandler := cors.Default().Handler(mux)
	log.Println("サーバーを開始します: http://localhost:8080")
	http.ListenAndServe("localhost:8080", corsHandler)
}
