import type { Magazine, Manga, Mangaka } from "@/db/types";

export const ROUTE = {
  HOME: () => "/",
  SEARCH: () => "/search",
  MANGA_LIST: () => "/manga",
  MANGA: (manga: Manga) => `/manga/${manga.id}/${manga.slug}`,
  MANGA_READ: (manga: Manga, chapterNum = 1, pageNum = 1) =>
    `/manga/${manga.id}/${manga.slug}/read?chapter=${chapterNum}&page=${pageNum}`,
  // to be replaced with cloudflare CDN
  MANGA_READ_SRC: (manga: Manga, chapterNum: number, pageNum: number) => {
    const pChapterNum = String(chapterNum).padStart(4, "0");
    const pPageNum = String(pageNum).padStart(3, "0");

    return `/imgs/${manga.id}/${pChapterNum}/${pPageNum}.jpg`;
  },
  MANGAKA_LIST: () => "/mangaka",
  MANGAKA: (mangaka: Mangaka) => `/mangaka/${mangaka.id}/${mangaka.slug}`,
  MAGAZINE_LIST: () => "/magazine",
  MAGAZINE: (magazine: Magazine) => `/magazine/${magazine.id}/${magazine.slug}`,
};
