BEGIN;
CREATE TABLE arcanas (
  id SERIAL
  --
  ,name   varchar(255) UNIQUE
  ,image_url  varchar(255)
  ,character_name  varchar(255)
  ,image_character_url  varchar(255)
  --
  ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
  ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
  -- soft delete
  ,deleted_at   timestamptz
  --
  ,PRIMARY KEY (id)
);

CREATE INDEX "idx_arcanas_deleted_at" ON "arcanas" ("deleted_at");
CREATE INDEX "idx_arcanas_name" ON "arcanas" ("name");
CREATE INDEX "idx_arcanas_character_name" ON "arcanas" ("character_name");
COMMIT;
