package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Facilities struct {
	db.ModelBase
	Id        uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name      *string            `json:"name,omitempty" column:"name:name;type:text;nullable"`
	AddressId *uuid.UUID         `json:"address_id,omitempty" column:"name:address_id;type:uuid;nullable"`
	Phone     *string            `json:"phone,omitempty" column:"name:phone;type:text;nullable"`
	CreatedAt *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"facilities" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Address                    *Addresses `json:"address,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:address_id"`
	DoctorFacilities           []*Doctor  `json:"doctor_facilities,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:facility_id"`
	PolisThroughDoctorFacility []*Poli    `json:"polis_through_doctor_facility,omitempty" join:"joinType:manyToMany;through:doctor;sourcePrimaryKey:id;sourceForeignKey:facility_id;targetPrimaryKey:id;targetForeign:facility_id"`
	UsersThroughDoctorFacility []*Users   `json:"users_through_doctor_facility,omitempty" join:"joinType:manyToMany;through:doctor;sourcePrimaryKey:id;sourceForeignKey:facility_id;targetPrimaryKey:id;targetForeign:facility_id"`
}
