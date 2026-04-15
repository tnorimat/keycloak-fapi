FROM docker:27.3.1-dind-alpine3.20
COPY --from=proxy-ca proxy-ca.crt /tmp/proxy-ca.crt
RUN if [ -s /tmp/proxy-ca.crt ]; then cp /tmp/proxy-ca.crt /usr/local/share/ca-certificates/proxy-ca.crt && update-ca-certificates; fi
ADD test-runner-entrypoint.sh .
RUN apk add --update --upgrade --no-cache curl
CMD [ "sh", "./test-runner-entrypoint.sh" ]