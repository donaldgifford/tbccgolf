# Twin Beach Country Club — Implementation Plan

## Hugo + Decap CMS + Cloudflare Pages

---

## Phase 1: Repository & Build Pipeline

**Goal:** Working Hugo site building and deploying to Cloudflare Pages from GitHub on every push.

### Tasks

- [ ] Create GitHub repository (`twinbeachcc-site`, private)
- [ ] Initialize Hugo project (`hugo new site . --format yaml`)
- [ ] Set Hugo version — pin in `netlify.toml` or `hugo.toml` (use latest stable, currently v0.142+)
- [ ] Create base directory structure:

  ```
  ├── archetypes/
  ├── assets/
  │   └── css/
  │       ├── main.scss
  │       └── _variables.scss
  ├── content/
  ├── data/
  ├── layouts/
  │   ├── _default/
  │   │   ├── baseof.html
  │   │   ├── list.html
  │   │   └── single.html
  │   ├── partials/
  │   │   ├── head.html
  │   │   ├── header.html
  │   │   ├── footer.html
  │   │   └── meta.html
  │   └── index.html
  ├── static/
  │   ├── fonts/
  │   └── images/
  └── hugo.toml
  ```

- [ ] Create `hugo.toml` with base configuration (baseURL, title, language, permalink structure, build settings, markup settings)
- [ ] Create minimal `layouts/_default/baseof.html` shell that loads — even if blank — to confirm Hugo builds
- [ ] Connect repo to Cloudflare Pages:
  - [ ] Log in to Cloudflare dashboard → Pages → Create a project → Connect to Git
  - [ ] Select the repo, set build command: `hugo --minify`
  - [ ] Set build output directory: `public`
  - [ ] Set environment variable `HUGO_VERSION` to pinned version
- [ ] Confirm Cloudflare Pages builds successfully on push to `main`
- [ ] Configure custom domain (`twinbeachcc.com`) in Cloudflare Pages
  - [ ] Add CNAME record pointing to `<project>.pages.dev`
  - [ ] Verify SSL provisioned automatically
- [ ] Set up `www` → apex redirect (or vice versa) via Cloudflare redirect rules
- [ ] Add `.gitignore` (public/, resources/, .hugo_build.lock, node_modules/)
- [ ] Confirm preview deploys work on pull requests (Cloudflare generates preview URLs automatically)

### Success Criteria

> A push to `main` triggers a Cloudflare Pages build, Hugo compiles without errors, and the site is accessible at `twinbeachcc.com` with valid HTTPS — even if it's just a blank page with a title.

---

## Phase 2: Design System & Layout Templates

**Goal:** Translate the Heritage Green HTML mockup into a reusable Hugo template system with the full design system (colors, typography, spacing, components) defined in SCSS.

### Tasks

- [ ] Extract design tokens from `design-a-revised-heritage-green.html` into `assets/css/_variables.scss`:
  - [ ] Colors: `$green-primary: #234d35`, `$green-deep: #1a3a2a`, `$green-mid: #2d6341`, `$gold: #c4a55a`, `$gold-light: #d4b96a`, `$cream: #faf8f4`, `$cream-dark: #f0ece4`, `$text-dark: #1a1a1a`
  - [ ] Typography: Cormorant Garamond (headings), Outfit (body), font sizes, letter-spacing, line-heights
  - [ ] Spacing scale (section padding, container max-widths, gutters)
  - [ ] Breakpoints (mobile, tablet, desktop)
- [ ] Build `assets/css/main.scss` with Hugo Pipes (SCSS → CSS, minified, fingerprinted)
- [ ] Create `layouts/_default/baseof.html` — the master wrapper:
  - [ ] DOCTYPE, html lang, head partial, body, header partial, `{{ block "main" . }}`, footer partial
  - [ ] Google Fonts link tags (Cormorant Garamond, Outfit)
  - [ ] Hugo Pipes CSS asset pipeline
  - [ ] Favicons and meta tags
