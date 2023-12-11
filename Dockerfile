# Use an official Golang runtime as a base image
FROM golang:1.18

# Set the working directory in the container
WORKDIR /go/src/app

# Copy the current directory contents into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Run the executable
CMD ["./main"]
