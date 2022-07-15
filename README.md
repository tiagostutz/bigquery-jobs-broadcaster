# Big Query Job Broadcaster

Cloud Run function that receives an eventarc from BQ Job Completed event and publishes it to a pub/sub topic

## Running tests locally

Bring up the pubsub emulator by creating the following docker-compose:

```yaml
version: "3.7"

services:
  gcp-pubsub-emulator:
    image: storytel/gcp-pubsub-emulator
    command: --host=0.0.0.0 --port=8262
    ports:
      - 8262:8262
```

Bring it up

```bash
docker-compose up
```

Then run the tests:

```
go test
```
