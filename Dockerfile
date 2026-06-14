FROM golang:alpine AS dist
WORKDIR /app
COPY . .
RUN make build

FROM alpine AS production
WORKDIR /app
# you can edit and install many things there...
# copy dist
COPY --from=dist /app/build/manboster /app/manboster

ENV MANBOSTER_HOMEDIR="/app/manboster"

EXPOSE 8080
# run build manboster
CMD ["./manboster"]