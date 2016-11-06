package main

import (
	songs "go_test/grpc_test/protoc/songs"

	"golang.org/x/net/context"
)

// server is used to implement songs.SongListServiceServer
type server struct{}

func (c *server) ListSongs(ctx context.Context, in *songs.SingerID) (*songs.SongList, error) {
    switch in.ID {
    case 1:
        rertun songs.SongList{}
    }
    return 
}
