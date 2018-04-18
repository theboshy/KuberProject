package main
import (
	"log"
	"net"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"KuberProject/mcs"
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