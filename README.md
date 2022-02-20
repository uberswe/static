# Static

A minimal base image for serving static files. I created 
this because I felt that it was a bit overkill to run nginx 
in a container to simply host some static files. I run this
for production sites where I have [Caddy](https://caddyserver.com/) 
handle traffic to all sites and makes a reverse proxy to 
the containers.

## Example

I have some static sites built with [Hugo](https://gohugo.io/) and I use the following Dockerfile to get them running in a container.

```
FROM klakegg/hugo:0.92.1-onbuild AS build

FROM ghcr.io/uberswe/static:main

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

If you are using a framework like ReactJS you might want all routes that are not found to be routed to the index file. I have added the `INDEX_IF_NOT_FOUND` variable to allow you to toggle this feature on, by default it's off and a basic 404 page will be shown.

```
version: "3.9"

services:
  app:
    container_name: app
    image: app
    build:
      context: .
      dockerfile: Dockerfile
    environment:
      INDEX_IF_NOT_FOUND: "true"
    restart: always
    ports:
      - "8080:80"
```