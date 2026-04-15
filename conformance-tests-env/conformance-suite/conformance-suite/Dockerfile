FROM maven:3.8.3-openjdk-17-slim AS builder
ARG OPENID_GIT_URL
ARG OPENID_GIT_TAG

COPY --from=proxy-ca proxy-ca.crt /tmp/proxy-ca.crt
RUN if [ -s /tmp/proxy-ca.crt ]; then \
      cp /tmp/proxy-ca.crt /usr/local/share/ca-certificates/proxy-ca.crt && \
      update-ca-certificates || true && \
      keytool -import -trustcacerts -keystore "$JAVA_HOME/lib/security/cacerts" -storepass changeit -noprompt -file /tmp/proxy-ca.crt -alias proxy-ca; \
    fi

RUN apt-get update && apt-get install -y git
RUN git clone -b ${OPENID_GIT_TAG} ${OPENID_GIT_URL} ./conformance-suite
ADD . ./conformance-suite/
RUN PROXY_OPTS="" && \
    PROXY_URL="${https_proxy:-${HTTPS_PROXY:-${http_proxy:-$HTTP_PROXY}}}" && \
    if [ -n "$PROXY_URL" ]; then \
      PROXY_HOST=$(echo "$PROXY_URL" | sed -e 's|https\?://||' -e 's|/||g' -e 's|:[0-9]*$||') && \
      PROXY_PORT=$(echo "$PROXY_URL" | sed -e 's|https\?://||' -e 's|/||g' -e 's|.*:||') && \
      echo "Configuring Maven proxy: host=$PROXY_HOST port=$PROXY_PORT" && \
      PROXY_OPTS="-Dhttps.proxyHost=$PROXY_HOST -Dhttps.proxyPort=$PROXY_PORT -Dhttp.proxyHost=$PROXY_HOST -Dhttp.proxyPort=$PROXY_PORT"; \
    fi && \
    cd conformance-suite && mv -f run-tests.sh ./.gitlab-ci && \
    MAVEN_OPTS="$PROXY_OPTS" mvn package -DskipTests

FROM eclipse-temurin:17

COPY --from=proxy-ca proxy-ca.crt /tmp/proxy-ca.crt
RUN if [ -s /tmp/proxy-ca.crt ]; then cp /tmp/proxy-ca.crt /usr/local/share/ca-certificates/proxy-ca.crt && update-ca-certificates; fi

# the server app requires redir to run
RUN apt-get update && apt-get install -y redir python3 python3-pip curl \
    && pip3 install httpx pyparsing --break-system-packages \
    && rm -rf /var/lib/apt/lists/*
COPY --from=builder /conformance-suite /conformance-suite
EXPOSE 8080
EXPOSE 9090
CMD [ "sh", "./conformance-suite/server-entrypoint.sh" ]
