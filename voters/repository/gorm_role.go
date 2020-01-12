package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/minas528/Online-voting-System/entities"
)

type RoleGormRepo struct {
	conn *gorm.DB
}

func NewRoleGormRepo(conn *gorm.DB) *RoleGormRepo  {
	return &RoleGormRepo{conn:conn}
}

func (roleRepo *RoleGormRepo) Roles() ([]entities.Role, []error){
	roles := []entities.Role{}
	errs := roleRepo.conn.Find(&roles).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return roles, errs
}

func(roleRepo *RoleGormRepo) Role(id uint) (*entities.Role, []error){
	role := entities.Role{}
	errs := roleRepo.conn.First(&role, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs
}
func (roleRepo *RoleGormRepo) RoleByName(name string) (*entities.Role, []error){
	role := entities.Role{}
	errs := roleRepo.conn.Find(&role, "name=?", name).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &role, errs
}
func (roleRepo *RoleGormRepo) UpdateRole(role *entities.Role) (*entities.Role, []error){
	r := role
	errs := roleRepo.conn.Save(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}
func (roleRepo *RoleGormRepo)DeleteRole(id uint) (*entities.Role, []error){
	r, errs := roleRepo.Role(id)
	if len(errs) > 0 {
		return nil, errs
	}
	errs = roleRepo.conn.Delete(r, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}
func (roleRepo *RoleGormRepo) StoreRole(role *entities.Role) (*entities.Role, []error){
	r := role
	errs := roleRepo.conn.Create(r).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return r, errs
}