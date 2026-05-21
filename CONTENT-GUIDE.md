# Content Guide

How to add and remove news articles and job postings on the linque website.

---

## News Articles

News articles are stored in two places:

- **`posts.json`** — the index that controls which articles appear and in what order
- **`posts/`** — one `.md` file per article containing the body text

### Adding an article

**Step 1 — Create the markdown file**

Create a new file in the `posts/` folder named `YYYY-MM-DD-your-article-slug.md`.  
The slug should be lowercase with hyphens, e.g. `2026-06-01-product-launch.md`.

Write the article body in standard markdown:

```markdown
Opening paragraph here.

## Section Heading

More content here. You can use **bold**, *italic*, lists, links, and images.

![Alt text](images/your-image.jpg)
```

No frontmatter or special headers are needed — all metadata goes in `posts.json`.

**Step 2 — Add an entry to `posts.json`**

Open `posts.json` and add a new object to the array. The most recent article should go at the top:

```json
[
  {
    "slug": "2026-06-01-product-launch",
    "title": "Product Launch",
    "date": "2026-06-01",
    "image": "images/your-image.jpg",
    "excerpt": "One or two sentences summarising the article, shown on the news listing page."
  }
]
```

| Field | Required | Description |
|-------|----------|-------------|
| `slug` | Yes | Must exactly match the filename (without `.md`) |
| `title` | Yes | Displayed as the article heading |
| `date` | Yes | ISO format `YYYY-MM-DD` |
| `image` | No | Path to a hero/card image; omit or leave empty to show no image |
| `excerpt` | Yes | Short summary shown on the news listing card |

### Removing an article

Remove the corresponding object from `posts.json`. The `.md` file in `posts/` can be left in place or deleted — it will no longer be linked anywhere.

---

## Job Postings

Job postings work the same way as news articles:

- **`jobs.json`** — the index that controls which jobs appear on the Careers page
- **`jobs/`** — one `.md` file per posting containing the job description

### Adding a posting

**Step 1 — Create the markdown file**

Create a new file in the `jobs/` folder named with a descriptive slug, e.g. `jobs/senior-hardware-engineer.md`.

Write the job description in standard markdown:

```markdown
Brief company introduction paragraph.

## Tasks

- Responsibility one
- Responsibility two

## Your Skills

- Required or desired skill one
- Required or desired skill two

## What We Offer

- Benefit one
- Benefit two

Closing paragraph with call to action.
```

**Step 2 — Add an entry to `jobs.json`**

Open `jobs.json` and add a new object to the array:

```json
[
  {
    "slug": "senior-hardware-engineer",
    "title": "Senior Hardware Engineer",
    "location": "Munich, Germany",
    "type": "Full-Time",
    "date": "2026-06-01",
    "image": "images/nano.jpg",
    "excerpt": "One sentence summary shown on the Careers page listing.",
    "subject": "Senior Hardware Engineer"
  }
]
```

| Field | Required | Description |
|-------|----------|-------------|
| `slug` | Yes | Must exactly match the filename (without `.md`) |
| `title` | Yes | Displayed as the job title |
| `location` | Yes | Shown on the listing and detail page |
| `type` | Yes | Either `Full-Time` or `Part-Time` — determines which column it appears in on Careers |
| `date` | Yes | Posting date in ISO format `YYYY-MM-DD` |
| `image` | No | Path to a hero image shown on the job detail page |
| `excerpt` | Yes | Short summary shown on the Careers page listing |
| `subject` | Yes | Pre-filled subject line for the application email |

### Removing a posting

Remove the corresponding object from `jobs.json`. The `.md` file in `jobs/` can be left for archival or deleted.

---

## Images

Place any images you reference in the `images/` folder. Common formats (`.jpg`, `.jpeg`, `.png`, `.avif`, `.webp`) all work. Use descriptive filenames and keep file sizes reasonable — images displayed as heroes or cards work best at around 1200–1600 px wide.
