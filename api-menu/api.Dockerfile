FROM golang:alpine AS builder

ENV PROJECT_PATH=/api-menu-d
WORKDIR $PROJECT_PATH

# ARG buildenv=dev
COPY . .
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
RUN go mod download

# COPY *.go ./
# RUN go build -o ./server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .

FROM scratch
COPY --from=builder ./server /
COPY ./configs ./configs
COPY ./serviceAccount.json ./serviceAccount.json


EXPOSE 8081

ENTRYPOINT ["./server"] 