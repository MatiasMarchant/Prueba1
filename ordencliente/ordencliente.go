package ordencliente

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) RedecirOrden(ctx context.Context, message *OrdenCliente) (*OrdenCliente, error) {
	log.Printf("Recibi orden de cliente:")
	log.Printf("Producto: %s", message.Producto)
	return &OrdenCliente{Producto: "m qro matar!"}, nil
}
