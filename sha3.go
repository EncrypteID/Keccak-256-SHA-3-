package keccak256sha3

import "hash"

// spongeDirection indicates the direction bytes are flowing through the sponge.
type spongeDirection int

// NewSHA3224 returns a new hash.Hash computing SHA3-224 as specified in the FIPS 202 draft.
func NewSHA3224() hash.Hash {
	return NewKeccak(224*2, 224, domainSHa3)
}

// NewSHA3256 returns a new hash.Hash computing SHA3-256 as specified in the FIPS 202 draft.
func NewSHA3256() hash.Hash {
	return NewKeccak(256*2, 256, domainSHa3)
}

// NewSHA3384 returns a new hash.Hash computing SHA3-384 as specified in the FIPS 202 draft.
func NewSHA3384() hash.Hash {
	return NewKeccak(384*2, 384, domainSHa3)
}

// NewSHA3512 returns a new hash.Hash computing SHA3-512 as specified in the FIPS 202 draft.
func NewSHA3512() hash.Hash {
	return NewKeccak(512*2, 512, domainSHa3)
}
