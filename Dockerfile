FROM golang:alpine AS dist
WORKDIR /app
COPY . .
RUN apk add make && make build

FROM alpine AS production
WORKDIR /app
# you can edit and install many things there...
# copy dist
COPY --from=dist /app/build/manboster /app/manboster

ENV MANBOSTER_HOME="/app/manboster"

EXPOSE 8080
# run build manboster
CMD ["./manboster"]