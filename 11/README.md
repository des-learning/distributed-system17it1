run server:
go run server.go

simulate client request (forever):
while true; do curl -XPOST http://localhost:8080/hello -d '{"message": "budi"}'; echo; done
