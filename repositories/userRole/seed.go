package userrole

import (
	"github.com/deyr02/bnzlcrm/graph/model"
)

var AdminOperation []model.Operation

func getAdminOperation() []model.Operation {
	AdminOperation = append(AdminOperation, model.OperationGet)
	AdminOperation = append(AdminOperation, model.OperationPost)
	AdminOperation = append(AdminOperation, model.OperationPut)
	AdminOperation = append(AdminOperation, model.OperationDel)
	return AdminOperation
}

var _systemRole = true
var AdminUserRole = &model.NewUserRole{
	RoleName:   "Admin",
	SystemRole: &_systemRole,
	Operations: getAdminOperation(),
}

var isSystemRole = false
var _operations []model.Operation

func getITSupportOperations() []model.Operation {
	_operations = append(_operations, model.OperationGet)
	_operations = append(_operations, model.OperationPut)
	return _operations
}

var ITSupportUserRole = &model.NewUserRole{
	RoleName:   "Admin",
	SystemRole: &isSystemRole,
	Operations: getITSupportOperations(),
}
