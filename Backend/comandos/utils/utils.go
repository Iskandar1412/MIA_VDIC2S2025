package utils

import (
	"Proyecto/Estructuras/structures"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var DirectorioDisco = "VDIC-MIA/Disks/"

func esEntero(valor string) (int32, bool, string) {
	i, err := strconv.Atoi(valor)
	if err != nil {
		return 0, true, "Error en la conversión a entero"
	}

	if i <= 0 {
		return 0, true, "Valor entero menor o igual a 0"
	}

	return int32(i), false, ""
}

func TieneSize(comando string, size string) (int32, bool, string) {
	salida, er, strmsg := esEntero(size)
	if er {
		return salida, er, fmt.Sprintf("%s - Comando: %s", strmsg, comando)
	}

	return salida, false, ""
}

func ObFechaInt() int32 {
	fecha := time.Now()
	timestamp := fecha.Unix()
	//fmt.Println(timestamp)
	return int32(timestamp)
}

func IntFechaToStr(fecha int32) string {
	conversion := int64(fecha)
	formato := "2006/01/02 (15:04:05)"
	fech := time.Unix(conversion, 0)
	fechaFormat := fech.Format(formato)
	//fmt.Println(fechaFormat)
	return fechaFormat
}

var unitRules = map[string]struct {
	Default byte
	Allowed map[string]bool
}{
	"mkdisk": {Default: 'M', Allowed: map[string]bool{"K": true, "M": true}},
	"fdisk":  {Default: 'K', Allowed: map[string]bool{"B": true, "K": true, "M": true}},
}

func TieneUnit(command string, unit string) (byte, bool, string) {
	command = strings.ToLower(command)

	temp, ok := unitRules[command]
	if !ok {
		return 0, false, fmt.Sprintf("El comando %s no maneja unit", command)
	}

	raw := strings.TrimSpace(unit)
	if raw == "" {
		return temp.Default, false, ""
	}

	u := strings.ToUpper(raw)
	if !temp.Allowed[u] {
		return temp.Default, true, fmt.Sprintf("[%s] unidad invalida: %s - (valores permitidos: %v)", command, u, temp.Allowed)
	}

	return u[0], false, ""
}

var fitRules = map[string]struct {
	Default string
	Allowed map[string]bool
}{
	"-": {Default: "FF", Allowed: map[string]bool{"BF": true, "WF": true, "FF": true}},
}

func TieneFit(command string, fit string) (byte, bool, string) {
	rule, ok := fitRules["-"]
	if !ok {
		// fmt.Println("err0")
		return 0, true, fmt.Sprintf("Valor no existente: %s", command)
	}

	fit = strings.ToUpper(strings.TrimSpace(fit))
	if fit == "" {
		// fmt.Println("erra")
		return 'F', true, "" // default "FF" => 'F'
	}

	if !rule.Allowed[fit] {
		// fmt.Println("errb")
		return 0, true, fmt.Sprintf("Fit invalido: %s", fit)
	}

	// (F/W/B)
	switch fit {
	case "FF":
		// fmt.Println("FF asegurado")
		return 'F', false, ""
	case "WF":
		// fmt.Println("WF asegurado")
		return 'W', false, ""
	case "BF":
		// fmt.Println("BF asegurado")
		return 'B', false, ""
	default:
		return 0, true, fmt.Sprintf("fit inválido: %s", fit)
	}
}

func ObtenerTamanioDisco(size int32, unidad byte) int32 {
	switch unidad {
	case 'B':
		return size
	case 'K':
		return size * 1024
	case 'M':
		return size * 1024 * 1024
	default:
		return 0
	}
}

func ObtenerDiskSignature() int32 {
	source := rand.NewSource(time.Now().UnixNano())
	numberR := rand.New(source)
	signature := numberR.Intn(1000000) + 1
	//fmt.Println(signature)
	return int32(signature)
}

func NuevaPartitionVacia() structures.Partition {
	var partition structures.Partition
	partition.Part_status = int8(-1)
	partition.Part_type = 'P'
	partition.Part_fit = 'F'
	partition.Part_start = -1
	partition.Part_s = -1
	for i := 0; i < len(partition.Part_name); i++ {
		partition.Part_name[i] = '\x00'
	}
	partition.Part_correlative = -1
	for i := 0; i < len(partition.Part_id); i++ {
		partition.Part_id[i] = '\x00'
	}
	return partition
}
