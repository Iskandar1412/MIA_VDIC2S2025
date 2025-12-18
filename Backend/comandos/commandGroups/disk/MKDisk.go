package disk

import (
	"Proyecto/Estructuras/structures"
	"Proyecto/comandos/utils"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

func mkdiskExecute(comando string, parametros map[string]string) (string, bool) {

	tamanio, er, msg := utils.TieneSize(comando, parametros["size"])
	if er {
		return msg, er
	}

	unidad, er, msg := utils.TieneUnit(comando, parametros["unit"])
	if er {
		return msg, er
	}

	fit, er, msg := utils.TieneFit(comando, parametros["fit"])
	// fmt.Println(er)
	if er {
		return msg, er
	}

	// fmt.Println(tamanio, unidad, fit)
	mkdisk_Create(tamanio, unidad, fit)

	return "Ok", true
}

func mkdisk_Create(_size int32, _unit byte, _fit byte) (string, bool) {
	for i := 0; i < 26; i++ {
		nombreDisco := fmt.Sprintf("VDIC-%c.mia", 'A'+i)
		archivo := utils.DirectorioDisco + nombreDisco
		if _, err := os.Stat(archivo); os.IsNotExist(err) {
			// Creamos archivo
			er, strmsg := createDiskFile(archivo, _size, _fit, _unit)
			if er {
				return strmsg, er
			}

			color.Green("[MKDISK]: Disco '" + nombreDisco + "' Creado -> " + strconv.Itoa(int(_size)) + string(_unit))
			return "", false
		} else {
			// Caso en el que el disco ya exista
			continue
		}
	}

	return "No hay mas para discos", true
}

func createDiskFile(archivo string, tamanio int32, fit byte, unidad byte) (bool, string) {
	file, err := os.Create(archivo)
	if err != nil {
		color.Red("Error al crear el archivo")
		return true, "Error al crear el archivo"
	}
	defer file.Close()

	var estructura structures.MBR

	tamanioDiscco := utils.ObtenerTamanioDisco(tamanio, unidad)
	estructura.Mbr_tamano = tamanioDiscco
	estructura.Mbr_fecha_creacion = utils.ObFechaInt()
	estructura.Mbr_disk_signature = utils.ObtenerDiskSignature()
	estructura.Dsk_fit = fit
	for i := 0; i < len(estructura.Mbr_partitions); i++ {
		estructura.Mbr_partitions[i] = utils.NuevaPartitionVacia()
	}

	bytes_llenar := make([]byte, int(tamanioDiscco))
	if _, err := file.Write(bytes_llenar); err != nil {
		color.Red("Error al escribir bytes en el disco")
		return true, "Error al escribir bytes en el disco"
	}

	// Cambio de posición del puntero
	if _, err := file.Seek(0, 0); err != nil {
		color.Red("Error al mover puntero del archivo")
		return true, "Error al mover puntero del archivo"
	}

	if err := binary.Write(file, binary.LittleEndian, &estructura); err != nil {
		color.Red("Error al escribir datos del MBR")
		return true, "Error al escribir datos del MBR"
	}

	return false, ""
}
