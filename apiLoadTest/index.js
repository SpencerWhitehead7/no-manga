'use strict'

const fs = require('fs')
const path = require('path')

const LoadTesting = require('easygraphql-load-tester')

const schema = fs.readFileSync(
  path.join(__dirname, '..', 'server', 'schema', 'schema.go'),
  'utf8'
).split("`")[1]

const mangaIDCount = 4012
const avgChapterCount = 180 // low enough to not incur too many misses
const mangakaIDCount = 2012
const magazineIDCount = 1009

const args = {
  manga: {
    ID: [...Array(mangaIDCount / 4).keys()].map(k => k + 1),
  },
  chapter: {
    mangaID: [...Array(mangaIDCount / 4).keys()].map(k => k + 1),
    chapterNum: [...Array(avgChapterCount / 90).keys()].map(k => k + 1),
  },
  mangaka: {
    ID: [...Array(mangakaIDCount / 4).keys()].map(k => k + 1),
  },
  magazine: {
    ID: [...Array((magazineIDCount - 1) / 4).keys()].map(k => k + 1),
  },
}

const easyGraphQLLoadTester = new LoadTesting(schema, args)

const customQueries = [
  // `
  //   BLEACH:manga(ID: 1) {
  //   id
  //   name
  //   otherNames
  //   description
  //   demo
  //   startDate
  //   endDate
  //   genres
  //   chapterList {
  //     id
  //     name
  //     pageCount
  //     updatedAt
  //     chapterNum
  //   }
  //   mangakaList {
  //     id
  //     name
  //     otherNames
  //     description
  //   }
  //   magazineList {
  //     id
  //     name
  //     otherNames
  //     description
  //     demo
  //   }
  // }
  // `,
]

const testCases = easyGraphQLLoadTester.artillery({
  // customQueries,
  // onlyCustomQueries: true,
  // queryFile: true,
  // withMutations: true,
  // selectedQueries: ["manga", "mangaList", "chapter", "chapterList", "mangaka", "mangakaList", "magazine", "magazineList"],
  selectedQueries: ["manga", "chapter", "mangaka", "magazine"],
  // selectedQueries: ["magazineList"],
})

module.exports = {
  testCases,
}