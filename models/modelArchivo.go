package models

//Esta es la estructura de la informacion que se tiene para poder ver y agregar archivos en la base de datos
type ModeloArchivo struct {
	IssueNum int    `bson:"issueNum" json:"issueNum,omitempty"`
	IssueUrl string `bson:"issueUrl" json:"issueUrl,omitempty"`
	IssueNom string `bson:"issueNom" json:"issueNom,omitempty"`
	Autor    string `bson:"autor" json:"autor,omitempty"`
	Labels   []struct {
		LNombre string `json:"name"`
	} `bson:"tags" json:"tags,omitempty"`
	Milstone struct {
		Nombre      string `bson:"nombre" json:"nombre,omitempty"`
		Descripcion string `bson:"descripcion" json:"descripcion,omitempty"`
	} `bson:"milestone" json:"milestone,omitempty"`
}
