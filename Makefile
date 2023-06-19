generate:
	find . -name "*easyjson.go" -type f -delete
	easyjson  -stubs */*/*/*/dto/*.go
	go generate ./...
func:
	curl -vvv -X POST http://localhost:5000/api/service/clear
	./technopark-dbms-forum func -u http://localhost:5000/api/ -r report.html
fill:
	curl -vvv -X POST http://localhost:5000/api/service/clear
	time ./technopark-dbms-forum fill -u http://localhost:5000/api/
perf:
	./technopark-dbms-forum perf -u http://localhost:5000/api/  --duration=60 --step=4 -v=1
