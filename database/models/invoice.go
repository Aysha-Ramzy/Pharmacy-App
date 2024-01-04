package models

type Invoice struct {
	ID          uint   ` json:"id"`
	Name        string `json:"name" validate:"required,min=3,max=40"`
	MobileNo    string `json:"mobileno" validate:"required"`
	Email       string `json:"email" validate:"required,email,min=6,max=32"`
	Address     string `json:"address" validate:"required"`
	BillingType string `json:"billingtype" validate:"required"`

}