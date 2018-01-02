FROM alpine:3.5

COPY AptCoPilot AptCoPilot
COPY assets/build/es5-bundled assets
COPY db_password /etc/db_password
ENTRYPOINT ["./AptCoPilot"]