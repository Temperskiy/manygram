package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// 1. Генерируем ключ прямо в текущей папке, если его нет
	if _, err := os.Stat("matrix_key.pem"); os.IsNotExist(err) {
		log.Println("Generating matrix_key.pem...")
		cmdKey := exec.Command("/opt/render/project/go/bin/generate-keys", "--private-key", "matrix_key.pem")
		cmdKey.Stdout = os.Stdout
		cmdKey.Stderr = os.Stderr
		_ = cmdKey.Run() // Пробуем сгенерировать
	}

	// 2. Запускаем Dendrite
	log.Println("Starting Dendrite...")
	cmdRun := exec.Command("/opt/render/project/go/bin/dendrite", "--config", "dendrite.yaml", "-http-bind-address", ":10000")
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		log.Fatalf("Dendrite crashed: %v", err)
	}
}
