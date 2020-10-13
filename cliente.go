package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.40.178:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Ordencliente{
		Id:          "Id",
		Producto:    "Producto",
		Valor:       15,
		Tienda:      "Tienda",
		Destino:     "Destino",
		Prioritario: true,
	}

	response, err := c.RedecirOrden(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Producto)
}
