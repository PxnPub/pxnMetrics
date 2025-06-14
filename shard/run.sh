clear
go mod tidy  || exit $?
go run . --bind tcp://127.0.0.1:9001 --broker tcp://127.0.0.1:9901     || exit $?
