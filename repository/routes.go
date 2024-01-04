package repository
import "github.com/gofiber/fiber/v2"

func(repo *Repository) SetupRoutes(app *fiber.App){
api :=app.Group("/api")

//routes for items
api.Get("/items",repo.GetItems)
api.Get("/items/:id",repo.GetItemByID)
api.Post("/items",repo.CreateItem)
api.Patch("/items/:id", repo.UpdateItem)
api.Delete("/items/:id", repo.DeleteItem)

//routes for invoices
api.Get("/invoices",repo.GetInvoices)
api.Get("/invoices/:id",repo.GetInvoiceByID)
api.Post("/invoices",repo.CreateInvoice)
api.Patch("/invoices/:id", repo.UpdateInvoice)
api.Delete("/invoices/:id", repo.DeleteInvoice)

}


