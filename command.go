package main

import "log"

const (
	INGRESO = "/ingreso"
	EGRESO  = "/egreso"
)

func egresoCommand() {
	log.Printf("cargo")
	getConnectionString()
}
