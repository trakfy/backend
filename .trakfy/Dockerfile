# Use the official Go image as the base image
FROM golang:1.19.3-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the module dependencies
RUN go mod download

# Copy the application source code to the working directory
COPY . .

# Copy the .env file to the working directory
COPY .env ./

# Copy the db/migrations directory to the working directory
COPY db/migrations ./db/migrations

# Build the application into an executable binary
RUN go build -o myapp

# Set the final base image
FROM alpine:latest

# Copy the executable binary from the build image to the final image
COPY --from=build /app/myapp /usr/local/bin/myapp

# Copy the .env file to the final image
COPY --from=build /app/.env /app/.env

# Copy the db/migrations directory to the final image
COPY --from=build /app/db/migrations /app/db/migrations

# Set the working directory inside the container
WORKDIR /app

# Load the environment variables from the .env file
ENV PATH="/usr/local/bin:${PATH}"
ENV ENV_FILE="/app/.env"

# Expose port 8000 to the outside
EXPOSE 8000

# Run the application when the container starts
CMD ["myapp"]
