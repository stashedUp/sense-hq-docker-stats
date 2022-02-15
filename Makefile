
   
EXE  := sense-stats

$(EXE): go.mod *.go
	go build -o $(EXE) main.go

