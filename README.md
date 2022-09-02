### HTTP-Server, that does the following:

- Gets order data from the channel (Nats-streaming);
- Validates it;
- Saves it to the Database (Postgres) and Cache (Redis);
- Provides an API for getting data by id (via a web page) - from cache or, if not found, from the database.

**Used:** *Postgres, Nats-streaming, Redis, HTML, CSS, REST API principles, gin-gonic framework.*

### To run an app:

```
make build && make run
```

If you are running the application for the first time, you must apply the migrations to the database:

```
make migrate
```
