CREATE OR REPLACE FUNCTION update_updated_at_column()
  RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$
language 'plpgsql';

CREATE TABLE "openevse"."chargers" (
  "id"         SERIAL8,
  "name"       VARCHAR(256)   NOT NULL,
  "host"       VARCHAR(256)   NOT NULL,
  "created_at" TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ(0) NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ(0),
  PRIMARY KEY ("id")
);

CREATE INDEX "idx_deleted_at"
  ON "openevse"."chargers" USING btree (
    "deleted_at" ASC NULLS LAST
  );

CREATE TRIGGER "tg_chargers_updated_at"
  BEFORE UPDATE
  ON "openevse"."chargers"
  FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
