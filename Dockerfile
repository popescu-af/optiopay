FROM scratch

WORKDIR /app
COPY ./main-svc $WORKDIR

EXPOSE 80

CMD ["/app/main-svc"]
