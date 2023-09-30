package models

// Product represents a data model for products with essential attributes.
type Product struct {
	ID   int    `db:"id,INT,AUTO_INCREMENT,PRIMARY KEY" json:"id"` // ID uniquely identifies the product.
	Name string `db:"name,VARCHAR(255),NOT NULL" json:"name"`      // Name is the name of the product.
}
