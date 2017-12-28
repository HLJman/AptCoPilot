FROM alpine:3.5

ENV DB_USERNAME admin
ENV DB_PASSWORD /etc/db_password
ENV DB_SERVER aptcopilot-test.ck7rzkuy8vhb.us-west-1.rds.amazonaws.com:3306
ENV DB_NAME aptcopilot_test

COPY AptCoPilot AptCoPilot
COPY assets/build/es5-bundled assets
COPY db_password /etc/db_password
ENTRYPOINT ["./AptCoPilot"]