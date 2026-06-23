package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Создаем ключ, если нет
	os.WriteFile("matrix_key.pem", []byte("SIGNING_KEY_MUST_BE_32_BYTES_LONG_!!!"), 0644)
	
	log.Println("Launching Dendrite directly...")
	// Запускаем напрямую
	cmd := exec.Command("/opt/render/project/go/bin/dendrite", "--config", "dendrite.yaml", "-http-bind-address", ":10000")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
