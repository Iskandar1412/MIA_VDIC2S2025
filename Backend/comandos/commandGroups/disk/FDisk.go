package disk

import (
	"Proyecto/Estructuras/size"
	"Proyecto/Estructuras/structures"
	"Proyecto/comandos/utils"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// P = Primario
// E = Extendido
// L = Logico
func fdiskExecute(comando string, parametros map[string]string) (string, bool) {
	tamanio, er, strError := utils.TieneSize(comando, parametros["size"])
	if er {
		return strError, er
	}

	unidad, er, strError := utils.TieneUnit(comando, parametros["unit"])
	if er {
		return strError, er
	}

	diskName, er, strError := utils.TieneDiskName(parametros["diskname"])
	if er {
		return strError, er
	}

	tipo, er, strError := utils.TieneType(parametros["type"])
	if er {
		return strError, er
	}

	tipoArreglo, er, strError := utils.TieneFit(comando, parametros["type"])
	if er {
		return strError, er
	}

	nombreParticion, er, strError := utils.TieneName(parametros["type"])
	if er {
		return strError, er
	}

	return fdiskCreate(tamanio, unidad, diskName, tipo, tipoArreglo, nombreParticion)
	// return "", true
}

func fdiskCreate(tamanio int32, unidad byte, diskName string, tipo byte, tipoFit byte, nombreParticion string) (string, bool) {
	var nombreSinExtension = ""
	var extensionArchivo = ""
	if strings.Contains(diskName, ".") {
		nombreSinExtension = strings.Split(diskName, ".")[0]
		extensionArchivo = strings.Split(diskName, ".")[1]
	}

	if extensionArchivo != "mia" {
		return "Extension del archivo no valida", true
	}

	nombreSinExtension = nombreSinExtension + ".mia"

	path := utils.DirectorioDisco + nombreSinExtension
	// tamanioDisco := utils.ObtenerTamanioDisco(tamanio, unidad)

	switch tipo {
	case 'P':
		fmt.Println("ParticionPrimaria")
		particionPrimaria(path, nombreParticion, tipo, tamanio, tipoFit, unidad)
		return "", false

	case 'E':
		fmt.Println("ParticionExtendida")
		return "", false

	case 'L':
		fmt.Println("ParticionLogica")
		return "", false

	default:
		return "Tipo para particion desconocido", true
	}
}

func particionPrimaria(ubicacionArchivo string, nombreParticion string, tipo byte, tamanioDisco int32, tipoArreglo byte, unidad byte) (string, bool) {
	if !utils.ExisteArchivo("FDISK", ubicacionArchivo) {
		color.Yellow("[FDISK]: Disco <<" + ubicacionArchivo + ">> no encontrado")
		return "Disco no encontrado: " + ubicacionArchivo, true
	}

	particion := utils.NuevaPartitionVacia()
	mbr, er, strError := utils.ObtenerEstructuraMBR(ubicacionArchivo)
	if er {
		color.Red(strError)
		return strError, er
	}

	pos := -1
	for i := range mbr.Mbr_partitions {
		if mbr.Mbr_partitions[i].Part_start == -1 {
			pos = i
			break
		}
	}

	// Continuación de la clase del 17
	// Verificamos que el nombre exista para evitar que se repita
	nombreExistente, strMensajeError := utils.ExisteNombreParticion(ubicacionArchivo, nombreParticion)
	if nombreExistente {
		return strMensajeError, nombreExistente
	}

	// ahora que pasamoss del nombre que es aceptado y no existe como tal
	// procedemos a ver si hay espacio como tal
	blnEspacioDisponible := utils.ExisteEspacioDisponible(tamanioDisco, ubicacionArchivo, unidad, int32(pos))
	if !blnEspacioDisponible {
		return "Espacio insuficiente", true
	}

	// if utils.ExisteNombreParticion()
	particion.Part_fit = tipoArreglo
	particion.Part_type = tipo
	particion.Part_name = [16]byte(utils.ConvertirStringAByte(string(nombreParticion), 16))
	particion.Part_status = -1
	particion.Part_correlative = utils.ObtenerDiskSignature()
	particion.Part_s = utils.ObtenerTamanioDisco(tamanioDisco, unidad)

	// Si la posición es la inicial (primer partición para hacer)
	// simplemente se toma desde el tamaño del mbr
	// de lo contrario se verá dondonde está la partición anterior más el tamaño que tiene
	if pos == 0 {
		particion.Part_start = size.SizeMBR()
	} else {
		particion.Part_start = mbr.Mbr_partitions[pos-1].Part_start + mbr.Mbr_partitions[pos-1].Part_s
	}

	mbr.Mbr_partitions[pos] = particion
	file, err := os.OpenFile(ubicacionArchivo, os.O_RDWR, 0666)
	if err != nil {
		return "[disk.line:144]: Error al abrir archivo", true
	}
	defer file.Close()

	if _, err := file.Seek(0, 0); err != nil {
		return "[disk.line:149]: Error al mover el puntero", true
	}

	if err := binary.Write(file, binary.LittleEndian, &mbr); err != nil {
		return "[disk.line:153]: Error al escribir el MBR", true
	}
	file.Close()

	comprobacion := structures.MBR{}
	file, err = os.OpenFile(ubicacionArchivo, os.O_RDWR, 0666)
	if err != nil {
		// color.Red("[FDISK]: Error al abrir archivo")
		return "[disk.line:161]: Error al escribir el MBR", true
	}
	defer file.Close()
	if _, err := file.Seek(0, 0); err != nil {
		// color.Red("[FDISK]: Error en mover puntero")
		return "[disk.line:166]: Error al mover el puntero", true
	}
	if err := binary.Read(file, binary.LittleEndian, &comprobacion); err != nil {
		return "[disk.line:169]: Error al escribir el MBR", true
	}
	file.Close()
	color.Green("-----------------------------------------------------------")
	color.Blue("Se creo la particion #" + strconv.Itoa(int(comprobacion.Mbr_partitions[pos].Part_correlative)))
	color.Blue("Particion: " + utils.ConvertirByteAString(comprobacion.Mbr_partitions[pos].Part_name[:]))
	color.Blue("Tipo Primaria")
	color.Blue("Inicio: " + strconv.Itoa(int(comprobacion.Mbr_partitions[pos].Part_start)))
	color.Blue("Tamaño: " + strconv.Itoa(int(comprobacion.Mbr_partitions[pos].Part_s)))
	color.Green("-----------------------------------------------------------")

	return "", false
}
