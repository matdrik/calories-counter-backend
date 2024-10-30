#cd ./pkg/migrate/ || exit
#go run migrate.go
migrate -source file://migrations -database postgresql://postgres:8563@localhost:5432/calories_counter up 2