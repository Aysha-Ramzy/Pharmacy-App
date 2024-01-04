package migrations

import (
	"gorm.io/gorm"
)
type Invoices_Items struct{
	ID_Invoice uint `gorm:"column:id_invoice;primaryKey" json:"id_invoice"`
	ID_Item    uint `gorm:"column:id_item;primaryKey" json:"id_item"`
	

}
func MigrateInvoices_Items(db *gorm.DB) error {
    err := db.Table("invoice_items").AutoMigrate(&Invoices_Items{})
    return err
}

