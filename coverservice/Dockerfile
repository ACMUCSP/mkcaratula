FROM ubuntu:16.04

EXPOSE 8000

RUN apt update && \
    apt install -y --no-install-recommends texlive-latex-base texlive-fonts-recommended musl && \
    rm -rf /var/lib/apt/lists/*

WORKDIR "/srv/coverservice"

COPY coverservice/docker-entrypoint.sh /entrypoint.sh
COPY img/ucsp.png img/
COPY templates/cover.tex templates/
COPY coverservice/coverservice-linux-amd64 server

ENTRYPOINT ["/entrypoint.sh"]
CMD ["./server"]