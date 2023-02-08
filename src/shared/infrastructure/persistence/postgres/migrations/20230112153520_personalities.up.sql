BEGIN;
CREATE TABLE personalities (
  id SERIAL
  --
  ,name   varchar(255) UNIQUE
  --
  ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
  ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
  -- soft delete
  ,deleted_at   timestamptz
  --
  ,PRIMARY KEY (id)
);

CREATE INDEX "idx_personalities_deleted_at" ON "personalities" ("deleted_at");
CREATE INDEX "idx_personalities_name" ON "personalities" ("name");
COMMIT;
