package models

import (
	"fmt"

	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/jinzhu/gorm"
)

//AnswerParticle ...
type AnswerParticle struct {
	gorm.Model
	QuestionID int `jgorm:"ForeignKey:id" json:"questionnaire_id"`
	ChoiceID   int `jgorm:"ForeignKey:id" json:"choice_id"`
	Text       string
}

//Validate ...
func (answerParticle *AnswerParticle) Validate() (map[string]interface{}, bool) {

	if answerParticle.QuestionID == 0 {
		return u.Message(false, "You have to associate question"), false
	}

	if answerParticle.ChoiceID == 0 || answerParticle.Text == "" {
		return u.Message(false, "You have to associate choice or plain text"), false
	}

	return u.Message(true, "success"), true
}

//Create ...
func (answerParticle *AnswerParticle) Create() map[string]interface{} {

	if resp, ok := answerParticle.Validate(); !ok {
		return resp
	}

	GetDB().Create(answerParticle)

	resp := u.Message(true, "success")
	resp["answer_particle"] = answerParticle
	return resp
}

//Update ...
func (answerParticle *AnswerParticle) Update() map[string]interface{} {

	if resp, ok := answerParticle.Validate(); !ok {
		return resp
	}

	GetDB().Save(answerParticle)

	resp := u.Message(true, "success")
	resp["answer_particle"] = answerParticle
	return resp
}

//GetAnswerParticle ...
func GetAnswerParticle(id uint) *AnswerParticle {

	answerParticle := &AnswerParticle{}
	err := GetDB().Table("answer_particles").Where("id = ?", id).First(answerParticle).Error
	if err != nil {
		return nil
	}
	return answerParticle
}

//GetAnswerParticles ...
func GetAnswerParticles(question uint) []*AnswerParticle {

	answerParticles := make([]*AnswerParticle, 0)
	err := GetDB().Table("answer_particles").Where("question_id = ?", question).Find(&answerParticles).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return answerParticles
}

//DeleteAnswerParticle ...
func DeleteAnswerParticle(id uint) *AnswerParticle {

	answerParticle := &AnswerParticle{}
	err := GetDB().Table("answer_particles").Where("id = ?", id).First(answerParticle).Error
	if err != nil {
		return nil
	}

	GetDB().Delete(answerParticle)

	return answerParticle
}
