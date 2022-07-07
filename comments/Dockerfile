FROM golang:1.18-alpine as build
WORKDIR /app
COPY . .
RUN ["go", "build"]

FROM golang:1.18-alpine as release
WORKDIR /bin
COPY --from=build /app/comments .
CMD ["comments"]