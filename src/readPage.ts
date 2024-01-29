import { ROUTE } from "src/consts";

import type { Manga } from "@/db/types";

export const getTitle = (manga: Manga, chapterNum: number, pageNum: number) =>
  `${manga.name} chapter ${chapterNum} page ${pageNum}`;

export const getPlaceInManga = (url: string) => {
  const { searchParams } = new URL(url);

  return {
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    chapterNum: Number(searchParams.get("chapter")!),
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    pageNum: Number(searchParams.get("page")!),
  };
};

type ChapterPagination = {
  chapterNum: number;
  pageCount: number;
};

export const getNextPageInfo = (
  manga: Manga,
  pageNum: number,
  chapter: ChapterPagination,
  nextChapter?: ChapterPagination,
) => {
  const isLastChapter = nextChapter === undefined;
  const isLastPage = pageNum === chapter.pageCount;

  return {
    isLastChapter,
    isLastPage,
    nextLink:
      isLastChapter && isLastPage
        ? ROUTE.MANGA(manga)
        : ROUTE.MANGA_READ(
            manga,
            // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
            ...getNextPlaceInManga(isLastPage, nextChapter!, chapter, pageNum),
          ),
    nextSrc:
      isLastChapter && isLastPage
        ? null
        : ROUTE.MANGA_READ_SRC(
            manga,
            // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
            ...getNextPlaceInManga(isLastPage, nextChapter!, chapter, pageNum),
          ),
  };
};

export const getNextPlaceInManga = (
  isLastPage: boolean,
  nextChapter: ChapterPagination,
  chapter: ChapterPagination,
  pageNum: number,
): [number, number] => [
  isLastPage ? nextChapter.chapterNum : chapter.chapterNum,
  isLastPage ? 1 : pageNum + 1,
];

export const getPrevPageInfo = (
  manga: Manga,
  pageNum: number,
  chapter: ChapterPagination,
  prevChapter?: ChapterPagination,
) => {
  const isFirstChapter = prevChapter === undefined;
  const isFirstPage = pageNum === 1;

  return {
    isFirstChapter,
    isFirstPage,
    prevLink:
      isFirstChapter && isFirstPage
        ? ROUTE.MANGA(manga)
        : ROUTE.MANGA_READ(
            manga,
            // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
            ...getPrevPlaceInManga(isFirstPage, prevChapter!, chapter, pageNum),
          ),
    prevSrc:
      isFirstChapter && isFirstPage
        ? null
        : ROUTE.MANGA_READ_SRC(
            manga,
            // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
            ...getPrevPlaceInManga(isFirstPage, prevChapter!, chapter, pageNum),
          ),
  };
};

export const getPrevPlaceInManga = (
  isFirstPage: boolean,
  prevChapter: ChapterPagination,
  chapter: ChapterPagination,
  pageNum: number,
): [number, number] => [
  isFirstPage ? prevChapter.chapterNum : chapter.chapterNum,
  isFirstPage ? prevChapter.pageCount : pageNum - 1,
];
