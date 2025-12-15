package disk

import (
	"fmt"
	"strings"
)

type Handler func(comando string, props map[string]string) (string, bool)

type CommandDef struct {
	Allowed  map[string]bool
	Required []string
	Defaults map[string]string
	Run      Handler
}

var commands = map[string]CommandDef{
	"mkdisk": {
		Allowed: map[string]bool{
			"size": true, "fit": true, "unit": true,
		},
		Required: []string{"size"},
		Defaults: map[string]string{"fit": "FF", "unit": "M"},
		Run:      mkdiskExecute,
	},
	"rmdisk": {
		Allowed: map[string]bool{
			"diskname": true,
		},
		Required: []string{"diskname"},
		Defaults: map[string]string{},
		Run:      nil,
	},
	"fdisk": {
		Allowed: map[string]bool{
			"size": true, "unit": true, "diskname": true, "type": true, "fit": true, "name": true,
		},
		Required: []string{"size", "diskname", "name"},
		Defaults: map[string]string{"unit": "K", "type": "P", "fit": "FF"},
		Run:      fdiskExecute,
	},
	"mount": {
		Allowed: map[string]bool{
			"diskname": true, "name": true,
		},
		Required: []string{"diskname", "name"},
		Defaults: map[string]string{},
		Run:      nil,
	},
	"mounted": {
		Allowed:  map[string]bool{},
		Required: []string{},
		Defaults: map[string]string{},
		Run:      nil,
	},
	"mkfs": {
		Allowed: map[string]bool{
			"id": true, "type": true,
		},
		Required: []string{"id"},
		Defaults: map[string]string{"type": "FULL"},
		Run:      nil,
	},
}

func diskCommandProps(comando string, instrucciones []string) (string, bool) {
	// fmt.Println(comando, instrucciones)
	cmd := strings.ToLower(comando)
	def, ok := commands[cmd]

	if !ok {
		// return nil, fmt.Sprintf("Comando no reconocido: %s", comando), true
		return fmt.Sprintf("Comando no reconocido: %s", comando), true
	}

	// fmt.Println(def)
	allowedLower := make(map[string]bool, len(def.Allowed))
	for k := range def.Allowed {
		allowedLower[strings.ToLower(k)] = true
	}

	props := make(map[string]string)
	for k, v := range def.Defaults {
		props[strings.ToLower(k)] = v
	}

	seen := make(map[string]bool)

	// parseamos los parámetros
	for _, token := range instrucciones {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}

		if !strings.Contains(token, "=") {
			// return nil, fmt.Sprintf("Parámetro inválido: '%v'", token), true
			return fmt.Sprintf("Parámetro inválido: '%v'", token), true
		}

		parts := strings.SplitN(token, "=", 2)
		key := strings.ToLower(strings.TrimSpace(parts[0]))
		val := strings.TrimSpace(parts[1])

		if key == "" {
			// return nil, fmt.Sprintf("Parámetro inválido: '%s'", token), true
			return fmt.Sprintf("Parámetro inválido: '%s'", token), true
		}

		if !allowedLower[key] {
			// return nil, fmt.Sprintf("Parámetro no permitido para '%s': '%s'", cmd, key), true
			return fmt.Sprintf("Parámetro no permitido para '%s': '%s'", cmd, key), true
		}

		if seen[key] {
			// return nil, fmt.Sprintf("Parámetro no permitido: %s", key), true
			return fmt.Sprintf("Parámetro no permitido: %s", key), true
		}

		seen[key] = true
		props[key] = val
	}

	// verificar valores mínimos
	for _, req := range def.Required {
		reqLower := strings.ToLower(req)
		if strings.TrimSpace(props[reqLower]) == "" {
			// return nil, fmt.Sprintf("Parámetro obligatorio faltante: %s", req), true
			return fmt.Sprintf("Parámetro obligatorio faltante: %s", req), true
		}
	}

	// spec, ok := commands[cmd]
	if def.Run == nil {
		return fmt.Sprintf("Comando que no tiene handler: %s", cmd), true
	}

	// fmt.Println(props)
	// fmt.Println(props["name"])

	// return props, "", false
	return def.Run(comando, props)
	// return fmt.Sprintf("Exito en la ejecución del comando: %s", comando), false
	// for _, tempParam := range instrucciones {
	// 	fmt.Println(comando, tempParam)

	// }
	// fmt.Println(instrucciones)
}

func DiskExecuteCommanWithProps(command string, parameters []string) {
	temp, ok := diskCommandProps(command, parameters)
	if !ok {
		return
	}

	fmt.Println(temp)
}
