package utils

import (
	"Proyecto/Estructuras/size"
	"Proyecto/Estructuras/structures"
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
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

func TieneType(tipo string) (byte, bool, string) {
	switch strings.ToUpper(tipo) {
	case "P":
		return 'P', false, ""
	case "E":
		return 'E', false, ""
	case "L":
		return 'E', false, ""
	default:
		return 0, true, "Tipo no reconocido"
	}
}

// diskname, blnError, strError
func TieneDiskName(diskName string) (string, bool, string) {
	if diskName == "" {
		return "", true, "Valor invalido para el DiskName"
	}
	return diskName, false, ""
}

func TieneName(name string) (string, bool, string) {
	if name == "" {
		return "", true, "Valor invalido para el Name"
	}
	return name, false, ""
}

func ExisteArchivo(comando string, pathArchivo string) bool {
	if _, err := os.Stat(pathArchivo); os.IsNotExist(err) {
		color.Red("[" + comando + "]: Archivo no encontrado")
		return false
	}

	return true
}

func ObtenerEstructuraMBR(pathDisco string) (structures.MBR, bool, string) {
	mbr := structures.MBR{}

	file, err := os.OpenFile(pathDisco, os.O_RDWR, 0666)
	if err != nil {
		color.Red("[utils.ln:205] Error en la lectura del disco")
		return structures.MBR{}, true, "[utils.ln:205] Error en la lectura del disco"
	}
	defer file.Close()

	if _, err := file.Seek(0, 0); err != nil {
		color.Red("[utils.ln:210]: Error al mover el puntero")
		return structures.MBR{}, true, "[utils.ln:210]: Error al mover el puntero"
	}

	if err := binary.Read(file, binary.LittleEndian, &mbr); err != nil {
		color.Red("[utils.ln:217]: Error en la lectura del MBR")
		return structures.MBR{}, true, "[utils.ln:217]: Error en la lectura del MBR"
	}

	return mbr, false, ""
}

func ConvertirByteAString(arreglo []byte) string {
	nullindex := bytes.IndexByte(arreglo, 0)
	if nullindex == -1 {
		return string(arreglo)
	}

	return string(arreglo[:nullindex])
}

func ExisteNombreParticion(pathDisco string, nombreParticion string) (bool, string) {
	mbr, er, strError := ObtenerEstructuraMBR(pathDisco)
	if er {
		return true, strError
	}

	for i := range mbr.Mbr_partitions {
		if ConvertirByteAString(mbr.Mbr_partitions[i].Part_name[:]) == nombreParticion {
			// Aquí continuamos posterior a lo de la clase
			// como el nombre es igual al que tenemos vamos a retornar que ya existe
			return true, "[utils.line:244]: Nombre ya existente"
		}

		// se va a iterrar en el mbr para ver el tipo de dico
		//En caso que sea uno extendido tendremos que leer lo que es el EBR si existe
		// e iterar para buscar todas las logicas en caso que existan
		if mbr.Mbr_partitions[i].Part_type == 'E' {
			// declaramoss una variable que tenga la estructura de EBR
			ebr := structures.EBR{}

			//vamos a leer el archivo y si hay error se retornara ello
			file, err := os.OpenFile(pathDisco, os.O_RDWR, 0666)
			if err != nil {
				return true, "[utils.line:257]: Error en abrir el archivo"
			}
			defer file.Close()

			// vamos a mover el puntero en la posición donde se encuentre
			// la partición extendida
			if _, err := file.Seek(int64(mbr.Mbr_partitions[i].Part_start), 0); err != nil {
				return true, "[utils.line:264]: Error al mover el puntero"
			}

			// se va a leer el archivo y asignar al ebr su valor
			if err := binary.Read(file, binary.LittleEndian, &ebr); err != nil {
				return true, "[utils.line:269]: Error en la lectura del EBR"
			}

			// ahora veremos si tiene ebr que le siga (otra particion lógica)
			if ebr.Part_next != -1 || ebr.Part_s != -1 {
				// vamos a comparar nombres de particiones
				if ConvertirByteAString(ebr.Name[:]) == nombreParticion {
					return true, "[utils.line:276]: El nombre de la partición a poner ya existe"
				}

				// verificamos si existe un siguiente ebr mediante ciclos
				for ebr.Part_next != -1 {
					if ConvertirByteAString(ebr.Name[:]) == nombreParticion {
						return true, "[utils.line:282]: Nombre de la partición existente"
					}

					// movemos puntero para el siguiente ebr
					if _, err := file.Seek(int64(ebr.Part_next), 0); err != nil {
						return true, "[utils.line:287]: Error al mover el puntero"
					}

					// leemos y asignamos al ebr el nuevo valor del siguiente
					if err := binary.Read(file, binary.LittleEndian, &ebr); err != nil {
						return true, "[utils.292]: Error en la lectura del EBR"
					}

					if ConvertirByteAString(ebr.Name[:]) == ConvertirByteAString([]byte(nombreParticion)) {
						return true, "[utils.line:296]: Nombre de la partición ya existente"
					}
				}
			}
		}
	}

	// Caso que el nombre por más que se busque
	// No exista, implica ello
	// Se puede proseguir con el nombre
	return false, ""
}

// Necesitamos ver si existe una partición extendida para ver si se puede
// crear o ver una lógica
func ExisteParticionExtendida(pathDisco string) bool {
	// se obtendrá todo el mbr
	mbr, err, strMensajeErr := ObtenerEstructuraMBR(pathDisco)
	if err {
		fmt.Println(strMensajeErr)
		return err
	}

	// vamos a iterar por las particiones para ver si existe una extendida
	for i := range mbr.Mbr_partitions {
		if mbr.Mbr_partitions[i].Part_type == 'E' {
			return true
		}
	}

	// Retronamos F ssi no existe
	return false
}

// Esta funcion es para ver si hay espacio para crear una partición
// Primaria/Extendída o lógica
func ExisteEspacioDisponible(tamanio int32, pathDisco string, unidad byte, posicion int32) bool {
	mbr, err, strMensajeErr := ObtenerEstructuraMBR(pathDisco)
	if err {
		fmt.Println(strMensajeErr)
		return err
	}

	// vamos a verificar que la posición inicial del disco no sea -1
	// -1 indicaría que no existe la partición
	if posicion <= -1 {
		return false
	}

	// si el tamaño es menor a 0
	tamanioDisco := ObtenerTamanioDisco(tamanio, unidad)
	if tamanioDisco <= 0 {
		return false
	}

	// verificar el tamaño del disco por donde empieza el nuevo y el tamaño
	// con respecto al tamaño total
	// en sí una resta entre una posición final con lo que es la posición de la partición
	espacioDisponible := 0
	if posicion == 0 {
		espacioDisponible = int(mbr.Mbr_tamano) - int(size.SizeMBR())
	} else {
		espacioDisponible = int(mbr.Mbr_tamano) - int(mbr.Mbr_partitions[posicion-1].Part_start) - int(mbr.Mbr_partitions[posicion-1].Part_s)
	}

	return espacioDisponible >= int(tamanioDisco)
}

// función para volver a bytes el arreglo string (nombres/etc)
func ConvertirStringAByte(contenido string, cantidadBytes int) []byte {
	textoDevolver := make([]byte, cantidadBytes)
	copy(textoDevolver[:], []byte(contenido))
	return textoDevolver
}
