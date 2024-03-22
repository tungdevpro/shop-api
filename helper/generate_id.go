package helper

import (
	"github.com/sqids/sqids-go"
)

func EncodeId(input int) string {
	s, _ := sqids.New(sqids.Options{
		MinLength: 10,
	})

	id, _ := s.Encode([]uint64{uint64(input)})
	return id
}

func DecodeId(input string) int {
	s, _ := sqids.New(sqids.Options{
		MinLength: 10,
	})

	result := s.Decode(input)
	return int(result[0])
}
