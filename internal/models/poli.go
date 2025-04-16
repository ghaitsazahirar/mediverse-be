package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Poli struct {
	db.ModelBase
	Id          uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name        string             `json:"name,omitempty" column:"name:name;type:text;nullable:false;unique"`
	Description *string            `json:"description,omitempty" column:"name:description;type:text;nullable"`
	CreatedAt   *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"poli" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorPolis                 []*Doctor     `json:"doctor_polis,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:poli_id"`
	FacilitiesThroughDoctorPoli []*Facilities `json:"facilities_through_doctor_poli,omitempty" join:"joinType:manyToMany;through:doctor;sourcePrimaryKey:id;sourceForeignKey:poli_id;targetPrimaryKey:id;targetForeign:poli_id"`
	UsersThroughDoctorPoli      []*Users      `json:"users_through_doctor_poli,omitempty" join:"joinType:manyToMany;through:doctor;sourcePrimaryKey:id;sourceForeignKey:poli_id;targetPrimaryKey:id;targetForeign:poli_id"`
}
