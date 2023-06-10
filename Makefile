.PHONY: bench
bench:
	@echo "Building Redis Bench Go binary..."
	go build -o redis-bench cmd/bench/main.go
