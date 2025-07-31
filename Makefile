

new:
	@echo "Creating new problem folder..."
	mkdir problems/leetcode/$(name)
	@echo "making go files in the created folder..."
	touch problems/leetcode/$(name)/$(name).go
	@echo "fill in the problem statement"
	

test:
	@echo "Running tests...after filling in the problem code"
	gotests -all -w -parallel problems/leetcode/$(name)/$(name).go
	@echo "test files generated successfully"

watch:
	@echo "Watching for changes..."
	go test problems/leetcode/$(name)/$(name).go problems/leetcode/$(name)/$(name)_test.go


help:
	@echo "Available commands:"
	@echo "NEW---> make new name=<problem_name> - Create a new problem folder and Go files"
	@echo "TEST---->  make test name=<problem_name> - Run tests for the specified problem"
	@echo "WATCH---->  make watch name=<problem_name> - Watch for changes and run tests automatically"
	@echo "HELP---->  make help - Show this help message"
