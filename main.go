package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Сразу запускаем уже скачанный в Build Command файл
	log.Println("Starting Dendrite...")
	// Путь к dendrite меняем на тот, куда его положил 'go install'
	cmdRun := exec.Command("/opt/render/project/go/bin/dendrite", "--config", "dendrite.yaml", "-http-bind-address", ":10000")
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		log.Fatalf("Dendrite crashed: %v", err)
	}
}
