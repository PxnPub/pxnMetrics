
while true; do
	clear
	go mod tidy              || break
	go run . --num-shards 5  || break
	sleep 2
done
