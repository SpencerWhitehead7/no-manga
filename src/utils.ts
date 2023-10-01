import slugify from "slugify";

export const getNameSlug = <T extends { name: string }>(v: T) => ({
  ...v,
  slug: slugify(v.name, {
    lower: true,
    locale: "eng",
  }),
});
