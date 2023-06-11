# Use an official GoLang runtime as the base image
FROM golang:1.16-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the GoLang source code into the container
COPY . .

# Build the Go application
RUN go build -o main

#multistage docker images 
FROM busybox as release

WORKDIR /app

COPY --from=builder /app/main .

# Expose the required port
EXPOSE 9000

# Define the command to run the application
CMD ["./main"]
