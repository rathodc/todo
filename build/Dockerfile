FROM golang:latest
RUN mkdir /app
COPY docker_install.sh /app
WORKDIR /app
RUN /app/docker_install.sh
EXPOSE 8000
CMD ["/app/run.sh"]
