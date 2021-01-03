-- CUSTOM CHECKS

CREATE OR REPLACE FUNCTION not_empty(v varchar) RETURNS boolean AS $$
  SELECT TRIM(v) <> ''
$$ LANGUAGE SQL STRICT IMMUTABLE;

CREATE OR REPLACE FUNCTION array_no_empties(v varchar[]) RETURNS boolean AS $$
  SELECT v <> '{}' AND (NOT EXISTS (SELECT 1 FROM unnest(v) s WHERE s = NULL OR TRIM(s) = ''))
$$ LANGUAGE SQL STRICT IMMUTABLE;

-- REFERENCE TABLES

CREATE TABLE "demo" (
  "name" varchar
    PRIMARY KEY CONSTRAINT no_empty_name CHECK (not_empty("name"))
);

CREATE TABLE "job" (
  "name" varchar
    PRIMARY KEY CONSTRAINT no_empty_name CHECK (not_empty("name"))
);

-- DATA TABLES

CREATE TABLE "magazine" (
  "id" int
    GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar
    NOT NULL CONSTRAINT no_empty_name CHECK (not_empty("name")),
  "other_names" varchar[]
    CONSTRAINT no_empties_other_names CHECK (array_no_empties("other_names")),
  "description" varchar
    NOT NULL CONSTRAINT no_empty_description CHECK (not_empty("description")),
  "demo" varchar REFERENCES "demo"("name")
    NOT NULL
);

CREATE TABLE "mangaka" (
  "id" int
    GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar
    NOT NULL CONSTRAINT no_empty_name CHECK (not_empty("name")),
  "other_names" varchar[]
    CONSTRAINT no_empties_other_names CHECK (array_no_empties("other_names")),
  "description" varchar
    NOT NULL CONSTRAINT no_empty_description CHECK (not_empty("description"))
);

CREATE TABLE "manga" (
  "id" int
    GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar
    NOT NULL CONSTRAINT no_empty_name CHECK (not_empty("name")),
  "other_names" varchar[]
    CONSTRAINT no_empties_other_names CHECK (array_no_empties("other_names")),
  "description" varchar
    NOT NULL CONSTRAINT no_empty_description CHECK (not_empty("description")),
  "demo" varchar REFERENCES "demo"("name")
    NOT NULL,
  "start_date" date
    NOT NULL,
  "end_date" date
);

CREATE TABLE "chapter" (
  "manga_id" int REFERENCES "manga"("id")
    NOT NULL,
  "chapter_num" float(4)
    NOT NULL,
  "name" varchar
    CONSTRAINT no_empty_name CHECK (not_empty("name")),
  "page_count" int
    NOT NULL,
  "updated_at" timestamptz
    NOT NULL,
  PRIMARY KEY ("manga_id", "chapter_num")
);

CREATE TABLE "genre" (
  "name" varchar
    PRIMARY KEY CONSTRAINT no_empty_name CHECK (not_empty("name"))
);

-- RELATIONS

CREATE TABLE "magazine_manga" (
  "magazine_id" int REFERENCES "magazine"("id") ON DELETE CASCADE
    NOT NULL,
  "manga_id" int REFERENCES "manga"("id") ON DELETE CASCADE
    NOT NULL
);

CREATE TABLE "manga_mangaka_job" (
  "manga_id" int REFERENCES "manga"("id") ON DELETE CASCADE
    NOT NULL,
  "mangaka_id" int REFERENCES "mangaka"("id") ON DELETE CASCADE
    NOT NULL,
  "job"  varchar REFERENCES "job"("name") ON DELETE CASCADE
    NOT NULL,
  UNIQUE("manga_id", "mangaka_id")
);

CREATE TABLE "manga_genre" (
  "manga_id" int REFERENCES "manga"("id") ON DELETE CASCADE
    NOT NULL,
  "genre" varchar REFERENCES "genre"("name") ON DELETE CASCADE
    NOT NULL,
  UNIQUE("manga_id", "genre")
);
