-- name: GetAuditOperations :many
select
    operation_name, created_at, updated_at 
from audit_operation ao;

-- name: GetAuditOperationByOperationName :one
select 
    operation_name, created_at, updated_at
from audit_operation
where operation_name = $1;
