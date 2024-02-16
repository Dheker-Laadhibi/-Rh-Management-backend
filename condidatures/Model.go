package condidatures



type Condidature struct {
	Idcondidature int    `json:"idcondidature" gorm:"primaryKey;autoIncrement"`
	Formation     string `json:"formation"`
	Competences   string `json:"compétences"`
	Cv            string `json:"cv"`
	 Github         string `json:"github"`
}
