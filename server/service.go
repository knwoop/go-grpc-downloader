package main

import (
	"io"
	"os"
	"path/filepath"

	pb "github.com/takahashi-ke/go_grpc_downloader/proto"
)

type fileService struct{}

func (s fileService) Download(req *pb.FileRequest,
	stream pb.FileService_DownloadServer) error {
	fp := filepath.Join(req.GetName())
	fs, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer fs.Close()
	buf := make([]byte, 1000*1024)
	for {
		n, err := fs.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		err = stream.Send(&pb.FileResponse{
			Data: buf[:n],
		})
		if err != nil {
			return err
		}
	}
	return nil
}