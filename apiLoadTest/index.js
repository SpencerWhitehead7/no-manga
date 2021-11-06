'use strict'

const fs = require('fs')
const path = require('path')

const LoadTesting = require('easygraphql-load-tester')

const schema = fs.readFileSync(
  path.join(__dirname, '..', 'server', 'schema', 'schema.go'),
  'utf8'
).split("`")[1]

const args = {
  manga: {
    ID: [...Array(24).keys()].map(k => k + 1),
  },
  chapter: {
    mangaID: [...Array(24).keys()].map(k => k + 1),
    chapterNum: [...Array(6).keys()].map(k => k + 1),
  },
  mangaka: {
    ID: [...Array(24).keys()].map(k => k + 1),
  },
  magazine: {
    ID: [...Array(18).keys()].map(k => k + 1),
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
})

module.exports = {
  testCases,
}