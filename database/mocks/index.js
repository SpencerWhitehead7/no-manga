const fs = require('fs')
const path = require('path')

// https://stackoverflow.com/questions/25582882/javascript-math-random-normal-distribution-gaussian-bell-curve
// /shrug it's all greek to me lol
const randnBM = (min, max, skew) => {
  let u = 0, v = 0;
  while (u === 0) u = Math.random() //Converting [0,1) to (0,1)
  while (v === 0) v = Math.random()
  let num = Math.sqrt(-2.0 * Math.log(u)) * Math.cos(2.0 * Math.PI * v)

  num = num / 10.0 + 0.5 // Translate to 0 -> 1
  if (num > 1 || num < 0) {
    num = randnBM(min, max, skew) // resample between 0 and 1 if out of range
  } else {
    num = Math.pow(num, skew) // Skew
    num *= max - min // Stretch to fill range
    num += min // offset to min
  }
  return num
}

const generateSQL = (tableName, rows) => `INSERT INTO ${tableName}\nVALUES\n(${rows.join("),\n(")});`

const generateChapters = (minMangaID, maxMangaID) => {
  const rows = []

  for (let mID = minMangaID; mID <= maxMangaID; mID++) {
    for (let cNum = 1; cNum < randnBM(1, 1400, 1.5); cNum++) {
      const name = Math.random() > .1
        ? `'mID: ${mID}, cNum: ${cNum}'`
        : "NULL"
      const pageCount = Math.random() > .1
        ? Math.round(randnBM(17, 22, 1))
        : Math.round(randnBM(30, 60, .8))
      rows.push(`${mID}, ${cNum}, ${name}, ${pageCount}`)
    }
  }

  return generateSQL("chapter", rows)
}

const generateGenres = () => {
  const names = []

  for (let i = 0; i < 10; i++) {
    for (let j = 0; j < 10; j++) {
      names.push(`genre-${i}${j}`)
    }
  }

  const genreSQL = generateSQL("genre", names.map(x => `'${x}'`))

  return [genreSQL, names]
}

generateMangaGenres = (minMangaID, maxMangaID, genreNames) => {
  const rows = []

  for (let mID = minMangaID; mID <= maxMangaID; mID++) {
    const genresPerManga = Math.floor(Math.random() * 6);
    const usedGenreNames = new Set()
    for (let j = 0; j < genresPerManga; j++) {
      let genreName = genreNames[Math.floor(Math.random() * genreNames.length)]
      while (usedGenreNames.has(genreName)) {
        genreName = genreNames[Math.floor(Math.random() * genreNames.length)]
      }
      rows.push(`${mID}, '${genreName}'`)
      usedGenreNames.add(genreName)
    }
  }

  return generateSQL("manga_genre", rows)
}

generateMagazineManga = (minMangaID, maxMangaID, minMagazineID, maxMagazineID) => {
  const rows = []

  for (let mID = minMangaID; mID <= maxMangaID; mID++) {
    const magazinesPerManga = randnBM(1, 4, 70)
    for (let j = 0; j < magazinesPerManga; j++) {
      magazineID = Math.floor(Math.random() * (maxMagazineID - minMagazineID) + 1) + minMagazineID
      rows.push(`${magazineID}, ${mID}`)
    }
  }

  return generateSQL("magazine_manga", rows)
}

generateMangaMangakaJob = (minMangaID, maxMangaID, minMangakaID, maxMangakaID) => {
  const rows = []

  for (let mID = minMangaID; mID <= maxMangaID; mID++) {
    const mangakaPerManga = randnBM(1, 3, 90)
    for (let j = 0; j < mangakaPerManga; j++) {
      const mangakaID = Math.floor(Math.random() * (maxMangakaID - minMangakaID + 1)) + minMangakaID
      const mangakaJob = mangakaPerManga === 1
        ? "'author_artist'"
        : j === 0
          ? "'author'"
          : j === 1
            ? "'artist'"
            : "'author_artist'"
      rows.push(`${mID}, ${mangakaID}, ${mangakaJob}`)
    }
  }

  return generateSQL("manga_mangaka_job", rows)
}


// MIN: last ID in the seed file for the respective table (if using seed file, 0 otherwise) + 1
// MAX: highest ID in the respective table
generate = ({ minMangaID, maxMangaID, minMagazineID, maxMagazineID, minMangakaID, maxMangakaID }) => {
  const chapterSQL = generateChapters(minMangaID, maxMangaID)
  const [genresSQL, genreNames] = generateGenres()
  const mangaGenresSQL = generateMangaGenres(minMangaID, maxMangaID, genreNames)
  const magazineMangaSQL = generateMagazineManga(minMangaID, maxMangaID, minMagazineID, maxMagazineID)
  const mangaMangakaJobSQL = generateMangaMangakaJob(minMangaID, maxMangaID, minMangakaID, maxMangakaID)

  const outputDir = path.join(__dirname, "scripts")
  if (!fs.existsSync(outputDir)) fs.mkdirSync(outputDir)

  fs.writeFileSync(path.join(outputDir, "chapter.sql"), chapterSQL)
  fs.writeFileSync(path.join(outputDir, "genre.sql"), genresSQL)
  fs.writeFileSync(path.join(outputDir, "manga_genre.sql"), mangaGenresSQL)
  fs.writeFileSync(path.join(outputDir, "magazine_manga.sql"), magazineMangaSQL)
  fs.writeFileSync(path.join(outputDir, "manga_mangaka_job.sql"), mangaMangakaJobSQL)
}

generate({
  minMangaID: 13,
  maxMangaID: 4012,
  minMagazineID: 10,
  maxMagazineID: 1009,
  minMangakaID: 13,
  maxMangakaID: 2012,
})
