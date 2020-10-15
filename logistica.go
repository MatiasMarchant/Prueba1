package main

import (
	"log"
	"net"

	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Falle al escuchar puerto 9000: %v", err)
	}

	var listaRegistro []Registro
	s := chat.Server{
		listaRegistro: listaRegistro,
		seguimiento:   "0",
	}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falle siendo un servidor gRPC en el puerto 9000: %v", err)
	}

}
