import Database from "better-sqlite3";
import { eq } from "drizzle-orm";
import type { BetterSQLite3Database } from "drizzle-orm/better-sqlite3";
import { drizzle } from "drizzle-orm/better-sqlite3";

import { getNameSlug } from "../../src/utils";
import * as schema from "../schema";
import {
  chapterSchema,
  magazineMangaSchema,
  magazineSchema,
  mangaGenreSchema,
  mangakaSchema,
  mangaMangakaJobSchema,
  mangaSchema,
} from "./schema";
import type {
  Chapter,
  FullMagazine,
  FullManga,
  FullMangaka,
  Genre,
  Job,
  Magazine,
  Manga,
  Mangaka,
  MangakaWithJob,
  MangaWithJob,
} from "./types";

export class BuildTimeRepo {
  private db: BetterSQLite3Database<typeof schema>;

  constructor(dbPath: string) {
    const sqlite = new Database(dbPath, { fileMustExist: true });
    this.db = drizzle(sqlite, { schema });
  }

  getAllMangas(): Manga[] {
    return this.db
      .select()
      .from(mangaSchema)
      .orderBy(mangaSchema.name)
      .all()
      .map(getNameSlug);
  }

  async getFullManga(manga: Manga): Promise<FullManga> {
    const [mangakas, magazines, chapters, genres] = await Promise.all([
      this.getMangakasByManga(manga),
      this.getMagazinesByManga(manga),
      this.getChaptersByManga(manga),
      this.getGenresByManga(manga),
    ]);

    return {
      ...manga,
      mangakas,
      magazines,
      chapters,
      genres,
    };
  }
  private getMangakasByManga(manga: Manga): Promise<MangakaWithJob[]> {
    return this.db
      .selectDistinct()
      .from(mangaMangakaJobSchema)
      .innerJoin(
        mangakaSchema,
        eq(mangaMangakaJobSchema.mangakaId, mangakaSchema.id),
      )
      .where(eq(mangaMangakaJobSchema.mangaId, manga.id))
      .orderBy(mangakaSchema.name)
      .then((rows) =>
        rows
          .map((row) => ({
            ...row.mangaka,
            job: row.manga_mangaka_job.job as unknown as Job,
          }))
          .map(getNameSlug),
      );
  }
  private getMagazinesByManga(manga: Manga): Promise<Magazine[]> {
    return this.db
      .selectDistinct()
      .from(magazineMangaSchema)
      .innerJoin(
        magazineSchema,
        eq(magazineMangaSchema.magazineId, magazineSchema.id),
      )
      .where(eq(magazineMangaSchema.mangaId, manga.id))
      .orderBy(magazineSchema.name)
      .then((rows) => rows.map((row) => row.magazine).map(getNameSlug));
  }
  private getChaptersByManga(manga: Manga): Promise<Chapter[]> {
    return this.db
      .select()
      .from(chapterSchema)
      .orderBy(chapterSchema.chapterNum)
      .where(eq(chapterSchema.mangaId, manga.id));
  }
  private getGenresByManga(manga: Manga): Promise<Genre["name"][]> {
    return this.db
      .selectDistinct()
      .from(mangaGenreSchema)
      .where(eq(mangaGenreSchema.mangaId, manga.id))
      .orderBy(mangaGenreSchema.genre)
      .then((rows) => rows.map((row) => row.genre));
  }

  getAllMangakas(): Mangaka[] {
    return this.db
      .select()
      .from(mangakaSchema)
      .orderBy(mangakaSchema.name)
      .all()
      .map(getNameSlug);
  }

  async getFullMangaka(mangaka: Mangaka): Promise<FullMangaka> {
    const [mangas, magazines] = await Promise.all([
      this.getMangasByMangaka(mangaka),
      this.getMagazinesByMangaka(mangaka),
    ]);

    return {
      ...mangaka,
      mangas,
      magazines,
    };
  }
  private getMangasByMangaka(mangaka: Mangaka): Promise<MangaWithJob[]> {
    return this.db
      .selectDistinct()
      .from(mangaMangakaJobSchema)
      .innerJoin(mangaSchema, eq(mangaMangakaJobSchema.mangaId, mangaSchema.id))
      .where(eq(mangaMangakaJobSchema.mangakaId, mangaka.id))
      .orderBy(mangaSchema.name)
      .then((rows) =>
        rows
          .map((row) => ({
            ...row.manga,
            job: row.manga_mangaka_job.job as unknown as Job,
          }))
          .map(getNameSlug),
      );
  }
  private getMagazinesByMangaka(mangaka: Mangaka): Promise<Magazine[]> {
    return this.db
      .selectDistinct()
      .from(mangaMangakaJobSchema)
      .innerJoin(mangaSchema, eq(mangaMangakaJobSchema.mangaId, mangaSchema.id))
      .innerJoin(
        magazineMangaSchema,
        eq(mangaSchema.id, magazineMangaSchema.mangaId),
      )
      .innerJoin(
        magazineSchema,
        eq(magazineMangaSchema.magazineId, magazineSchema.id),
      )
      .where(eq(mangaMangakaJobSchema.mangakaId, mangaka.id))
      .orderBy(magazineSchema.name)
      .then((rows) => rows.map((row) => row.magazine).map(getNameSlug));
  }

  getAllMagazines(): Magazine[] {
    return this.db
      .select()
      .from(magazineSchema)
      .orderBy(magazineSchema.name)
      .all()
      .map(getNameSlug);
  }

  async getFullMagazine(magazine: Magazine): Promise<FullMagazine> {
    const [mangas, mangakas] = await Promise.all([
      this.getMangasByMagazine(magazine),
      this.getMangakasByMagazine(magazine),
    ]);

    return {
      ...magazine,
      mangas,
      mangakas,
    };
  }
  private getMangasByMagazine(magazine: Magazine): Promise<Manga[]> {
    return this.db
      .selectDistinct()
      .from(magazineMangaSchema)
      .innerJoin(mangaSchema, eq(magazineMangaSchema.mangaId, mangaSchema.id))
      .where(eq(magazineMangaSchema.magazineId, magazine.id))
      .orderBy(mangaSchema.name)
      .then((rows) => rows.map((row) => row.manga).map(getNameSlug));
  }
  private getMangakasByMagazine(magazine: Magazine): Promise<MangakaWithJob[]> {
    return this.db
      .selectDistinct()
      .from(magazineMangaSchema)
      .innerJoin(mangaSchema, eq(magazineMangaSchema.mangaId, mangaSchema.id))
      .innerJoin(
        mangaMangakaJobSchema,
        eq(mangaSchema.id, mangaMangakaJobSchema.mangaId),
      )
      .innerJoin(
        mangakaSchema,
        eq(mangaMangakaJobSchema.mangakaId, mangakaSchema.id),
      )
      .where(eq(magazineMangaSchema.magazineId, magazine.id))
      .orderBy(mangakaSchema.name)
      .then((rows) =>
        rows
          .map((row) => ({
            ...row.mangaka,
            job: row.manga_mangaka_job.job as unknown as Job,
          }))
          .map(getNameSlug),
      );
  }

  getAllChapters(): Chapter[] {
    return this.db
      .select()
      .from(chapterSchema)
      .orderBy(chapterSchema.updatedAt)
      .all();
  }
}
