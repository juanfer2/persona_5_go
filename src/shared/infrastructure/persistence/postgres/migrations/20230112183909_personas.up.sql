BEGIN;
CREATE TABLE personas (
  id SERIAL
  --
  ,name       varchar(255) UNIQUE
  ,level      INTEGER
  ,rare       boolean
  ,image_url  text

  ,arcana_id INTEGER REFERENCES arcanas (id)
  ,personality_id INTEGER REFERENCES personalities (id)
  --
  ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
  ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
  -- soft delete
  ,deleted_at   timestamptz
  --
  ,PRIMARY KEY (id)
);

CREATE INDEX "idx_personas_deleted_at" ON "personas" ("deleted_at");
CREATE INDEX "idx_personas_name" ON "personas" ("name");
COMMIT;
