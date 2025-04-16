package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Schedules struct {
	db.ModelBase
	Id        uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	DoctorId  *uuid.UUID         `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable"`
	Date      *postgres.DateTime `json:"date,omitempty" column:"name:date;type:date;nullable"`
	Time      *postgres.DateTime `json:"time,omitempty" column:"name:time;type:time;nullable"`
	ServiceId *uuid.UUID         `json:"service_id,omitempty" column:"name:service_id;type:uuid;nullable"`
	CreatedAt *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"schedules" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor                             *Doctor         `json:"doctor,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	DoctorsThroughReservationsSchedule []*Doctor       `json:"doctors_through_reservations_schedule,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
	ReservationSchedules               []*Reservations `json:"reservation_schedules,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:schedule_id"`
	Service                            *Services       `json:"service,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:service_id"`
	UsersThroughReservationsSchedule   []*Users        `json:"users_through_reservations_schedule,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
}
