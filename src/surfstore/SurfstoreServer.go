package surfstore

import (
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
	BlockStore BlockStoreInterface
	MetaStore  MetaStoreInterface
}

func (s *Server) GetFileInfoMap(succ *bool, serverFileInfoMap *map[string]FileMetaData) error {
	// panic("todo")
	err := s.MetaStore.GetFileInfoMap(succ, serverFileInfoMap)
	if err != nil {
		*succ = false
	}
	return err
}

func (s *Server) UpdateFile(fileMetaData *FileMetaData, latestVersion *int) error {
	// panic("todo")
	err := s.MetaStore.UpdateFile(fileMetaData, latestVersion)
	return err
}

func (s *Server) GetBlock(blockHash string, blockData *Block) error {
	// panic("todo")
	err := s.BlockStore.GetBlock(blockHash, blockData)
	return err
}

func (s *Server) PutBlock(blockData Block, succ *bool) error {
	// panic("todo")
	err := s.BlockStore.PutBlock(blockData, succ)
	if err != nil {
		*succ = false
	}
	return err
}

func (s *Server) HasBlocks(blockHashesIn []string, blockHashesOut *[]string) error {
	// panic("todo")
	err := s.BlockStore.HasBlocks(blockHashesIn, blockHashesOut)
	return err
}

// This line guarantees all method for surfstore are implemented
var _ Surfstore = new(Server)

func NewSurfstoreServer() Server {
	blockStore := BlockStore{BlockMap: map[string]Block{}}
	metaStore := MetaStore{FileMetaMap: map[string]FileMetaData{}}

	return Server{
		BlockStore: &blockStore,
		MetaStore:  &metaStore,
	}
}

func ServeSurfstoreServer(hostAddr string, surfstoreServer Server) error {
	// panic("todo")

	rpc.Register(&surfstoreServer)
	rpc.HandleHTTP()
	ln, err := net.Listen("tcp", hostAddr)
	if err != nil {
		return err
	}

	http.Serve(ln, nil)

	for {
	}
}
