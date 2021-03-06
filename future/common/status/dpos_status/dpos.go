package dpos_status

import (
	"fmt"
	"github.com/Futuremine-chain/future/common/config"
	"github.com/Futuremine-chain/future/future/db/status/dpos_db"
	fmctypes "github.com/Futuremine-chain/future/future/types"
	"github.com/Futuremine-chain/future/tools/arry"
	"github.com/Futuremine-chain/future/types"
)

const dPosDB = "dpos_db"

type DPosStatus struct {
	db IDPosDB
}

func NewDPosStatus() (*DPosStatus, error) {
	db, err := dpos_db.Open(config.Param.Data + "/" + dPosDB)
	if err != nil {
		return nil, err
	}
	return &DPosStatus{db: db}, nil
}

func (d *DPosStatus) SetTrieRoot(hash arry.Hash) error {
	return d.db.SetRoot(hash)
}

func (d *DPosStatus) TrieRoot() arry.Hash {
	return d.db.Root()
}

func (d *DPosStatus) Commit() (arry.Hash, error) {
	return d.db.Commit()
}

// If the current number of candidates is less than or equal to the
// number of super nodes, it is not allowed to withdraw candidates.
func (d *DPosStatus) CheckMessage(msg types.IMessage) error {
	switch fmctypes.MessageType(msg.Type()) {
	case fmctypes.Cancel:
		if d.db.CandidatesCount() <= config.Param.SuperSize {
			return fmt.Errorf("candidate nodes are already in the minimum number. Cannot cancel the candidate status now, please wait")
		}
	}
	return nil
}

func (d *DPosStatus) CycleSupers(cycle uint64) (types.ICandidates, error) {
	return d.db.CycleSupers(cycle)
}

func (d *DPosStatus) SaveCycle(cycle uint64, supers types.ICandidates) {
	ss := supers.(*fmctypes.Supers)
	d.db.SaveCycle(cycle, ss)
}

func (d *DPosStatus) Candidates() (types.ICandidates, error) {
	return d.db.Candidates()
}

func (d *DPosStatus) Voters() map[arry.Address][]arry.Address {
	return d.db.Voters()
}

func (d *DPosStatus) Confirmed() (uint64, error) {
	return d.db.Confirmed()
}

func (d *DPosStatus) SetConfirmed(height uint64) {
	d.db.SetConfirmed(height)
}

func (d *DPosStatus) AddCandidate(msg types.IMessage) error {
	body := msg.MsgBody().(*fmctypes.CandidateBody)
	candidate := &fmctypes.Member{
		Signer: msg.From(),
		PeerId: body.Peer.String(),
		Weight: 0,
	}
	d.db.AddCandidate(candidate)
	d.db.Voter(msg.From(), msg.From())
	return nil
}

func (d *DPosStatus) CancelCandidate(msg types.IMessage) error {
	d.db.CancelCandidate(msg.From())
	return nil
}

func (d *DPosStatus) Voter(msg types.IMessage) error {
	d.db.Voter(msg.From(), msg.MsgBody().MsgTo())
	return nil
}

func (d *DPosStatus) AddSuperBlockCount(cycle uint64, signer arry.Address) {
	d.db.AddSuperBlockCount(cycle, signer)
}

func (d *DPosStatus) SuperBlockCount(cycle uint64, signer arry.Address) uint32 {
	return d.db.SuperBlockCount(cycle, signer)
}
