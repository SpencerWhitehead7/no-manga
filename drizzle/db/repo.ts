import Database from "better-sqlite3";
import type { BetterSQLite3Database } from "drizzle-orm/better-sqlite3";
import { drizzle } from "drizzle-orm/better-sqlite3";

import { getNameSlug } from "../../src/utils";
import * as schema from "../schema";
import {
  chapterSchema,
  magazineSchema,
  mangakaSchema,
  mangaSchema,
} from "./schema";
import type { Chapter, Magazine, Manga, Mangaka } from "./types";

export class BuildTimeRepo {
  private db: BetterSQLite3Database<typeof schema>;

  constructor(dbPath: string) {
    const sqlite = new Database(dbPath, { fileMustExist: true });
    this.db = drizzle(sqlite, { schema });
  }

  getAllMangas(): Manga[] {
    return this.db.select().from(mangaSchema).all().map(getNameSlug);
  }

  getAllMangakas(): Mangaka[] {
    return this.db.select().from(mangakaSchema).all().map(getNameSlug);
  }

  getAllMagazines(): Magazine[] {
    return this.db.select().from(magazineSchema).all().map(getNameSlug);
  }

  getAllChapters(): Chapter[] {
    return this.db.select().from(chapterSchema).all();
  }
}
