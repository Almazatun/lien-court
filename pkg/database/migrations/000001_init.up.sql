CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "user" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar NULL,
  "email" varchar NOT NULL,
  "pass" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "link" (
  "id" UUID PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "short" varchar NOT NULL,
  "original" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" UUID NOT NULL
);
-- one to many
ALTER TABLE "link" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

COMMENT ON COLUMN "link"."short" IS 'can be empty string';