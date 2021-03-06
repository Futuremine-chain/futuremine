package types

import (
	"errors"
	"fmt"
	"github.com/Futuremine-chain/future/common/config"
	"github.com/Futuremine-chain/future/future/common/kit"
	"github.com/Futuremine-chain/future/tools/amount"
	"github.com/Futuremine-chain/future/tools/arry"
)

type MessageType uint8

const (
	Transaction MessageType = iota
	Token
	Candidate
	Cancel
	Vote
)

const (
	minFees = 1e4
	maxFees = 1e9
)

type MsgHeader struct {
	Type      MessageType
	Hash      arry.Hash
	From      arry.Address
	Nonce     uint64
	Fee       uint64
	Time      uint64
	Signature *Signature
}

func (m *MsgHeader) Check() error {
	if err := m.checkType(); err != nil {
		return err
	}

	if err := m.checkFrom(); err != nil {
		return err
	}

	if err := m.checkFee(); err != nil {
		return err
	}

	if err := m.checkSinger(); err != nil {
		return err
	}
	return nil
}

func (m *MsgHeader) checkType() error {
	switch m.Type {
	case Transaction:
		return nil
	case Token:
		return nil
	case Candidate:
		return nil
	case Cancel:
		return nil
	case Vote:
		return nil
	}
	return fmt.Errorf("there are no messages of type %d", m.Type)
}

func (m *MsgHeader) checkFrom() error {
	if !kit.CheckAddress(config.Param.Name, m.From.String()) {
		return fmt.Errorf("%s address illegal", m.From.String())
	}
	return nil
}

func (m *MsgHeader) checkFee() error {
	if m.Fee < minFees {
		return fmt.Errorf("fee %.8f is less than the minimum poundage allowed %.8f", amount.Amount(m.Fee).ToCoin(), amount.Amount(minFees).ToCoin())
	}
	if m.Fee > maxFees {
		return fmt.Errorf("fee %.8f greater is greater than the maximum poundage allowed %.8f", amount.Amount(m.Fee).ToCoin(), amount.Amount(maxFees).ToCoin())
	}
	return nil
}

func (m *MsgHeader) checkSinger() error {
	if !Verify(m.Hash, m.Signature) {
		return errors.New("signature verification failed")
	}

	if !VerifySigner(config.Param.Name, m.From, m.Signature.PubKey) {
		return errors.New("signer and sender do not match")
	}
	return nil
}
