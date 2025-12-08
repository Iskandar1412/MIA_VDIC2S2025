package general

type Resultado struct {
	StrMensajeError string
	BlnError        bool
	Respuesta       interface{}
}

type SalidaComandoEjecutado struct {
	LstComandos []string
}

type ResultadoAPI struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResultadoSalida(message string, isError bool, data interface{}) ResultadoAPI {
	return ResultadoAPI{
		Message: message,
		Error:   isError,
		Data:    data,
	}
}
