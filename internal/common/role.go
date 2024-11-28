package common

import (
	deps "auth/pkg/user_v1"
)

func RoleToValue(str string) deps.Role {
	role := deps.Role(deps.Role_value[str])
	return role
}
