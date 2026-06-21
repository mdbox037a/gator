# Gator

an RSS feed aggregator for linux CLI
-> learning project at boot.dev

## basic usage

start the aggregator process in one terminal window:
`gator agg <feed-update-interval>`

then try some commands in another terminal window:

```
# register new user
gator register <name>

# login to user account
gator login <user>

# view all registered users
gator users

# add a new RSS feed to track
gator addfeed <arbitrary-name> <http-or-https-URL>

# see all feeds
gator feeds

# follow a feed for the current logged in user
gator follow <existing-feed-URL>

# see feeds followed by current logged in user
gator followed

# stop following a feed
gator unfollow <existing-feed-url>

# see latest X posts from followed feeds (default 2)
gator browse [positive-integer]
```

## requirements

- Go 1.22 or newer
- Postgresql 14 or newer

## set up .gatorconfig.json file in user home dir

```
{
	"db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
	"current_user_name": "will-be-overwritten-on-login"
}

```
