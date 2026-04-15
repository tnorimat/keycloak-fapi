FROM openresty/openresty:alpine-fat

RUN apk update && apk add ca-certificates

COPY --from=proxy-ca proxy-ca.crt /tmp/proxy-ca.crt
RUN if [ -s /tmp/proxy-ca.crt ]; then cp /tmp/proxy-ca.crt /usr/local/share/ca-certificates/proxy-ca.crt && update-ca-certificates; fi

RUN /usr/local/openresty/luajit/bin/luarocks install lua-resty-openidc \
  && /usr/local/openresty/luajit/bin/luarocks install lua-resty-jit-uuid

COPY *.template /usr/local/openresty/nginx/conf/
COPY *.lua /usr/local/openresty/

COPY entrypoint.sh /usr/local/sbin/
RUN chmod +x /usr/local/sbin/entrypoint.sh

EXPOSE 443

ENTRYPOINT ["/usr/local/sbin/entrypoint.sh"]
CMD []
