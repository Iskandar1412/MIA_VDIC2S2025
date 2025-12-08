package structures

// ESTRUCTURAS DISCOS

type EBR struct { //30 bytes
	Part_mount int8     //indica si esta montada o no
	Part_fit   byte     //tipo de ajuste B (best) F (first) W (worst)
	Part_start int32    //indica en que byte del disco inicia particion
	Part_s     int32    //contiene tamaño total de particion en bytes
	Part_next  int32    //byte en el que esta proximo EBR -1 si no hay siguiente
	Name       [16]byte //nombre de particion
}

type Partition struct { //38 bytes
	Part_status      int8     //indica si particion esta montada o no
	Part_type        byte     //indica el tipo de particion P (primaria) E (extendida)
	Part_fit         byte     //tipo de ajuste de particion B (best) F (first) W (worst)
	Part_start       int32    //indica en que byte del disco inicia la particion
	Part_s           int32    //contiene el tamaño total de la particion en bytes
	Part_name        [16]byte //nombre de particion
	Part_correlative int32    //correlativo de particion
	Part_id          [4]byte  //indica el id de particion generada al montar la misma
}

type MBR struct { //153bytes
	Mbr_tamano         int32        //tamaño total del disco en bytes
	Mbr_fecha_creacion int32        //fecha y hora de creacion del disco (time)
	Mbr_disk_signature int32        //numero random que identifica de forma unica cada disco
	Dsk_fit            byte         //tipo de ajuste de particion B (best) F (first) W (worst)
	Mbr_partitions     [4]Partition //estructura con informacion de las 4 particiones
}

type Particion_Enviar struct {
	Particion  string
	Type       string
	Status     int8
	Id_mounted string
}

type MBR_Obtener struct {
	Disco          string
	Disco_Path     string
	Mbr_partitions [4]Particion_Enviar
}
