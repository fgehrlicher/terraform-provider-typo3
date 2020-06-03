package api

import (
	"database/sql"
)

type Typo3UserGroup struct {
	Uid int
	Pid int

	Title       string
	Description string

	Deleted     int
	Hidden      int
	HideInLists int

	Tstamp int
	Crdate int

	WorkspacePerms int
	CruserID       int

	LockToDomain      string
	TSconfig          sql.NullString
	AllowedLanguages  string
	CategoryPerms     sql.NullString
	CustomOptions     sql.NullString
	DbMountpoints     sql.NullString
	ExplicitAllowdeny sql.NullString
	FileMountpoints   sql.NullString
	FilePermissions   sql.NullString
	GroupMods         sql.NullString
	NonExcludeFields  sql.NullString
	PagetypesSelect   string
	Subgroup          sql.NullString
	TablesModify      sql.NullString
	TablesSelect      sql.NullString
}
