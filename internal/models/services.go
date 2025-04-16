package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Services struct {
	db.ModelBase
	Id          uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Description *string            `json:"description,omitempty" column:"name:description;type:text;nullable"`
	Price       *int32             `json:"price,omitempty" column:"name:price;type:integer;nullable"`
	CreatedAt   *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	Type        *string            `json:"type,omitempty" column:"name:type;type:text;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"services" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorsThroughSchedulesService []*Doctor    `json:"doctors_through_schedules_service,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:service_id;targetPrimaryKey:id;targetForeign:service_id"`
	ScheduleServices               []*Schedules `json:"schedule_services,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:service_id"`
}
