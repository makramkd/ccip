package node

import (
	"bytes"
	_ "embed"
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/utils/ptr"

	"github.com/smartcontractkit/chainlink/integration-tests/types/config/node"
	itutils "github.com/smartcontractkit/chainlink/integration-tests/utils"
	evmcfg "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config/toml"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	"github.com/smartcontractkit/chainlink/v2/core/utils/config"
)

func NewConfigFromToml(tomlConfig []byte, opts ...node.NodeConfigOpt) (*chainlink.Config, error) {
	var cfg chainlink.Config
	err := config.DecodeTOML(bytes.NewReader(tomlConfig), &cfg)
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	return &cfg, nil
}

func WithPrivateEVMs(networks []blockchain.EVMNetwork, commonChainConfig *evmcfg.Chain, chainSpecificConfig map[int64]evmcfg.Chain) node.NodeConfigOpt {
	var evmConfigs []*evmcfg.EVMConfig
	for _, network := range networks {
		var evmNodes []*evmcfg.Node
		for i := range network.URLs {
			evmNodes = append(evmNodes, &evmcfg.Node{
				Name:    ptr.Ptr(fmt.Sprintf("%s-%d", network.Name, i)),
				WSURL:   itutils.MustURL(network.URLs[i]),
				HTTPURL: itutils.MustURL(network.HTTPURLs[i]),
			})
		}
		evmConfig := &evmcfg.EVMConfig{
			ChainID: utils.NewBig(big.NewInt(network.ChainID)),
			Nodes:   evmNodes,
		}
		if commonChainConfig != nil {
			evmConfig.Chain = *commonChainConfig
		}
		if chainSpecificConfig == nil {
			if overriddenChainCfg, ok := chainSpecificConfig[network.ChainID]; ok {
				evmConfig.Chain = overriddenChainCfg
			}
		}

		evmConfigs = append(evmConfigs, evmConfig)
	}
	return func(c *chainlink.Config) {
		c.EVM = evmConfigs
	}
}
