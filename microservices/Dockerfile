# Start from the latest golang base image
FROM golang:1.17-buster

# Add Maintainer Info
LABEL maintainer="Your Name <your_email@domain.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the api folder.
COPY ./api ./api

# Copy go mod and sum files
ARG service
COPY ./${service} ./${service}

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN cd ./${service} && go mod download





# Build the Go app
RUN cd ./${service}/cmd && go build -o main


# Stage 2: Create a minimal runtime image



# Copy the executable from the builder stage

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable

