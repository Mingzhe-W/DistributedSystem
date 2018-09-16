package main
import(
  "fmt"
  "proto"
  "os"
  "net"
  "net/rpc"
  "log"
	"regexp"
  "bufio"
	"net/tcp"
)



func report_and_exit(msg string, err error){
  log.Fatal(msg,err)
  os.Exit(1)
}

type Logging int


func (t *Logging)fetchLog(byteArgs []byte, byteReply *[]byte)error{
  //Decoding
  args :=&pb.ClientQuery{}
  err:= proto.Unmarshal(byteArgs, args)

  //place holder
  reply:=pb.ServerResponse{}

  //Open file
  file, err 