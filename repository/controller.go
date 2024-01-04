package repository

import (
	"net/http"

	"github.com/Aysha-Ramzy/Pharmacy-App/database/migrations"
	"github.com/Aysha-Ramzy/Pharmacy-App/database/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gopkg.in/go-playground/validator.v9"
)
type ErrorResponse struct{
	FailedField string
	Tag string
	Value string

}
var validate= validator.New()
//For Items
//validate items
func ValidateStruct(item models.Item) []*ErrorResponse{
var errors []*ErrorResponse
err := validate.Struct(item)
if err != nil {
	for _, err := range err.(validator.ValidationErrors){
		var element ErrorResponse
		element.FailedField= err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = err.Param()
		errors = append(errors,&element )
	}
}
return errors
}

//getitems
func (r *Repository) GetItems(context *fiber.Ctx) error {
db := r.DB
model := db.Model(&migrations.Items{})
pg:= paginate.New(&paginate.Config{
	DefaultSize: 20,
	CustomParamEnabled: true,

})
page := pg.With(model).Request(context.Request()).Response(&[]migrations.Items{})
context.Status(http.StatusOK).JSON(&fiber.Map{
	"data":page,
})
return nil
}

//create items
func (r *Repository) CreateItem(context *fiber.Ctx) error {
	item := models.Item{}
	err := context.BodyParser(&item)

	if err !=nil{
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message":"Request Failed"})
		return err
	}
	errors := ValidateStruct(item)
	if errors != nil{
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}
	if err := r.DB.Create(&item).Error; err !=nil{
		return context.Status(http.StatusBadRequest).JSON(fiber.Map{"status":"error","message":"Couldn't create user","data": err})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Item has been added successfully!","data":item})
 return nil
 }
 
//Edit item details
 func (r *Repository) UpdateItem(context *fiber.Ctx) error {
	item := models.Item{}
	err := context.BodyParser(&item)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}
	errors := ValidateStruct(item)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := r.DB
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}
	if db.Model(&item).Where("id = ?", id).Updates(&item).RowsAffected == 0 {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get item with given id"})
	}

	return context.JSON(fiber.Map{"status": "success", "message": "Item successfully updated"})
}


//delete an item
func (r *Repository) DeleteItem(context *fiber.Ctx) error {
	itemModel := migrations.Items{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Delete(itemModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Item deleted successfully"})
	return nil
}

//get an item
func (r *Repository) GetItemByID(context *fiber.Ctx) error {
	id := context.Params("id")
	itemModel := &migrations.Items{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(itemModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get the item"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Item id fetched successfully", "data": itemModel})
	return nil
}


//For Invoices

//validate invoice
func ValidateStructInvoice(invoice models.Invoice) []*ErrorResponse{
var errors []*ErrorResponse
err := validate.Struct(invoice)
if err != nil {
	for _, err := range err.(validator.ValidationErrors){
		var element ErrorResponse
		element.FailedField= err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = err.Param()
		errors = append(errors,&element )
	}
}

return errors
}

//getinvoices
// GetInvoices retrieves a list of invoices from the database

func (r *Repository) GetInvoices(context *fiber.Ctx) error {
    db := r.DB
    var invoices []migrations.Invoices

    // Fetch invoices from the database
    if err := db.Find(&invoices).Error; err != nil {
        context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "Failed to fetch invoices"})
        return err
    }

    context.Status(http.StatusOK).JSON(&fiber.Map{
        "data": map[string]interface{}{
            "invoices":     invoices,
            
        },
    })
    return nil
}


 
//Edit invoice details
 func (r *Repository) UpdateInvoice(context *fiber.Ctx) error {
	invoice := models.Invoice{}
	err := context.BodyParser(&invoice)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})

		return err
	}
	errors := ValidateStructInvoice(invoice)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := r.DB
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}
	if db.Model(&invoice).Where("id = ?", id).Updates(&invoice).RowsAffected == 0 {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get invoice with given id"})
	}

	return context.JSON(fiber.Map{"status": "success", "message": "Invoice successfully updated"})
}


//delete an invoice
func (r *Repository) DeleteInvoice(context *fiber.Ctx) error {
	invoiceModel := migrations.Invoices{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Delete(invoiceModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Invoice deleted successfully"})
	return nil
}

//get an invoice
func (r *Repository) GetInvoiceByID(context *fiber.Ctx) error {
	id := context.Params("id")
	invoiceModel := &migrations.Invoices{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "ID cannot be empty"})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(invoiceModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "Could not get the invoice"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Invoice id fetched successfully", "data": invoiceModel})
	return nil
}


//create invoices
func (r *Repository) CreateInvoice(context *fiber.Ctx) error {
    invoiceRequest := models.Invoice_Request{}
    err := context.BodyParser(&invoiceRequest)
    if err != nil {
        return context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "Request Failed"})
    }

    // Validate the Invoice structure
    errors := ValidateStructInvoice(invoiceRequest.Invoice)
    if errors != nil {
        return context.Status(fiber.StatusBadRequest).JSON(errors)
    }

    // Create the invoice
    if err := r.DB.Create(&invoiceRequest.Invoice).Error; err != nil {
        return context.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't create invoice", "data": err})
    }

    // Create entries in the Invoice_Item table for the invoice and associated items
    for _, itemID := range invoiceRequest.ItemIDs {
        invoiceItem := models.Invoice_Item{
            ID_Invoice: invoiceRequest.Invoice.ID,
            ID_Item:    itemID,
        }

        if err := r.DB.Create(&invoiceItem).Error; err != nil {
            return context.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Couldn't add items to invoice", "data": err})
        }
    }

    context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Invoice has been added successfully!", "data": invoiceRequest.Invoice})
    return nil
}