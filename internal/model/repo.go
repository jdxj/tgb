package model

import "fmt"

type Repository struct {
	Owner string
	Name  string
}

func (r *Repository) Id() string {
	return fmt.Sprintf("%s/%s", r.Owner, r.Name)
}
