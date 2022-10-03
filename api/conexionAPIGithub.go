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

// Esta funcion es la que conecta con la api de github para extraer las issues necesarias
func ConectarAPIGithub() []models.ResponseModel {
	var responseObject []models.ResponseModel
	var resul []models.ResponseModel
	var resultados []models.ResponseModel
	/*
		Este ciclo nos permite revisar varias de las paginas de la api de github https://api.github.com/repos/golang/go/issues
		para extraer la informacion necesario, ya que que la api sin una llave de autentificacion solo permite 60 llamadas por dia
		se realizara la busqueda a 30 de las paginas para ver la informacion, si lo desea puede aumentar o disminuir la cantidad de paginas
		en el ciclo, puede aumentar el numero de paginas a no mas de 60
	*/
	for i := 1; i <= 30; i++ {
		url := "https://api.github.com/repos/golang/go/issues?page=" + strconv.Itoa(i)
		response, err := http.Get(url)
		if err != nil {
			fmt.Print(err.Error())
		}

		/*
			Esta linea es para leer lo que regresa la llamda a la api de manera correcta, regresando una lista de enteros
			en esta variable solo se estan agregando los valores que hay en la pagina actual no guarda la informacion de las
			paginas anteriores
		*/
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		/*
			Esta linea es para agrega todos los archivos del responseData al responseObject que es un modelo de ResponseModel
			para asi poder tener la informacion bien organizada y asi poder accesar a esa informacion de una manera mas facil
			en esta variable tambien se estan agregando los valores que hay en la pagina actual no guarda la informacion de las
			paginas anteriores
		*/
		jsonErr := json.Unmarshal(responseData, &responseObject)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		/*
			En esta variable result se estan guardando todos los registros de cada pagina que se busca para despues de todos esos valores
			solo retornar los que tengan la etiqueta Go2
		*/
		resul = append(resul, responseObject...)
	}

	/*
		Esta ciclo se encarga que despues de haber visitado todas y cada una de las paginas en el ciclo anterior se busque en los datos
		guardados en la variable resul se busque en el apartado de etiquetas la de Go2 y la guarde en nuestra variable resultados ya que
		los resultados que regrese seran los que se guardaran en la base de datos
	*/
	for i := 0; i < len(resul); i++ {
		for _, j := range resul[i].Labels {
			if j.LNombre == "Go2" {
				resultados = append(resultados, resul[i])
			}
		}
	}

	return resultados
}
