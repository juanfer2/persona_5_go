BEGIN;
CREATE TABLE stats (
  id SERIAL
  --
  ,name       varchar(255) UNIQUE
  ,image_url  text

  --
  ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
  ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
  -- soft delete
  ,deleted_at   timestamptz
  --
  ,PRIMARY KEY (id)
);

CREATE INDEX "idx_stats_deleted_at" ON "stats" ("deleted_at");
CREATE INDEX "idx_stats_name" ON "stats" ("name");
COMMIT;
