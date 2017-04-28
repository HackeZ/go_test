package main

import "go_test/grpc/protoc/songs"

import (
	"errors"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Server

// server is used to implement songs.SongListServiceServer
type server struct{}

var Singer1 = &songs.Singer{
	ID:     1,
	Name:   "Beethoven",
	Gender: true,
	Height: 1.72,
}

var Singer2 = &songs.Singer{
	ID:     2,
	Name:   "Mozart",
	Gender: true,
	Height: 1.82,
}

var Song1 = &songs.Song{
	ID:     42863,
	Name:   "今夜无人入眠",
	Singer: Singer1,
}

var Song2 = &songs.Song{
	ID:     56383,
	Name:   "莫斯科郊外的晚上",
	Singer: Singer2,
}

var Song3 = &songs.Song{
	ID:     85348,
	Name:   "欢乐颂",
	Singer: Singer2,
}

func (s *server) ListSongs(ctx context.Context, in *songs.SingerID) (*songs.SongList, error) {
	switch in.ID {
	case Singer1.ID:
		return &songs.SongList{
			Songs: []*songs.Song{
				Song1,
			},
		}, nil
	case Singer2.ID:
		return &songs.SongList{
			Songs: []*songs.Song{
				Song2,
				Song3,
			},
		}, nil
	default:
		return nil, errors.New("Unkown Singer ID")
	}
}

func (s *server) GetSongs(in *songs.SingerID, stream songs.SongListService_GetSongsServer) error {
	SongBelongSinger2 := []*songs.Song{Song2, Song3}
	switch in.ID {
	case Singer1.ID:
		if err := stream.Send(Song1); err != nil {
			return err
		}
	case Singer2.ID:
		for _, song := range SongBelongSinger2 {
			if err := stream.Send(song); err != nil {
				return err
			}
		}
	default:
		return errors.New("Unknow Singer ID")
	}
	return nil
}

func main() {
	l, err := net.Listen(Net, ServiceAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	songs.RegisterSongListServiceServer(s, &server{})
	s.Serve(l)
}
