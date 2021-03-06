package common

import (
	"github.com/MixinNetwork/mixin/crypto"
)

type UTXO struct {
	Input
	Output
	Asset crypto.Hash `msgpack:"C"`
}

type UTXOWithLock struct {
	UTXO
	LockHash  crypto.Hash `msgpack:"LH"`
	LockUntil uint64      `msgpack:"LU"`
}

type UTXOLocker func(hash crypto.Hash, index int, tx crypto.Hash, until uint64) (*UTXO, error)

type GhostChecker func(key crypto.Key) (bool, error)

func (s *Snapshot) UnspentOutputs() []*UTXO {
	tx := s.Transaction

	var utxos []*UTXO
	for i, out := range tx.Outputs {
		utxo := &UTXO{
			Input: Input{
				Hash:  tx.Hash(),
				Index: i,
			},
			Output: Output{
				Type:   out.Type,
				Amount: out.Amount,
				Keys:   out.Keys,
				Script: out.Script,
				Mask:   out.Mask,
			},
			Asset: tx.Asset,
		}
		utxos = append(utxos, utxo)
	}
	return utxos
}
