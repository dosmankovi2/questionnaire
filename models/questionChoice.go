package models

import (
	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/jinzhu/gorm"
)

//QuestionChoice ...
type QuestionChoice struct {
	gorm.Model
	Expression string `json:"expression"`
}

//Validate ...
func (questionChoice *QuestionChoice) Validate() (map[string]interface{}, bool) {

	if questionChoice.Expression == "" {
		return u.Message(false, "You must define expression"), false
	}

	return u.Message(true, "success"), true
}

//Create ...
func (questionChoice *QuestionChoice) Create() map[string]interface{} {

	if resp, ok := questionChoice.Validate(); !ok {
		return resp
	}

	GetDB().Create(questionChoice)

	resp := u.Message(true, "success")
	resp["question_choice"] = questionChoice
	return resp
}

//Update ...
func (questionChoice *QuestionChoice) Update() map[string]interface{} {

	if resp, ok := questionChoice.Validate(); !ok {
		return resp
	}

	GetDB().Save(questionChoice)

	resp := u.Message(true, "success")
	resp["question_choice"] = questionChoice
	return resp
}

//GetQuestionChoice...
func GetQuestionChoice(id uint) *QuestionChoice {

	questionChoice := &QuestionChoice{}
	err := GetDB().Table("question_choices").Where("id = ?", id).First(questionChoice).Error
	if err != nil {
		return nil
	}
	return questionChoice
}

//DeleteQuestionChoice ...
func DeleteQuestionChoice(id uint) *QuestionChoice {

	questionChoice := &QuestionChoice{}
	err := GetDB().Table("question_choices").Where("id = ?", id).First(questionChoice).Error
	if err != nil {
		return nil
	}

	GetDB().Delete(questionChoice)

	return questionChoice
}
