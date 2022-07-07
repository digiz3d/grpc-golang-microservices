FROM golang:1.18-alpine as build
WORKDIR /app
COPY . .
WORKDIR /app/comments
RUN ["go", "build"]

FROM golang:1.18-alpine as release
WORKDIR /bin
COPY --from=build /app/comments/comments .
CMD ["comments"]