package common

import deps "auth/pkg/user_v1"

func ToPointer(str string) *string {
	return &str
}

func RoleToPointer(role deps.Role) *deps.Role {
	return &role
}
