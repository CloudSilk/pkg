package constants

var (
	//超级管理员角色ID
	SuperAdminRoleID = ""
	//平台租户ID
	PlatformTenantID = ""
	//是否启用租户
	EnabelTenant = false
	//默认用户角色
	DefaultRoleID = ""
)

func SetSuperAdminRoleID(roleID string) {
	SuperAdminRoleID = roleID
}

func SetPlatformTenantID(tenantID string) {
	PlatformTenantID = tenantID
}

func SetEnabelTenant(enabelTenant bool) {
	EnabelTenant = enabelTenant
}
func SetDefaultRoleID(roleID string) {
	DefaultRoleID = roleID
}
