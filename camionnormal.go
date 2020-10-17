package main

import (
	"bufio"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
)

//RegistroCamion es
type RegistroCamion struct {
	idpaquete    string
	tipo         string
	valor        string
	origen       string
	destino      string
	intentos     string
	fechaentrega time.Time
}

func preguntasinicialescamion() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Preguntas iniciales")
	fmt.Println("Tiempo en segundos de espera por segundo paquete")
	fmt.Printf("> ")
	tiempoespera, _ := reader.ReadBytes('\n')
	fmt.Println("Tiempo en segundos demora envio paquete")
	fmt.Printf("> ")
	tiempodemora, _ := reader.ReadBytes('\n')
	return string(tiempoespera), string(tiempodemora)
}

//Entregarpaquete es
func Entregarpaquete(paquete *chat.ColaPaquete, ListaRegistroCamion []RegistroCamion) chat.ColaPaquete {
	fmt.Println(paquete)
	paquete.Estado = "En camino"
	exito := rand.Float64()
	if exito < 0.8 {
		// Se entrega paquete
		// Se modifica registro - fechaentrega
		for _, elem := range ListaRegistroCamion {
			if elem.idpaquete == paquete.Idpaquete {
				elem.fechaentrega = time.Now()
			}
		}
		paquete.Estado = "Recibido"
		// Se notifica a logistica

	} else {
		// No se entrega paquete - depende de tipo ver q se hace
		// Aumentar atributo intentos, etc

		// quizas se notifica a logistica

		// PORMIENTRAS SE DEJA ESTO
		for _, elem := range ListaRegistroCamion {
			if elem.idpaquete == paquete.Idpaquete {
				elem.fechaentrega = time.Now()
			}
		}
		paquete.Estado = "No Recibido"

	}
	return *paquete
}

func main() {
	var ListaRegistroCamion []RegistroCamion
	var tiempoespera, tiempodemora string
	tiempoespera, tiempodemora = preguntasinicialescamion()
	tiempodemoraint, _ := strconv.Atoi(strings.TrimSuffix(tiempodemora, "\r\n")) // CUIDADO LINUX
	tiempoesperaint, _ := strconv.Atoi(strings.TrimSuffix(tiempoespera, "\r\n")) // CUIDADO LINUX

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No me pude conectar al puerto 9001: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	// time.Sleep(time.Second * time.Duration(int64(tiempoesperaint)))
	idcamion := chat.IdCamion{
		Idcamion: "1",
	}

	for true {
		paquete, _ := c.EntregarPaqueteCamionNormal(context.Background(), &idcamion)
		if paquete.Idpaquete == "NoPaquetes" { // Si no encuentra paquetes, dormir
			time.Sleep(time.Second * time.Duration(int64(tiempoesperaint)))
		} else { // Si encontró paquete, dormir para esperar el 2do y si no, marchar
			time.Sleep(time.Second * time.Duration(int64(tiempoesperaint)))
			paquete2, _ := c.EntregarPaqueteCamionNormal(context.Background(), &idcamion)
			if paquete2.Idpaquete == "Nopaquetes" { // Solo paquete

				// Primero se ingresa a su registro
				nuevoRegistro := RegistroCamion{
					idpaquete:    paquete.Idpaquete,
					tipo:         paquete.Tipo,
					valor:        paquete.Valor,
					origen:       paquete.Origen,
					destino:      paquete.Destino,
					intentos:     paquete.Intentos,
					fechaentrega: time.Time{},
				}

				ListaRegistroCamion = append(ListaRegistroCamion, nuevoRegistro)

				// Marchar solo con paquete
				*paquete = Entregarpaquete(paquete, ListaRegistroCamion)
				time.Sleep(time.Second * time.Duration(int64(tiempodemoraint)))

				//fmt.Println(paquete)
				//os.Exit(0)
			} else {

				// Primero se ingresa a su registro
				nuevoRegistro1 := RegistroCamion{
					idpaquete:    paquete.Idpaquete,
					tipo:         paquete.Tipo,
					valor:        paquete.Valor,
					origen:       paquete.Origen,
					destino:      paquete.Destino,
					intentos:     paquete.Intentos,
					fechaentrega: time.Time{},
				}

				nuevoRegistro2 := RegistroCamion{
					idpaquete:    paquete2.Idpaquete,
					tipo:         paquete2.Tipo,
					valor:        paquete2.Valor,
					origen:       paquete2.Origen,
					destino:      paquete2.Destino,
					intentos:     paquete2.Intentos,
					fechaentrega: time.Time{},
				}

				ListaRegistroCamion = append(ListaRegistroCamion, nuevoRegistro1)
				ListaRegistroCamion = append(ListaRegistroCamion, nuevoRegistro2)

				// Marchar con paquete y paquete2 (ver cual es mas caro)
				paquete.Estado = "En camino"
				paquete2.Estado = "En camino"

				valor1, _ := strconv.Atoi(paquete.Valor)
				valor2, _ := strconv.Atoi(paquete2.Valor)

				if valor1 > valor2 {
					// Se entrega paquete primero

					*paquete = Entregarpaquete(paquete, ListaRegistroCamion)
					*paquete2 = Entregarpaquete(paquete2, ListaRegistroCamion)
				} else {
					// Se entrega paquete2 primero
					*paquete2 = Entregarpaquete(paquete2, ListaRegistroCamion)
					*paquete = Entregarpaquete(paquete, ListaRegistroCamion)

				}

				//os.Exit(0)
			}
		}

	}

}
