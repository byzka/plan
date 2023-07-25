package models

import (
	"time"
)

type Search_type struct {
	Type  string
	Name  string
	Local string
}
type Answer struct {
	Type   string
	Name   string
	Price  string
	Local  string
	TimeAT time.Time
}

func Find_name(name string) (a []Answer, err error) {
	rows, err := DB.Query("SELECT type, name, local, price FROM search WHERE name = ?", name)
	if err != nil {
		return
	}
	for rows.Next() {
		an := Answer{TimeAT: time.Now()}
		if err := rows.Scan(&an.Type, &an.Name, &an.Local, &an.Price); err != nil {
			return
		}
		a = append(a, an)
	}
	rows.Close()
	return
}

func Find_local(Local string) (a []Answer, err error) {
	rows, err := DB.Query("SELECT type, name, local, price FROM search WHERE local = ?", Local)
	if err != nil {
		return
	}
	for rows.Next() {
		an := Answer{TimeAT: time.Now()}
		if err := rows.Scan(&an.Type, &an.Name, &an.Local, &an.Price); err != nil {
			return
		}
		a = append(a, an)
	}
	rows.Close()
	return
}
