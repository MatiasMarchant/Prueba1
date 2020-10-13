package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) RedecirOrden(ctx context.Context, message *Ordencliente) (*Ordencliente, error) {
	log.Printf("Recibi la orden")
	log.Printf("")
	log.Printf("Producto: %s", message.Producto)
	return message, nil
}
