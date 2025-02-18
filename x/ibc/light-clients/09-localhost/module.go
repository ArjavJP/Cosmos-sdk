package localhost

import (
	"github.com/ArjavJP/Cosmos-sdk/x/ibc/light-clients/09-localhost/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
