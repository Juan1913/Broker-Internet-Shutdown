package main

import (
	"fmt"
	"log"
	"os/exec"
)

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func blockInternet() {
	// Ejecutar iptables para bloquear todo el tráfico saliente
	output, err := runCommand("sudo", "iptables", "-A", "OUTPUT", "-o", "eth0", "-j", "REJECT")
	if err != nil {
		log.Fatalf("Error al bloquear Internet: %v\n", err)
	}
	fmt.Println("Acceso a Internet bloqueado.")
	fmt.Println(output)
}

func restoreInternet() {

	output, err := runCommand("sudo", "iptables", "-D", "OUTPUT", "-o", "eth0", "-j", "REJECT")
	if err != nil {
		log.Fatalf("Error al restaurar Internet: %v\n", err)
	}
	fmt.Println("Acceso a Internet restaurado.")
	fmt.Println(output)
}

func printBanner() {
	banner := `
  _____  _             _____               _       _     _             
 |_   _|(_)           |  __ \             | |     | |   (_)            
   | |   _  ___ _ __  | |  | | __ _ _ __  | | ___ | |__  _ _ __   __ _ 
   | |  | |/ _ \ '_ \ | |  | |/ _' | '_ \ | |/ _ \| '_ \| | '_ \ / _' |
  _| |_ | |  __/ | | || |__| | (_| | | | || | (_) | | | | | | | | (_| |
 |_____|/ |\___|_| |_||_____/ \__,_|_| |_||_|\___/|_| |_|_|_| |_|\__, |
       |__/                                                     __/ |
                                                                  |___/ 
                                 By JuanThorRes
`
	fmt.Println(banner)
}

func main() {

	printBanner()

	var action string
	fmt.Println("Escribe 'block' para cortar Internet o 'restore' para restaurar el acceso.")
	fmt.Scanln(&action)

	if action == "block" {
		blockInternet()
	} else if action == "restore" {
		restoreInternet()
	} else {
		fmt.Println("Acción no reconocida. Por favor, escribe 'block' o 'restore'.")
	}
}
