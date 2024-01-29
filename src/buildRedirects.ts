import { BuildTimeRepo } from "../drizzle/db/repo";

const createAssetRedirects = (
  baseRoute: string,
  assets: { id: number; slug: string }[],
) =>
  assets.reduce<Record<string, string>>((redirects, { id, slug }) => {
    redirects[`/${baseRoute}/${id}`] = `/${baseRoute}/${id}/${slug}`;

    return redirects;
  }, {});

const createReaderRedirects = (
  baseRoute: string,
  assets: { id: number; slug: string }[],
) =>
  assets.reduce<Record<string, string>>((redirects, { id, slug }) => {
    redirects[
      `/${baseRoute}/${id}/read`
    ] = `/${baseRoute}/${id}/${slug}/read?chapter=1&page=1`;

    return redirects;
  }, {});

export const buildRedirects = () => {
  const repo = new BuildTimeRepo(import.meta.env.VITE_DB_PATH);

  return {
    ...createAssetRedirects("manga", repo.getAllMangas()),
    ...createReaderRedirects("manga", repo.getAllMangas()),
    ...createAssetRedirects("mangaka", repo.getAllMangakas()),
    ...createAssetRedirects("magazine", repo.getAllMagazines()),
  };
};
