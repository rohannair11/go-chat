# Use an official GoLang runtime as the base image
FROM golang:1.16-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the GoLang source code into the container
COPY . .

# Build the Go application
RUN go build -o main

# Expose the required port
EXPOSE 8080

# Define the command to run the application
CMD ["./main"]
