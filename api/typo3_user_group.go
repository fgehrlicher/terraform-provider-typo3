package api

import (
	"database/sql"
)

type Typo3UserGroup struct {
	Uid int `db:"uid"`
	Pid int `db:"pid"`

	Title       string `db:"title"`
	Description string `db:"description"`

	Deleted      int    `db:"deleted"`
	Hidden       int    `db:"hidden"`
	HideInLists  int    `db:"hide_in_lists"`
	LockToDomain string `db:"lockToDomain"`

	Tstamp int `db:"tstamp"`
	Crdate int `db:"crdate"`

	WorkspacePerms int `db:"workspace_perms"`
	CruserID       int `db:"cruser_id"`

	TSconfig          sql.NullString `db:"TSconfig"`
	AllowedLanguages  string         `db:"allowed_languages"`
	CategoryPerms     sql.NullString `db:"category_perms"`
	CustomOptions     sql.NullString `db:"custom_options"`
	DbMountpoints     sql.NullString `db:"db_mountpoints"`
	ExplicitAllowdeny sql.NullString `db:"explicit_allowdeny"`
	FileMountpoints   sql.NullString `db:"file_mountpoints"`
	FilePermissions   sql.NullString `db:"file_permissions"`
	GroupMods         sql.NullString `db:"groupMods"`
	NonExcludeFields  sql.NullString `db:"non_exclude_fields"`
	PagetypesSelect   string         `db:"pagetypes_select"`
	Subgroup          sql.NullString `db:"subgroup"`
	TablesModify      sql.NullString `db:"tables_modify"`
	TablesSelect      sql.NullString `db:"tables_select"`
}
