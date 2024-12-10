package main

import "time"

type user struct {
	Username string
	Password string
}

type Patient struct {
	Id           string       `json:"id" bson:"_id,omitempty"`
	Name         string       `json:"name"`
	Birthdate    time.Time    `json:"birthdate"`
	PreferedName string       `json:"preferedname"`
	Pronouns     string       `json:"pronouns"`
	Height       string       `json:"height"`
	Weight       int          `json:"weight"`
	Bpsys        int          `json:"bpsys"`
	Bpdys        int          `json:"bpdys"`
	Hr           int          `json:"hr"`
	Alergies     []string     `json:"alergies"`
	Medications  []Medication `json:"medications"`
	Unit         string       `json:"unit"`
	Histories    []History    `json:"histories"`
}

type Medication struct {
	Name      string
	dosage    int
	frequency string
}

type History struct {
	Date     time.Time
	Recorder string
	Body     string
}
