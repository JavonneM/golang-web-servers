-- +goose Up
CREATE TABLE AUDIT_OPERATION(
    operation_name text primary key,
    created_at timestamptz,
    updated_at timestamptz
);

CREATE TABLE AUDIT_CALL (
    id uuid primary key,
    result text,
    exception text,
    created_at timestamptz,
    updated_at timestamptz
);


-- +goose Down

Drop table AUDIT_CALL;
Drop table AUDIT_OPERATION;


