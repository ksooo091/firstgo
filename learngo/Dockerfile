FROM golang as builder

WORKDIR /app
COPY go.mod ./
COPY *.go ./


RUN CGO_ENABLED=0 GOOS=linux go build \
#-a 
    -ldflags="-s -w" -o /ksooo-study
CMD ["/ksooo-study"]

FROM alpine

COPY --from=builder /ksooo-study /ksooo-study



ENTRYPOINT ["/ksooo-study"]