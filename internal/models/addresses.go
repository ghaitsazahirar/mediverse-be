package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Addresses struct {
	db.ModelBase
	Id         uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Street     *string            `json:"street,omitempty" column:"name:street;type:text;nullable"`
	District   *string            `json:"district,omitempty" column:"name:district;type:text;nullable"`
	City       *string            `json:"city,omitempty" column:"name:city;type:text;nullable"`
	Province   *string            `json:"province,omitempty" column:"name:province;type:text;nullable"`
	PostalCode *string            `json:"postal_code,omitempty" column:"name:postal_code;type:text;nullable"`
	CreatedAt  *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"addresses" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	FacilityAddresses []*Facilities `json:"facility_addresses,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:address_id"`
}
