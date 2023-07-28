package model

import "last/init_DB"

type Search struct {
	Type  string
	Name  string
	Local string
	Price string
}

func (s *Search) All() (search []Search, err error) {
	rows, err := init_DB.DB.Query("SELECT type, name, local, price FROM search")
	if err != nil {
		return
	}
	for rows.Next() {
		ss := Search{}
		if err = rows.Scan(&ss.Type, &ss.Name, &ss.Local, &ss.Price); err != nil {
			return
		}
		search = append(search, ss)
	}
	rows.Close()
	return
}
func (s *Search) T_S() (search []Search, err error) {
	rows, err := init_DB.DB.Query("SELECT type, name, local, price FROM search WHERE type = ?", s.Type)
	if err != nil {
		return
	}
	for rows.Next() {
		ss := Search{}
		if err := rows.Scan(&ss.Type, &ss.Name, &ss.Local, &ss.Price); err != nil {
			return
		}
		search = append(search, ss)
	}
	rows.Close()
	return
}

func (s *Search) N_S() (search []Search, err error) {
	rows, err := init_DB.DB.Query("SELECT type, name, local, price FROM search WHERE name = ?", s.Name)
	if err != nil {
		return
	}
	for rows.Next() {
		ss := Search{}
		if err := rows.Scan(&ss.Type, &ss.Name, &ss.Local, &ss.Price); err != nil {
			return
		}
		search = append(search, ss)
	}
	rows.Close()
	return
}
func (s *Search) L_S() (search []Search, err error) {
	rows, err := init_DB.DB.Query("SELECT type, name, local, price FROM search WHERE local = ?", s.Local)
	if err != nil {
		return
	}
	for rows.Next() {
		ss := Search{}
		if err := rows.Scan(&ss.Type, &ss.Name, &ss.Local, &ss.Price); err != nil {
			return
		}
		search = append(search, ss)
	}
	rows.Close()
	return
}

func (s *Search) Search_sth() (answer []Search, err error) {
	if s.Type != "" {
		answer, err = s.T_S()
		return answer, err
	} else if s.Name != "" {
		answer, err = s.N_S()
		return answer, err
	} else if s.Local != "" {
		answer, err = s.L_S()
		return answer, err
	} else {
		answer, err = s.All()
		return answer, err
	}
}
