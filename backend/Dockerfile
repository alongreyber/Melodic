# Start from the latest golang base image
FROM golang:latest as build-stage

# Add Maintainer Info
LABEL maintainer="Alon Greyber <alongreyber@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install air to run with hot reload
RUN curl -fLo /bin/air \
	https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air && \
    chmod +x /bin/air 

# Build the Go app
RUN go build -o main .

# Expose port 5000 to the outside world
EXPOSE 5000

# Command to run the executable
CMD ["/bin/air", "-c", "air.conf"]

#
# Production code!
# 

FROM alpine:3.7 as production-stage

COPY --from=build-stage /app /app

EXPOSE 5000

ENTRYPOINT /app/main
