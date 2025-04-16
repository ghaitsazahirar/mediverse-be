package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type EPrescriptions struct {
	db.ModelBase
	Id            uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	ReservationId *uuid.UUID         `json:"reservation_id,omitempty" column:"name:reservation_id;type:uuid;nullable"`
	DoctorId      *uuid.UUID         `json:"doctor_id,omitempty" column:"name:doctor_id;type:uuid;nullable"`
	PatientId     *uuid.UUID         `json:"patient_id,omitempty" column:"name:patient_id;type:uuid;nullable"`
	Diagnosis     *string            `json:"diagnosis,omitempty" column:"name:diagnosis;type:text;nullable"`
	Medications   *string            `json:"medications,omitempty" column:"name:medications;type:text;nullable"`
	Notes         *string            `json:"notes,omitempty" column:"name:notes;type:text;nullable"`
	CreatedAt     *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"e_prescriptions" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor      *Doctor       `json:"doctor,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	Patient     *Patients     `json:"patient,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:patient_id"`
	Reservation *Reservations `json:"reservation,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:reservation_id"`
}
