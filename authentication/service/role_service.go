package service

import (
	"github.com/minase528/Online-voting-System/entities"
	"github.com/minase528/Online-voting-System/authentication"
)

// RoleService implements menu.RoleService interface
type RoleService struct {
	roleRepo user.RoleRepository
}

// NewRoleService  returns new RoleService
func NewRoleService(RoleRepo user.RoleRepository) user.RoleService {
	return &RoleService{roleRepo: RoleRepo}
}

// Roles returns all stored roles
func (rs *RoleService) Roles() ([]entity.Role, []error) {

	rls, errs := rs.roleRepo.Roles()
	if len(errs) > 0 {
		return nil, errs
	}
	return rls, errs

}

// RoleByName returns a role identified by its name
func (rs *RoleService) RoleByName(name string) (*entity.Role, []error) {
	role, errs := rs.roleRepo.RoleByName(name)
	if len(errs) > 0 {
		return nil, errs
	}
	return role, errs
}

// Role retrievs a given user role by its id
func (rs *RoleService) Role(id uint) (*entity.Role, []error) {
	rl, errs := rs.roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

// UpdateRole updates a given user role
func (rs *RoleService) UpdateRole(role *entity.Role) (*entity.Role, []error) {
	rl, errs := rs.roleRepo.UpdateRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs

}

// DeleteRole deletes a given user role
func (rs *RoleService) DeleteRole(id uint) (*entity.Role, []error) {

	rl, errs := rs.roleRepo.DeleteRole(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}

// StoreRole stores a given user role
func (rs *RoleService) StoreRole(role *entity.Role) (*entity.Role, []error) {

	rl, errs := rs.roleRepo.StoreRole(role)
	if len(errs) > 0 {
		return nil, errs
	}
	return rl, errs
}
