package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Esta funcion servira para eliminar todos los archivos de la base de datos para asi poder probar la funcion de agregar archivos
func EliminarArchivosDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	//Esta linea crea la conexion con la base de datos en mongodb en la que se tienen las colecciones en donde se guardara la informacion
	db := MongoCN.Database("pruebatecnica")

	//Esta linea crea la conexion a la coleccion en donde se guarda la informacion
	col := db.Collection("DataRepository")

	condicion := bson.M{}

	_, err := col.DeleteMany(ctx, condicion)

	return err
}
