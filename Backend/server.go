package main

import (
	"Proyecto/comandos/controllers"
	"Proyecto/comandos/general"
	"Proyecto/middlewares"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	puerto := 9700
	// Configurar CORS
	c := cors.AllowAll()

	// Manejar las rutas
	mux.HandleFunc("/commands", controllers.HandleCommand)
	// mux.HandleFunc("/login", handleLogin)
	// mux.HandleFunc("/logout", handleLogout)
	// mux.HandleFunc("/obtainmbr", handleObtainMBR)
	// mux.HandleFunc("/reportesobtener", handleReportsObtener)
	// mux.HandleFunc("/graphs", handleGraph)
	// mux.HandleFunc("/obtain-carpetas-archivos", handleObtainCarpetasArchivos)
	// mux.HandleFunc("/cat", handleCat) //-------Nuevo

	// handler := c.Handler(mux)
	handler := middlewares.RecoverMiddleware(c.Handler(mux))

	fmt.Println("" + fmt.Sprintf("Backend server is on %v", puerto))
	general.CrearCarpeta()
	// obtencionpf.ObtenerMBR_Mounted()
	// obtencionpf.MostrarParticionesMontadas()
	// http.ListenAndServe(":8080", handler)
	err := http.ListenAndServe(":"+fmt.Sprintf("%v", puerto), handler)
	if err != nil {
		fmt.Println("ERROR al iniciar servidor:", err)
	}
}
