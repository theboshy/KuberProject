package main
import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"KuberProject/mcs"
	"io/ioutil"
	"strconv"
)

type server struct{}

const port  = ":3001"
const netConn  = "tcp"

func main() {
	lis, err := net.Listen(netConn, port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	mcs.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *mcs.GCDRequest) (*mcs.GCDResponse, error) {
	a, b := r.A, r.B
	for b != 0 {
		a, b = b, a%b
	}
	return &mcs.GCDResponse{Result: a}, nil
}

func (c *server) SaveFile(ctx context.Context, r *mcs.FileRequest) (*mcs. FileResponse, error) {
	const outputpath ="C:/Users/Software1/Desktop/"
error := ioutil.WriteFile(outputpath+r.FileName,r.BinaryFile,0644)
check(error)
	return &mcs.FileResponse{OutPath: outputpath +r.FileName,Message:"image saved, image size : "+strconv.Itoa(int(r.FileSize))}, nil
}



func check(e error) {
	if e != nil {
		panic(e)
	}
}