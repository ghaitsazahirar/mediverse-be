package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"github.com/sev-2/raiden/pkg/postgres"
)

type Doctor struct {
	db.ModelBase
	Id             uuid.UUID          `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name           *string            `json:"name,omitempty" column:"name:name;type:text;nullable"`
	Specialization *string            `json:"specialization,omitempty" column:"name:specialization;type:text;nullable"`
	FacilityId     *uuid.UUID         `json:"facility_id,omitempty" column:"name:facility_id;type:uuid;nullable"`
	PhotoUrl       *string            `json:"photo_url,omitempty" column:"name:photo_url;type:text;nullable"`
	CreatedAt      *postgres.DateTime `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable;default:now()"`
	PoliId         *uuid.UUID         `json:"poli_id,omitempty" column:"name:poli_id;type:uuid;nullable"`
	UserId         *uuid.UUID         `json:"user_id,omitempty" column:"name:user_id;type:uuid;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctor" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	EPrescriptionDoctors                    []*EPrescriptions `json:"e_prescription_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	Facility                                *Facilities       `json:"facility,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:facility_id"`
	Poli                                    *Poli             `json:"poli,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:poli_id"`
	PatientsThroughEPrescriptionsDoctor     []*Patients       `json:"patients_through_e_prescriptions_doctor,omitempty" join:"joinType:manyToMany;through:e_prescriptions;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	ReservationDoctors                      []*Reservations   `json:"reservation_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	ReservationsThroughEPrescriptionsDoctor []*Reservations   `json:"reservations_through_e_prescriptions_doctor,omitempty" join:"joinType:manyToMany;through:e_prescriptions;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	ScheduleDoctors                         []*Schedules      `json:"schedule_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	SchedulesThroughReservationsDoctor      []*Schedules      `json:"schedules_through_reservations_doctor,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	ServicesThroughSchedulesDoctor          []*Services       `json:"services_through_schedules_doctor,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	User                                    *Users            `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	UsersThroughReservationsDoctor          []*Users          `json:"users_through_reservations_doctor,omitempty" join:"joinType:manyToMany;through:reservations;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
}
