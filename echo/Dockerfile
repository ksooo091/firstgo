FROM golang as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /ksooo-study
CMD ["/ksooo-study"]

FROM alpine

COPY --from=builder /ksooo-study /ksooo-echo
COPY /index.html ./echo/index.html


ENTRYPOINT ["/ksooo-echo"]