package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Создаем пустой файл ключа, чтобы Dendrite его увидел при запуске
	if _, err := os.Stat("matrix_key.pem"); os.IsNotExist(err) {
		log.Println("Creating matrix_key.pem...")
		os.WriteFile("matrix_key.pem", []byte(""), 0644)
	}

	log.Println("Starting Dendrite...")
	// Запускаем Dendrite с портом 10000
	cmdRun := exec.Command("/opt/render/project/go/bin/dendrite", "--config", "dendrite.yaml", "-http-bind-address", ":10000")
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		log.Fatalf("Dendrite crashed: %v", err)
	}
}
