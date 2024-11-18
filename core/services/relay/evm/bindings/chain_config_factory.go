package bindings

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

func NewChainReaderConfig() types.ChainReaderConfig {
	chainReaderConfig := types.ChainReaderConfig{
		Contracts: map[string]types.ChainContractReader{
			"ChainReaderTester": types.ChainContractReader{
				Configs: map[string]*types.ChainReaderDefinition{
					"GetAlterablePrimitiveValue": &types.ChainReaderDefinition{
						CacheEnabled:      false,
						ChainSpecificName: "getAlterablePrimitiveValue",
						ReadType:          0,
					}, "GetDifferentPrimitiveValue": &types.ChainReaderDefinition{
						CacheEnabled:      false,
						ChainSpecificName: "getDifferentPrimitiveValue",
						ReadType:          0,
					}, "GetElementAtIndex": &types.ChainReaderDefinition{
						CacheEnabled:      false,
						ChainSpecificName: "getElementAtIndex",
						ReadType:          0,
					}, "GetPrimitiveValue": &types.ChainReaderDefinition{
						CacheEnabled:      false,
						ChainSpecificName: "getPrimitiveValue",
						ReadType:          0,
					}, "GetSliceValue": &types.ChainReaderDefinition{
						CacheEnabled:      false,
						ChainSpecificName: "getSliceValue",
						ReadType:          0,
					}, "ReturnSeen": &types.ChainReaderDefinition{
						CacheEnabled:      false,
						ChainSpecificName: "returnSeen",
						ReadType:          0,
					},
				},
				ContractABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"StaticBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"nestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"nestedStaticStruct\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"struct AccountStruct\",\"name\":\"accountStruct\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"}],\"name\":\"Triggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fieldHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"}],\"name\":\"TriggeredEventWithDynamicTopic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field1\",\"type\":\"int32\"},{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field2\",\"type\":\"int32\"},{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field3\",\"type\":\"int32\"}],\"name\":\"TriggeredWithFourTopics\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"field1\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint8[32]\",\"name\":\"field2\",\"type\":\"uint8[32]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"field3\",\"type\":\"bytes32\"}],\"name\":\"TriggeredWithFourTopicsWithHashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"accountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"nestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"nestedStaticStruct\",\"type\":\"tuple\"}],\"name\":\"addTestStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAlterablePrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDifferentPrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getElementAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"Field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"DifferentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"OracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"OracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"AccountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"BigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"NestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"NestedStaticStruct\",\"type\":\"tuple\"}],\"internalType\":\"struct TestStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSliceValue\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"accountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"nestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"nestedStaticStruct\",\"type\":\"tuple\"}],\"name\":\"returnSeen\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"Field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"DifferentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"OracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"OracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"AccountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"BigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"NestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"NestedStaticStruct\",\"type\":\"tuple\"}],\"internalType\":\"struct TestStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"setAlterablePrimitiveValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"nestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"nestedStaticStruct\",\"type\":\"tuple\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"accountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"}],\"name\":\"triggerEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"}],\"name\":\"triggerEventWithDynamicTopic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"val1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"val2\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"val3\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"val4\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"val5\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"val6\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"val7\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"raw\",\"type\":\"bytes\"}],\"name\":\"triggerStaticBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field1\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"field2\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"field3\",\"type\":\"int32\"}],\"name\":\"triggerWithFourTopics\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field1\",\"type\":\"string\"},{\"internalType\":\"uint8[32]\",\"name\":\"field2\",\"type\":\"uint8[32]\"},{\"internalType\":\"bytes32\",\"name\":\"field3\",\"type\":\"bytes32\"}],\"name\":\"triggerWithFourTopicsWithHashed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
				ContractPollingFilter: types.ContractPollingFilter{
					PollingFilter: types.PollingFilter{
						LogsPerBlock: 0,
						MaxLogsKept:  0,
						Retention:    0,
					},
				},
			},
		},
	}
	return chainReaderConfig
}

