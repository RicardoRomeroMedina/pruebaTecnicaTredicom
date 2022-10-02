package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"strconv"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/models"
)

func ConectarAPIGithub() []models.ResponseModel {
	var responseObject []models.ResponseModel
	var resul []models.ResponseModel
	var resultados []models.ResponseModel
	for i := 1; i <= 30; i++ {
		url := "https://api.github.com/repos/golang/go/issues?page=" + strconv.Itoa(i)
		response, err := http.Get(url)
		if err != nil {
			fmt.Print(err.Error())
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		// var resp = string(responseData)
		jsonErr := json.Unmarshal(responseData, &responseObject)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		resul = append(resul, responseObject...)
	}
	for i := 0; i < len(resul); i++ {
		for _, j := range resul[i].Labels {
			if j.LNombre == "Go2" {
				resultados = append(resultados, resul[i])
			}
		}
	}

	return resultados
}
