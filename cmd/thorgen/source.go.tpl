// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package {{.Package}}

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"
	"fmt"

	"github.com/darrenvechain/thorgo/contracts"
	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.Is
	_ = big.NewInt
	_ = strings.ReplaceAll
	_ = abi.ConvertType
	_ = hexutil.Decode
	_ = context.Background
	_ = tx.NewClause
	_ = blocks.New
	_ = time.Sleep
	_ = transactions.New
	_ = fmt.Errorf
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
        func Deploy{{.Type}}(ctx context.Context, thor *thorest.Client, sender contracts.TxManager, opts *transactions.Options{{range .Constructor.Inputs}}, {{.Name}} {{bindtype .Type $structs}}{{end}}) (common.Hash, *{{.Type}}, error) {
            parsed, err := {{.Type}}MetaData.GetAbi()
            if err != nil {
                return common.Hash{}, nil, err
            }
            {{range $pattern, $name := .Libraries}}
                {{decapitalise $name}}Addr, _, _, _ := Deploy{{capitalise $name}}(auth, backend)
                {{$contract.Type}}MetaData.Bin = strings.ReplaceAll({{$contract.Type}}Bin, "__${{$pattern}}$__", {{decapitalise $name}}Addr.String()[2:])
            {{end}}
            bytes, err := hexutil.Decode({{.Type}}MetaData.Bin)
            if err != nil {
                return common.Hash{}, nil, err
            }
            contract, txID, err := contracts.NewDeployer(thor, bytes, parsed).Deploy(ctx, sender, opts{{range .Constructor.Inputs}}, {{.Name}}{{end}})
            if err != nil {
                return common.Hash{}, nil, err
            }
            return txID, &{{.Type}}{thor: thor, contract: contract}, nil
        }
    {{end}}

	// {{.Type}} is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
	type {{.Type}} struct {
		thor     *thorest.Client // Thor client connection to use
		contract *contracts.Contract // Generic contract wrapper for the low level calls
	}

	// New{{.Type}} creates a new instance of {{.Type}}, bound to a specific deployed contract.
	func New{{.Type}}(address common.Address, thor *thorest.Client) (*{{.Type}}, error) {
		parsed, err := {{.Type}}MetaData.GetAbi()
		if err != nil {
			return nil, err
		}
		contract := contracts.New(thor, address, parsed)
		return &{{.Type}}{ thor: thor, contract: contract }, nil
	}

	// Address returns the address of the contract.
	func (_{{$contract.Type}} *{{$contract.Type}}) Address() common.Address {
        return _{{$contract.Type}}.contract.Address
    }

	// ==================== View Functions ====================

	{{range .Calls}}
		// {{.Normalized.Name}} is a free data retrieval call binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) {{.Normalized.Name}}({{range .Normalized.Inputs}}{{.Name}} {{bindtype .Type $structs}}, {{end}}) *{{$contract.Type}}{{.Normalized.Name}}Caller {
			return &{{$contract.Type}}{{.Normalized.Name}}Caller{caller: _{{$contract.Type}}.contract.Call("{{.Original.Name}}"{{range .Normalized.Inputs}}, {{.Name}}{{end}})}
		}
	{{end}}

	// ==================== Transaction Functions ====================

	{{range .Transacts}}
		// {{.Normalized.Name}} is a paid mutator transaction binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		{{- if eq .Normalized.StateMutability "payable" }}
        //
        // Setting the value in options is replaced by the vetValue argument.
        {{- end }}
		func (_{{$contract.Type}} *{{$contract.Type}}) {{.Normalized.Name}}({{range .Normalized.Inputs}} {{.Name}} {{bindtype .Type $structs}}, {{end}} {{- if eq .Normalized.StateMutability "payable" }}vetValue *big.Int, {{end}}) *contracts.Sender {
            {{- if eq .Normalized.StateMutability "payable" }}
            return contracts.NewSender(_{{$contract.Type}}.contract, "{{.Original.Name}}"{{range .Normalized.Inputs}}, {{.Name}}{{end}}).WithVET(vetValue)
            {{- else }}
            return contracts.NewSender(_{{$contract.Type}}.contract, "{{.Original.Name}}"{{range .Normalized.Inputs}}, {{.Name}}{{end}})
            {{- end }}
		}
	{{end}}

	// ==================== Event Functions ====================

	{{range .Events}}
		{{ $indexedArgCount := 0 }}
        {{ range .Normalized.Inputs }}
            {{- if .Indexed }}
                {{ $indexedArgCount = add $indexedArgCount 1 }}
            {{ end }}
        {{ end }}

		// Filter{{.Normalized.Name}} is a free log retrieval operation binding the contract event 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		func (_{{$contract.Type}} *{{$contract.Type}}) Filter{{.Normalized.Name}}({{ if gt $indexedArgCount 0 }}criteria []{{$contract.Type}}{{.Normalized.Name}}Criteria{{ end }}) *{{$contract.Type}}{{.Normalized.Name}}Filterer {
			filterer := _{{$contract.Type}}.contract.Filter("{{.Normalized.Name}}")
			
			{{ if gt $indexedArgCount 0 }}
			// Add criteria to the filterer
			for _, c := range criteria {
				eventCriteria := &contracts.EventCriteria{}
				{{- range $index, $element := .Normalized.Inputs }}
					{{- if .Indexed }}
						if c.{{capitalise .Name}} != nil {
							{{- $type := bindtype .Type $structs }}
							{{- if (eq (slice $type 0 1) "*") }}
								eventCriteria.Topic{{add $index 1}} = c.{{capitalise .Name}}
							{{- else }}
								eventCriteria.Topic{{add $index 1}} = *c.{{capitalise .Name}}
							{{- end }}
						}
					{{- end }}
				{{- end }}
				filterer.AddCriteria(eventCriteria)
			}
			{{ end }}
			
			return &{{$contract.Type}}{{.Normalized.Name}}Filterer{filterer: filterer, contract: _{{$contract.Type}}.contract}
		}
	{{end}}

	// ==================== Event Types and Criteria ====================

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
			Log *thorest.EventLog
		}

        {{ if gt $indexedArgCount 0 }}
            type {{$contract.Type}}{{.Normalized.Name}}Criteria struct {
                {{- range .Normalized.Inputs }}
                    {{- if .Indexed }}
                        {{- $type := bindtype .Type $structs }}
                        {{capitalise .Name}} {{if (eq (slice $type 0 1) "*")}}{{$type}}{{else}}*{{$type}}{{end}}
                    {{- end }}{{- end }}
            }
        {{ end }}
	{{end}}

	// ==================== Call Result Types ====================

	{{range .Calls}}
		{{if gt (len .Normalized.Outputs) 1}}
		// {{$contract.Type}}{{.Normalized.Name}}Result is a free data retrieval call binding the contract method 0x{{printf "%x" .Original.ID}}.
		//
		// Solidity: {{.Original.String}}
		type {{$contract.Type}}{{.Normalized.Name}}Result struct {
			{{range $i, $output := .Normalized.Outputs}}{{if .Name}}{{.Name}}{{else}}Result{{$i}}{{end}} {{bindtype .Type $structs}}
			{{end}}
		}
		{{end}}
	{{end}}

	// ==================== Caller Types and Methods ====================

	{{range .Calls}}
		// {{$contract.Type}}{{.Normalized.Name}}Caller provides typed access to the {{.Normalized.Name}} method
		type {{$contract.Type}}{{.Normalized.Name}}Caller struct {
			caller *contracts.Caller
		}

        // WithRevision sets the revision for the call to the contract method 0x{{printf "%x" .Original.ID}}.
		func (c *{{$contract.Type}}{{.Normalized.Name}}Caller) WithRevision(rev thorest.Revision) *{{$contract.Type}}{{.Normalized.Name}}Caller {
			c.caller.WithRevision(rev)
			return c
		}

        // Call executes the raw call to the contract method 0x{{printf "%x" .Original.ID}}.
		func (c *{{$contract.Type}}{{.Normalized.Name}}Caller) Call() (*thorest.InspectResponse, error) {
			return c.caller.Call()
		}

		{{if gt (len .Normalized.Outputs) 1}}
        // Execute executes the contract method 0x{{printf "%x" .Original.ID}} and returns the result.
		func (c *{{$contract.Type}}{{.Normalized.Name}}Caller) Execute() (*{{$contract.Type}}{{.Normalized.Name}}Result, error) {
			data, err := c.caller.Execute()
			if err != nil {
				return nil, err
			}
			if len(data) != {{len .Normalized.Outputs}} {
				return nil, errors.New("invalid number of return values")
			}
			out := new({{$contract.Type}}{{.Normalized.Name}}Result)
			{{range $i, $t := .Normalized.Outputs}}out.{{if .Name}}{{.Name}}{{else}}Result{{$i}}{{end}} = *abi.ConvertType(data[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}})
			{{end}}
			return out, nil
		}
		{{else}}
        // Execute executes the contract method 0x{{printf "%x" .Original.ID}} and returns the result.
		func (c *{{$contract.Type}}{{.Normalized.Name}}Caller) Execute() ({{range .Normalized.Outputs}}{{bindtype .Type $structs}}{{end}}, error) {
			data, err := c.caller.Execute()
			if err != nil {
				var zero {{range .Normalized.Outputs}}{{bindtype .Type $structs}}{{end}}
				return zero, err
			}
			if len(data) != 1 {
				var zero {{range .Normalized.Outputs}}{{bindtype .Type $structs}}{{end}}
				return zero, errors.New("expected single return value")
			}
			if result, ok := data[0].({{range .Normalized.Outputs}}{{bindtype .Type $structs}}{{end}}); ok {
				return result, nil
			}
			var zero {{range .Normalized.Outputs}}{{bindtype .Type $structs}}{{end}}
			return zero, fmt.Errorf("unexpected type returned: %T", data[0])
		}
		{{end}}
	{{end}}

	// ==================== Event Filterer Types and Methods ====================

	{{range .Events}}
		// {{$contract.Type}}{{.Normalized.Name}}Filterer provides typed access to filtering {{.Normalized.Name}} events
		type {{$contract.Type}}{{.Normalized.Name}}Filterer struct {
			filterer *contracts.Filterer
			contract *contracts.Contract
		}

		// Unit sets the range type for the filterer. It can be `block` or `time`
		func (f *{{$contract.Type}}{{.Normalized.Name}}Filterer) Unit(unit string) *{{$contract.Type}}{{.Normalized.Name}}Filterer {
            f.filterer.RangeUnit(unit)
            return f
        }

        // Range sets the range for the filterer. It can be a block range or a time range.
		func (f *{{$contract.Type}}{{.Normalized.Name}}Filterer) Range(from, to int64) *{{$contract.Type}}{{.Normalized.Name}}Filterer {
			f.filterer.Range(from, to)
			return f
		}

        // From sets the start time or block number for the filterer.
		func (f *{{$contract.Type}}{{.Normalized.Name}}Filterer) From(from int64) *{{$contract.Type}}{{.Normalized.Name}}Filterer {
			f.filterer.From(from)
			return f
		}

        // To sets the end time or block number for the filterer.
		func (f *{{$contract.Type}}{{.Normalized.Name}}Filterer) To(to int64) *{{$contract.Type}}{{.Normalized.Name}}Filterer {
			f.filterer.To(to)
			return f
		}

        // Offset sets the offset for the filterer, allowing you to skip a number of events.
		func (f *{{$contract.Type}}{{.Normalized.Name}}Filterer) Offset(offset int64) *{{$contract.Type}}{{.Normalized.Name}}Filterer {
			f.filterer.Offset(offset)
			return f
		}

        // Limit sets the maximum number of events to return.
		func (f *{{$contract.Type}}{{.Normalized.Name}}Filterer) Limit(limit int64) *{{$contract.Type}}{{.Normalized.Name}}Filterer {
			f.filterer.Limit(limit)
			return f
		}

        // Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
		func (f *{{$contract.Type}}{{.Normalized.Name}}Filterer) Order(order string) *{{$contract.Type}}{{.Normalized.Name}}Filterer {
			f.filterer.Order(order)
			return f
		}

        // Execute the query and return the events matching the filter criteria.
		func (f *{{$contract.Type}}{{.Normalized.Name}}Filterer) Execute() ([]{{$contract.Type}}{{.Normalized.Name}}, error) {
			logs, err := f.filterer.Execute()
			if err != nil {
				return nil, err
			}

			events := make([]{{$contract.Type}}{{.Normalized.Name}}, len(logs))
			for i, log := range logs {
				event := new({{$contract.Type}}{{.Normalized.Name}})
				if err := f.contract.UnpackLog(event, "{{.Normalized.Name}}", log); err != nil {
					return nil, err
				}
				event.Log = log
				events[i] = *event
			}

			return events, nil
		}
	{{end}}
{{end}}
