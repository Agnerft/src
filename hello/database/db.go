package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hello/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var clienteConfig []models.ClienteConfig

func BuscaPorDoc(doc int) (string, error) {
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

	return string(body), nil
}

func AtualizarINUSE(resourceID int) {
	jsons, _ := BuscaPorDoc(12310400000182)

	// Defina a URL do servidor JSON Server e o JSON de atualização
	serverURL := "http://localhost:3004" // Substitua pela URL correta do seu servidor JSON Server

	for _, test := range clienteConfig {

		fmt.Println(test.QuantRamaisOpen[0].Ramal)
		fmt.Println("Passou aq")

	}

	//resourceID                           // Substitua pelo ID do recurso que você deseja atualizar
	updateData := map[string]interface{}{

		// Defina os campos que você deseja atualizar
	}

	fmt.Println(jsons)
	fmt.Println(serverURL)

	updateJSON, err := json.Marshal(updateData)
	if err != nil {
		fmt.Println("Erro ao codificar os dados de atualização em JSON:", err)
		return
	}

	// Crie uma solicitação HTTP PATCH para atualizar o recurso
	url := fmt.Sprintf("%s/clientes/%d", serverURL, resourceID)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(updateJSON))
	if err != nil {
		fmt.Println("Erro ao criar a solicitação HTTP PATCH:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Faça a solicitação HTTP PATCH
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro na solicitação HTTP PATCH:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Recurso atualizado com sucesso.")
	} else {
		fmt.Println("Erro ao atualizar o recurso. Status code:", resp.StatusCode)
	}

}
