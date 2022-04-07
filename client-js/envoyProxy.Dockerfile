FROM envoyproxy/envoy-dev:latest
COPY envoy.yml /etc/envoy/
RUN chmod go+r /etc/envoy/envoy.yml
