package db

import (
	"context"
	"time"

	// "github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Esta funcion es la que se llama para agregar todos los archivos en la base de datos
func NuevoArchivo(a []interface{}) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//Esta linea crea la conexion con la base de datos en mongodb en la que se tienen las colecciones en donde se guardara la informacion
	db := MongoCN.Database("pruebatecnica")

	//Esta linea crea la conexion a la coleccion en donde se guarda la informacion
	col := db.Collection("DataRepository")

	//Esta linea se encargara de insertar la informacion en la base de datos
	result, err := col.InsertMany(ctx, a)

	//Esta linea verifica si hay algun error al guardar la informacion se es asi se imprimira en error
	if err != nil {
		return "", false, err
	}

	//Esta linea se encargara de regresar la informacion necesaria si no hay ningun error
	ObjID, _ := result.InsertedIDs[len(result.InsertedIDs)-1].(primitive.ObjectID)
	return ObjID.String(), true, nil
}
