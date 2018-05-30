package main
import (
	"log"
	"time"
	"net"
	pb "math/math"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func printQR(client pb.MathClient,request *pb.Request){
	ctx , cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	resp,err := client.Div(ctx,request)
	if err != nil{
		log.Fatalf("%v.Div(_) = _, %v", client , err)
	}
	log.Println(resp)
}
func main(){
	serverAddr := net.JoinHostPort("localhost","8080")

	//setup insecure connection
	conn,err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewMathClient(conn)
	printQR(client,&pb.Request{ Dividend:5 , Divisor: 3})
}