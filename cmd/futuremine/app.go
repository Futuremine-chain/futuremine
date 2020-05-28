package main

const (
	// Mainnet logo
	MainNet = "MN"
	// Testnet logo
	TestNet = "TN"
	// Token name
	Token = "FMC"
)

var (
	// Program name
	AppName = "future-mine-chain"
	// Current network
	NetWork = MainNet
	// Network logo
	P2pNetWork = "_FUTURE_MINE_CHAIN"

	Version = "0.0.0"
)

const (
	// Maximum amount limit for second-level tokens
	MaxCoinCount = 90000000000
	// The minimum amount limit of the second-level token
	MinCoinCount = 1000
)

type FMC struct {
}

func (F *FMC) AppName() string {
	return AppName
}

func (F *FMC) Version() string {
	return Version
}

func (F *FMC) NetWork() string {
	return NetWork
}

func (F *FMC) P2pNetWork() string {
	return P2pNetWork
}

func (F *FMC) TestNet() string {
	return TestNet
}

func (F *FMC) MainNet() string {
	return MainNet
}

func (F *FMC) InitTestNet() {
	NetWork = TestNet
}

func (F *FMC) InitP2pNet() {
	P2pNetWork = NetWork + P2pNetWork
}