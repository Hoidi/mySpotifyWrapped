general:
	go build main.go helper.go setup.go stats.go
	./main general
	rm -f main helper

artistUnique:
	go build main.go helper.go setup.go stats.go
	./main artist unique
	rm -f main helper

artistTime:
	go build main.go helper.go setup.go stats.go
	./main artist time
	rm -f main helper

trackUnique:
	go build main.go helper.go setup.go stats.go
	./main track unique
	rm -f main helper

trackTime:
	go build main.go helper.go setup.go stats.go
	./main track time
	rm -f main helper

all:
	go build main.go helper.go setup.go stats.go
	./main all
	rm -f main helper