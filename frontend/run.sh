clear
go mod tidy                             || exit $?
go run . --broker tcp://127.0.0.1:9901  || exit $?
