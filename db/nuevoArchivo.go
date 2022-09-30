package db

import (
	"context"
	"time"

	"github.com/RicardoRomeroMedina/pruebaTecnicaTredicom/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NuevoArchivo(a models.ModeloArchivo) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("pruebatecnica")
	col := db.Collection("DataRepository")

	result, err := col.InsertOne(ctx, a)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
