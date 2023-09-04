package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	jsonData, _ := BuscaPorDoc(12310400000182)

	fmt.Println(string(jsonData))

	cliente := map[byte]interface{}{byte(jsonData)}

	for _, v := range cliente {

	}

	// quantRamaisOpen := make(map[string]interface{ jsonData })

	// // Exemplo de como acessar valores no mapa
	// clientes := quantRamaisOpen["clientes"].([]map[string]interface{})
	// primeiroCliente := clientes[1]
	// fmt.Println("ID do primeiro cliente:", primeiroCliente["id"])

	// // Exemplo de como acessar valores em quantRamaisOpen
	// quantRamais := primeiroCliente["quantRamaisOpen"].([]map[string]interface{})
	// fmt.Println("Estado do ramal 7801:", quantRamais[0]["INUSE"])
}

func BuscaPorDoc(doc int) ([]byte, error) {
	//doc := 12310400000182
	// ajustar porta
	url := "http://localhost:3004/clientes?doc=" + strconv.Itoa(doc)
	method := "GET"

	//fmt.Println(url)

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		//return

	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		//return

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		//return

	}
	fmt.Println(string(body))

	return body, nil
}

func adicionarCliente(mapa map[string]interface{}, id int, doc int, cliente string, grupoRecurso string, linkGvc string, porta string, ramal interface{}, senha string) {
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

func adicionarRamal(mapa map[string]interface{}, clienteID int, ramalNum int, inUse bool) {
	clientes := mapa["clientes"].([]map[string]interface{})
	for i := range clientes {
		if clientes[i]["id"].(int) == clienteID {
			ramais := clientes[i]["quantRamaisOpen"].([]map[string]interface{})
			ramal := map[string]interface{}{
				"ramal": ramalNum,
				"INUSE": inUse,
			}
			clientes[i]["quantRamaisOpen"] = append(ramais, ramal)
			break
		}
	}
}
