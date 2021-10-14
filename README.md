# Dockwork GraphQL API

## Configuration

All configuration for this application is done through the following environment variables:

| Name                   | Description |
| ---------------------- | ----------- |
| `ADDRESS`              | (default: `""` all IPs) The IP Address for the server to listen on |
| `PORT`                 | (default: `5000`) The port number for the server to listen on |
| `DATABASE_URL`         | (default: `mongodb://localhost`) The URL to use to connect to MongoDB |
| `DATABASE_NAME`        | (default: `dockwork`) The default database name to use (when not specified in the URL) |
| `GRAPHQL_PLAY_ENABLED` | (default: `true`) Enable or disable the GraphQL Playground endpoint |

## Development

Development can be done by using Docker Compose. The application and MongoDB instance will be started and their containers
will be linked to allow for quickly getting started with development of this service.

### Start Development Server (Docker Compose)

To start the development server run:

```bash
docker compose up
```

You should then be able to access the site by going to [http://localhost:5000](http://localhost:5000).

**Note:** It may take a minute to initialize the MongoDB instance before a request can be properly processed by the API.

**Note:** The MongoDB instance is not exposed by default. To expose it you will need to add the port binding in the `docker-compose.yml` file.

### Stop Development Server (Docker Compose)

To stop the development server run:

```bash
docker compose down
```

## Deployment

This service is designed to be deployed in a Docker environment and has a Dockerfile to build the image.

To build the image run:

```bash
docker build -t nmhackathon.azurecr.io/dockwork .
```

Then, push and deploy to your favorite Docker orchestration platform!
