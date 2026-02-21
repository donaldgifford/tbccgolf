# Twin Beach Country Club — Project Plan

## Hugo Static Site + Decap CMS on Cloudflare

---

## Overview

Rebuild the Twin Beach Country Club website as a Hugo static site using the "Heritage Green" design direction. Content management is handled by Decap CMS (git-based, browser UI). The site is hosted on Cloudflare Pages/Workers with a Cloudflare Worker handling the password-protected Members Only area. Source code and content live in a GitHub repository.

---

## Design Direction

**Heritage Green** — premium editorial layout with TBCC's forest green and gold palette on light cream backgrounds. Cormorant Garamond + Outfit typography. Solid green hero band (no gradient over photos), alternating green/cream feature sections, card-based content areas, deep green footer. Fully responsive.

The approved HTML mockup (`design-a-revised-heritage-green.html`) serves as the source of truth for all layout, color, spacing, and typography decisions. It translates directly into Hugo templates.

---

## Architecture

```
┌─────────────────────────────────────────────────┐
│                   GitHub Repo                    │
│                                                  │
│  content/          ← Markdown files (pages)      │
│  static/admin/     ← Decap CMS admin UI          │
│  layouts/          ← Hugo Go templates            │
│  assets/css/       ← Stylesheets (SCSS)           │
│  static/images/    ← Uploaded media               │
│  functions/        ← CF Pages Functions (OAuth)   │
│  hugo.toml         ← Site configuration           │
└──────────────┬──────────────────────────────────┘
               │ git push / Decap CMS commit
               ▼
┌──────────────────────────────────────────────────┐
│              Cloudflare Pages                     │
│                                                   │
│  Build: hugo --minify                             │
│  Output: /public (static HTML)                    │
│                                                   │
│  Pages Functions:                                 │
│    /api/auth   ← GitHub OAuth for Decap CMS       │
│    /api/callback                                  │
│                                                   │
│  Workers (middleware):                             │
│    /members/*  ← Shared password auth + cookie     │
└──────────────────────────────────────────────────┘
               │
               ▼
        twinbeachcc.com
```

---

## Tech Stack

| Component | Technology | Purpose |
|-----------|-----------|---------|
| Static site generator | Hugo | Builds HTML from Markdown + Go templates |
| CSS | SCSS via Hugo Pipes | Compiled at build time, no runtime CSS framework |
| CMS | Decap CMS | Browser-based content editing, commits to GitHub |
| CMS Auth | GitHub OAuth via CF Pages Functions | Authenticates CMS editors against GitHub |
| Members area | Cloudflare Workers (middleware) | Shared password check, signed cookie session |
| Hosting | Cloudflare Pages | Static file serving, automatic deploys from GitHub |
| DNS / SSL | Cloudflare | Domain management, automatic HTTPS |
| Analytics | Cloudflare Web Analytics | Privacy-friendly, no cookies, free |
| Source control | GitHub | Repository, triggers Cloudflare builds |
| Image optimization | Hugo image processing | Resize, crop, WebP conversion at build time |

---

## Content Architecture

All content is Markdown with YAML front matter, stored in `content/`.

```
content/
├── _index.md                  # Homepage
├── golf/
│   ├── _index.md              # Golf landing page
│   ├── tee-times.md
│   ├── pro-shop.md
│   ├── lessons.md
│   ├── scorecard.md
│   ├── course-record.md
│   ├── rates.md
│   ├── practice-range.md
│   ├── leagues.md
│   ├── junior-golf.md
│   ├── club-storage.md
│   ├── amenities.md
│   ├── course-policies.md
│   ├── outings.md
│   └── halfway-house.md
├── dining/
│   ├── _index.md              # Dining landing page
│   ├── menus.md
│   ├── bar-menu.md
│   ├── clubhouse.md
│   ├── banquets.md
│   └── social-events.md
├── membership/
│   ├── _index.md              # Membership info
│   └── join.md
├── members/                    # Password-protected
│   ├── _index.md
│   ├── announcements.md
│   └── documents.md
├── about/
│   ├── _index.md
│   ├── history.md
│   ├── staff.md
│   ├── gallery.md
│   └── employment.md
├── events/
│   └── _index.md              # Events calendar
└── contact.md
```

---

## Decap CMS Configuration

Decap CMS runs at `/admin/` and provides a browser-based editing interface. It authenticates via GitHub OAuth (not Netlify Identity, since we're on Cloudflare).

**Authentication flow:**

1. Editor navigates to `twinbeachcc.com/admin/`
2. Clicks "Login with GitHub"
3. Redirected to GitHub OAuth → Cloudflare Pages Function handles the OAuth exchange
4. On success, Decap CMS uses the GitHub API to read/write content files in the repo
5. Edits create commits (or pull requests in editorial workflow mode)
6. Cloudflare Pages auto-rebuilds on push to `main`

**Who needs GitHub access:** Only content editors (likely 1-2 club staff). They need a GitHub account with write access to the repo. This is the tradeoff of git-based CMS — editors need GitHub accounts, but in return you get version history, rollback, and no database.

**Collections defined in Decap:**

- Golf pages (folder collection)
- Dining pages (folder collection)
- Membership pages (folder collection)
- About pages (folder collection)
- Members-only content (folder collection)
- Events (folder collection)
- Site settings (file collection — for things like phone number, address, announcement bar text)

---

## Members Only Area

The members section is password-protected with a single shared password (matching current Wix behavior).

**Implementation:**

- A Cloudflare Workers middleware intercepts requests to `/members/*`
- If no valid session cookie exists, redirect to `/members/login/`
- Login page has a simple password form
- Worker validates the password against a hashed value stored in a CF environment secret
- On success, sets a signed HTTP-only cookie with a configurable TTL
- Subsequent requests to `/members/*` are allowed if the cookie is valid

This requires zero database — just a Worker and an environment variable.

---

## Deployment Pipeline

```
Editor makes change in Decap CMS
        │
        ▼
Decap commits to GitHub (main branch)
        │
        ▼
Cloudflare Pages detects push, runs build
        │
        ▼
Hugo builds site (~1-3 seconds)
        │
        ▼
Static files deployed to Cloudflare edge
        │
        ▼
Site is live (typically <60 seconds end-to-end)
```

Preview deployments are available on pull requests — useful if editorial workflow mode is enabled in Decap (edits create PRs instead of direct commits).

---

## Domain & DNS

- Move `twinbeachcc.com` DNS to Cloudflare (if not already there)
- Cloudflare Pages provides automatic SSL
- Set up `www` redirect to apex domain (or vice versa)
- Old Wix URLs should be mapped with redirect rules to avoid broken links

---

## Cost

| Service | Monthly Cost |
|---------|-------------|
| Cloudflare Pages (free tier) | $0 |
| Cloudflare Workers (free tier, 100k req/day) | $0 |
| Cloudflare Web Analytics | $0 |
| Cloudflare DNS + SSL | $0 |
| GitHub (free tier, private repo) | $0 |
| Hugo (open source) | $0 |
| Decap CMS (open source) | $0 |
| **Total** | **$0/month** |

The only potential cost is if the site exceeds Cloudflare's free tier limits (500 builds/month, 100k Worker requests/day), which is extremely unlikely for a club website.