- [ ] Create `layouts/partials/header.html` — top nav:
  - [ ] Logo/club name
  - [ ] Navigation links from `hugo.toml` menu config
  - [ ] Mobile hamburger menu with JS toggle
  - [ ] Translucent-on-cream scroll behavior (optional — can defer)
- [ ] Create `layouts/partials/footer.html`:
  - [ ] Deep green background
  - [ ] Club address, phone, hours
  - [ ] Navigation links
  - [ ] Copyright line
  - [ ] Cloudflare Web Analytics snippet
- [ ] Create `layouts/index.html` — homepage:
  - [ ] Hero section: large course photo + solid green band below with headline/subtitle
  - [ ] "The Course" feature section (green background, two-column image + text)
  - [ ] "The Clubhouse" feature section (cream background, reversed layout)
  - [ ] Three-card grid (Lessons, Leagues, Banquets — images with green overlay + text)
  - [ ] Membership CTA section (photo background with green overlay)
  - [ ] Announcement bar (if enabled via site params)
- [ ] Create `layouts/_default/single.html` — generic inner page template:
  - [ ] Page hero (smaller, section-colored header with title and optional subtitle)
  - [ ] Content body with `{{ .Content }}` rendered from Markdown
  - [ ] Sidebar or sub-navigation (if applicable to section)
- [ ] Create `layouts/_default/list.html` — section index template:
  - [ ] Section hero with title
  - [ ] List of child pages as cards or linked items
- [ ] Create section-specific layouts as needed:
  - [ ] `layouts/golf/list.html` — golf section landing
  - [ ] `layouts/dining/list.html` — dining section landing
  - [ ] `layouts/about/list.html` — about section landing
- [ ] Configure navigation menus in `hugo.toml`:

  ```yaml
  [menus]
    [[menus.main]]
      name = "Golf"
      pageRef = "/golf"
      weight = 10
    [[menus.main]]
      name = "Dining"
      pageRef = "/dining"
      weight = 20
    # ... etc
  ```

- [ ] Implement responsive styles for all templates (mobile-first):
  - [ ] Navigation collapses to hamburger at mobile breakpoint
  - [ ] Hero and feature sections stack vertically on mobile
  - [ ] Card grid goes from 3-column → 2-column → 1-column
  - [ ] Typography scales down appropriately
  - [ ] Test on 375px, 768px, 1024px, 1440px widths
- [ ] Set up Hugo image processing pipeline:
  - [ ] Store source images in `assets/images/` (not `static/`) for Hugo Pipes processing
  - [ ] Create image partial that generates responsive `srcset` with WebP + JPEG fallback
  - [ ] Configure quality settings in `hugo.toml` (`[imaging]` section)
- [ ] Verify full homepage renders correctly in browser matching the mockup

### Success Criteria

> The homepage at `twinbeachcc.com` visually matches the Heritage Green mockup — correct colors, typography, layout, spacing, and responsive behavior at all breakpoints. Inner pages have a consistent look using the shared base template. All styles come from SCSS compiled through Hugo Pipes (no CDN CSS frameworks).

---

## Phase 3: Content Migration

**Goal:** All existing TBCC website content ported into Hugo Markdown files with proper front matter, organized in the content directory structure.

### Tasks

- [ ] Audit current Wix site — inventory every page and its content:
  - [ ] Golf section (~14 pages): Tee Times, Pro Shop, Lessons, Scorecard, Course Record, Rates, Practice Range, Leagues, Junior Golf, Club Storage, Amenities, Course Policies, Outings, Halfway House
  - [ ] Dining section (~5 pages): Menus, Bar Menu, Clubhouse, Banquets, Social Events
  - [ ] Membership section (~2 pages): Info, Join/Apply
  - [ ] About section (~4 pages): History, Staff, Gallery, Employment
  - [ ] Events/Calendar
  - [ ] Contact page
  - [ ] Members Only section (protected pages)
