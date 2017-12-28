export DB_USERNAME=jadmin
export DB_PASSWORD=dev/db_password_local
export DB_SERVER=localhost:8889
export DB_NAME=copilot

go build -ldflags "-s" -o AptCoPilot

./AptCoPilot