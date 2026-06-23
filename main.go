package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// 1. Скачиваем официальный Dendrite прямо во время запуска в облаке
	log.Println("Downloading Dendrite...")
	cmdDownload := exec.Command("go", "install", "github.com/matrix-org/dendrite/cmd/dendrite@latest")
	cmdDownload.Stdout = os.Stdout
	cmdDownload.Stderr = os.Stderr
	if err := cmdDownload.Run(); err != nil {
		log.Fatalf("Failed to download Dendrite: %v", err)
	}

	// 2. Запускаем Dendrite с нашим файлом конфигурации
	log.Println("Starting Dendrite...")
	// Путь куда Go устанавливает бинарники по умолчанию в Linux
	cmdRun := exec.Command("/opt/render/project/go/bin/dendrite", "--config", "dendrite.yaml")
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	if err := cmdRun.Run(); err != nil {
		log.Fatalf("Dendrite crashed: %v", err)
	}
}
