# Use an official Go runtime as a parent image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app/src/github.com/bubblepays
RUN echo "This is first docker bubblepays application"

# Copy the current directory contents into the container at /app
COPY . .

# Set environment variables
# ENV DB_CONNECTION=root:root@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local
# ENV TEST_ENV='Hello environment'

# Download dependencies
RUN go mod tidy 

# Build the Go app
RUN go build -o bubblepays .

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Run the binary program
CMD ["./bubblepays"]