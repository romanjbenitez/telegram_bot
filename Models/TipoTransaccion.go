package Models

type TipoTransaccion string

const (
	Ingreso   TipoTransaccion = "Ingreso"
	Egreso    TipoTransaccion = "Egreso"
	Ahorro    TipoTransaccion = "Ahorro"
	Inversion TipoTransaccion = "Inversion"
)
