// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.Package}}

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/darrenvechain/thorgo"
	"github.com/darrenvechain/thorgo/accounts"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/client"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = hexutil.MustDecode
)

{{$structs := .Structs}}
{{range $structs}}
	// {{.Name}} is an auto generated low-level Go binding around a user-defined struct.
	type {{.Name}} struct {
		{{range $field := .Fields}}
		{{$field.Name}} {{$field.Type}}
		{{end}}
	}
{{end}}

{{range $contract := .Contracts}}
	// {{.Type}}MetaData contains all meta data concerning the {{.Type}} contract.
	var {{.Type}}MetaData = &bind.MetaData{
		ABI: "{{.InputABI}}",
		{{if $contract.FuncSigs -}}
		Sigs: map[string]string{
			{{range $strsig, $binsig := .FuncSigs}}"{{$binsig}}": "{{$strsig}}",
			{{end}}
		},
		{{end -}}
		{{if .InputBin -}}
		Bin: "0x{{.InputBin}}",
		{{end}}
	}
	// {{.Type}}ABI is the input ABI used to generate the binding from.
	// Deprecated: Use {{.Type}}MetaData.ABI instead.
	var {{.Type}}ABI = {{.Type}}MetaData.ABI

	{{if $contract.FuncSigs}}
		// Deprecated: Use {{.Type}}MetaData.Sigs instead.
		// {{.Type}}FuncSigs maps the 4-byte function signature to its string representation.
		var {{.Type}}FuncSigs = {{.Type}}MetaData.Sigs
	{{end}}

    {{if .InputBin}}
        // {{.Type}}Bin is the compiled bytecode used for deploying new contracts.
        // Deprecated: Use {{.Type}}MetaData.Bin instead.
        var {{.Type}}Bin = {{.Type}}MetaData.Bin

        // Deploy{{.Type}} deploys a new Ethereum contract, binding an instance of {{.Type}} to it.
        func Deploy{{.Type}}(thor *thorgo.Thor, sender accounts.TxManager{{range .Constructor.Inputs}}, {{.Name}} {{bindtype .Type $structs}}{{end}}) (common.Hash, *{{.Type}}, error) {
            parsed, err := {{.Type}}MetaData.GetAbi()
            if err != nil {
                return common.Hash{}, nil, err
            }
            if parsed == nil {
                return common.Hash{}, nil, errors.New("GetABI returned nil")
            }

            {{range $pattern, $name := .Libraries}}
                {{decapitalise $name}}Addr, _, _, _ := Deploy{{capitalise $name}}(auth, backend)
                {{$contract.Type}}Bin = strings.ReplaceAll({{$contract.Type}}Bin, "__${{$pattern}}$__", {{decapitalise $name}}Addr.String()[2:])
            {{end}}

            bytes, err := hexutil.Decode({{.Type}}MetaData.Bin)
            if err != nil {
                return common.Hash{}, nil, err
            }


            contract, txID, err := thor.Deployer(bytes, parsed).Deploy(sender{{range .Constructor.Inputs}}, {{.Name}}{{end}})
            if err != nil {
                return common.Hash{}, nil, err
            }

            return txID, &{{.Type}}{
                thor:     thor,
                contract: contract,
            }, nil
        }
    {{end}}

	// {{.Type}} is an auto generated Go binding around an Ethereum contract.
	type {{.Type}} struct {
		thor     *thorgo.Thor // Thor connection to use
		contract *accounts.Contract // Generic contract wrapper for the low level calls
	}

	// New{{.Type}} creates a new instance of {{.Type}}, bound to a specific deployed contract.
	func New{{.Type}}(address common.Address, thor *thorgo.Thor) (*{{.Type}}, error) {
		parsed, err := {{.Type}}MetaData.GetAbi()
		if err != nil {
			return nil, err
		}
		contract := thor.Account(address).Contract(parsed)
		if err != nil {
			return nil, err
		}
		return &{{.Type}}{ thor: thor, contract: contract }, nil
	}

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_{{$contract.Type}} *{{$contract.Type}}) Call(result *[]interface{}, method string, params ...interface{}) error {
		return _{{$contract.Type}}.contract.Call(method, result, params)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_{{$contract.Type}} *{{$contract.Type}}) Transact(sender accounts.TxManager, method string, params ...interface{}) (*transactions.Visitor, error) {
		return _{{$contract.Type}}.contract.Send(sender, method, params)
	}

	{{range .Calls}}
		// {{.Normalized.Name}} is a free data retrieval call binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) {{.Normalized.Name}}(opts *bind.CallOpts {{range .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
			var out []interface{}
			err := _{{$contract.Type}}.Call(&out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
			{{if .Structured}}
			outstruct := new(struct{ {{range .Normalized.Outputs}} {{.Name}} {{bindtype .Type $structs}}; {{end}} })
			if err != nil {
				return *outstruct, err
			}
			{{range $i, $t := .Normalized.Outputs}}
			outstruct.{{.Name}} = *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}

			return *outstruct, err
			{{else}}
			if err != nil {
				return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
			}
			{{range $i, $t := .Normalized.Outputs}}
			out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}

			return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err {{end}}
		}
	{{end}}

	{{range .Transacts}}
		// {{.Normalized.Name}} is a paid mutator transaction binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) {{.Normalized.Name}}(sender accounts.TxManager {{range .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) (*transactions.Visitor, error) {
			return _{{$contract.Type}}.contract.Send(sender, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
		}

		// {{.Normalized.Name}}AsClause is a transaction clause generator 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) {{.Normalized.Name}}AsClause({{range .Normalized.Inputs}}{{.Name}} {{bindtype .Type $structs}}, {{end}}) (*tx.Clause, error) {
			return _{{$contract.Type}}.contract.AsClause("{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
		}
	{{end}}

	{{range .Events}}
        // {{$contract.Type}}{{.Normalized.Name}} represents a {{.Normalized.Name}} event raised by the {{$contract.Type}} contract.
		type {{$contract.Type}}{{.Normalized.Name}} struct { {{- range .Normalized.Inputs }}
			{{capitalise .Name}} {{if .Indexed}}{{bindtopictype .Type $structs}}{{else}}{{bindtype .Type $structs}}{{end}}{{- end }}
			Log client.EventLog
		}

        type {{$contract.Type}}{{.Normalized.Name}}Criteria struct {
            {{- range .Normalized.Inputs }}
                {{- if .Indexed }}
                    {{- $type := bindtype .Type $structs }}
                    {{capitalise .Name}} {{if (eq (slice $type 0 1) "*")}}{{$type}} `abi:"{{.Name}}"`{{else}}*{{$type}} `abi:"{{.Name}}"`{{end}}
                {{- end }}{{- end }}
        }

		// Filter{{.Normalized.Name}} is a free log retrieval operation binding the contract event 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) Filter{{.Normalized.Name}}(opts *client.FilterOptions, rang *client.FilterRange, criteria []{{$contract.Type}}{{.Normalized.Name}}Criteria) ([]{{$contract.Type}}{{.Normalized.Name}}, error) {
			topicHash := _{{$contract.Type}}.contract.ABI.Events["{{.Normalized.Name}}"].ID

			criteriaSet := make([]client.EventCriteria, len(criteria))
			for i, c := range criteria {
			    crteria := client.EventCriteria{
            		Address: &_{{$contract.Type}}.contract.Address,
            		Topic0:  &topicHash,
            	}
            	{{- range $index, $element := .Normalized.Inputs }}
                    {{- if .Indexed }}
                        if c.{{capitalise .Name}} != nil {
                            {{- $type := bindtype .Type $structs }}
                            {{- if (eq (slice $type 0 1) "*") }}
                                matcher := c.{{capitalise .Name}}
                            {{- else }}
                                matcher := *c.{{capitalise .Name}}
                            {{- end }}
                            topics, err := abi.MakeTopics([]interface{}{matcher})
                            if err != nil {
                            	return nil, err
                            }

                            {{- if eq $index 0}}
                                crteria.Topic1 = &topics[0][0]
                            {{- end}}
                            {{- if eq $index 1}}
                                crteria.Topic2 = &topics[0][0]
                            {{- end}}
                            {{- if eq $index 2}}
                                crteria.Topic3 = &topics[0][0]
                            {{- end}}
                            {{- if eq $index 3}}
                                crteria.Topic4 = &topics[0][0]
                            {{- end}}
                        }
                    {{- end }}
                {{- end }}

                criteriaSet[i] = crteria
			}

            if len(criteriaSet) == 0 {
                criteriaSet = append(criteriaSet, client.EventCriteria{
                    Address: &_{{$contract.Type}}.contract.Address,
                    Topic0: &topicHash, // Add Topic0 here
                })
            }

			filter := &client.EventFilter{
            		Range: rang,
            		Options: opts,
            		Criteria: &criteriaSet,
            }

            logs, err := _{{$contract.Type}}.thor.Client.FilterEvents(filter)
			if err != nil {
				return nil, err
			}

			inputs := _{{$contract.Type}}.contract.ABI.Events["{{.Normalized.Name}}"].Inputs
			var indexed abi.Arguments
			for _, arg := range inputs {
				if arg.Indexed {
					indexed = append(indexed, arg)
				}
			}

			events := make([]{{$contract.Type}}{{.Normalized.Name}}, len(logs))
			for i, log := range logs {
				event := new({{$contract.Type}}{{.Normalized.Name}})
                if err := _{{$contract.Type}}.contract.UnpackLog(event, "{{.Normalized.Name}}", log); err != nil {
                    return nil, err
                }
				event.Log = log
				events[i] = *event
			}

			return events, nil
		}
	{{end}}
{{end}}
