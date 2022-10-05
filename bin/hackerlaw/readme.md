# Hackers Law

This little Go app embeds everything you need for a webserver
into a single package/executable.

This was code yoinked from the following [website](https://hackandsla.sh/posts/2021-06-18-embed-vuejs-in-go/#updates)
## Build

First, `cd hackerlaw/frontend` and run

`yarn build`

Then into hackerlaw and run:

`go run hackerlaw.go`

Or you can build it:

`go build`


# Builds

Here is how to build these jawns.

```bash
# On Mack
GOOS=darwin GOARCH=amd64  go build -o /tmp/hackermac hackerlaw.go 

# Linux
GOOS=linux GOARCH=amd64  go build -o /tmp/hackerlin hackerlaw.go 

# Windows
GOOS=windows GOARCH=amd64  go build -o /tmp/hackerwin.exe hackerlaw.go
```