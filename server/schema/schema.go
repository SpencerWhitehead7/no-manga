package schema

const Schema string = `
scalar Time

type Manga {
  id: ID!
  name: String!
  otherNames: [String!]!
  description: String!
  demo: String!
  startDate: Time!
  endDate: Time
  genres: [String!]!
  chapterCount: Int!
  chapterList: [Chapter!]!
  mangakaList: [SeriesMangaka!]!
  magazineList: [Magazine!]!
}

type Chapter {
  id: ID!
  chapterNum: Float!
  name: String
  pageCount: Int!
  updatedAt: Time!
  manga: Manga!
}

# Mangaka + job field (job refers to their role on a specific series)
# so it only makes sense in the context of a Manga
# GQL doesn't have a native "extends" functionality, unfortunately
type SeriesMangaka {
  id: ID!
  name: String!
  otherNames: [String!]!
  description: String!
  job: String!
  mangaList: [Manga!]!
}

type Mangaka {
  id: ID!
  name: String!
  otherNames: [String!]!
  description: String!
  mangaList: [Manga!]!
}

type Magazine {
  id: ID!
  name: String!
  otherNames: [String!]!
  description: String!
  demo: String!
  mangaList: [Manga!]!
}

type Query {
  manga(ID: ID!): Manga
  mangaList: [Manga!]!

  chapter(mangaID: ID!, chapterNum: Float!): Chapter
  chapterList: [Chapter!]!

  mangaka(ID: ID!): Mangaka
  mangakaList: [Mangaka!]!

  magazine(ID: ID!): Magazine
  magazineList: [Magazine!]!
}
`
