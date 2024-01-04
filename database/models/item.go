package models
type Item struct{
	Name string `json:"name" validate:"required,min=3,max=40"`
	 UnitPrice int `json:"unitprice" validate:"required"`
	Category string `json:"category" validate:"required"`

	
}