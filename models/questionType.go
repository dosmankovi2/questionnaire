package models

import (
	u "github.com/dosmankovi2/questionnaire/utils"
	"github.com/jinzhu/gorm"
)

//QuestionType ...
type QuestionType struct {
	gorm.Model
	Expression string `json:"expression"`
}

//Validate ...
func (questionType *QuestionType) Validate() (map[string]interface{}, bool) {

	if questionType.Expression == "" {
		return u.Message(false, "Expression must be defined"), false
	}

	return u.Message(true, "success"), true
}

//Create ...
func (questionType *QuestionType) Create() map[string]interface{} {

	if resp, ok := questionType.Validate(); !ok {
		return resp
	}

	GetDB().Create(questionType)

	resp := u.Message(true, "success")
	resp["questionType"] = questionType
	return resp
}

//Update ...
func (questionType *QuestionType) Update() map[string]interface{} {

	if resp, ok := questionType.Validate(); !ok {
		return resp
	}

	GetDB().Save(questionType)

	resp := u.Message(true, "success")
	resp["questionType"] = questionType
	return resp
}

//GetQuestionType ...
func GetQuestionType(id uint) *QuestionType {

	questionType := &QuestionType{}
	err := GetDB().Table("question_types").Where("id = ?", id).First(questionType).Error
	if err != nil {
		return nil
	}
	return questionType
}

//DeleteQuestionType ...
func DeleteQuestionType(id uint) *QuestionType {

	questionType := &QuestionType{}
	err := GetDB().Table("question_types").Where("id = ?", id).First(questionType).Error
	if err != nil {
		return nil
	}

	GetDB().Delete(questionType)

	return questionType
}
