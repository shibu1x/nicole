# CLAUDE.md

Nicole is a Niconico video ranking viewer. Backend (Go) scrapes rankings hourly → S3. Frontend (Vue.js 3) displays them.

## Development Commands

```bash
docker compose up -d node                       # frontend dev server
docker compose run --rm golang go run main.go   # scraper → dist/ranking.json
docker compose run --rm terraform plan          # preview infrastructure changes
```

## Architecture

**Backend**: Scrapes https://www.nicovideo.jp/ranking/custom, extracts JSON from `<meta name="server-response">` tag via goquery/gjson (`$getCustomRankingRanking`). Outputs `ranking.json` + `ranking.json.gz`.

**Frontend**: Fetches gzipped JSON from S3 → `Rankings = Video[][]` (2D array, one inner array per category).

**Filtering** (computed client-side):
- `IsMuted = IsMutedByOwner || IsMutedByTitle || IsPaymentRequired`
- `IsMutedByOwner`: OwnerId in `block_owner_ids` (localStorage)
- `IsMutedByTitle`: title contains keyword from `block_titles` (localStorage)

**CI**: Hourly backend run → S3 upload; frontend build → GitHub Pages.
