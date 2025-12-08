package size

import (
	// "proyecto/estructuras/structures"
	"Proyecto/Estructuras/structures"
	"unsafe"
)

func SizeEBR() int32 { //30 bytes
	a01 := unsafe.Sizeof(structures.EBR{}.Part_mount)
	a01 += unsafe.Sizeof(structures.EBR{}.Part_fit)
	a01 += unsafe.Sizeof(structures.EBR{}.Part_start)
	a01 += unsafe.Sizeof(structures.EBR{}.Part_s)
	a01 += unsafe.Sizeof(structures.EBR{}.Part_next)
	a01 += unsafe.Sizeof(structures.EBR{}.Name)
	return int32(a01)
}

func SizePartition() int32 { //35 bytes
	a01 := unsafe.Sizeof(structures.Partition{}.Part_status)
	a01 += unsafe.Sizeof(structures.Partition{}.Part_type)
	a01 += unsafe.Sizeof(structures.Partition{}.Part_fit)
	a01 += unsafe.Sizeof(structures.Partition{}.Part_start)
	a01 += unsafe.Sizeof(structures.Partition{}.Part_s)
	a01 += unsafe.Sizeof(structures.Partition{}.Part_name)
	a01 += unsafe.Sizeof(structures.Partition{}.Part_correlative)
	a01 += unsafe.Sizeof(structures.Partition{}.Part_id)
	return int32(a01)
}

func SizeMBR() int32 { //153 bytes
	a01 := unsafe.Sizeof(structures.MBR{}.Mbr_tamano)
	a01 += unsafe.Sizeof(structures.MBR{}.Mbr_fecha_creacion)
	a01 += unsafe.Sizeof(structures.MBR{}.Mbr_disk_signature)
	a01 += unsafe.Sizeof(structures.MBR{}.Dsk_fit)
	a01 += uintptr(SizePartition() * 4)
	return int32(a01)
}

func SizeMBR_NotPartitions() int32 {
	a01 := unsafe.Sizeof(structures.MBR{}.Mbr_tamano)
	a01 += unsafe.Sizeof(structures.MBR{}.Mbr_fecha_creacion)
	a01 += unsafe.Sizeof(structures.MBR{}.Mbr_disk_signature)
	a01 += unsafe.Sizeof(structures.MBR{}.Dsk_fit)
	return int32(a01)
}
