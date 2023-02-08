BEGIN;
CREATE TABLE elements (
  id SERIAL
  --
  ,name   varchar(255) UNIQUE
  ,image_url  varchar(255)
  --
  ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
  ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
  -- soft delete
  ,deleted_at   timestamptz
  --
  ,PRIMARY KEY (id)
);

CREATE INDEX "idx_elements_deleted_at" ON "elements" ("deleted_at");
CREATE INDEX "idx_elements_name" ON "elements" ("name");
COMMIT;
