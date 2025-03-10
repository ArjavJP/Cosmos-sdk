package types

import (
	"fmt"

	codectypes "github.com/ArjavJP/Cosmos-sdk/codec/types"
	gov "github.com/ArjavJP/Cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeSoftwareUpgrade       string = "SoftwareUpgrade"
	ProposalTypeCancelSoftwareUpgrade string = "CancelSoftwareUpgrade"
)

func NewSoftwareUpgradeProposal(title, description string, plan Plan) gov.Content {
	return &SoftwareUpgradeProposal{title, description, plan}
}

// Implements Proposal Interface
var _ gov.Content = &SoftwareUpgradeProposal{}
var _ codectypes.UnpackInterfacesMessage = SoftwareUpgradeProposal{}

func init() {
	gov.RegisterProposalType(ProposalTypeSoftwareUpgrade)
	gov.RegisterProposalTypeCodec(&SoftwareUpgradeProposal{}, "cosmos-sdk/SoftwareUpgradeProposal")
	gov.RegisterProposalType(ProposalTypeCancelSoftwareUpgrade)
	gov.RegisterProposalTypeCodec(&CancelSoftwareUpgradeProposal{}, "cosmos-sdk/CancelSoftwareUpgradeProposal")
}

func (sup *SoftwareUpgradeProposal) GetTitle() string       { return sup.Title }
func (sup *SoftwareUpgradeProposal) GetDescription() string { return sup.Description }
func (sup *SoftwareUpgradeProposal) ProposalRoute() string  { return RouterKey }
func (sup *SoftwareUpgradeProposal) ProposalType() string   { return ProposalTypeSoftwareUpgrade }
func (sup *SoftwareUpgradeProposal) ValidateBasic() error {
	if err := sup.Plan.ValidateBasic(); err != nil {
		return err
	}
	return gov.ValidateAbstract(sup)
}

func (sup SoftwareUpgradeProposal) String() string {
	return fmt.Sprintf(`Software Upgrade Proposal:
  Title:       %s
  Description: %s
`, sup.Title, sup.Description)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (sup SoftwareUpgradeProposal) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return sup.Plan.UnpackInterfaces(unpacker)
}

func NewCancelSoftwareUpgradeProposal(title, description string) gov.Content {
	return &CancelSoftwareUpgradeProposal{title, description}
}

// Implements Proposal Interface
var _ gov.Content = &CancelSoftwareUpgradeProposal{}

func (csup *CancelSoftwareUpgradeProposal) GetTitle() string       { return csup.Title }
func (csup *CancelSoftwareUpgradeProposal) GetDescription() string { return csup.Description }
func (csup *CancelSoftwareUpgradeProposal) ProposalRoute() string  { return RouterKey }
func (csup *CancelSoftwareUpgradeProposal) ProposalType() string {
	return ProposalTypeCancelSoftwareUpgrade
}
func (csup *CancelSoftwareUpgradeProposal) ValidateBasic() error {
	return gov.ValidateAbstract(csup)
}

func (csup CancelSoftwareUpgradeProposal) String() string {
	return fmt.Sprintf(`Cancel Software Upgrade Proposal:
  Title:       %s
  Description: %s
`, csup.Title, csup.Description)
}
