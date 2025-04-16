package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Patients struct {
	db.ModelBase
	Id               uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	UserId           *uuid.UUID         `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable;unique"`
	BirthDate        *postgres.DateTime `json:"birth_date,omitempty" column:"name:birth_date;type:date;nullable"`
	Gender           *string            `json:"gender,omitempty" column:"name:gender;type:text;nullable"`
	BloodType        *string            `json:"blood_type,omitempty" column:"name:blood_type;type:text;nullable"`
	EmergencyContact *string            `json:"emergency_contact,omitempty" column:"name:emergency_contact;type:text;nullable"`
	CreatedAt        *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"patients" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorsThroughEPrescriptionsPatient      []*Doctor         `json:"doctors_through_e_prescriptions_patient,omitempty" join:"joinType:manyToMany;through:e_prescriptions;sourcePrimaryKey:id;sourceForeignKey:patient_id;targetPrimaryKey:id;targetForeign:patient_id"`
	EPrescriptionPatients                    []*EPrescriptions `json:"e_prescription_patients,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:patient_id"`
	ReservationsThroughEPrescriptionsPatient []*Reservations   `json:"reservations_through_e_prescriptions_patient,omitempty" join:"joinType:manyToMany;through:e_prescriptions;sourcePrimaryKey:id;sourceForeignKey:patient_id;targetPrimaryKey:id;targetForeign:patient_id"`
	User                                     *Users            `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
