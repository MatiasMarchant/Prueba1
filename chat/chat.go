package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) RedecirOrdenPymes(ctx context.Context, message *Ordenclientepymes) (*Ordenclientepymes, error) {
	log.Println("Recibi la orden")
	log.Printf("Producto pymes: %s", message.Producto)
	return message, nil
}

func (s *Server) RedecirOrdenRetail(ctx context.Context, message *Ordenclienteretail) (*Ordenclienteretail, error) {
	log.Println("Recibi la orden retail")
	log.Printf("Producto retail: %s", message.Producto)
	return message, nil
}
