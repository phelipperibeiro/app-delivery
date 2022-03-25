FROM docker.elastic.co/elasticsearch/elasticsearch:7.11.2

USER root
RUN mkdir -p /usr/share/elasticsearch/data
RUN chmod 777 -R /usr/share/elasticsearch/
RUN chown -R elasticsearch:elasticsearch /usr/share/elasticsearch
USER elasticsearch


# rm -Rf es01
# mkdir -p es01/usr/share/elasticsearch/data
# RUN chown 1000:1000 -R /usr/share/elasticsearch
# USER 1000