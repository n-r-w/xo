package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// ColumnDomainUsage represents a row from 'information_schema.column_domain_usage'.
type ColumnDomainUsage struct {
	DomainCatalog pgtypes.SQLIdentifier `json:"domain_catalog"` // domain_catalog
	DomainSchema  pgtypes.SQLIdentifier `json:"domain_schema"`  // domain_schema
	DomainName    pgtypes.SQLIdentifier `json:"domain_name"`    // domain_name
	TableCatalog  pgtypes.SQLIdentifier `json:"table_catalog"`  // table_catalog
	TableSchema   pgtypes.SQLIdentifier `json:"table_schema"`   // table_schema
	TableName     pgtypes.SQLIdentifier `json:"table_name"`     // table_name
	ColumnName    pgtypes.SQLIdentifier `json:"column_name"`    // column_name
}