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
$ make mgratedown
```