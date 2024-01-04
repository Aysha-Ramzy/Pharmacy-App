package models

type Invoice_Request struct {
	
	Invoice Invoice `json:"invoice"`
	ItemIDs     []uint `json:"itemIDs"`
}