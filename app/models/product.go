package models

type Product struct {
	ID   int    `db:"id,INT,AUTO_INCREMENT,PRIMARY KEY" json:"id"`
	Name string `db:"name,VARCHAR(255),NOT NULL" json:"name"`
}
