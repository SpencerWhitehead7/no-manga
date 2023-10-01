import type {
  Chapter,
  Demo,
  Genre,
  Job,
  Magazine,
  Manga,
  Mangaka,
} from "../schema";
import {
  chapter,
  demo,
  genre,
  job,
  magazine,
  magazineManga,
  manga,
  mangaGenre,
  mangaka,
  mangaMangakaJob,
} from "../schema";

export type { Chapter, Demo, Genre, Job, Magazine, Manga, Mangaka };

export {
  chapter as chapterSchema,
  demo as demoSchema,
  genre as genreSchema,
  job as jobSchema,
  magazineManga as magazineMangaSchema,
  magazine as magazineSchema,
  mangaGenre as mangaGenreSchema,
  mangaka as mangakaSchema,
  mangaMangakaJob as mangaMangakaJobSchema,
  manga as mangaSchema,
};
