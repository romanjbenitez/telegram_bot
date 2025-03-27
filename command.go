package main

import (
	"log"
	"telegram_bot/Config"
)

const (
	INGRESO = "/ingreso"
	EGRESO  = "/egreso"
)

func egresoCommand() {
	log.Printf("cargo")
	Config.GetConnection()
}
