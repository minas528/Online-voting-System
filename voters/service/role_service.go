package service

import (
	"github.com/minas528/Online-voting-System/entities"
	"github.com/minas528/Online-voting-System/voters"
)

type RoleService struct {
	roleRepo voters.RoleRepository
}

// NewRoleService  returns new RoleService
func NewRoleService(RoleRepo voters.RoleRepository) voters.RoleService {
	return &RoleService{roleRepo: RoleRepo}
}

// Roles returns all stored roles
func (rs *RoleService) Roles() ([]entities.Role, []error) {

	rls, errs := rs.roleRepo.Roles()
	if len(errs) > 0 {
		return nil, errs
	}
	return rls, errs

}

// RoleByName returns a role identified by its name
func (rs *RoleService) RoleByName(name string) (*entities.Role, []error) {
	role, errs := rs.roleRepo.RoleByName(name)
	if len(errs) > 0 {
		return nil, errs
	}
	return role, errs
}

// Role retrievs a given user role by its id
func (rs *RoleService) Role(id uint) (*entities.Role, []error) {
	rl, errs := rs.roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

// UpdateRole updates a given user role
func (rs *RoleService) UpdateRole(role *entities.Role) (*entities.Role, []error) {
	rl, errs := rs.roleRepo.UpdateRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

// DeleteRole deletes a given user role
func (rs *RoleService) DeleteRole(id uint) (*entities.Role, []error) {

	rl, errs := rs.roleRepo.DeleteRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

// StoreRole stores a given user role
func (rs *RoleService) StoreRole(role *entities.Role) (*entities.Role, []error) {

	rl, errs := rs.roleRepo.StoreRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}
