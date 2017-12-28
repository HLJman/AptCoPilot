FROM alpine:3.5

COPY AptCoPilot AptCoPilot
COPY assets/build/es5-bundled assets
COPY dev/db_password_local /etc/db_password
ENTRYPOINT ["./AptCoPilot"]