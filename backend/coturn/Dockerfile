FROM debian:bullseye

RUN apt-get update && apt-get install -y coturn && rm -rf /var/lib/apt/lists/*

COPY turnserver.conf /etc/turnserver.conf

EXPOSE 3478/udp 3478/tcp 5349/udp 5349/tcp

ENTRYPOINT ["turnserver"]
CMD ["-c", "/etc/turnserver.conf", "--no-cli"]
