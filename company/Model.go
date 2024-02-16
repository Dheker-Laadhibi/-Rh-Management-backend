package company

import   "Test/users"

type Company struct {
	Idcompany       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Namecompany     string `json:"name"`
	Locationcompany string `json:"locationcompany"`
	Industry        string `json:"industry"`
	Empnumber       uint   `json:"empnumber"`
	Empfax          uint   `json:"empfax"`
	Empemail        string `json:"empemail"`
	Founder         string `json:"founder"`
	Datefounded     string `json:"datefounded"`
	Website         string `json:"website"`
	Description     string `json:"description"`
	// les employees
	Users []users.User `json:"users"`
}


