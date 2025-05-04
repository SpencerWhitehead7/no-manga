import { eq } from "drizzle-orm";
import type { AnyD1Database, DrizzleD1Database } from "drizzle-orm/d1";
import { drizzle } from "drizzle-orm/d1";

import * as schema from "../../drizzle/schema";
import { getNameSlug } from "../utils";
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

export class Repo {
  private db: DrizzleD1Database<typeof schema>;

  constructor(d1Db: AnyD1Database) {
    this.db = drizzle(d1Db, { schema });
  }

  getAllMangas(): Promise<Manga[]> {
    return this.db
      .select()
      .from(mangaSchema)
      .orderBy(mangaSchema.name)
      .all()
      .then((allMangas) => allMangas.map(getNameSlug));
  }

  getManga(id: number): Promise<Manga | null> {
    return this.db
      .select()
      .from(mangaSchema)
      .where(eq(mangaSchema.id, id))
      .get()
      .then((manga) => (manga ? getNameSlug(manga) : null));
  }

  async getFullManga(id: number): Promise<FullManga | null> {
    const [manga, mangakas, magazines, chapters, genres] = await Promise.all([
      this.getManga(id),
      this.getMangakasByManga(id),
      this.getMagazinesByManga(id),
      this.getChaptersByManga(id),
      this.getGenresByManga(id),
    ]);

    return manga
      ? {
          ...manga,
          mangakas,
          magazines,
          chapters,
          genres,
        }
      : null;
  }

  private getMangakasByManga(id: number): Promise<MangakaWithJob[]> {
    return this.db
      .selectDistinct()
      .from(mangaMangakaJobSchema)
      .innerJoin(
        mangakaSchema,
        eq(mangaMangakaJobSchema.mangakaId, mangakaSchema.id),
      )
      .where(eq(mangaMangakaJobSchema.mangaId, id))
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

  private getMagazinesByManga(id: number): Promise<Magazine[]> {
    return this.db
      .selectDistinct()
      .from(magazineMangaSchema)
      .innerJoin(
        magazineSchema,
        eq(magazineMangaSchema.magazineId, magazineSchema.id),
      )
      .where(eq(magazineMangaSchema.mangaId, id))
      .orderBy(magazineSchema.name)
      .then((rows) => rows.map((row) => row.magazine).map(getNameSlug));
  }

  private getChaptersByManga(id: number): Promise<Chapter[]> {
    return this.db
      .select()
      .from(chapterSchema)
      .where(eq(chapterSchema.mangaId, id))
      .orderBy(chapterSchema.chapterNum);
  }

  private getGenresByManga(id: number): Promise<Genre["name"][]> {
    return this.db
      .selectDistinct()
      .from(mangaGenreSchema)
      .where(eq(mangaGenreSchema.mangaId, id))
      .orderBy(mangaGenreSchema.genre)
      .then((rows) => rows.map((row) => row.genre));
  }

  getAllMangakas(): Promise<Mangaka[]> {
    return this.db
      .select()
      .from(mangakaSchema)
      .orderBy(mangakaSchema.name)
      .all()
      .then((allMangakas) => allMangakas.map(getNameSlug));
  }

  getMangaka(id: number): Promise<Mangaka | null> {
    return this.db
      .select()
      .from(mangakaSchema)
      .where(eq(mangakaSchema.id, id))
      .get()
      .then((mangaka) => (mangaka ? getNameSlug(mangaka) : null));
  }

  async getFullMangaka(id: number): Promise<FullMangaka | null> {
    const [mangaka, mangas, magazines] = await Promise.all([
      this.getMangaka(id),
      this.getMangasByMangaka(id),
      this.getMagazinesByMangaka(id),
    ]);

    return mangaka
      ? {
          ...mangaka,
          mangas,
          magazines,
        }
      : null;
  }

  private getMangasByMangaka(id: number): Promise<MangaWithJob[]> {
    return this.db
      .selectDistinct()
      .from(mangaMangakaJobSchema)
      .innerJoin(mangaSchema, eq(mangaMangakaJobSchema.mangaId, mangaSchema.id))
      .where(eq(mangaMangakaJobSchema.mangakaId, id))
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

  private getMagazinesByMangaka(id: number): Promise<Magazine[]> {
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
      .where(eq(mangaMangakaJobSchema.mangakaId, id))
      .orderBy(magazineSchema.name)
      .then((rows) => rows.map((row) => row.magazine).map(getNameSlug));
  }

  getAllMagazines(): Promise<Magazine[]> {
    return this.db
      .select()
      .from(magazineSchema)
      .orderBy(magazineSchema.name)
      .all()
      .then((allMagazines) => allMagazines.map(getNameSlug));
  }

  getMagazine(id: number): Promise<Magazine | null> {
    return this.db
      .select()
      .from(magazineSchema)
      .where(eq(magazineSchema.id, id))
      .get()
      .then((magazine) => (magazine ? getNameSlug(magazine) : null));
  }

  async getFullMagazine(id: number): Promise<FullMagazine | null> {
    const [magazine, mangas, mangakas] = await Promise.all([
      this.getMagazine(id),
      this.getMangasByMagazine(id),
      this.getMangakasByMagazine(id),
    ]);

    return magazine
      ? {
          ...magazine,
          mangas,
          mangakas,
        }
      : null;
  }

  private getMangasByMagazine(id: number): Promise<Manga[]> {
    return this.db
      .selectDistinct()
      .from(magazineMangaSchema)
      .innerJoin(mangaSchema, eq(magazineMangaSchema.mangaId, mangaSchema.id))
      .where(eq(magazineMangaSchema.magazineId, id))
      .orderBy(mangaSchema.name)
      .then((rows) => rows.map((row) => row.manga).map(getNameSlug));
  }

  private getMangakasByMagazine(id: number): Promise<MangakaWithJob[]> {
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
      .where(eq(magazineMangaSchema.magazineId, id))
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

  getAllChapters(): Promise<Chapter[]> {
    return this.db
      .select()
      .from(chapterSchema)
      .orderBy(chapterSchema.updatedAt)
      .all();
  }
}
