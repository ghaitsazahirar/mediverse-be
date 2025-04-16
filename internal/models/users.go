package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Users struct {
	db.ModelBase
	Id        uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name      *string            `json:"name,omitempty" column:"name:name;type:text;nullable"`
	Email     string             `json:"email,omitempty" column:"name:email;type:text;nullable:false;unique"`
	Password  string             `json:"password,omitempty" column:"name:password;type:text;nullable:false"`
	Phone     *string            `json:"phone,omitempty" column:"name:phone;type:text;nullable"`
	CreatedAt *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	RoleId    *uuid.UUID         `json:"role_id,omitempty" column:"name:role_id;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"users" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorUsers                      []*Doctor       `json:"doctor_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	DoctorsThroughReservationsUser   []*Doctor       `json:"doctors_through_reservations_user,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	FacilitiesThroughDoctorUser      []*Facilities   `json:"facilities_through_doctor_user,omitempty" join:"joinType:manyToMany;through:doctor;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	PatientUsers                     []*Patients     `json:"patient_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	PolisThroughDoctorUser           []*Poli         `json:"polis_through_doctor_user,omitempty" join:"joinType:manyToMany;through:doctor;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	ReservationUsers                 []*Reservations `json:"reservation_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	Role                             *Roles          `json:"role,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:role_id"`
	SchedulesThroughReservationsUser []*Schedules    `json:"schedules_through_reservations_user,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
}
