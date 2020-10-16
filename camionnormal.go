package main

import (
	"bufio"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/MatiasMarchant/Prueba1/tree/master/chat"
)

func preguntasiniciales() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Preguntas iniciales")
	fmt.Println("Tiempo en segundos de espera por segundo paquete")
	fmt.Printf("> ")
	tiempoespera, _ := reader.ReadBytes('\n')
	return string(tiempoespera)
}

func main() {
	var tiempoespera string
	tiempoespera = preguntasiniciales()
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
			if paquete2.Idpaquete == "Nopaquetes" {
				// Marchar solo con paquete
				os.Exit(0)
			} else {
				// Marchar con paquete y paquete2
				os.Exit(0)
			}
		}

	}

}
