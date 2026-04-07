package postgresql

import "database/sql"

func InitDB(db *sql.DB) error {
	sqlStmt := `
create table if not exists users
(
    id              serial  not null primary key,
    android_version text    null,
    device_model    text    null,
    manufacturer    text    null,
    total_ram_gb    integer null,
    app_version     text    null
);
`
	_, err := db.Exec(sqlStmt)
	return err
}
