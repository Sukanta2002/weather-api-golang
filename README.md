# Weather API

A high-performance weather API built with Go that fetches weather data from Visual Crossing Weather API and implements Redis caching for improved response times.

## Tech Stack

- Go (1.21+)
- Redis
- Docker
- Visual Crossing Weather API

## Features

- Real-time weather data retrieval
- Redis caching for optimized performance
- Dockerized deployment
- Environment-based configuration
- RESTful API endpoints

## Prerequisites

- Docker and Docker Compose
- Go 1.23+ (for local development)
- Redis (automatically handled via Docker)

## Quick Start

1. Clone the repository:

```bash
git clone https://github.com/yourusername/weather-api.git
cd weather-api
```

2. Set environment variables:

```bash
cp .env.example .env
# Edit .env with your Visual Crossing API key
```

3. Run with Docker:

```bash
docker compose up -d
```

## API Endpoints

### Get Current Weather

Endpoint: /weather/current/{location} Method: GET Params: location (string, required)

Example:

```bash
curl -X GET "http://localhost:8080/weather/current/london"
```

## Caching Implementation

The API implements Redis caching with the following features:

- Cache duration: 1 minute
- Automatic cache invalidation
- Location-based cache keys

The caching logic is implemented in the [`WeatherController`](./controllers/weather_controller.go):

- Checks Redis cache before making API calls
- Stores new weather data in Redis
- Returns cached data when available

## Docker Configuration

The application uses Docker Compose to run two services:

1. API Server ([`go-api`](./controllers/weather_controller.go)):

- Built from Golang 1.23
- Exposes port 8000
- Connects to Redis service

2. Redis Server ([`redis`](./controllers/weather_controller.go)):

- Uses official Redis image
- Persistent storage with Docker volumes
- Exposes port 6379
