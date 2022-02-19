# Static

A minimal base image for serving static files. I created 
this because I felt that it was a bit overkill to run nginx 
in a container to simply host some static files. I run this
for production sites where I have [Caddy](https://caddyserver.com/) 
handle traffic to all sites and makes a reverse proxy to 
the containers.

[Available on Docker Hub: tenghamn/static](https://hub.docker.com/r/tenghamn/static)

## Example

I have some static sites built with [Hugo](https://gohugo.io/) and I use the following Dockerfile to get them running in a container.

```
FROM klakegg/hugo:0.92.1-onbuild AS build

FROM tenghamn/static:latest

WORKDIR /app

COPY --from=build /target /app

EXPOSE 80
```

Then I have a docker-compose.yml file which will run the container and serve the static files on port 8080

```
version: "3.9"

services:
  app:
    container_name: app
    image: app
    build:
      context: .
      dockerfile: Dockerfile
    restart: always
    ports:
      - "8080:80"
```