- [ ] Create Markdown files for each page with YAML front matter:

  ```yaml
  ---
  title: "Tee Times"
  description: "Book your tee time at Twin Beach Country Club"
  section: "golf"
  menu:
    golf:
      weight: 10
  hero_image: "/images/golf/tee-times-hero.jpg"
  ---
  ```

- [ ] Migrate text content from Wix — rewrite or clean up as needed
- [ ] Download and organize photography from current Wix site:
  - [ ] Save to `assets/images/` organized by section (golf/, dining/, about/, etc.)
  - [ ] Rename files with descriptive slugs (`fairway-hole-7.jpg`, not `IMG_4032.jpg`)
  - [ ] Note any images that are low quality and need replacement
- [ ] Replace all stock photos in templates with actual TBCC photography
- [ ] Create `data/site.yaml` for global site data:

  ```yaml
  club_name: "Twin Beach Country Club"
  phone: "(248) 683-2400"
  address: "2600 Cass Lake Rd, Keego Harbor, MI 48320"
  hours:
    clubhouse: "11am - 9pm"
    pro_shop: "7am - 7pm"
  announcement:
    enabled: false
    text: ""
  ```

- [ ] Create `content/contact.md` with embedded contact form (handled by CF Worker, Phase 5)
- [ ] Verify all pages render correctly with migrated content
- [ ] Check for broken internal links between pages
- [ ] Validate that no Wix-specific markup or embedded code carried over

### Success Criteria

> Every page from the current Wix site has an equivalent Hugo page with the same (or improved) content. All navigation links work. All section landing pages show their child pages. Photography is organized and displayed correctly.

---

## Phase 4: Decap CMS Integration

**Goal:** Content editors can log in at `/admin/`, edit any page, upload images, and publish changes — all from a browser UI.

### Tasks

- [ ] Create `static/admin/index.html`:

  ```html
  <!DOCTYPE html>
  <html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>TBCC Admin</title>
  </head>
  <body>
    <script src="https://unpkg.com/decap-cms@^3.0.0/dist/decap-cms.js"></script>
  </body>
  </html>
  ```

- [ ] Create `static/admin/config.yml` with all collections:
  - [ ] Backend: `name: github`, `repo: <org>/twinbeachcc-site`, `branch: main`
  - [ ] `base_url` set to site URL (for OAuth callback)
  - [ ] `media_folder: "static/images/uploads"`
  - [ ] `public_folder: "/images/uploads"`
  - [ ] Collections for each content section (golf, dining, membership, about, members, events)
  - [ ] File collection for site settings (`data/site.yaml`)
  - [ ] Define fields per collection: title, description, hero_image (image widget), body (markdown widget), plus any section-specific fields
