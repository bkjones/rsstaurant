rsstaurant - an RSS aggregator back end in golang

This was just a learning exercise. It's not intended to be run by anyone in any environment for any reason :)

The database back end is postgres. I used sqlc to handle db schema migrations, and 'goose' to generate golang client
code for queries.

To start the daemon:
`go build && ./rsstaurant`

This does the job more or less right now, but as it was done in a rush, there are a few outstanding issues:
 1. Some feeds throw XML errors & I haven't bothered to look into that.
 2. I parse timestamps using only one hard-coded method, so if the feed doesn't happen to use that format, it fails.
 3. I didn't implement authentication on a few endpoints that should have it if this were going to be used in any kind of
    "production" way: /users, and /feeds in particular.

At some point, it'd be good to add a front end for this using React/Typescript. Here's a quick non-exhaustive api doc.
See main.go for a list of all of the endpoints.

```
### Get a list of users
GET http://localhost:8080/users

### Get all posts for all followed feeds for auth'd user
GET http://localhost:8080/v1/posts
Accept: application/json
Authorization: ApiKey <apikey>

### show all feeds. This is not auth'd right now
GET http://localhost:8080/v1/feeds

### add a new feed subscription for the auth'd user
POST http://localhost:8080/v1/feed_follows
Authorization: ApiKey <apikey>

{
  "feed_id": "480d1d87-9a34-4c7c-827c-b70d3ac65efd"
}

### Get a list of feeds being followed by the auth'd user
GET http://localhost:8080/v1/feed_follows
Authorization: ApiKey <apikey>
```
