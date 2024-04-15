# Use Ubuntu as base image
FROM ubuntu:latest

# Install Rust
RUN apt-get update && apt-get install -y \
    build-essential \
    curl

RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

# Set up your Rust application
WORKDIR /usr/src/app
COPY . .

# Build your Rust application
RUN cargo build --release

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["target/release/memcache_client_rust"]