- [ ] Set up GitHub OAuth for Decap CMS on Cloudflare Pages:
  - [ ] Register a GitHub OAuth App (Settings → Developer Settings → OAuth Apps):
    - Application name: `TBCC CMS`
    - Homepage URL: `https://twinbeachcc.com`
    - Callback URL: `https://twinbeachcc.com/api/callback`
  - [ ] Copy Client ID and Client Secret
  - [ ] Add as Cloudflare Pages environment variables:
    - `GITHUB_CLIENT_ID`
    - `GITHUB_CLIENT_SECRET`
  - [ ] Create Cloudflare Pages Functions for OAuth:
    - `functions/api/auth.js` — redirects to GitHub authorization URL
    - `functions/api/callback.js` — exchanges auth code for token, sends to Decap via `postMessage`
  - [ ] Reference: [SubhenduX/decap-cms-cloudflare-pages](https://github.com/SubhenduX/decap-cms-cloudflare-pages) or [sterlingwes/decap-proxy](https://github.com/sterlingwes/decap-proxy)
- [ ] Create GitHub account(s) for TBCC staff editors and grant repo write access
- [ ] Test the full CMS flow:
  - [ ] Navigate to `twinbeachcc.com/admin/`
  - [ ] Click "Login with GitHub" — OAuth popup, authorize, return
  - [ ] Browse collections, open a page, edit content
  - [ ] Upload an image via the media library
  - [ ] Save / publish the change
  - [ ] Confirm the commit appears in GitHub
  - [ ] Confirm Cloudflare Pages auto-builds and the change is live
- [ ] (Optional) Enable editorial workflow (`publish_mode: editorial_workflow`):
  - [ ] Edits create pull requests instead of direct commits
  - [ ] Preview deploy links show the change before merging
  - [ ] Decide if this adds value or just friction for the staff editors
- [ ] Configure Decap CMS local backend for development:
  - [ ] Install `decap-server` for local proxy: `npx decap-server`
  - [ ] Set `local_backend: true` in config (dev only, not production)
  - [ ] Verify CMS works at `localhost:1313/admin/` during development
- [ ] Write a short CMS user guide for staff (1-2 pages): how to log in, edit a page, upload an image, publish

### Success Criteria

> A non-technical staff member can navigate to `twinbeachcc.com/admin/`, log in with GitHub, edit any content page (text + images), publish the change, and see it live on the site within 2 minutes — without touching a code editor or terminal.

---

## Phase 5: Members Area & Contact Form

**Goal:** Password-protected `/members/` section and a working contact form — both powered by Cloudflare Workers.

### Tasks

#### Members Area

- [ ] Create a Cloudflare Pages Function or Worker middleware for `/members/*` routes:
  - [ ] Check for a signed session cookie on each request
  - [ ] If no valid cookie → redirect to `/members/login/`
  - [ ] If valid cookie → pass request through to static content
- [ ] Create the members login page (`content/members/login.md` or a dedicated layout):
  - [ ] Simple form: password input + submit button
  - [ ] Styled consistently with the rest of the site
  - [ ] Form POSTs to a Worker endpoint (e.g., `/api/members/auth`)
- [ ] Create the auth Worker endpoint (`functions/api/members/auth.js`):
  - [ ] Receives the submitted password
  - [ ] Compares against a bcrypt hash stored in CF environment secret (`MEMBERS_PASSWORD_HASH`)
  - [ ] On match: sets a signed, HTTP-only cookie with configurable TTL (e.g., 7 days)
  - [ ] On failure: redirects back to login with error message
- [ ] Create members-only content pages:
  - [ ] `content/members/_index.md` — members landing/dashboard
  - [ ] `content/members/announcements.md` — club announcements
  - [ ] `content/members/documents.md` — downloadable files (bylaws, schedules, etc.)
- [ ] Store member documents in `static/members-files/` (also protected by the middleware)
- [ ] Add a "Members" link in the main nav that goes to `/members/`
- [ ] Add a logout mechanism (clear cookie, redirect to homepage)
- [ ] Test the flow: visit `/members/`, get redirected to login, enter password, access content, close browser, reopen — cookie should persist

#### Contact Form

- [ ] Create contact form layout in `layouts/contact/single.html` or as a partial:
  - [ ] Fields: name, email, phone (optional), subject, message
  - [ ] Styled with Heritage Green design system
  - [ ] Client-side validation (required fields, email format)
- [ ] Create form handler Worker (`functions/api/contact.js`):
  - [ ] Receives form POST data
  - [ ] Basic validation and honeypot spam check
  - [ ] Sends email via Cloudflare Email Workers, Resend API, or Mailgun
  - [ ] Returns success/error response
  - [ ] Rate limiting (optional — prevent spam floods)
- [ ] Configure email delivery service:
  - [ ] Store API key as CF environment secret
  - [ ] Set recipient address(es) as environment variable
- [ ] Add success/error feedback UI on the contact page
- [ ] Test end-to-end: submit form → receive email at club's address

### Success Criteria

> Visiting `/members/announcements/` without logging in redirects to a login page. Entering the correct password grants access for the session. The contact form submits successfully and an email arrives at the configured club email address. Both features work on mobile.

---

## Phase 6: SEO, Performance & Analytics

**Goal:** Site is optimized for search engines, scores 90+ on Lighthouse across all categories, and analytics are tracking.

### Tasks

#### SEO

- [ ] Add `<title>` and `<meta name="description">` to every page via front matter + head partial
- [ ] Create `layouts/partials/meta.html` with Open Graph and Twitter Card meta tags:
  - [ ] `og:title`, `og:description`, `og:image`, `og:url`, `og:type`
  - [ ] `twitter:card`, `twitter:title`, `twitter:description`, `twitter:image`
- [ ] Generate `sitemap.xml` (Hugo does this automatically — verify it's correct)
- [ ] Generate `robots.txt` (Hugo built-in — configure in `hugo.toml`)
- [ ] Add structured data (JSON-LD) for the organization:
  - [ ] `LocalBusiness` schema with name, address, phone, hours, geo coordinates
  - [ ] `GolfCourse` schema if applicable
- [ ] Create a 404 page (`layouts/404.html`) styled with the design system
- [ ] Set up Wix → Hugo URL redirect map:
  - [ ] Audit all indexed Wix URLs (use Google Search Console or `site:twinbeachcc.com`)
  - [ ] Create redirect rules in Cloudflare (bulk redirects or `_redirects` file)
  - [ ] Ensure no 404s for previously indexed pages

#### Performance

- [ ] Verify Hugo image processing is generating responsive images with `srcset`
- [ ] Add `loading="lazy"` to below-fold images
- [ ] Verify SCSS is minified and fingerprinted (cache-busting)
- [ ] Confirm no render-blocking resources (Google Fonts can be `font-display: swap`)
- [ ] Test with Lighthouse — target 90+ on all four categories (Performance, Accessibility, Best Practices, SEO)
- [ ] Test with WebPageTest for real-world load times
- [ ] Verify Cloudflare caching headers are correct (static assets cached at edge)

#### Analytics

- [ ] Add Cloudflare Web Analytics to `layouts/partials/footer.html`
- [ ] Verify analytics data appears in Cloudflare dashboard
- [ ] (Optional) Add Google Search Console — submit sitemap

### Success Criteria

> Lighthouse scores 90+ across all categories. Every page has unique title/description and Open Graph tags. Sitemap is submitted. No broken links from old Wix URLs. Analytics are recording page views.

---

## Phase 7: Launch & DNS Cutover

**Goal:** Live site on `twinbeachcc.com` serving from Cloudflare Pages, Wix subscription cancelled.

### Tasks

- [ ] Final content review — walk through every page with club stakeholders
- [ ] Final responsive testing — verify on real iOS and Android devices
- [ ] Cross-browser testing — Chrome, Safari, Firefox, Edge
- [ ] Verify Decap CMS works in production (`/admin/` login, edit, publish)
- [ ] Verify members area works in production (login, access, logout)
- [ ] Verify contact form delivers emails
- [ ] DNS cutover:
  - [ ] If DNS is not already on Cloudflare: migrate nameservers
  - [ ] Update A/CNAME records to point to Cloudflare Pages
  - [ ] Remove any Wix DNS records
  - [ ] Verify SSL certificate provisions correctly (may take up to 24 hours)
- [ ] Test site on `twinbeachcc.com` after DNS propagation
- [ ] Set up Cloudflare redirect rules for `www` → apex (or vice versa)
- [ ] Monitor for errors in Cloudflare analytics for first 48 hours
- [ ] Confirm old Wix URLs redirect properly
- [ ] Train staff on CMS usage (walkthrough session or recorded video)
- [ ] Document the password/credentials:
  - [ ] GitHub repo access
  - [ ] Cloudflare account access
  - [ ] Members area password (and how to change it)
  - [ ] GitHub OAuth app credentials
- [ ] Cancel Wix subscription (after confirming everything works and DNS has fully propagated)

### Success Criteria

> `twinbeachcc.com` serves the new Hugo site. All pages load correctly. CMS editing works. Members area works. Contact form works. No 404s from old URLs. Wix is fully decommissioned. Staff can independently manage content.

---

## Open Questions

These need answers before or during implementation. Decisions here may affect the tasks above.

### Content & Design

1. **Exact brand colors** — The green (`#234d35`) was inferred from the TBCC logo. Does the club have official brand guidelines with exact hex values? If not, the inferred values are close enough to proceed.

2. **Photography assets** — Does the club have a library of high-resolution course, clubhouse, and event photos? The stock photos need to be replaced, and photo quality will make or break the design. If not, is a photography session in scope?

3. **Announcement bar** — The current Wix site has seasonal announcements. Should the new site have a dismissible announcement banner at the top? If so, should it be editable via Decap CMS?

4. **Events/calendar** — The current site has an events section. Is this a static list of upcoming events maintained via CMS, or does the club use an external calendar service (Google Calendar embed, etc.) that needs to be integrated?

5. **Online tee times** — Does the club use a third-party tee time booking system (ForeUP, GolfNow, etc.)? If so, we need to know the embed/link details. If it's just a phone number, no integration needed.

### Technical — Decap CMS

1. **Number of CMS editors** — How many staff members need CMS access? Each needs a GitHub account. If the answer is more than 2-3 people, we should evaluate whether Sveltia CMS (a Decap-compatible drop-in with better UX) would be a better fit, or whether editorial workflow mode is needed to prevent conflicts.

2. **Editorial workflow vs. direct publish** — Decap supports two modes: (a) edits commit directly to `main` and go live immediately, or (b) edits create a pull request that can be reviewed before publishing. For a small club site with 1-2 editors, direct publish is simpler. Do stakeholders want a review step?

3. **Content that lives outside Markdown** — Are there any pages on the current site that contain embedded widgets, iframes, or third-party integrations beyond tee times? (e.g., weather widgets, social media feeds, payment forms). These need special handling in Hugo templates.

### Technical — Infrastructure

1. **Current DNS provider** — Where is `twinbeachcc.com` registered and where are nameservers currently pointed? If Wix manages DNS, we need to transfer it to Cloudflare before or during launch. If a third-party registrar manages it, we just update records.

2. **Members area content scope** — What specifically lives behind the members-only wall today? Just announcements and documents, or are there additional features (forums, directories, payment portals)? This determines how much the Workers middleware needs to protect.

3. **Members area password rotation** — How often does the shared password change? Should we build a simple admin UI for changing it, or is updating a Cloudflare environment variable (with a redeploy) acceptable?

4. **Contact form recipients** — Which email address(es) should receive contact form submissions? Is there a preference for email delivery service (Resend, Mailgun, or Cloudflare Email Routing)?

5. **Email service for contact form** — Cloudflare Workers can't send email natively. Options are: (a) Resend (free tier: 100 emails/day), (b) Mailgun (free tier: 1,000/month first 3 months then paid), (c) Cloudflare Email Routing + Workers (free but more complex setup), or (d) a simple mailto link instead of a form. Which is preferred?

### Process

1. **Stakeholder review process** — Who needs to sign off on the design and content before launch? Is there a formal approval process, or is it the decision of one person?

2. **Timeline pressure** — Is there a hard deadline (e.g., start of golf season, a scheduled event) driving the launch date? The phases above are scoped for roughly 4-6 weeks of part-time effort.

3. **Ongoing maintenance** — After launch, who handles: (a) content updates (staff via CMS), (b) design/template changes (developer), (c) domain/hosting/Cloudflare account (who owns the account)? This should be documented clearly.
