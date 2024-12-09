CREATE TYPE device AS (
    id INT,
    cabinet_id INT,
    team_id INT,
    serial TEXT,
    status TEXT,
    profile device_profile
);

CREATE TYPE device_profile AS (
    id INT,
    name TEXT,
    type TEXT
);

CREATE TABLE devices (
    id SERIAL PRIMARY KEY,
    cabinet_id INT,
    team_id INT,
    serial TEXT NOT NULL,
    status TEXT DEFAULT 'active',
    profile device_profile
);

CREATE TABLE device_profiles (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL
);

CREATE OR REPLACE FUNCTION create_device(serial TEXT, cabinet_id INT, team_id INT, profile device_profile)
RETURNS device AS $$
DECLARE
    new_device device;
BEGIN
    INSERT INTO devices (serial, cabinet_id, team_id, profile)
    VALUES (serial, cabinet_id, team_id, profile)
    RETURNING id
    INTO new_device;

    RETURN new_device;
END;
$$ LANGUAGE plpgsql;
