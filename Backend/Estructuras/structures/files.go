package structures

//ESTRUCTURAS PARA CARPETAS Y ARCHIVOS

type SuperBloque struct { //68 bytes
	S_filesistem_type   int32 //guarda numero que identifica el sistema de archivos utilizado
	S_inodes_count      int32 //guarda numero total de inodos
	S_blocks_count      int32 //guarda numero total de bloques
	S_free_blocks_count int32 //contiene numero de bloques libres
	S_free_inodes_count int32 //contiene numero de inodos libres
	S_mtime             int32 //ultima fecha en el que el sistema fue montado (time)
	S_umtime            int32 //ultima fecha en el que el sistema fue desmontado (time)
	S_mnt_count         int32 //indica cuantas veces se ha montado el sistema
	S_magic             int32 //valor que identifica el sistema de archivos, tendra valor 0xEF53
	S_inode_s           int32 //tamaño del inodo
	S_block_s           int32 //tamaño del bloque
	S_first_ino         int32 //primer inodo libre
	S_first_blo         int32 //primer bloque libre
	S_bm_inode_start    int32 //guarda el inicio del bitmap de inodos
	S_bm_block_start    int32 //guarda el inicio del bitmap de bloques
	S_inode_start       int32 //guarda el inicio de la tabla de inodos
	S_block_start       int32 //guarda inicio de la tala de bloques
}

type TablaInodo struct { //92 bytes
	I_uid   int32     //UID del usuario propietario del archivo o carpeta
	I_gid   int32     //GID del grupo al que pertenece el archivo o carpeta
	I_s     int32     //tamaño del archivo en bytes
	I_atime int32     //ultima fecha en que se leyo el inodo sin modificarlo
	I_ctime int32     //fecha en la que se creo el inodo
	I_mtime int32     //ultima fecha en la que se modifica el inodo
	I_block [15]int32 //array en los que los primeros 12 registros son bloques directos si no son utilizados valor -1
	I_type  int32     //indica si es archivo o carpeta (1 = archivo, 2 = carpeta)
	I_perm  int32     //guarda los permisos del archivo R (permiso de lectura) W (permiso escritura) X (permiso ejecucion)
}
