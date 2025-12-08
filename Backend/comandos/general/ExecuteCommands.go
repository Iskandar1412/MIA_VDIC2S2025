package general

import (
	"strings"

	"github.com/fatih/color"
)

func GlobalCom(lista []string) bool {
	for _, comm := range lista {
		// Administracion de Discos
		if (strings.HasPrefix(strings.ToLower(comm), "mkdisk")) || (strings.HasPrefix(strings.ToLower(comm), "fdisk")) || (strings.HasPrefix(strings.ToLower(comm), "rmdisk")) || (strings.HasPrefix(strings.ToLower(comm), "mount")) || (strings.HasPrefix(strings.ToLower(comm), "mounted")) || (strings.HasPrefix(strings.ToLower(comm), "mkfs")) {
			comandos := ObtenerComandos(comm)
			// command := getCommand(strings.ToLower(comm), "mkdisk", "fdisk", "rmdisk", "mount", "unmount", "mkfs")
			command := getCommand(strings.ToLower(comm), "mkdisk", "fdisk", "rmdisk", "mount", "mounted", "mkfs")
			// admindisk.DiskCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
			// Reportes
		} else if strings.HasPrefix(strings.ToLower(comm), "rep") {
			comandos := ObtenerComandos(comm)
			// report.ReportCommandProps("REP", comandos) // Comandos a Usar
			// } else if strings.HasPrefix(strings.ToLower(comm), "pause") {
			// 	Pause()
			// Files
		} else if (strings.HasPrefix(strings.ToLower(comm), "mkfile")) || (strings.HasPrefix(strings.ToLower(comm), "mkdir")) {
			comandos := ObtenerComandos(comm)
			// command := getCommand(strings.ToLower(comm), "mkfile", "remove", "edit", "rename", "mkdir", "copy", "move", "find")
			command := getCommand(strings.ToLower(comm), "mkfile", "mkdir")
			// filesystem.FilesCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
			// Permisos
		} else if strings.HasPrefix(strings.ToLower(comm), "cat") { // solo para cat
			comandos := ObtenerComandos(comm)
			command := getCommand(strings.ToLower(comm), "cat")
			// filesystem.FilesCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
			// return filesystem.FilesCommandProps(strings.ToUpper(command), comandos)

			// } else if (strings.HasPrefix(strings.ToLower(comm), "chown")) || (strings.HasPrefix(strings.ToLower(comm), "chgrp")) || (strings.HasPrefix(strings.ToLower(comm), "chmod")) {
			// 	comandos := ObtenerComandos(comm)
			// 	command := getCommand(strings.ToLower(comm), "chown", "chgrp", "chmod")
			// 	permitions.PermissionsCommandProps(strings.ToUpper(command), comandos)
			// Usuarios
		} else if (strings.HasPrefix(strings.ToLower(comm), "login")) || (strings.HasPrefix(strings.ToLower(comm), "logout")) {
			comandos := ObtenerComandos(comm)
			command := getCommand(strings.ToLower(comm), "login", "logout")
			// adminusers.UserCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
			// return adminusers.UserCommandProps(strings.ToUpper(command), comandos)
			// Grupo
		} else if (strings.HasPrefix(strings.ToLower(comm), "mkgrp")) || (strings.HasPrefix(strings.ToLower(comm), "mkusr")) {
			comandos := ObtenerComandos(comm)
			// command := getCommand(strings.ToLower(comm), "mkgrp", "rmgrp", "mkusr", "rmusr")
			command := getCommand(strings.ToLower(comm), "mkgrp", "mkusr")
			// adminusers.GroupCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
		} else {
			color.Red("Comando no Reconocido")
		}
	}

	return true
}
