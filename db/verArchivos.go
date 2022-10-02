package db

import (
	"context"
	"log"
	"time"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Esta funcion es la que se llamara para poder traer la informacion de la base de datos
func VerArchivos() ([]*models.ModeloArchivo, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//Esta linea crea la conexion con la base de datos en mongodb en la que se tienen las colecciones en donde se guardara la informacion
	db := MongoCN.Database("pruebatecnica")
	//Esta linea crea la conexion a la coleccion en donde se guarda la informacion
	col := db.Collection("DataRepository")

	//Esta linea es una condicion que se pide para filtrar la informacion a buscar
	//Por el momento esta esta vacia ya que no se filtrara ninguna informacion por el momento pero es necesaria en la funcion de buscar
	condicion := bson.M{}

	var results []*models.ModeloArchivo

	//Esta es la linea encargada de buscar la informacion en la base de datos
	cursor, err := col.Find(ctx, condicion)

	//Esta linea se encarga de verificar que no ocurra un error al buscar de lo contrario regresara el error
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	//Esta funcion se encarga de agregar cada uno de los datos en la coleccion para poder asi mostrarlos despues en la consola
	for cursor.Next(context.TODO()) {
		var registro models.ModeloArchivo

		err := cursor.Decode(&registro)

		if err != nil {
			return results, false
		}

		results = append(results, &registro)
	}

	return results, true
}
