import slugify from "slugify";

export const getNameSlug = <T extends { name: string }>(v: T) => ({
  ...v,
  slug: slugify(v.name, {
    lower: true,
    locale: "eng",
  }),
});

export const getItemAndAdjacents = <T>(
  items: T[],
  findIndexCb: (v: T) => boolean,
): [T | undefined, T | undefined, T | undefined] => {
  const itemIdx = items.findIndex(findIndexCb);
  if (itemIdx === -1) return [undefined, undefined, undefined];

  return [items[itemIdx - 1], items[itemIdx], items[itemIdx + 1]];
};
