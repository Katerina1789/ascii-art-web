#We download the correct go version and naming this stage as builder because we are going to use it later
# Katerina added: Changed the Go version to match the go.mod file.
FROM golang:1.25.7 AS builder

#Creating our working directory witch stores our programme 
WORKDIR /app

#Copies go.mod file 
COPY go.mod ./

#Download the dependencies of go mod if they exist 
RUN go mod download

#We copy the whole
COPY . .

#We build the excecutable without the libriries and linkers of C select the operating 
#system our app will run and the CPU architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ascii-art ./cmd/main.go


# Stage 2 - Removing unessesary steps

#Installing minimal OS like alpine in order to save space (takes around ~5mb)
FROM alpine:3.19

# Katerina added: Installing certificates for HTTPS support as part of good practice since we use HTTP and the code downloads banner files from URL.
RUN apk add --no-cache ca-certificates

#Creating the working directory again cause we basically create a new instance of an image 
WORKDIR /app

#Copy binary from builder 

COPY --from=builder /app/ascii-art .

#Copy required folders

COPY --from=builder /app/templates ./templates

COPY --from=builder /app/static ./static

COPY --from=builder /app/internal/utilities/ascii/banners ./internal/utilities/ascii/banners

#Open port for docker to communicate with our programme
EXPOSE 8080

#Creating arguments for better tracking builds
ARG BUILD_DATE
ARG MAINTAINER

LABEL build_date=${BUILD_DATE}
LABEL maintainer=${MAINTAINER}

#Starting the app

LABEL version="1.1.0"
CMD ["./ascii-art"]
