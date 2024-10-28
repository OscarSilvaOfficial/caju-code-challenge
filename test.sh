echo "-- Running unit tests --"
go test ./tests/unit/**/*

echo "\n"
echo "-- Running integration tests --"
go test ./tests/integration/**