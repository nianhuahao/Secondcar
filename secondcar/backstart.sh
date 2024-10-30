kill -9 `lsof -ti:8000`
nohup  go run main.go >log.txt &

