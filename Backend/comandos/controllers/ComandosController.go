package controllers

import (
	"Proyecto/comandos/general"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func HandleCommand(w http.ResponseWriter, r *http.Request) {
	// command
	if r.Method == http.MethodOptions {
		// Establecer encabezados CORS para las solicitudes OPTIONS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Solo permitir solicitudes POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Configurar encabezados CORS para las solicitudes POST
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Decodificar el cuerpo JSON de la solicitud
	var requestBody struct {
		Comandos *string `json:"Comandos"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&requestBody)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			general.ResultadoSalida("JSON inválido o campos no permitidos", true, nil),
		)
		return
	}

	if requestBody.Comandos == nil || strings.TrimSpace(*requestBody.Comandos) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			general.ResultadoSalida("El campo 'Comandos' es obligatorio y no puede ser nulo", true, nil),
		)
		return
	}

	var ejecutar []string
	ejecutar = append(ejecutar, *requestBody.Comandos)
	comando := strings.Split(ejecutar[0], "\n")
	tempComandos := general.ExecuteCommandList(comando)
	salida, ok := tempComandos.Respuesta.(general.SalidaComandoEjecutado)
	if !ok {
		// http.Error(w, "Error al obtener comandos", http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			general.ResultadoSalida("Error al obtener comandos", true, nil),
		)
		return
		// return
	}

	errores, contadorErrorres := general.GlobalCom(salida.LstComandos)
	fmt.Println(errores, contadorErrorres)

	// comandos.GlobalCom(ejecutar)
	// obtencionpf.ObtenerMBR_Mounted()
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(general.ResultadoSalida("", false, salida.LstComandos))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			general.ResultadoSalida("Error al leer el cuerpo de la solicitud", true, nil),
		)
		return
	}
}
