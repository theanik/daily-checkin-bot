### Daily Check-in bot
A CLI chat bot app using go for update daily check-in status and show reports.

### Installation 
```
// Pull the repo and cd to it.
cd {directory}
// Setup .env file
cp .env.example .env
// Build docker containers
docker compose up
// Go inside docker conainer and run migrations
docker exec -it GO_APP bash
// Run migrations
go run app/migrations/create_daily_report_migration.go
go run app/migrations/create_user_migration.go
```

### Run App
```
// Run the command from inside GO_APP docker container
go run app/chat.go
```