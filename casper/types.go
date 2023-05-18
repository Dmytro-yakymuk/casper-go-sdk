package casper

import (
	"github.com/make-software/casper-go-sdk/types"
	"github.com/make-software/casper-go-sdk/types/clvalue"
	"github.com/make-software/casper-go-sdk/types/clvalue/cltype"
)

type (
	AuctionState    = types.AuctionState
	Args            = types.Args
	ValidatorBid    = types.ValidatorBid
	AuctionBid      = types.AuctionBid
	Block           = types.Block
	BlockHeader     = types.BlockHeader
	BlockBody       = types.BlockBody
	ContractPackage = types.ContractPackage
	ContractVersion = types.ContractVersion
	Contract        = types.Contract
	Deploy          = types.Deploy
	DeployHeader    = types.DeployHeader
	DeployApproval  = types.Approval
	Entrypoint      = types.EntryPoint
	ExecutionResult = types.ExecutionResultStatus
	NamedKeys       = types.NamedKeys
	NamedKey        = types.NamedKey
	TransformKey    = types.TransformKey
	Transform       = types.Transform
	Argument        = types.Argument
	Account         = types.Account
	Reward          = types.Reward
	WriteTransfer   = types.WriteTransfer
	UnbondingPurse  = types.UnbondingPurse
)

type (
	ExecutableDeployItem = types.ExecutableDeployItem
	ModuleBytes          = types.ModuleBytes
)

var (
	DefaultHeader   = types.DefaultHeader
	MakeDeploy      = types.MakeDeploy
	StandardPayment = types.StandardPayment
)

type (
	CLType  = cltype.CLType
	CLValue = clvalue.CLValue
	CLMap   = clvalue.Map
)

type (
	StoredContractByHash          = types.StoredContractByHash
	StoredContractByName          = types.StoredContractByName
	StoredValue                   = types.StoredValue
	StoredVersionContractByHash   = types.StoredVersionContractByHash
	StoredVersionedContractByName = types.StoredVersionedContractByName
)

var (
	ErrArgumentNotFound = types.ErrArgumentNotFound
)
