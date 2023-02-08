BEGIN;
CREATE TABLE persona_stats (
  id SERIAL
  --
  ,value      INTEGER
  ,image_url  text

  ,stat_id INTEGER REFERENCES stats (id)
  ,persona_id INTEGER REFERENCES personas (id)
  --
  ,created_at timestamptz  NOT NULL  DEFAULT current_timestamp
  ,updated_at timestamptz  NOT NULL  DEFAULT current_timestamp
  -- soft delete
  ,deleted_at   timestamptz
  --
  ,PRIMARY KEY (id)
);

CREATE INDEX "idx_persona_stats_deleted_at" ON "persona_stats" ("deleted_at");
COMMIT;
