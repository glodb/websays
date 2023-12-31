# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 for the Go application to listen on
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]