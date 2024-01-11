FROM neo4j:latest

COPY ./neo4j_imports /var/lib/neo4j/import
COPY ./neo4j_plugins /var/lib/neo4j/plugins
COPY ./neo4j.conf /var/lib/neo4j/conf/neo4j.conf

EXPOSE 7687 7474
