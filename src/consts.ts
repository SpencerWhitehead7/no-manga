import type { Magazine, Manga, Mangaka } from "@/db/types";

export const ROUTE = {
  HOME: () => "/",
  SEARCH: () => "/search",
  MANGA_LIST: () => "/manga",
  MANGA: (manga: Manga) => `/manga/${manga.id}/${manga.slug}`,
  MANGAKA_LIST: () => "/mangaka",
  MANGAKA: (mangaka: Mangaka) => `/mangaka/${mangaka.id}/${mangaka.slug}`,
  MAGAZINE_LIST: () => "/magazine",
  MAGAZINE: (magazine: Magazine) => `/magazine/${magazine.id}/${magazine.slug}`,
};
