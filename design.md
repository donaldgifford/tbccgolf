# Design Doc

## Intro

tbcc golf app is an app to allow tracking games to replace the google sheets used now. Mainly this is to be used for non official games since golf genius is used for those.

## Features

### Games

Games will be easier to manage since there will only be one scorecard used. Calculating net/gross etc is way simpler this way as well.

Not even sure if it's worth adding tees to the scoring.

#### Match Play

Supporting 9 and 18 hole match play games. 2 players only.

#### Stroke Play

Standard 9 and 18 hole stroke play games. 2-4 players

#### Multi group tournaments

Setup for allowing tournament structured stroke and match play games. Possibly have to support adding byes for rounds, seeding, etc.

#### Net and Gross

Gross is simpler, net will require handicaps being updated. Not sure how to get those atm, will probably be manual. Net support for 100, 90, 85 percent handicaps.

### Scoring and results

Each hole in a match is updated for each player. 1 scorekeeper can be used. Once all holes are updated the match still is in progress until a player completes the round.

#### Scoreboard

At any time a user can see the current matches, previous matches, and upcoming matches. They can click into the scoreboard to see the by hole results.

#### results

A results page showing scores, players, and possibly wagers?

#### Matchup history

Users can see match histories of players in each game types, win percentages, and wager balance between them? Maybe not public.

### Players

Simple info like name, handicap, and list results.

#### Handicaps

Will have to be manually entered for the time being.

#### Payouts

Show a users venmo username. Also show any pending transactions. IE when a payment request is send the receiver has to accept it in the app to show its paid.

### Ledger/Banker

When a match is created, the creator sets a banker as to who the money is paid to before the round. This person then is then given the payout structure and to who it goes to once the round is complete. Can be venmo or cash.

### Side Bets

Maybe build a side betting for each match or round to also be tracked.

## Development and Deployment

### Languages and tools

Tooling used to build and deploy. Local Development will be handled in docker compose.

#### Golang

Using golang with a combo of echo web server, gorm for sqlite. Easiest to deploy etc.

#### SQLite

Cloudflare offers D1 which is a hosted free sqlite. Used for storing all data.

#### HTMX

Will be used to creating the frontend with go templ.

### Cloudflare

#### Workers

What the go binary is deployed to.

#### D1

Hosted sqlite.

### Identity

#### Okta

Probably easiest to use for user management. Can also use for permissions with auth server.
