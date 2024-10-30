kill -9 `lsof -ti:8000`
cd fixtures && docker-compose down -v && docker-compose up -d
cd ..

go run main.go
