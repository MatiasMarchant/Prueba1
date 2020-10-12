package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) RedecirOrden(ctx context.Context, message *OrdenCliente) (*OrdenCliente, error) {
	log.Printf("Recibi la orden")
	log.Printf("")
	log.Printf("Producto: %s", message.Producto)
	return , nil
}