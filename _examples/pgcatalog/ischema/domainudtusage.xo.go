package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// DomainUdtUsage represents a row from 'information_schema.domain_udt_usage'.
type DomainUdtUsage struct {
	UdtCatalog    pgtypes.SQLIdentifier `json:"udt_catalog"`    // udt_catalog
	UdtSchema     pgtypes.SQLIdentifier `json:"udt_schema"`     // udt_schema
	UdtName       pgtypes.SQLIdentifier `json:"udt_name"`       // udt_name
	DomainCatalog pgtypes.SQLIdentifier `json:"domain_catalog"` // domain_catalog
	DomainSchema  pgtypes.SQLIdentifier `json:"domain_schema"`  // domain_schema
	DomainName    pgtypes.SQLIdentifier `json:"domain_name"`    // domain_name
}