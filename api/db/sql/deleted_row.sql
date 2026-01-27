-- name: InsertDeletedRow :exec
insert into deleted_row (source_table, source_id, data)
  values (?, ?, ?);