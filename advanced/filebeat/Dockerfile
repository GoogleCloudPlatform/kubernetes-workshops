FROM debian:jessie

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    wget

RUN curl -L -O https://download.elastic.co/beats/filebeat/filebeat_1.2.2_amd64.deb && \
    dpkg -i filebeat_1.2.2_amd64.deb

COPY filebeat.yml /etc/filebeat/filebeat.yml

CMD ["/usr/bin/filebeat", "-c", "/etc/filebeat/filebeat.yml"]
