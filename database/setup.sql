-- REFERENCE TABLES

CREATE TABLE "demo" (
  "name" varchar
    PRIMARY KEY
);

CREATE TABLE "job" (
  "name" varchar
    PRIMARY KEY
);

-- DATA TABLES

CREATE TABLE "magazine" (
  "id" int
    GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "other_names" varchar[],
  "description" varchar,
  "demo" varchar REFERENCES "demo"("name")
);

CREATE TABLE "mangaka" (
  "id" int
    GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "other_names" varchar[],
  "description" varchar
);

CREATE TABLE "manga" (
  "id" int
    GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  "name" varchar,
  "other_names" varchar[],
  "description" varchar,
  "demo" varchar REFERENCES "demo"("name"),
  "start_date" date,
  "end_date" date
);

CREATE TABLE "chapter" (
  "manga_id" int REFERENCES "manga"("id"),
  "chapter_num" float(4),
  "name" varchar,
  "page_count" int,
  "updated_at" timestamp without time zone,
  PRIMARY KEY ("manga_id", "chapter_num")
);

CREATE TABLE "genre" (
  "name" varchar
    PRIMARY KEY
);
