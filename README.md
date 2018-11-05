# questionnaire
Questionnaire API

Technology:
  - GO
    - GORM
  - JWT
  - PostgreSQL
  - Docker - (Introduced Dockerfile and docker-compose)
  

Database:
  - Questionnaire
  - Question
  - QuestionType
  - QuestionChoice
  - Answer
  - AnswerParticle

Tried to mimic user API for the questionnaire. The solution is not completed and it is a part of the interview task using
completely new technology. There are a few major features not implemented, and this only represents the basis of architecture.

Whole API set works with domain models

	/api/user/new - POST
	/api/user/login - POST

  	/api/questionnaire/{id} - GET
	/api/questionnaire - POST
	/api/questionnaire - PUT 
	/api/questionnaire/{id} - DELETE

  	/api/question/{id} - GET
	/api/question - POST
	/api/question - PUT 
	/api/question/{id} - DELETE

	/api/type/{id} - GET
	/api/type - POST
	/api/type - PUT 
	/api/type/{id} - DELETE

  	/api/choice/{id} - GET
	/api/choice - POST
	/api/choice - PUT 
	/api/choicie/{id} - DELETE

	/api/answer/{id} - GET
	/api/answer - POST
	/api/answer - PUT 
	/api/answer/{id} - DELETE
