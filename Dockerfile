# syntax=docker/dockerfile:1.3

FROM golang:1.23.6-bookworm AS build

RUN apt-get update

WORKDIR /build

# Copy full source
ADD . /build

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y \
    nodejs \
    npm \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /build/bin/performer /usr/local/bin/performer
COPY --from=build /build/scripts/js-runner.js /app/scripts/js-runner.js
COPY --from=build /build/scripts/py-runner.py /app/scripts/py-runner.py

RUN chmod +x /app/scripts/js-runner.js
RUN chmod +x /app/scripts/py-runner.py

CMD ["/usr/local/bin/performer"]
