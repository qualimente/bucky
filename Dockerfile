FROM scratch

COPY bin/linux/amd64/bucky /bucky

ENTRYPOINT [ "/bucky" ]
