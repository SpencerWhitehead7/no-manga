# no-manga.com

## What is this

Manga reading sites are a website subgenre I've spent a lot of time on, and frankly, most of them are slow, ugly, and bloated. I'd like to write a better one, following these guiding principles:

- Usable/simple features
- Fast/responsive UX
- Minimal aesthetic
- Prodution ready

These inform the stories:

- Discovery

  - Homepage
    - Latest updates
      - Within each page, chunked by day for the last 7 days: each page is prior 7 day stretch
    - V2 hottest chapters? Hottest series? Hotness over time? Hotness weighted by recency of said hotness?
    - V2 "staff picks"/"Community picks/polls"?
  - Search
    - For manga/mangaka/magazines by title/name, demo, genre, years
    - V2 get into that real crazy query shit
  - Exploration
    - Find manga by same mangaka, collaborators, magazine, similar demos or genres by embedding links in their pages
    - Filterable by demo and genre?

- Navigation

  - Transparent URLs
  - Get to any root route from sidebar
  - Navigate within series from sidebar (chapter to chapter)
  - Navigate within chapter from sidebar (page to page) (or maybe that should be in the page indicator)

- Reading
  - Most efficient possible use of screen real estate for page viewing experience
    - Conserve vertical space
  - Maximize reading flow
    - Optimistic pre-loading
    - Aggressive caching (still have to invalidate on page update though)
  - Easy to navigate around series/chapter
    - See navigation section

## Plans

Of course, I don't have hundreds of licensed series to fill the site with and I'm not starting a real pirate site. However, my goal is build it out so that all you'd have to do to get a fully functioning site would be fill it with content and turn on the servers. I'll add dummy data to demo it.

This repo will someday contain:

- Backend system for storing and serving data and static resources
- Frontend app
- Admin tool for managing the site's content - probably a CLI

## Stack

From back to front:

- Cloudflare D1 Sqlite database
- Cloudflare CDN serving images
- Cloudflare pages serving SSR Astro site
- Cloudflare cache doing ISG for the SSRed pages
- Webtorrent serving images peer to peer (I am real excited about this bit)

### Weird side effects

D1 is in early-ish development and is not stable. It can only be connected to from wrangler and Cloudflare workers, although hopefully Cloudflare plans to add the ability to connect programmatically from anywhere. Because of this, to do local development and build the static site, I'm keeping 3 synced copies of the database; a filesystem db, a local wrangler copy, and a remote wrangler copy. If Cloudflare ever adds the ability to connect to D1 from outside of workers, I'll be able to do local dev and builds using the remote database and delete the other two.

This would be a tremendous hassle if I was ever writing to the DB, but at this point it's read-only because I can put off writing the update tool until Cloudflare has better support. If every page was SSGed, I'd be able to ditch D1 and rely solely the filesystem DB, but because the search page needs to be SSRed, I'll eventually need a real running database.

## Logistics

It will be completed very, very, very slowly in my spare time.
