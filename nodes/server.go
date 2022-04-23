package nodes

import (
	"github.com/iktefish/binary-helix/utils"
	"io"
	"net"
	"os"
)

func Server() {
	target := "../tmp/"
	source := "../tmp/test-cli-ser-exchange"
	address := "172.17.0.2:4040"

	/* Server start listenning */
	server, err := net.Listen("tcp", address)
	utils.CheckError(err)
	defer server.Close()

	recieveFile(server, target)
	sendFile(server, source)
}

func recieveFile(server net.Listener, target string) {

	/* Accept connection */
	conn, err := server.Accept()
	utils.CheckError(err)

	/* Create new file */
	fo, err := os.Create(target)
	utils.CheckError(err)
	defer fo.Close()

	/* Accept file from client & write to new file */
	_, err = io.Copy(fo, conn)
	utils.CheckError(err)
}

func sendFile(server net.Listener, source string) {
	/* Accept connection */
	conn, err := server.Accept()
	utils.CheckError(err)

	/* Open file to send */
	fi, err := os.Open(source)
	utils.CheckError(err)
	defer fi.Close()

	/* Send file to client */
	_, err = io.Copy(conn, fi)
	utils.CheckError(err)
}
