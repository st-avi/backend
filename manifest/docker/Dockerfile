FROM alpine:latest

ENV WORKDIR /app

ADD ./main $WORKDIR/main
RUN chmod +x $WORKDIR/main

WORKDIR $WORKDIR
CMD ./main