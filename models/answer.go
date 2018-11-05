package models

import (
	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/jinzhu/gorm"
)

//Answer ...
type Answer struct {
	gorm.Model
	QuestionnaireID int              `jgorm:"ForeignKey:id" json:"questionnaire_id"`
	Questionnaire   Questionnaire    `json:"questionnaire"`
	Particles       []AnswerParticle `json:"answer_particles" gorm:"many2many:answer_elements"`
}

//Validate ...
func (answer *Answer) Validate() (map[string]interface{}, bool) {

	if answer.QuestionnaireID == 0 {
		return u.Message(false, "You have to provide questionnaire"), false
	}

	return u.Message(true, "success"), true
}

//Create ...
func (answer *Answer) Create() map[string]interface{} {

	if resp, ok := answer.Validate(); !ok {
		return resp
	}

	GetDB().Create(answer)

	resp := u.Message(true, "success")
	resp["question"] = answer
	return resp
}

//Update ...
func (answer *Answer) Update() map[string]interface{} {

	if resp, ok := answer.Validate(); !ok {
		return resp
	}

	GetDB().Save(answer)

	resp := u.Message(true, "success")
	resp["answer"] = answer
	return resp
}

//GetAnswer ...
func GetAnswer(id uint) *Answer {

	answer := &Answer{}
	err := GetDB().Table("answers").Where("id = ?", id).First(answer).Error
	if err != nil {
		return nil
	}
	return answer
}

//DeleteAnswer ...
func DeleteAnswer(id uint) *Answer {

	answer := &Answer{}
	err := GetDB().Table("answers").Where("id = ?", id).First(answer).Error
	if err != nil {
		return nil
	}

	GetDB().Delete(answer)

	return answer
}
