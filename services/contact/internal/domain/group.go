package domain

import "errors"

type Group struct {
	Id   uint
	Name Name
}

type Name string

func NewGroup(id uint, name Name) (*Group, error) {
	group := Group{
		Id:   id,
		Name: name,
	}

	if len(group.Name) > 250 {
		return nil, errors.New("name must not exceed 250 symbols")
	}

	return &group, nil
}
