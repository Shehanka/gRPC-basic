package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	"github.com/campoy/justforfunc/31-grpc/todo"
	"github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

type taskServer struct {

}

func (s taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	return nil, fmt.Errorf("Not implemented")
}

func main() {
	srv := grpc.NewServer()
	var tasks taskServer
	todo.RegisterTasksServer(srv, tasks)
	l,err := net.Listen("tcp",":8888")
	if err != nil {
		log.Fatalf("Could not listen to :8888 %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type length int64