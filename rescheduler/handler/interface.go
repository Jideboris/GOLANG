package handler

type ComputeHandler interface {
	GenerateTableData() (bool, error)
	GenerateTable() (bool, error)
}
