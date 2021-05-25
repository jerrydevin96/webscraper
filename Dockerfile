FROM ubuntu
RUN apt-get update
RUN apt-get install -y ca-certificates
RUN mkdir /webscraper
WORKDIR /webscraper
COPY webscraper /webscraper/
RUN chmod 777 webscraper
EXPOSE 8080

CMD ["./webscraper"]