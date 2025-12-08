package structures

//BLOQUES

type Content struct { //16 bytes
	B_name  [12]byte //nombe de carpeta o archivo
	B_inodo int32    //apuntador hacia un inodo asociado al archivo o carpeta
}

type BloqueCarpeta struct { //64 bytes
	B_content [4]Content //array con contenido de carpeta
}

type BloqueArchivo struct { //64 bytes
	B_content [64]byte //array con contenido del archivo
}

type BloqueApuntador struct { //64 bytes
	B_pointers [16]int32 //array con apuntadores a bloques (archivos o carpeta)
}
