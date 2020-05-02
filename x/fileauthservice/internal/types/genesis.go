package types

import "fmt"

// GenesisState - all fileauthservice state that must be provided at genesis
type GenesisState struct {
	// TODO: Fill out what is needed by the module for genesis
	FileAuthRecords []Fileauth `json:"fileauth_record"`
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(FileAuthRecords []Fileauth) GenesisState {
	return GenesisState{FileAuthRecords: nil}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() GenesisState {
	return GenesisState{
		// TODO: Fill out according to your genesis state, these values will be initialized but empty
		FileAuthRecords: []Fileauth{},
	}
}

// ValidateGenesis validates the fileauthservice genesis parameters
func ValidateGenesis(data GenesisState) error {
	// TODO: Create a sanity check to make sure the state conforms to the modules needs
	for _, record := range data.FileAuthRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid fileauth, file name: %s. Error: Missing Owner", record.Name)
		}
		if record.Name == "" {
			return fmt.Errorf("invalid fileauth, file name: %s. Error: Missing Name", record.Name)
		}
		if record.Hash == "" {
			return fmt.Errorf("invalid fileauth, file name: %s. Error: Missing Hash", record.Name)
		}
		if record.Auth == "" {
			return fmt.Errorf("invalid fileauth, file name: %s. Error: Missing Auth", record.Name)
		}
	}
	return nil
}
