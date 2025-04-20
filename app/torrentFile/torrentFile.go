package torrentFile

import (
	"io"

	"github.com/jackpal/bencode-go"
)

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type bencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

func open(reader io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}

	err := bencode.Unmarshal(reader, &bto)

	if err != nil {
		return nil, err
	}

	return &bto, nil
}
