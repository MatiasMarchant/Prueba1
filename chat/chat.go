package chat

import (
	"golang.org/x/net/context"
	"log"
	"strconv"
	"time"
)

//Registro es
type Registro struct {
	timestamp   time.Time
	idpaquete   string
	tipo        string
	nombre      string
	valor       string
	origen      string
	destino     string
	seguimiento string
}

//Server es
type Server struct {
	listaRegistro []Registro
	seguimiento   string
}

//RecibirOrdenPymes es
func (s *Server) RecibirOrdenPymes(ctx context.Context, message *Ordenclientepymes) (*Ordenseguimiento, error) {
	var tipo string
	messageOrdenseguimiento := Ordenseguimiento{
		Nordenseguimiento: s.seguimiento,
	}
	if message.Prioritario == true {
		tipo = "prioritario"
	} else {
		tipo = "normal"
	}
	nuevaEntrada := Registro{
		timestamp:   time.Now(),
		idpaquete:   message.Id,
		tipo:        tipo,
		nombre:      message.Producto,
		valor:       string(message.Valor),
		origen:      message.Tienda,
		destino:     message.Destino,
		seguimiento: s.seguimiento,
	}

	numerosiguiente, _ := strconv.Atoi(s.seguimiento)
	numerosiguiente++

	s.seguimiento = string(numerosiguiente)
	s.listaRegistro = append(s.listaRegistro, nuevaEntrada)
	return &messageOrdenseguimiento, nil
}

//RedecirOrdenPymes es
func (s *Server) RedecirOrdenPymes(ctx context.Context, message *Ordenclientepymes) (*Ordenclientepymes, error) {
	log.Println("Recibi la orden")
	log.Printf("Producto pymes: %s", message.Producto)
	return message, nil
}

//RedecirOrdenRetail es
func (s *Server) RedecirOrdenRetail(ctx context.Context, message *Ordenclienteretail) (*Ordenclienteretail, error) {
	log.Println("Recibi la orden retail")
	log.Printf("Producto retail: %s", message.Producto)
	return message, nil
}
