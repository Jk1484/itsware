CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE cabinet_type AS (
    name TEXT,
    location TEXT
);

CREATE TABLE cabinets (
    id UUID PRIMARY KEY DEFAULT get_random_uuid(),
    name TEXT,
    location TEXT
);

CREATE OR REPLACE FUNCTION create_cabinet(cabinet cabinet_type)
RETURNS TABLE(id UUID, name TEXT, location TEXT) AS $$
BEGIN
    RETURN QUERY
    INSERT INTO cabinets (name, location)
    VALUES (cabinet.name, cabinet.location)
    RETURNING id, name, location;
END;
$$ LANGUAGE plpgsql;
