package general

import (
	"strings"

	"github.com/fatih/color"
)

var commandGroups = map[string][]string{
	"disk":    {"mkdisk", "fdisk", "rmdisk", "mount", "mounted", "mkfs"},
	"reports": {"rep"},
	"files":   {"mkfile", "mkdir"},
	"cat":     {"cat"},
	"users":   {"login", "logout"},
	"groups":  {"mkgrp", "mkusr"},
}

func detectGroup(cmd string) (string, string, bool, string) {
	cmdLower := strings.ToLower(cmd)

	for group, cmds := range commandGroups {
		for _, prefix := range cmds {
			if strings.HasPrefix(cmdLower, prefix) {
				return group, prefix, false, ""
			}
		}
	}

	return "", "", true, "Comando no reconocido"
}

// error, mssgEror, comandos
func GlobalCom(lista []string) ([]string, int) {
	var errores []string
	var contErrores = 0

	for _, comm := range lista {
		// Administracion de Discos
		// if (strings.HasPrefix(strings.ToLower(comm), "mkdisk")) || (strings.HasPrefix(strings.ToLower(comm), "fdisk")) || (strings.HasPrefix(strings.ToLower(comm), "rmdisk")) || (strings.HasPrefix(strings.ToLower(comm), "mount")) || (strings.HasPrefix(strings.ToLower(comm), "mounted")) || (strings.HasPrefix(strings.ToLower(comm), "mkfs")) {
		// 	comandos := ObtenerComandos(comm)
		// 	// command := getCommand(strings.ToLower(comm), "mkdisk", "fdisk", "rmdisk", "mount", "unmount", "mkfs")
		// 	command := getCommand(strings.ToLower(comm), "mkdisk", "fdisk", "rmdisk", "mount", "mounted", "mkfs")
		// 	println(command, comandos)
		// 	// admindisk.DiskCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
		// 	// Reportes
		// } else if strings.HasPrefix(strings.ToLower(comm), "rep") {
		// 	comandos := ObtenerComandos(comm)
		// 	// report.ReportCommandProps("REP", comandos) // Comandos a Usar
		// 	// } else if strings.HasPrefix(strings.ToLower(comm), "pause") {
		// 	// 	Pause()
		// 	// Files
		// } else if (strings.HasPrefix(strings.ToLower(comm), "mkfile")) || (strings.HasPrefix(strings.ToLower(comm), "mkdir")) {
		// 	comandos := ObtenerComandos(comm)
		// 	// command := getCommand(strings.ToLower(comm), "mkfile", "remove", "edit", "rename", "mkdir", "copy", "move", "find")
		// 	command := getCommand(strings.ToLower(comm), "mkfile", "mkdir")
		// 	// filesystem.FilesCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
		// 	// Permisos
		// } else if strings.HasPrefix(strings.ToLower(comm), "cat") { // solo para cat
		// 	comandos := ObtenerComandos(comm)
		// 	command := getCommand(strings.ToLower(comm), "cat")
		// 	// filesystem.FilesCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
		// 	// return filesystem.FilesCommandProps(strings.ToUpper(command), comandos)

		// 	// } else if (strings.HasPrefix(strings.ToLower(comm), "chown")) || (strings.HasPrefix(strings.ToLower(comm), "chgrp")) || (strings.HasPrefix(strings.ToLower(comm), "chmod")) {
		// 	// 	comandos := ObtenerComandos(comm)
		// 	// 	command := getCommand(strings.ToLower(comm), "chown", "chgrp", "chmod")
		// 	// 	permitions.PermissionsCommandProps(strings.ToUpper(command), comandos)
		// 	// Usuarios
		// } else if (strings.HasPrefix(strings.ToLower(comm), "login")) || (strings.HasPrefix(strings.ToLower(comm), "logout")) {
		// 	comandos := ObtenerComandos(comm)
		// 	command := getCommand(strings.ToLower(comm), "login", "logout")
		// 	// adminusers.UserCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
		// 	// return adminusers.UserCommandProps(strings.ToUpper(command), comandos)
		// 	// Grupo
		// } else if (strings.HasPrefix(strings.ToLower(comm), "mkgrp")) || (strings.HasPrefix(strings.ToLower(comm), "mkusr")) {
		// 	comandos := ObtenerComandos(comm)
		// 	// command := getCommand(strings.ToLower(comm), "mkgrp", "rmgrp", "mkusr", "rmusr")
		// 	command := getCommand(strings.ToLower(comm), "mkgrp", "mkusr")
		// 	// adminusers.GroupCommandProps(strings.ToUpper(command), comandos) // Comandos a Usar
		// } else {
		// 	color.Red("Comando no Reconocido")
		// }

		group, command, blnError, strError := detectGroup(comm)
		if blnError {
			color.Red("Comando no reconocido %v", command)
			errores = append(errores, strError)
			contErrores++
		}

		comandos := ObtenerComandos(comm)
		// fmt.Println("00000000000000000000000000000")
		// fmt.Println(group, command)
		switch group {
		case "disk":
			color.Cyan("Administración de discos: %v", command)

		case "reports":
			color.Red("Administración de reportes: %v", command)

		case "files":
			color.Green("Administración de Archivos: %v", command)

		case "cat":
			color.Blue("Comando CAT")

		case "users":
			color.Yellow("Administración de Usuarios: %v", command)

		case "groups":
			color.White("Administración de Grrupos: %v", command)
		}
	}

	return errores, contErrores
}
