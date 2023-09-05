package models

type QuantRamaisOpen struct {
	Ramal int  `json:"ramal"`
	INUSE bool `json:"INUSE"`
}

type ClienteConfig struct {
	ID              int               `json:"id"`
	Doc             int64             `json:"doc"`
	Cliente         string            `json:"cliente"`
	QuantRamaisOpen []QuantRamaisOpen `json:"quantRamaisOpen"`
	GrupoRecurso    string            `json:"grupoRecurso"`
	LinkGvc         string            `json:"linkGvc"`
	Porta           string            `json:"porta"`
	Ramal           string            `json:"ramal"`
	Senha           string            `json:"senha"`
}

type Cliente struct {
	Clientes []ClienteConfig `json:"clientes"`
}

func AdicionarCliente(mapa map[string]interface{}, id int, doc int, cliente string, grupoRecurso string, linkGvc string, porta string, ramal interface{}, senha string) {
	clientes := []map[string]interface{}{
		{
			"id":              id,
			"doc":             doc,
			"cliente":         cliente,
			"grupoRecurso":    grupoRecurso,
			"linkGvc":         linkGvc,
			"porta":           porta,
			"ramal":           ramal,
			"senha":           senha,
			"quantRamaisOpen": []map[string]interface{}{},
		},
	}

	mapa["clientes"] = clientes
}

func AdicionarRamal(mapa map[string]interface{}, clienteID int, ramalNum int, inUse bool) {
	clientes := mapa["clientes"].([]map[string]interface{})
	for i := range clientes {
		if clientes[i]["id"].(int) == clienteID {
			ramais := clientes[i]["quantRamaisOpen"].([]map[string]interface{})
			ramal := map[string]interface{}{
				"ramal": 780 + ramalNum,
				"INUSE": inUse,
			}
			clientes[i]["quantRamaisOpen"] = append(ramais, ramal)
			break
		}
	}
}
