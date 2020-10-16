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

//Cola es
type Cola struct {
	idpaquete   string
	seguimiento string
	tipo        string
	valor       string
	intentos    string
	estado      string
}

//Server es
type Server struct {
	ListaRegistro   []Registro
	Seguimiento     string
	ColaRetail      []Cola
	ColaPrioritario []Cola
	ColaNormal      []Cola
}

//EntregarPaqueteCamionNormal es
func (s *Server) EntregarPaqueteCamionNormal(ctx context.Context, message *IdCamion) (*ColaPaquete, error) {
	log.Println("Entre aca")
	messageColaPaqueteError := ColaPaquete{
		Idpaquete:   "NoPaquetes",
		Seguimiento: "",
		Tipo:        "",
		Valor:       "",
		Intentos:    "",
		Estado:      "",
	}
	if len(s.ColaPrioritario) < 1 && len(s.ColaNormal) < 1 { // NoPaquetes ninguna cola
		return &messageColaPaqueteError, nil
	}

	indicemascaro := 0
	elemmascaro := 0
	iterador := 0

	if len(s.ColaPrioritario) > 1 {
		// Busqueda paquete mas caro
		for _, elem := range s.ColaPrioritario {
			value, _ := strconv.Atoi(elem.valor)
			// Buscar elemento más caro
			if value > elemmascaro {
				indicemascaro = iterador
				elemmascaro = value
			}
			iterador++
		}

		// Preparacion retorno paquete prioritario mas caro

		messageColaPaquete := ColaPaquete{
			Idpaquete:   s.ColaPrioritario[indicemascaro].idpaquete,
			Seguimiento: s.ColaPrioritario[indicemascaro].seguimiento,
			Tipo:        s.ColaPrioritario[indicemascaro].tipo,
			Valor:       s.ColaPrioritario[indicemascaro].valor,
			Intentos:    s.ColaPrioritario[indicemascaro].intentos,
			Estado:      s.ColaPrioritario[indicemascaro].estado,
		}

		// Quitar elemento mas caro de la cola

		// https://yourbasic.org/golang/delete-element-slice/
		s.ColaPrioritario[indicemascaro] = s.ColaPrioritario[len(s.ColaPrioritario)-1]
		s.ColaPrioritario = s.ColaPrioritario[:len(s.ColaPrioritario)-1]
		log.Println(&messageColaPaquete)

		return &messageColaPaquete, nil
	}

	if len(s.ColaNormal) > 1 {
		// Busqueda paquete mas caro
		for _, elem := range s.ColaNormal {
			value, _ := strconv.Atoi(elem.valor)
			// Buscar elemento más caro
			if value > elemmascaro {
				indicemascaro = iterador
				elemmascaro = value
			}
			iterador++
		}

		// Preparacion retorno paquete normal mas caro

		messageColaPaquete := ColaPaquete{
			Idpaquete:   s.ColaPrioritario[indicemascaro].idpaquete,
			Seguimiento: s.ColaPrioritario[indicemascaro].seguimiento,
			Tipo:        s.ColaPrioritario[indicemascaro].tipo,
			Valor:       s.ColaPrioritario[indicemascaro].valor,
			Intentos:    s.ColaPrioritario[indicemascaro].intentos,
			Estado:      s.ColaPrioritario[indicemascaro].estado,
		}

		// Quitar elemento mas caro de la cola

		// https://yourbasic.org/golang/delete-element-slice/
		s.ColaPrioritario[indicemascaro] = s.ColaPrioritario[len(s.ColaPrioritario)-1]
		s.ColaPrioritario = s.ColaPrioritario[:len(s.ColaPrioritario)-1]
		log.Println(&messageColaPaquete)

		return &messageColaPaquete, nil
	}
	return &messageColaPaqueteError, nil
}

//RecibirOrdenPymes es
func (s *Server) RecibirOrdenPymes(ctx context.Context, message *Ordenclientepymes) (*Ordenseguimiento, error) {
	var tipo string
	messageOrdenseguimiento := Ordenseguimiento{
		Nordenseguimiento: s.Seguimiento,
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
		valor:       strconv.Itoa(int(message.Valor)),
		origen:      message.Tienda,
		destino:     message.Destino,
		seguimiento: s.Seguimiento,
	}

	nuevaCola := Cola{
		idpaquete:   message.Id,
		seguimiento: s.Seguimiento,
		tipo:        tipo,
		valor:       strconv.Itoa(int(message.Valor)),
		intentos:    "0",
		estado:      "En bodega",
	}
	// Agregar paquete a una de las tres colas de Server
	if tipo == "prioritario" {
		// Add cola prioritario
		s.ColaPrioritario = append(s.ColaPrioritario, nuevaCola)
	} else {
		// Add cola normal
		s.ColaNormal = append(s.ColaNormal, nuevaCola)
	}

	// Actualizacion codigo de seguimiento (atributo Server)
	numerosiguiente, _ := strconv.Atoi(s.Seguimiento)
	numerosiguiente++
	s.Seguimiento = strconv.Itoa(numerosiguiente)
	s.ListaRegistro = append(s.ListaRegistro, nuevaEntrada)

	//log.Println(s.ColaNormal)
	//log.Println(s.ListaRegistro)
	return &messageOrdenseguimiento, nil
}

//RedecirOrdenPymes es
func (s *Server) RedecirOrdenPymes(ctx context.Context, message *Ordenclientepymes) (*Ordenclientepymes, error) {
	log.Println("Recibi la orden")
	log.Printf("Producto pymes: %s", message.Producto)
	return message, nil
}

//RecibirOrdenRetail es
func (s *Server) RecibirOrdenRetail(ctx context.Context, message *Ordenclienteretail) (*Ordenseguimiento, error) {
	messageOrdenseguimiento := Ordenseguimiento{
		Nordenseguimiento: s.Seguimiento,
	}

	nuevaEntrada := Registro{
		timestamp:   time.Now(),
		idpaquete:   message.Id,
		tipo:        "retail",
		nombre:      message.Producto,
		valor:       strconv.Itoa(int(message.Valor)),
		origen:      message.Tienda,
		destino:     message.Destino,
		seguimiento: s.Seguimiento,
	}

	nuevaCola := Cola{
		idpaquete:   message.Id,
		seguimiento: s.Seguimiento,
		tipo:        "retail",
		valor:       strconv.Itoa(int(message.Valor)),
		intentos:    "0",
		estado:      "En bodega",
	}

	// Agregar paquete a una de las tres colas de Server
	s.ColaRetail = append(s.ColaRetail, nuevaCola)
	//log.Println(s.ColaRetail)

	// Actualizacion codigo de seguimiento (atributo Server)
	numerosiguiente, _ := strconv.Atoi(s.Seguimiento)
	numerosiguiente++
	s.Seguimiento = strconv.Itoa(numerosiguiente)
	s.ListaRegistro = append(s.ListaRegistro, nuevaEntrada)
	//log.Println(s.ListaRegistro)
	return &messageOrdenseguimiento, nil
}

//RedecirOrdenRetail es
func (s *Server) RedecirOrdenRetail(ctx context.Context, message *Ordenclienteretail) (*Ordenclienteretail, error) {
	log.Println("Recibi la orden retail")
	log.Printf("Producto retail: %s", message.Producto)
	return message, nil
}

//CodigoSeguimiento es
func (s *Server) CodigoSeguimiento(ctx context.Context, message *Ordenseguimiento) (*Estado, error) {
	mensajeError := Estado{
		Estado: "No se pudo encontrar código",
	}
	return &mensajeError, nil
}
