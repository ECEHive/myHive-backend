package model

type WorkbenchCreationRequest struct {
	Name string `validation:"required"`
}

type WorkbenchItemCreationRequest struct {
	Name string `validation:"required"`
}
