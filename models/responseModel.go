package models

type ResponseModel struct {
	Number  int    `json:"number"`
	HtmlUrl string `json:"html_url"`
	Nombre  string `json:"title"`
	User    struct {
		Login string `json:"login"`
	}
	Labels []struct {
		LNombre string `json:"name"`
	} `json:"labels"`
	Milestone struct {
		Nombre      string `json:"title"`
		Descripcion string `json:"description"`
	} `json:"milestone"`
}
