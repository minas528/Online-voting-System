package repository


import (
	"errors"
	"github.com/jinzhu/gorm"
"github.com/minas528/Online-voting-System/entities"
)

type MockRoleGormRepo struct {
	conn *gorm.DB
}

func NewMockRoleGormRepo(conn *gorm.DB) *MockRoleGormRepo  {
	return &MockRoleGormRepo{conn:conn}
}

func (mockRoleRepo *MockRoleGormRepo) Roles() ([]entities.Role, []error){
	roles := []entities.Role{entities.RoleMock}

	return roles, nil
}

func (mockRoleRepo *MockRoleGormRepo) Role(id uint) (*entities.Role, []error){
	role := entities.RoleMock

	return &role, nil
}
func (mockRoleRepo *MockRoleGormRepo) RoleByName(name string) (*entities.Role, []error){
	role := entities.RoleMock
	return &role, nil
}
func (mockRoleRepo *MockRoleGormRepo) UpdateRole(role *entities.Role) (*entities.Role, []error){
	r := entities.RoleMock

	return &r, nil
}
func (mockRoleRepo *MockRoleGormRepo) DeleteRole(id uint) (*entities.Role, []error){
	r := entities.RoleMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &r, nil
}
func (mockRoleRepo *MockRoleGormRepo) StoreRole(role *entities.Role) (*entities.Role, []error){
	r := role

	return r,nil
}
