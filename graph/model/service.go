package model

import (
	"context"
	"errors"
)

type Service interface {
	CreateU(ctx context.Context, input NewUser) (*User, error)
	CreateC(ctx context.Context, input *NewCompany) (*Company, error)
	CreateJ(ctx context.Context, input *NewJob) (*Job, error)
	GetCompanyById(id int) (*Company, error)
	GetAllCompanies() ([]*Company, error)
	GetJobById(JobId int) (*Job, error)
	GetAllJobs() ([]*Job, error)
	GetUserById(id int) (*User, error)
	GetAllUsers() ([]*User, error)
}

func (s *Conn) CreateU(ctx context.Context, input NewUser) (*User, error) {
	u := User{
		Name:  input.Name,
		Email: input.Email,
	}
	tx := s.Dbc.WithContext(ctx).Create(&u)
	if tx.Error != nil {
		return &User{}, errors.New("error while creating user")
	}
	return &u, nil
}

func (s *Conn) CreateC(ctx context.Context, input *NewCompany) (*Company, error) {
	c := Company{
		Name:     input.Name,
		Location: input.Location,
	}
	tx := s.Dbc.WithContext(ctx).Create(&c)
	if tx.Error != nil {
		return &Company{}, errors.New("error while cvreating a company")
	}
	return &c, nil
}

func (s *Conn) CreateJ(ctx context.Context, input *NewJob) (*Job, error) {
	j := Job{
		Tittle:      input.Tittle,
		Description: input.Description,
		CompanyID:   input.CompanyID,
	}
	tx := s.Dbc.WithContext(ctx).Create(&j)
	if tx.Error != nil {
		return &Job{}, errors.New("error while creating job")
	}
	return &j, nil
}
func (s *Conn) GetCompanyById(id int) (*Company, error) {
	var com Company
	tx := s.Dbc.Where("id = ?", id).Find(&com)
	if tx.Error != nil {
		return &Company{}, errors.New("no company with the given id")
	}
	return &com, nil
}
func (s *Conn) GetAllCompanies() ([]*Company, error) {
	var com []*Company
	tx := s.Dbc.Find(&com)
	if tx.Error != nil {
		return []*Company{}, errors.New("error while fetching all companies")
	}
	return com, nil

}
func (s *Conn) GetJobById(JobId int) (*Job, error) {
	var job Job
	tx := s.Dbc.Where("job_id =?", JobId).Find(&job)
	if tx.Error != nil {
		return &Job{}, errors.New("error while fetching job by id")
	}
	return &job, nil

}
func (s *Conn) GetAllJobs() ([]*Job, error) {
	var j []*Job
	tx := s.Dbc.Find(&j)
	if tx.Error != nil {
		return []*Job{}, errors.New("error fetching all jobs")
	}
	return j, nil
}
func (s *Conn) GetUserById(id int) (*User, error) {
	var user User
	tx := s.Dbc.Where("id =?", id).Find(&user)
	if tx.Error != nil {
		return &User{}, errors.New("error while fetching user by id")
	}
	return &user, nil

}
func (s *Conn) GetAllUsers() ([]*User, error) {
	var j []*User
	tx := s.Dbc.Find(&j)
	if tx.Error != nil {
		return []*User{}, errors.New("error fetching all users")
	}
	return j, nil
}
