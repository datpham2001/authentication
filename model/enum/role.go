package enum

type AccountRoleEnumValue string

type accountRole struct {
	Admin AccountRoleEnumValue
	Guest AccountRoleEnumValue
}

var AccountRole = &accountRole{
	Admin: "ADMIN",
	Guest: "GUEST",
}

type AccountStatusValue string

type accountStatus struct {
	Active   AccountStatusValue
	Inactive AccountStatusValue
}

var AccountStatus = &accountStatus{
	Active:   "ACTIVE",
	Inactive: "INACTIVE",
}
