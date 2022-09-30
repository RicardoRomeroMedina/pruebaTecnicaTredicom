package models

type ModeloArchivo struct {
	IssueNum int      `bson:"issueNum" json:"issueNum,omitempty"`
	IssueUrl string   `bson:"issueUrl" json:"issueUrl,omitempty"`
	IssueNom string   `bson:"issueNom" json:"issueNom,omitempty"`
	Autor    string   `bson:"autor" json:"autor,omitempty"`
	Tags     []string `bson:"tags" json:"tags,omitempty"`
	Milstone struct {
		Nombre      string `bson:"nombre" json:"nombre,omitempty"`
		Descripcion string `bson:"descripcion" json:"descripcion,omitempty"`
	} `bson:"milestone" json:"milestone,omitempty"`
}
