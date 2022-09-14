package models

import (
	"fmt"
	"regexp"
	"strings"
)

type Field struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Operation string `json:"operation"`
	Order     bool   `json:"order"` // true for ASC false for DESC
}

func (fld *Field) FilConcat() string {
	if fld.Operation == "LIKE" {
		fld.Value = "%" + fld.Value + "%"
	}
	return fld.Name + " " + fld.Operation + " " + fld.Value
}

func (fld *Field) OrdConcat() string {
	if fld.Order {
		return fmt.Sprintf("%s ASC", fld.Name)
	}
	return fmt.Sprintf("%s DESC", fld.Name)
}

func (fld *Field) IsValid() bool {
	if fld.Name == "" {
		return false
	}
	switch fld.Operation {
	case "LIKE":
	case ">":
	case "=>":
	case "<":
	case "<=":
	case "=":
	default:
		return false
	}
	return regexp.MustCompile("^[a-zA-Zа-яА-Я0-9-/]+$").MatchString(fld.Value)
}

type PageRequest struct {
	PageNumber   int `json:"pg_number"`
	PageLength   int `json:"pg_length"`
	Fields       *[]Field
	TotalRecords int
}

func (pgReq PageRequest) New() *PageRequest {
	Fields := make([]Field, 0)
	pgReq.Fields = &Fields
	return &pgReq
}

func (pgReq *PageRequest) Filters() string {
	filterArray := make([]string, 0)
	for _, field := range *pgReq.Fields {
		if field.IsValid() {
			filterArray = append(filterArray, field.FilConcat())
		}
	}
	if len(filterArray) == 0 {
		return ""
	}
	return fmt.Sprintf(" WHERE %s ", strings.Join(filterArray, "and"))
}

func (pgReq *PageRequest) Order() string {
	orderArray := make([]string, 0)
	for _, field := range *pgReq.Fields {
		if field.IsValid() {
			orderArray = append(orderArray, field.OrdConcat())
		}
	}
	fmt.Println("orderArray length", len(orderArray))
	if len(orderArray) == 0 {
		return ""
	}
	return fmt.Sprintf(" ORDER BY %s ", strings.Join(orderArray, "and"))
}

func (pgReq *PageRequest) Offset() string {
	return fmt.Sprintf(" LIMIT %d OFFSET %d",
		pgReq.PageLength,
		(pgReq.PageNumber-1)*pgReq.PageLength)
}

func (pgReq *PageRequest) PageReq() string {
	return pgReq.Filters() + pgReq.Order() + pgReq.Offset()
}
