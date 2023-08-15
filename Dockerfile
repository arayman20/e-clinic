# Start from a base Go image
FROM golang:alpine AS build



# Set the working directory inside the container
WORKDIR /app

# Copy source code
COPY . .

RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/images /home/images
COPY --from=build /app/main .
CMD /app/main