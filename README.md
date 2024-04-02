Starting postgres docker container
```bash
docker run --name postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
```

Create casdoor database
```bash
docker exec -it postgres psql -U postgres
```
```sql
CREATE DATABASE casdoor;
```

Starting casdoor & postgres docker container
```bash
docker compose up -d
```