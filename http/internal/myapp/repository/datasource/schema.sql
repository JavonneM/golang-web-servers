CREATE TABLE AUDIT_OPERATION(
    operation_name text primary key,
    created_at timestampz,
    updated_at timestampz,
);

CREATE TABLE AUDIT_CALL (
    id uuid primary key,
    result text,
    exception text,
    created_at timestampz,
    updated_at timestampz,
);

