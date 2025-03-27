package Models

type FormaDePago string

const (
	Debito                FormaDePago = "Debito"
	Credito_1             FormaDePago = "Credito un pagos"
	Credito_3             FormaDePago = "Credito tres pagos"
	Credito_6             FormaDePago = "Credito seis pagos"
	Credito_9             FormaDePago = "Credito nueve pagos"
	Credito_12            FormaDePago = "Credito doce pagos"
	Credito_18            FormaDePago = "Credito diesciocho pagos"
	Transferencia         FormaDePago = "Transferencia"
	Transferencia_Credito FormaDePago = "Transferencia a pagar"
)
