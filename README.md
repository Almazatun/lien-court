# 🦊 Lien-court

## Installation
```bash
# Clone the repository
git clone https://github.com/Almazatun/lien-court.git
# Enter into the directory
cd lien-court/
# Install the dependencies
go mod download | go mod tidy | make install
```

### Build app
```bash
$ make build
```
### Run app
```bash
$ make run
```

### Database
```bash
# Migration up
$ make migrateup
# Migration down
$ make migratedown
```
### API
```bash
# Create link
curl -d '{ "link":"original_url" }' -H "Content-Type: application/json" -H "Authorization: Bearer ${TOKEN}" -X POST http://localhost:${PORT}/api/v1/links
```
