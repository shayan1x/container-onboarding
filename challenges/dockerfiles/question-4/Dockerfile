# Use a base image for Go
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the entire application code into the container
COPY . .

# Install dependencies and build the Go application
RUN go mod download && go build -o myapp

# Set default environment variables
ENV PORT=8080
ENV MESSAGE="Hello, Docker!"

# Expose the port the application will run on
EXPOSE 8080

# Define the command to run the application
CMD ["/app/myapp"]