func NewChainWriterConfig(maxGasPrice assets.Wei, defaultGasPrice uint64, fromAddress common.Address) types.ChainWriterConfig {
	chainWriterConfig := types.ChainWriterConfig{
		Contracts: map[string]*types.ContractConfig{
			"ChainReaderTester": &types.ContractConfig{
				Configs: map[string]*types.ChainWriterDefinition{
					"AddTestStruct": &types.ChainWriterDefinition{
						ChainSpecificName: "addTestStruct",
						Checker:           "simulate",
						FromAddress:       fromAddress,
						GasLimit:          0,
					}, "SetAlterablePrimitiveValue": &types.ChainWriterDefinition{
						ChainSpecificName: "setAlterablePrimitiveValue",
						Checker:           "simulate",
						FromAddress:       fromAddress,
						GasLimit:          0,
					}, "TriggerEvent": &types.ChainWriterDefinition{
						ChainSpecificName: "triggerEvent",
						Checker:           "simulate",
						FromAddress:       fromAddress,
						GasLimit:          0,
					}, "TriggerEventWithDynamicTopic": &types.ChainWriterDefinition{
						ChainSpecificName: "triggerEventWithDynamicTopic",
						Checker:           "simulate",
						FromAddress:       fromAddress,
						GasLimit:          0,
					}, "TriggerStaticBytes": &types.ChainWriterDefinition{
						ChainSpecificName: "triggerStaticBytes",
						Checker:           "simulate",
						FromAddress:       fromAddress,
						GasLimit:          0,
					}, "TriggerWithFourTopics": &types.ChainWriterDefinition{
						ChainSpecificName: "triggerWithFourTopics",
						Checker:           "simulate",
						FromAddress:       fromAddress,
						GasLimit:          0,
					}, "TriggerWithFourTopicsWithHashed": &types.ChainWriterDefinition{
						ChainSpecificName: "triggerWithFourTopicsWithHashed",
						Checker:           "simulate",
						FromAddress:       fromAddress,
						GasLimit:          0,
					},
				},
				ContractABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"StaticBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"nestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"nestedStaticStruct\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"struct AccountStruct\",\"name\":\"accountStruct\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"}],\"name\":\"Triggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fieldHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"}],\"name\":\"TriggeredEventWithDynamicTopic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field1\",\"type\":\"int32\"},{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field2\",\"type\":\"int32\"},{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field3\",\"type\":\"int32\"}],\"name\":\"TriggeredWithFourTopics\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"field1\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint8[32]\",\"name\":\"field2\",\"type\":\"uint8[32]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"field3\",\"type\":\"bytes32\"}],\"name\":\"TriggeredWithFourTopicsWithHashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"accountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"nestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"nestedStaticStruct\",\"type\":\"tuple\"}],\"name\":\"addTestStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAlterablePrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDifferentPrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getElementAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"Field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"DifferentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"OracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"OracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"AccountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"BigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"NestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"NestedStaticStruct\",\"type\":\"tuple\"}],\"internalType\":\"struct TestStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSliceValue\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"accountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"nestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"nestedStaticStruct\",\"type\":\"tuple\"}],\"name\":\"returnSeen\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"Field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"DifferentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"OracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"OracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"AccountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"BigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"NestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"NestedStaticStruct\",\"type\":\"tuple\"}],\"internalType\":\"struct TestStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"setAlterablePrimitiveValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"struct InnerDynamicTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelDynamicTestStruct\",\"name\":\"nestedDynamicStruct\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"address\",\"name\":\"A\",\"type\":\"address\"}],\"internalType\":\"struct InnerStaticTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"struct MidLevelStaticTestStruct\",\"name\":\"nestedStaticStruct\",\"type\":\"tuple\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"AccountStr\",\"type\":\"address\"}],\"internalType\":\"struct AccountStruct\",\"name\":\"accountStruct\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"}],\"name\":\"triggerEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"}],\"name\":\"triggerEventWithDynamicTopic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"val1\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"val2\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"val3\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"val4\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"val5\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"val6\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"val7\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"raw\",\"type\":\"bytes\"}],\"name\":\"triggerStaticBytes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field1\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"field2\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"field3\",\"type\":\"int32\"}],\"name\":\"triggerWithFourTopics\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field1\",\"type\":\"string\"},{\"internalType\":\"uint8[32]\",\"name\":\"field2\",\"type\":\"uint8[32]\"},{\"internalType\":\"bytes32\",\"name\":\"field3\",\"type\":\"bytes32\"}],\"name\":\"triggerWithFourTopicsWithHashed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
			},
		},
	}
	chainWriterConfig.MaxGasPrice = &maxGasPrice
	for _, contract := range chainWriterConfig.Contracts {
		for _, chainWriterDefinition := range contract.Configs {
			chainWriterDefinition.GasLimit = defaultGasPrice
		}
	}
	return chainWriterConfig
}