---
permalink: /:path/:basename:output_ext
title: 0002-Matches
layout: adr
---

# Matches

| Status | Author         | Date       |
| ------ | -------------- | ---------- |
| Draft  | @donaldgifford | 2024-06-13 |

## Context and Problem Statement

Matches is a feature for tbccgolf to allow players to create, track, and manage matches between themselves.

Main components of matches:

- players can create a match, and add users.
- matches show up on matches screen of running, open, completed.
- home page shows current matches.

## Tech choices

### Links

- [SSE Echo](https://echo.labstack.com/docs/cookbook/sse)
- [SSE blogs](https://threedots.tech/post/live-website-updates-go-sse-htmx/)
- [SSE Mozilla](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#event_stream_format)

### Background

Using new Web API for Server Sent Events we dont have to reinvent anything crazy to manage them. This way we can create event streams and listeners to pages for views like match status, etc.

### Flow

- Match created
- Scores can be updated for each hole.
- Each match creates a new stream.
- The stream updates are written to the match DB.
- Any viewers pulls from a view stream created for each match? Or maybe a main match screen gets updated for it.

```
                                                                  ┌──────────────────────────┐
                                   ┌──────────────────────┐       │         Results:         │
                            ┌──────┤  DB record created   ├───────┤  Pulls records from DB   │
                            │      └──────────────────────┘       └──────────────────────────┘
                            │
┌──────────────────────┐    │
│    Match Created     ├────┤                                     ┌───────────────────────────┐
└──────────────────────┘    │                                     │         Screens:          │
                            │      ┌──────────────────────┐       │HomePageMatch Page updates │
                            └──────┤  Match event stream  ├───────┤  Just listens for event   │
                                   │       created        │       │  updates to show scores.  │
                                   └───────────┬──────────┘       └───────────────────────────┘
                                               │
         Saving score submits                  │                 Event Types:Update - writes
         the event update        ┌─────────────┴──────────┐      to the DB, in case app poops
                                 │                        │      out and needs to be resumed.
                                 │                        │
                       ┌─────────┴───────────┐ ┌──────────┴──────────┐
                       │    Player Screen    │ │    Player Screen    │
                       └─────────────────────┘ └─────────────────────┘
```

### Notes

#### DB Notes

Need a match DB.

```go
type Match struct {
  ID int
  NetScore bool // if match is net or gross
  Players []Player
  Length int // 9 or 18
  StartingHole int
  StartDate time.Date
  FinishDate time.Date
  Completed bool
  Scores []Score
}
```

Scores DB:

```go
type Score struct {
  ID int
  Length int
  PlayerID Player
  MatchID Match
  RawScore int
  ScoreByHoles []ScoreByHole // Json blob of scores
}
```

```go
type ScoreByHole struct {
  HoleID int
  Score int
}
```
