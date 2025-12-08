package size

import (
	// "proyecto/estructuras/structures"
	"Proyecto/Estructuras/structures"
	"unsafe"
)

func SizeContent() int32 { //16bytes
	a01 := unsafe.Sizeof(structures.Content{}.B_name)
	a01 += unsafe.Sizeof(structures.Content{}.B_inodo)
	// result := a01
	return int32(a01)
}

func SizeBloqueCarpeta() int32 { //64 bytes

	a01 := unsafe.Sizeof(structures.BloqueCarpeta{}.B_content[0].B_name) + unsafe.Sizeof(structures.BloqueCarpeta{}.B_content[0].B_inodo)
	a01 += unsafe.Sizeof(structures.BloqueCarpeta{}.B_content[1].B_name) + unsafe.Sizeof(structures.BloqueCarpeta{}.B_content[1].B_inodo)
	a01 += unsafe.Sizeof(structures.BloqueCarpeta{}.B_content[2].B_name) + unsafe.Sizeof(structures.BloqueCarpeta{}.B_content[2].B_inodo)
	a01 += unsafe.Sizeof(structures.BloqueCarpeta{}.B_content[3].B_name) + unsafe.Sizeof(structures.BloqueCarpeta{}.B_content[3].B_inodo)
	return int32(a01)
}

func SizeBloqueArchivo() int32 { //64 bytes
	a01 := unsafe.Sizeof(structures.BloqueArchivo{}.B_content)
	return int32(a01)
}

func SizeBloqueApuntador() int32 { //64 bytes
	a01 := unsafe.Sizeof(structures.BloqueApuntador{}.B_pointers)
	return int32(a01)
}
