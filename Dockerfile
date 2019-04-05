FROM golang:1.12 AS builder

WORKDIR /opt/risk-rating

COPY . .

RUN make get clean build-alpine-scratch
# --------
FROM scratch

WORKDIR /

COPY --from=builder /opt/risk-rating/bin/amd64/scratch .

ENTRYPOINT [ "/risk-rating" ]
CMD ["--help"]
