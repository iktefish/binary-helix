package nodes

import (
	"github.com/iktefish/binary-helix/utils"
	"io"
	"net"
	"os"
)

func Client(source string) {
	// source := "../tmp/aFile"
	target := "../tmp/file1"
	address := "localhost:4040"
	uploadFile(source, address)
	downloadFile(target, address)
}

func uploadFile(source, address string) {

	/* Connect to server */
	conn, err := net.Dial("tcp", address)
	utils.CheckError(err)
	defer conn.Close()

	/* Open file to upload */
	fi, err := os.Open(source)
	utils.CheckError(err)
	defer fi.Close()

	/* Upload */
	_, err = io.Copy(conn, fi)
	utils.CheckError(err)
}

func downloadFile(target, address string) {

	/* Create new file to hold response */
	fo, err := os.Create(target)
	utils.CheckError(err)
	defer fo.Close()

	/* Connect to server */
	conn, err := net.Dial("tcp", address)
	utils.CheckError(err)
	defer conn.Close()

	_, err = io.Copy(fo, conn)
	utils.CheckError(err)
}
