import type {
  Chapter as ChapterSchema,
  Demo as DemoSchema,
  Genre as GenreSchema,
  Job as JobSchema,
  Magazine as MagazineSchema,
  Manga as MangaSchema,
  Mangaka as MangakaSchema,
} from "./schema";

export type Chapter = ChapterSchema;
export type Demo = DemoSchema;
export type Genre = GenreSchema;
export type Job = JobSchema;
export type Manga = MangaSchema & { slug: string };
export type Mangaka = MangakaSchema & { slug: string };
export type Magazine = MagazineSchema & { slug: string };
