FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    build-essential \
    curl

RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

WORKDIR /usr/src/app
COPY . .

RUN cargo build --release

EXPOSE 8080

CMD ["target/release/memcache_client_rust"]