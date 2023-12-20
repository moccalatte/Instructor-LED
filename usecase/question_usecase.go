package usecase

import (
	"final-project-kelompok-1/model"
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/repository"
	"fmt"
)

type QuestionUseCase interface {
	AddQuestion(payload dto.QuestionRequestDto) (model.Question, error)
	FindQuestionById(id string) (model.Question, error)
	FindQuestionByStudentId(id string) (model.Question, error)
	GetAllQuestion() ([]model.Question, error)
	Update(payload dto.QuestionRequestDto, id string) (model.Question, error)
	Delete(id string) (model.Question, error)
	Answer(payload dto.QuestionRequestDto, id string) (model.Question, error)
	GetImagePath(questionID string) (string, error)
}

type questionUseCase struct {
	repo repository.QuestionRepository
}

func (s *questionUseCase) AddQuestion(payload dto.QuestionRequestDto) (model.Question, error) {
	newSession := model.Question{
		SessionID:   payload.SessionID,
		StudentID:   payload.StudentID,
		TrainerID:   payload.TrainerID,
		Title:       payload.Title,
		Description: payload.Description,
		CourseID:    payload.CourseID,
		Image:       payload.Image,
		Answer:      payload.Answer,
		Status:      payload.Status,
		// ImagePath: payload.ImagePath,
	}

	createsQuestion, err := s.repo.Create(newSession)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to save data : %s", err.Error())
	}
	return createsQuestion, nil
}

func (s *questionUseCase) FindQuestionById(id string) (model.Question, error) {
	Question, err := s.repo.GetById(id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to get data question by id : %s", err.Error())
	}
	return Question, nil
}

func (s *questionUseCase) GetAllQuestion() ([]model.Question, error) {
	var sliceQuest []model.Question
	Question, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get data all data : %s", err.Error())
	}
	return append(sliceQuest, Question...), nil
}

func (s *questionUseCase) FindQuestionByStudentId(id string) (model.Question, error) {
	Question, err := s.repo.GetByStudentId(id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to get data question by id : %s", err.Error())
	}
	return Question, nil
}

// func (s *questionUseCase) FindQuestionByStudentId(id string) (model.Question, error) {
// 	Question, err := s.repo.GetByStudentId(id)

// 	if err != nil {
// 		return model.Question{}, fmt.Errorf("failed to get data by id in question use case student id : %s", err.Error())
// 	}
// 	return Question, nil
// }

// func (s *questionUseCase) GetAllQuestion() ([]model.Question, error) {
// 	questionAll, err := s.repo.GetAll()

// 	if err != nil {
// 		fmt.Println("Error Get All Data in use case : ", err.Error())
// 		return questionAll, fmt.Errorf("failed to find data : %v", err.Error())
// 	}

// 	return questionAll, nil
// }

func (s *questionUseCase) Update(payload dto.QuestionRequestDto, id string) (model.Question, error) {
	question := model.Question{
		SessionID:   payload.SessionID,
		StudentID:   payload.StudentID,
		TrainerID:   payload.TrainerID,
		Title:       payload.Title,
		Description: payload.Description,
		CourseID:    payload.CourseID,
		Image:       payload.Image,
		Answer:      payload.Answer,
		Status:      payload.Status,
	}

	question, err := s.repo.Update(question, id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to Update Question : %s", err.Error())
	}

	return question, nil
}

func (s *questionUseCase) Delete(id string) (model.Question, error) {
	question, err := s.repo.Delete(id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to delete data : %s", err.Error())
	}

	return question, nil
}

func (s *questionUseCase) Answer(payload dto.QuestionRequestDto, id string) (model.Question, error) {
	answer := model.Question{
		Answer: payload.Answer,
	}
	answered, err := s.repo.Answer(answer, id)

	if err != nil {
		return model.Question{}, fmt.Errorf("failed to answer : %s", err.Error())
	}

	return answered, nil

}

func (q *questionUseCase) GetImagePath(questionID string) (string, error) {
	// Logika untuk mendapatkan path gambar dari repository
	question, err := q.repo.GetById(questionID)
	if err != nil {
		return "", fmt.Errorf("failed to get question: %s", err.Error())
	}

	// Menggunakan path gambar dari model.Question
	imagePath := question.Image

	return imagePath, nil
}

func NewQuestion(repo repository.QuestionRepository) QuestionUseCase {
	return &questionUseCase{repo: repo}
}
