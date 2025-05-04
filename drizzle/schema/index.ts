import { sql } from "drizzle-orm";
import {
  index,
  integer,
  primaryKey,
  real,
  sqliteTable,
  text,
  unique,
} from "drizzle-orm/sqlite-core";

// -- ENUMS START --
export const demo = sqliteTable("demo", {
  name: text("name", {
    enum: ["kodomo", "shonen", "shojo", "seinen", "josei", "seijin", "mina"],
  }).primaryKey(),
});
export type Demo = typeof demo.$inferSelect;
export type NewDemo = typeof demo.$inferInsert;

export const job = sqliteTable("job", {
  name: text("name", {
    enum: ["author", "artist", "author_artist"],
  }).primaryKey(),
});
export type Job = typeof job.$inferSelect;
export type NewJob = typeof job.$inferInsert;
// -- ENUMS END --

// -- DATA START --
export const magazine = sqliteTable(
  "magazine",
  {
    id: integer("id").primaryKey({ autoIncrement: true }),
    name: text("name").notNull(),
    otherNames: text("other_names", { mode: "json" })
      .notNull()
      .$type<string[]>(),
    description: text("description").notNull(),
    demo: text("demo")
      .references(() => demo.name)
      .notNull(),
  },
  (table) => [
    index("idx_magazine_name").on(table.name),
    index("idx_magazine_demo").on(table.demo),
  ],
);
export type Magazine = typeof magazine.$inferSelect;
export type NewMagazine = typeof magazine.$inferInsert;

export const mangaka = sqliteTable(
  "mangaka",
  {
    id: integer("id").primaryKey({ autoIncrement: true }),
    name: text("name").notNull(),
    otherNames: text("other_names", { mode: "json" })
      .notNull()
      .$type<string[]>(),
    description: text("description").notNull(),
  },
  (table) => [index("idx_mangaka_name").on(table.name)],
);
export type Mangaka = typeof mangaka.$inferSelect;
export type NewMangaka = typeof mangaka.$inferInsert;

export const manga = sqliteTable(
  "manga",
  {
    id: integer("id").primaryKey({ autoIncrement: true }),
    name: text("name").notNull(),
    otherNames: text("other_names", { mode: "json" })
      .notNull()
      .$type<string[]>(),
    description: text("description").notNull(),
    demo: text("demo")
      .references(() => demo.name)
      .notNull(),
    startDate: integer("start_date", { mode: "timestamp" }).notNull(),
    endDate: integer("end_date", { mode: "timestamp" }),
  },
  (table) => [
    index("idx_manga_name").on(table.name),
    index("idx_manga_demo").on(table.demo),
  ],
);
export type Manga = typeof manga.$inferSelect;
export type NewManga = typeof manga.$inferInsert;

export const chapter = sqliteTable(
  "chapter",
  {
    mangaId: integer("manga_id")
      .references(() => manga.id)
      .notNull(),
    chapterNum: real("chapter_num").notNull(),
    name: text("name"),
    pageCount: integer("page_count").notNull(),
    updatedAt: integer("updated_at", { mode: "timestamp" }).default(
      sql`CURRENT_TIMESTAMP`,
    ),
  },
  (table) => [
    primaryKey({ columns: [table.mangaId, table.chapterNum] }),
    index("idx_chapter_manga_id").on(table.mangaId),
    index("idx_chapter_updated_at").on(table.updatedAt),
  ],
);
export type Chapter = typeof chapter.$inferSelect;
export type NewChapter = typeof chapter.$inferInsert;

export const genre = sqliteTable("genre", {
  name: text("name").primaryKey(),
});
export type Genre = typeof genre.$inferSelect;
export type NewGenre = typeof genre.$inferInsert;
// -- DATA END --

// -- RELATIONS START --
export const magazineManga = sqliteTable(
  "magazine_manga",
  {
    magazineId: integer("magazine_id")
      .references(() => magazine.id)
      .notNull(),
    mangaId: integer("manga_id")
      .references(() => manga.id)
      .notNull(),
  },
  (table) => [
    index("idx_magazine_manga_magazine_id").on(table.magazineId),
    index("idx_magazine_manga_manga_id").on(table.mangaId),
  ],
);

export const mangaMangakaJob = sqliteTable(
  "manga_mangaka_job",
  {
    mangaId: integer("manga_id")
      .references(() => manga.id)
      .notNull(),
    mangakaId: integer("mangaka_id")
      .references(() => mangaka.id)
      .notNull(),
    job: text("job")
      .references(() => job.name)
      .notNull(),
  },
  (table) => [
    unique("unique_manga_mangaka").on(table.mangaId, table.mangakaId),
    index("idx_manga_mangaka_job_manga_id").on(table.mangaId),
    index("idx_manga_mangaka_job_mangaka_id").on(table.mangakaId),
  ],
);

export const mangaGenre = sqliteTable(
  "manga_genre",
  {
    mangaId: integer("manga_id")
      .references(() => manga.id)
      .notNull(),
    genre: text("genre")
      .references(() => genre.name)
      .notNull(),
  },
  (table) => [
    unique("unique_manga_genre").on(table.mangaId, table.genre),
    index("idx_manga_genre_manga_id").on(table.mangaId),
    index("idx_manga_genre_genre").on(table.genre),
  ],
);
// -- RELATIONS END --
