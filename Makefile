build:
	go build -v -o bin/taxi-fare.exe cmd/*.go

run: 
	@echo " >> build taxi-fare"
	@make build
	@echo " >> taxi-fare built."
	@echo " >> executing taxi-fare"
	@./bin/taxi-fare.exe