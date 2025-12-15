package general

import (
	"Proyecto/comandos/commandGroups/disk"
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
		group, command, blnError, strError := detectGroup(comm)
		if blnError {
			color.Red("Comando no reconocido %v", command)
			errores = append(errores, strError)
			contErrores++
			continue
		}

		comandos := ObtenerParametros(comm)
		// fmt.Println("00000000000000000000000000000")
		// fmt.Println(group, command)
		switch group {
		case "disk":
			color.Cyan("Administración de discos: %v", command)
			disk.DiskExecuteCommanWithProps(command, comandos)

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
