# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Nicole is a Niconico video ranking viewer that fetches hourly rankings from Niconico (ニコニコ動画) and displays them in a browser. The architecture consists of:

- **Backend (Go)**: Scrapes Niconico ranking data and generates JSON output
- **Frontend (Vue.js 3)**: Displays rankings with filtering capabilities
- **Deployment**: Automated via GitHub Actions to GitHub Pages (frontend) and AWS S3 (ranking data)

Live site: https://shibu1x.github.io/nicole/

## Development Commands

### Frontend

```bash
cd front
npm install          # Install dependencies
npm run dev          # Start dev server with hot reload
npm run build        # Build for production
npm run preview      # Preview production build
```

### Backend

```bash
cd back
go mod download      # Download dependencies
go run main.go       # Run scraper (outputs to dist/ranking.json)
```

## Architecture

### Backend Data Flow

1. **Scraping**: `fetchVideoLists()` scrapes https://www.nicovideo.jp/ranking/custom
   - Parses HTML using goquery
   - Extracts JSON from `<meta name="server-response">` tag
   - Uses gjson to parse ranking data from `$getCustomRankingRanking`

2. **Filtering**: Excludes "その他" (Other) category from rankings

3. **Output**: Generates both compressed (`ranking.json.gz`) and uncompressed (`ranking.json`) files

### Frontend Architecture

**Data Structure**: Rankings are represented as `Rankings = Video[][]` - a 2D array where each inner array represents a category's ranking

**State Management**:
- `rankings`: Raw ranking data from S3
- `blockedOwnerIds`: Set of blocked content creator IDs (stored in localStorage)
- `filteredRankings`: Computed property that applies `IsMuted` flag based on blocked IDs

**Key Components**:
- `App.vue`: Root component that fetches data from S3, manages block list
- `VideoList.vue`: Displays a single category column of videos
- `Slider.vue`: Toggle for blocking/unblocking creators

**Data Source**: Rankings fetched from https://shibu1x-public.s3.ap-northeast-1.amazonaws.com/test.json (gzipped JSON)

**Layout**: Desktop displays 4 category columns side-by-side (responsive grid)

### Video Filtering Logic

Videos are hidden when:
1. `IsMuted` is true (creator is blocked by user)
2. `IsPaymentRequired` is true (paid content)

User block preferences persist in localStorage under key `block_owner_ids`

### GitHub Actions Workflows

1. **Scheduled Run** (hourly at minute 1): Runs backend container, uploads ranking data to S3
2. **Build Backend**: Builds Docker image and pushes to GitHub Container Registry
3. **Build Frontend**: Builds Vue app and deploys to GitHub Pages

## Important Notes

- Backend scraping relies on Niconico's page structure - specifically the `server-response` meta tag format
- Frontend expects gzipped JSON served with `content-encoding: gzip` header
- Blocking is per-creator (`OwnerId`), not per-video
- The `IsMuted` field is computed client-side and not part of the API response
