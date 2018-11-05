package models

import (
	"fmt"

	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/jinzhu/gorm"
)

//Question ...
type Question struct {
	gorm.Model
	Expression     string           `json:"expression"`
	QuestionTypeID int              `jgorm:"ForeignKey:id" json:"question_type_id"`
	QuestionType   QuestionType     `json:"question_type"`
	Choices        []QuestionChoice `json:"question_choices" gorm:"many2many:choices_elements"`
}

//Validate ...
func (question *Question) Validate() (map[string]interface{}, bool) {

	if question.QuestionTypeID == 0 {
		return u.Message(false, "Must specify type"), false
	}

	if question.Expression == "" {
		return u.Message(false, "Must specify expression"), false
	}

	return u.Message(true, "success"), true
}

//Create ...
func (question *Question) Create() map[string]interface{} {

	if resp, ok := question.Validate(); !ok {
		return resp
	}

	GetDB().Create(question)

	resp := u.Message(true, "success")
	resp["question"] = question
	return resp
}

//Update ...
func (question *Question) Update() map[string]interface{} {

	if resp, ok := question.Validate(); !ok {
		return resp
	}

	GetDB().Save(question)

	resp := u.Message(true, "success")
	resp["question"] = question
	return resp
}

//GetQuestion ...
func GetQuestion(id uint) *Question {

	question := &Question{}
	err := GetDB().Table("questions").Where("id = ?", id).First(question).Error
	if err != nil {
		return nil
	}
	return question
}

//GetQuestions ...
func GetQuestions(questionnarie uint) []*Question {

	questions := make([]*Question, 0)
	err := GetDB().Table("contacts").Where("questionnaire_id = ?", questionnarie).Find(&questions).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return questions
}

//DeleteQuestion ...
func DeleteQuestion(id uint) *Question {

	question := &Question{}
	err := GetDB().Table("questions").Where("id = ?", id).First(question).Error
	if err != nil {
		return nil
	}

	GetDB().Delete(question)

	return question
}
