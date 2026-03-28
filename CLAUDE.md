# CLAUDE.md

Nicole is a Niconico video ranking viewer. Backend (Go) scrapes rankings hourly → S3. Frontend (Vue.js 3) displays them.

Live site: https://shibu1x.github.io/nicole/

## Development Commands

```bash
cd front && npm run dev   # frontend dev server
cd back && go run main.go # scraper → dist/ranking.json
```

## Architecture

**Backend**: Scrapes https://www.nicovideo.jp/ranking/custom, extracts JSON from `<meta name="server-response">` tag via goquery/gjson (`$getCustomRankingRanking`). Excludes "その他" category. Outputs `ranking.json` + `ranking.json.gz`.

**Frontend data flow**: Fetches gzipped JSON from https://shibu1x-public.s3.ap-northeast-1.amazonaws.com/test.json → `Rankings = Video[][]` (2D array, one inner array per category).

**Filtering** (computed client-side, not in API):
- `IsMuted = IsMutedByOwner || IsMutedByTitle || IsPaymentRequired`
- `IsMutedByOwner`: OwnerId in `block_owner_ids` (localStorage)
- `IsMutedByTitle`: title contains keyword from `block_titles` (localStorage)

**GitHub Actions**: Hourly backend run → S3 upload; frontend build → GitHub Pages.
