// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.Package}}

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/darrenvechain/thorgo"
	"github.com/darrenvechain/thorgo/accounts"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = bind.Bind
	_ = common.Big1
	_ = abi.ConvertType
	_ = hexutil.MustDecode
	_ = context.Background
	_ = tx.NewClause
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

    {{if .InputBin}}
        // Deploy{{.Type}} deploys a new Ethereum contract, binding an instance of {{.Type}} to it.
        func Deploy{{.Type}}(thor *thorgo.Thor, sender accounts.TxManager{{range .Constructor.Inputs}}, {{.Name}} {{bindtype .Type $structs}}{{end}}) (common.Hash, *{{.Type}}Transactor, error) {
            parsed, err := {{.Type}}MetaData.GetAbi()
            if err != nil {
                return common.Hash{}, nil, err
            }
            if parsed == nil {
                return common.Hash{}, nil, errors.New("GetABI returned nil")
            }
            {{range $pattern, $name := .Libraries}}
                {{decapitalise $name}}Addr, _, _, _ := Deploy{{capitalise $name}}(auth, backend)
                {{$contract.Type}}MetaData.Bin = strings.ReplaceAll({{$contract.Type}}Bin, "__${{$pattern}}$__", {{decapitalise $name}}Addr.String()[2:])
            {{end}}
            bytes, err := hexutil.Decode({{.Type}}MetaData.Bin)
            if err != nil {
                return common.Hash{}, nil, err
            }
            contract, txID, err := thor.Deployer(bytes, parsed).Deploy(sender{{range .Constructor.Inputs}}, {{.Name}}{{end}})
            if err != nil {
                return common.Hash{}, nil, err
            }
            return txID, &{{.Type}}Transactor{&{{.Type}}{ thor: thor, contract: contract }, sender }, nil
        }
    {{end}}

	// {{.Type}} is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
	type {{.Type}} struct {
		thor     *thorgo.Thor // Thor connection to use
		contract *accounts.Contract // Generic contract wrapper for the low level calls
	}

	// {{.Type}}Transactor is an auto generated Go binding around an Ethereum, allowing you to transact with the contract.
    type {{.Type}}Transactor struct {
            *{{.Type}}
    		manager accounts.TxManager // TxManager to use
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

	// New{{.Type}}Transactor creates a new instance of {{.Type}}Transactor, bound to a specific deployed contract.
    func New{{.Type}}Transactor(address common.Address, thor *thorgo.Thor, manager accounts.TxManager) (*{{.Type}}Transactor, error) {
        base, err := New{{.Type}}(address, thor)
        if err != nil {
            return nil, err
        }
        return &{{.Type}}Transactor{ {{.Type}}: base, manager: manager }, nil
    }

	// Address returns the address of the contract.
	func (_{{$contract.Type}} *{{$contract.Type}}) Address() common.Address {
        return _{{$contract.Type}}.contract.Address
    }

	// Call invokes the (constant) contract method with params as input values and
	// sets the output to result. The result type might be a single field for simple
	// returns, a slice of interfaces for anonymous returns and a struct for named
	// returns.
	func (_{{$contract.Type}} *{{$contract.Type}}) Call(revision thorest.Revision, result *[]interface{}, method string, params ...interface{}) error {
		return _{{$contract.Type}}.contract.CallAt(revision, method, result, params...)
	}

	// Transact invokes the (paid) contract method with params as input values.
	func (_{{$contract.Type}}Transactor *{{$contract.Type}}Transactor) Transact(vetValue *big.Int, method string, params ...interface{}) (*transactions.Visitor, error) {
		return _{{$contract.Type}}Transactor.contract.SendWithVET(_{{$contract.Type}}Transactor.manager, vetValue, method, params...)
	}

	{{range .Calls}}
		// {{.Normalized.Name}} is a free data retrieval call binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) {{.Normalized.Name}}({{range .Normalized.Inputs}} {{.Name}} {{bindtype .Type $structs}}, {{end}} revision ...thorest.Revision) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
		    var rev thorest.Revision
		    if len(revision) > 0 {
                rev = revision[0]
            } else {
                rev = thorest.RevisionBest()
            }

			var out []interface{}
			err := _{{$contract.Type}}.Call(rev, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
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
		func (_{{$contract.Type}}Transactor *{{$contract.Type}}Transactor) {{.Normalized.Name}}({{range .Normalized.Inputs}} {{.Name}} {{bindtype .Type $structs}}, {{end}} vetValue ... *big.Int) (*transactions.Visitor, error) {
		    var val *big.Int
		    if len(vetValue) > 0 {
                val = vetValue[0]
            } else {
                val = big.NewInt(0)
            }
			return _{{$contract.Type}}Transactor.Transact( val, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
		}

		// {{.Normalized.Name}}AsClause is a transaction clause generator 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) {{.Normalized.Name}}AsClause({{range .Normalized.Inputs}}{{.Name}} {{bindtype .Type $structs}}, {{end}} vetValue ... *big.Int) (*tx.Clause, error) {
            var val *big.Int
            if len(vetValue) > 0 {
                val = vetValue[0]
            } else {
                val = big.NewInt(0)
            }
			return _{{$contract.Type}}.contract.AsClauseWithVET(val, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
		}
	{{end}}

	{{range .Events}}

		{{ $indexedArgCount := 0 }}
        {{ range .Normalized.Inputs }}
            {{- if .Indexed }}
                {{ $indexedArgCount = add $indexedArgCount 1 }}
            {{ end }}
        {{ end }}

        // {{$contract.Type}}{{.Normalized.Name}} represents a {{.Normalized.Name}} event raised by the {{$contract.Type}} contract.
		type {{$contract.Type}}{{.Normalized.Name}} struct { {{- range .Normalized.Inputs }}
			{{capitalise .Name}} {{if .Indexed}}{{bindtopictype .Type $structs}}{{else}}{{bindtype .Type $structs}}{{end}}{{- end }}
			Log thorest.EventLog
		}


        {{ if gt $indexedArgCount 0 }}
            type {{$contract.Type}}{{.Normalized.Name}}Criteria struct {
                {{- range .Normalized.Inputs }}
                    {{- if .Indexed }}
                        {{- $type := bindtype .Type $structs }}
                        {{capitalise .Name}} {{if (eq (slice $type 0 1) "*")}}{{$type}} `abi:"{{.Name}}"`{{else}}*{{$type}} `abi:"{{.Name}}"`{{end}}
                    {{- end }}{{- end }}
            }
        {{ end }}

		// Filter{{.Normalized.Name}} is a free log retrieval operation binding the contract event 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) Filter{{.Normalized.Name}}({{ if gt $indexedArgCount 0 }}criteria []{{$contract.Type}}{{.Normalized.Name}}Criteria, {{ end }}filters *thorest.LogFilters) ([]{{$contract.Type}}{{.Normalized.Name}}, error) {
			topicHash := _{{$contract.Type}}.contract.ABI.Events["{{.Normalized.Name}}"].ID

            {{ if gt $indexedArgCount 0 }}
                criteriaSet := make([]thorest.EventCriteria, len(criteria))
                for i, c := range criteria {
                    crteria := thorest.EventCriteria{
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
                    criteriaSet = append(criteriaSet, thorest.EventCriteria{
                        Address: &_{{$contract.Type}}.contract.Address,
                        Topic0: &topicHash,
                    })
                }
            {{ else }}
                criteriaSet := []thorest.EventCriteria{
                    thorest.EventCriteria{
                        Address: &_{{$contract.Type}}.contract.Address,
                        Topic0: &topicHash,
                    },
                }
            {{ end }}

            logs, err := _{{$contract.Type}}.thor.Client.FilterEvents(criteriaSet, filters)
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

        // Watch{{.Normalized.Name}} listens for on chain events binding the contract event 0x{{printf "%x" .Original.ID}}.
        //
        // Solidity: {{.Original.String}}
        func (_{{$contract.Type}} *{{$contract.Type}}) Watch{{.Normalized.Name}}({{ if gt $indexedArgCount 0 }}criteria []{{$contract.Type}}{{.Normalized.Name}}Criteria, {{ end }} ctx context.Context, bufferSize int) (chan *{{$contract.Type}}{{.Normalized.Name}}, error) {
            topicHash := _{{$contract.Type}}.contract.ABI.Events["{{.Normalized.Name}}"].ID

            {{ if gt $indexedArgCount 0 }}
            criteriaSet := make([]thorest.EventCriteria, len(criteria))
            {{ else }}
            criteriaSet := make([]thorest.EventCriteria, 0)
            {{ end }}

            {{ if gt $indexedArgCount 0 }}

                for i, c := range criteria {
                    crteria := thorest.EventCriteria{
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
            {{ end }}

            eventChan := make(chan *{{$contract.Type}}{{.Normalized.Name}}, bufferSize)
            blockSub := _{{$contract.Type}}.thor.Blocks.Subscribe(ctx, bufferSize)

            go func() {
                defer close(eventChan)

                for {
                    select {
                    case block := <-blockSub:
                        for _, tx := range block.Transactions {
                            for index, outputs := range tx.Outputs {
                                for _, event := range outputs.Events {
                                    if event.Address != _{{$contract.Type}}.contract.Address {
                                        continue
                                    }
                                    if topicHash != event.Topics[0] {
                                        continue
                                    }
                                    for _, c := range criteriaSet {
                                        if c.Topic1 != nil && *c.Topic1 != event.Topics[1] {
                                            continue
                                        }
                                        if c.Topic2 != nil && *c.Topic2 != event.Topics[2] {
                                            continue
                                        }
                                        if c.Topic3 != nil && *c.Topic3 != event.Topics[3] {
                                            continue
                                        }
                                        if c.Topic4 != nil && *c.Topic4 != event.Topics[4] {
                                            continue
                                        }
                                    }

                                    log := thorest.EventLog{
                                        Address: &_{{$contract.Type}}.contract.Address,
                                        Topics:  event.Topics,
                                        Data:    event.Data,
                                        Meta: thorest.LogMeta{
                                            BlockID:     block.ID,
                                            BlockNumber: block.Number,
                                            BlockTime:   block.Timestamp,
                                            TxID:        tx.ID,
                                            TxOrigin:    tx.Origin,
                                            ClauseIndex: int64(index),
                                        },
                                    }

                                    ev := new({{$contract.Type}}{{.Normalized.Name}})
                                    if err := _{{$contract.Type}}.contract.UnpackLog(ev, "{{.Normalized.Name}}", log); err != nil {
                                        continue
                                    }
                                    ev.Log = log
                                    eventChan <- ev
                                }
                            }
                        }
                    case <-ctx.Done():
                        return
                    }
                }
            }()

            return eventChan, nil
        }
	{{end}}
{{end}}
