package uuid

import satoriuuid "github.com/satori/go.uuid"

// UUID represents a 128 bit universally unique identifier
type UUID [16]byte

// NewUUID will return a universally unique identifier
func NewUUID() UUID {
	var bytes [16]byte
	copy(bytes[:], satoriuuid.NewV4().Bytes())
	return UUID(bytes)
}

func (u UUID) bytes() []byte {
	return u[:]
}

func (u UUID) String() string {
	uuid, err := satoriuuid.FromBytes(u.bytes())

	if err != nil {
		panic(err)
	}

	return uuid.String()
}

// UUID represents globally unique identifier to use in wt project
// type UUID [16]byte

// var (
// 	zeroUUID = [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
// )

// // NewUUID returns new v4 uuid
// func NewUUID() UUID {
// 	var bytes [16]byte
// 	copy(bytes[:], uuid.NewV4().Bytes())
// 	return UUID(bytes)
// }

// // ZeroUUID returns zero initialized UUID
// func ZeroUUID() UUID {
// 	return zeroUUID
// }

// // IsEmpty returns if the bytes are all zeros
// func IsEmpty(arg UUID) bool {
// 	return bytes.Equal(arg[:], zeroUUID[:])
// }

// func (u UUID) bytes() []byte {
// 	return u[:]
// }

// func (u UUID) String() string {
// 	uuid, err := uuid.FromBytes(u.bytes())

// 	if err != nil {
// 		panic(err)
// 	}

// 	return uuid.String()
// }

// // FromString returns UUID parsed from string input
// func FromString(input string) (u UUID, err error) {
// 	uuid, err := uuid.FromString(input)

// 	if err != nil {
// 		return zeroUUID, err
// 	}

// 	var bytes [16]byte
// 	copy(bytes[:], uuid.Bytes())

// 	return UUID(bytes), nil
// }

// // MustFromString returns UUID parsed from string or it panics
// func MustFromString(input string) (u UUID) {
// 	uuid, err := uuid.FromString(input)

// 	if err != nil {
// 		panic(err)
// 	}

// 	var bytes [16]byte
// 	copy(bytes[:], uuid.Bytes())

// 	return UUID(bytes)
// }
