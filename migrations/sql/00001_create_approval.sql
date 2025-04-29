-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- Enable the uuid-ossp extension to use uuid_generate_v4()
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the approvals table
CREATE TABLE IF NOT EXISTS approvals (
    "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    "flow_id" UUID NOT NULL,
    "flow_name" VARCHAR(100) NOT NULL DEFAULT '',
    "status" VARCHAR(20) NOT NULL DEFAULT '',
    "created_by" UUID NOT NULL,
    "updated_by" UUID,
    "comments" JSONB DEFAULT '[]',
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),          
    CONSTRAINT unique_flow_id UNIQUE (flow_id) -- Add unique constraint on flow_id
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS approvals;