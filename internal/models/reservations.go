package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Reservations struct {
	db.ModelBase
	Id         uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId     *uuid.UUID         `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable"`
	DoctorId   *uuid.UUID         `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable"`
	ScheduleId *uuid.UUID         `json:"schedule_id,omitempty" column:"name:schedule_id;type:uuid;nullable"`
	Status     *string            `json:"status,omitempty" column:"name:status;type:text;nullable"`
	Notes      *string            `json:"notes,omitempty" column:"name:notes;type:text;nullable"`
	CreatedAt  *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservations" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor                                   *Doctor           `json:"doctor,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	DoctorsThroughEPrescriptionsReservation  []*Doctor         `json:"doctors_through_e_prescriptions_reservation,omitempty" join:"joinType:manyToMany;through:e_prescriptions;sourcePrimaryKey:id;sourceForeignKey:reservation_id;targetPrimaryKey:id;targetForeign:reservation_id"`
	EPrescriptionReservations                []*EPrescriptions `json:"e_prescription_reservations,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:reservation_id"`
	PaymentReservations                      []*Payments       `json:"payment_reservations,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:reservation_id"`
	PatientsThroughEPrescriptionsReservation []*Patients       `json:"patients_through_e_prescriptions_reservation,omitempty" join:"joinType:manyToMany;through:e_prescriptions;sourcePrimaryKey:id;sourceForeignKey:reservation_id;targetPrimaryKey:id;targetForeign:reservation_id"`
	Schedule                                 *Schedules        `json:"schedule,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:schedule_id"`
	User                                     *Users            `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
