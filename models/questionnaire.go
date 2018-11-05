package models

import (
	"fmt"

	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/jinzhu/gorm"
)

//Questionnaire ...
type Questionnaire struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserID      uint       `json:"user_id"`
	Questions   []Question `json:"questions" gorm:"many2many:questionnaire_questions"`
}

//Validate ...
func (questionnaire *Questionnaire) Validate() (map[string]interface{}, bool) {

	if questionnaire.Title == "" {
		return u.Message(false, "You must define title"), false
	}

	if questionnaire.Description == "" {
		return u.Message(false, "You must define description"), false
	}

	if questionnaire.UserID == 0 {
		return u.Message(false, "User is not defined"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

//Create ..
func (questionnaire *Questionnaire) Create() map[string]interface{} {

	if resp, ok := questionnaire.Validate(); !ok {
		return resp
	}

	GetDB().Create(questionnaire)

	resp := u.Message(true, "success")
	resp["questionnaire"] = questionnaire
	return resp
}

//Update ...
func (questionnaire *Questionnaire) Update() map[string]interface{} {

	if resp, ok := questionnaire.Validate(); !ok {
		return resp
	}

	GetDB().Save(questionnaire)

	resp := u.Message(true, "success")
	resp["questionnaire"] = questionnaire
	return resp
}

//GetQuestionnaire ...
func GetQuestionnaire(id uint) *Questionnaire {

	questionnaire := &Questionnaire{}
	err := GetDB().Table("questionnaires").Where("id = ?", id).First(questionnaire).Error

	if err != nil {
		return nil
	}

	err = GetDB().Table("questions").Where("questionnaire_id = ?", questionnaire.ID).Find(&questionnaire.Questions).Error

	if err != nil {
		return nil
	}

	return questionnaire
}

//GetQuestionnairesForUser ...
func GetQuestionnairesForUser(user uint) []*Questionnaire {

	questionnaires := make([]*Questionnaire, 0)
	err := GetDB().Table("questionnnaire").Where("user_id = ?", user).Find(&questionnaires).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return questionnaires
}

//DeleteQuestionnaire ...
func DeleteQuestionnaire(id uint) *Questionnaire {

	questionnaire := &Questionnaire{}
	err := GetDB().Table("questionnaire").Where("id = ?", id).First(questionnaire).Error
	if err != nil {
		return nil
	}

	GetDB().Delete(questionnaire)

	return questionnaire
}
