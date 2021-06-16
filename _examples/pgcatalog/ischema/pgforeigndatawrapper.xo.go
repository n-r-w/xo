package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"

	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// PgForeignDataWrapper represents a row from 'information_schema._pg_foreign_data_wrappers'.
type PgForeignDataWrapper struct {
	Oid                        pgtypes.Oid           `json:"oid"`                           // oid
	Fdwowner                   pgtypes.Oid           `json:"fdwowner"`                      // fdwowner
	Fdwoptions                 []sql.NullString      `json:"fdwoptions"`                    // fdwoptions
	ForeignDataWrapperCatalog  pgtypes.SQLIdentifier `json:"foreign_data_wrapper_catalog"`  // foreign_data_wrapper_catalog
	ForeignDataWrapperName     pgtypes.SQLIdentifier `json:"foreign_data_wrapper_name"`     // foreign_data_wrapper_name
	AuthorizationIdentifier    pgtypes.SQLIdentifier `json:"authorization_identifier"`      // authorization_identifier
	ForeignDataWrapperLanguage pgtypes.CharacterData `json:"foreign_data_wrapper_language"` // foreign_data_wrapper_language
}