# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Twin Beach Country Club (TBCC) website — a Hugo static site with Decap CMS, hosted on Cloudflare Pages. Replaces an existing Wix site. The project is currently in pre-development (planning docs and a design mockup exist, no Hugo scaffolding yet).

## Build & Dev Commands

```bash
# Local development server
hugo server

# Production build
hugo --minify

# Create new content page
hugo new golf/new-page.md

# Local CMS development (when Decap CMS is set up)
npx decap-server
```

Build output goes to `public/`. Cloudflare Pages builds automatically on push to `main` with `hugo --minify`.

## Architecture

```
GitHub repo → push/Decap CMS commit → Cloudflare Pages build (hugo --minify) → static site at twinbeachcc.com
```

**Key technology choices:**
- **Hugo** — static site generator with Go templates in `layouts/`
- **SCSS via Hugo Pipes** — no npm/bundler, Hugo compiles CSS at build time
- **Decap CMS** at `/admin/` — browser-based git CMS, authenticates via GitHub OAuth
- **Cloudflare Pages Functions** (`functions/`) — OAuth handlers, contact form, members auth
- **No JS framework** — vanilla JavaScript only

**Content flow:** Markdown files with YAML front matter in `content/` → Hugo templates in `layouts/` → static HTML in `public/`

## Project Structure (Planned)

```
content/           # Markdown pages (golf/, dining/, membership/, members/, about/, events/, contact)
layouts/           # Hugo Go templates (_default/, partials/, section overrides)
assets/css/        # SCSS source (_variables.scss, main.scss) compiled via Hugo Pipes
static/admin/      # Decap CMS (index.html + config.yml)
static/images/     # Uploaded media
functions/api/     # Cloudflare Pages Functions (auth.js, callback.js, contact.js, members/auth.js)
data/site.yaml     # Global site data (phone, address, hours)
hugo.toml          # Hugo configuration
docs/plan/         # Project planning documents
docs/impl/         # Implementation roadmap
docs/design/       # HTML design mockup (index.html) — visual source of truth for the Hugo site
```

## Design System — Heritage Green

The design mockup at `docs/design/index.html` is the visual source of truth for the Hugo site. Key tokens:

- **Colors:** Deep green `#1a3a2a`, primary green `#234d35`, gold `#c4a55a`, cream `#faf8f4`
- **Typography:** Cormorant Garamond (headings), Outfit (body) — loaded from Google Fonts
- **Layout:** Mobile-first responsive. Breakpoints at 768px and 1024px
- **Pattern:** Alternating green/cream sections, card grids, solid green hero bands (no gradient over photos)

## Members Area

Password-protected `/members/*` routes via Cloudflare Workers middleware. Single shared password (bcrypt hash in CF environment variable). Session managed with signed HTTP-only cookies.

## Implementation Phases

The project follows a 7-phase plan detailed in `docs/impl/mvp.md`:
1. Repository & build pipeline
2. Design system & layout templates
3. Content migration from Wix
4. Decap CMS integration with GitHub OAuth
5. Members area & contact form (Cloudflare Workers)
6. SEO, performance & analytics
7. Launch & DNS cutover

Architecture decisions and open questions are documented in `docs/plan/mvp.md`.
