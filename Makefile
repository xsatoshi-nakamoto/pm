migrateup:
	migrate -path ./db/migration -database "postgresql://admin:HTrang@091297@localhost:5432/firmsone_pm?sslmode=disable" -verbose up
migratedown:
	migrate -path ./db/migration -database "postgresql://admin:HTrang@091297@localhost:5432/firmsone_pm?sslmode=disable" -verbose down