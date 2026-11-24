# gator

gator is a simple to use CLI RSS feed aggregator.

## Requirements

* Postgres
* Go (version 1.27 and up)

## Installation

```bash
go install github.com/william-jce/gator@latest
```

Once you've installed the package, create a JSON file in your home directory: `~/.gatorconfig.json`.
In the JSON file, set your postgres database URL: `{"db_url":"postgres://..."}`

## Getting Started

1. Register your desired username: `gator register <name>`
2. Add whatever feed you'd like, for example: `gator addfeed "Tech Crunch" <url>`

## Managing Feeds

- Follow an existing feed: `gator follow <url>`
- Unfollow a feed: `gator unfollow <url>`
- List all feeds: `gator feeds`

## Viewing Posts

- Browse posts: `gator browse [limit]` (defaults to 2)

## Aggregating

- Start the aggregator: `gator agg <time_between_requests>` (example: `gator agg 1m`)
This runs the aggregator continuously until cancelled with ctrl+c.

## Bookmarks

Posts in `browse` are shown with a short ID, e.g. `[3f9a2c1b] Post title`.
Use that ID (or any unique prefix of it) to bookmark posts:

- Bookmark a post: `gator bookmark <post id>`
- Remove a bookmark: `gator unbookmark <post id>`
- List your bookmarks: `gator bookmarks`

Bookmarks are kept even if you unfollow the post's feed.

## User Management

- Log in as an existing user: `gator login <name>`
- List all users: `gator users`
