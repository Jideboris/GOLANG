package handler

import (
	"umotif.com/rescheduler/app/usecase/questionnaires_usecase"
)

type computerHandler struct {
	questionnaires_uc questionnaires_usecase.UseCase
}

func NewHandler(questusecase questionnaires_usecase.UseCase) ComputeHandler {
	return &computerHandler{
		questionnaires_uc: questusecase,
	}
}

// GenerateTable implements ComputeHandler.
func (rest *computerHandler) GenerateTable() (bool, error) {
	_, err := rest.questionnaires_uc.CreateTable()
	if err != nil {
		return false, err
	}
	return true, nil
}

// GenerateTableData implements ComputeHandler.
func (rest *computerHandler) GenerateTableData() (bool, error) {
	_, err := rest.questionnaires_uc.CreateData()
	if err != nil {
		return false, err
	}
	return true, nil
}