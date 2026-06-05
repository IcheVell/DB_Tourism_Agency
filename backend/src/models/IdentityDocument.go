package models

import "time"

type IdentityDocument struct {
	ID             int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	DocumentType   string     `gorm:"column:document_type;not null" json:"document_type"`
	DocumentSeries string     `gorm:"column:document_series;not null" json:"document_series"`
	DocumentNumber string     `gorm:"column:document_number;not null" json:"document_number"`
	ExpirationDate *time.Time `gorm:"column:expiration_date;type:date" json:"expiration_date,omitempty"`
	IssuedBy       string     `gorm:"column:issued_by;not null" json:"issued_by"`
	IssueDate      time.Time  `gorm:"column:issue_date;not null;type:date" json:"issue_date"`
	Citizenship    string     `gorm:"column:citizenship;not null" json:"citizenship"`
	TouristID      int64      `gorm:"column:tourist_id;not null" json:"tourist_id"`

	Tourist *Tourist `gorm:"foreignKey:TouristID;references:ID" json:"tourist,omitempty"`
}

func (IdentityDocument) TableName() string {
	return "identity_documents"
}
