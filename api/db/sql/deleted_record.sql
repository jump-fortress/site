-- name: InsertDeletedRecord :exec
insert into deleted_record (source_table, source_id, data)
  values (?, ?, ?);