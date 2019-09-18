package model

type WorkbenchCreationRequest struct {
	Name string `validation:"required"`
}
