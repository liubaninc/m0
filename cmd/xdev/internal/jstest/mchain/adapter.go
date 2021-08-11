package mchain

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/robertkrimen/otto"

	"github.com/liubaninc/m0/cmd/xdev/internal/jstest"
)

type contractObject struct {
	env  *environment
	Name string
}

// func (c *contractObject) Invoke(method string, args map[string]string, option InvokeOptions) *contract.Response {
func (c *contractObject) Invoke(call otto.FunctionCall) otto.Value {
	var args invokeArgs

	method := call.Argument(0).String()
	args.Method = method

	if !call.Argument(1).IsObject() {
		jstest.Throws("expect method args with object type")
	}
	export, _ := call.Argument(1).Export()
	err := mapstructure.Decode(export, &args.Args)
	if err != nil {
		jstest.Throw(err)
	}

	if call.Argument(2).IsObject() {
		export, _ := call.Argument(2).Export()
		err := mapstructure.Decode(export, &args.Options)
		if err != nil {
			jstest.Throw(err)
		}
	}

	resp, err := c.env.Invoke(c.Name, args)
	if err != nil {
		jstest.Throw(err)
	}
	v, err := call.Otto.ToValue(resp)
	if err != nil {
		jstest.Throw(err)
	}
	return v
}

type mchainObject struct {
	env *environment
}

func newXchainObject() (*mchainObject, error) {
	env, err := newEnvironment()
	if err != nil {
		return nil, err
	}
	return &mchainObject{
		env: env,
	}, nil
}

func (x *mchainObject) Contract(name string) *contractObject {
	if !x.env.ContractExists(name) {
		jstest.Throw(fmt.Errorf("contract %s not found", name))
	}
	return &contractObject{
		Name: name,
		env:  x.env,
	}
}

func (x *mchainObject) Deploy(args deployArgs) *contractObject {
	_, err := x.env.Deploy(args)
	if err != nil {
		jstest.Throw(err)
	}
	return &contractObject{
		env:  x.env,
		Name: args.Name,
	}
}

type mchainAdapter struct {
}

// NewAdapter is the mchain adapter
func NewAdapter() jstest.Adapter {
	return new(mchainAdapter)
}

func (x *mchainAdapter) OnSetup(r *jstest.Runner) {
	r.GlobalObject().Set("Xchain", func() *mchainObject {
		x, err := newXchainObject()
		if err != nil {
			jstest.Throw(err)
		}
		return x
	})
}

func (x *mchainAdapter) OnTeardown(r *jstest.Runner) {
}

func (x *mchainAdapter) OnTestCase(r *jstest.Runner, test jstest.TestCase) jstest.TestCase {
	body := func(t *testing.T) {
		xctx, err := newXchainObject()
		if err != nil {
			t.Fatal(err)
		}
		defer xctx.env.Close()

		if !r.Option.Quiet {
			// TODO: add log output
		}
		// reset mchain environment
		r.GlobalObject().Set("mchain", xctx)

		test.F(t)
	}
	return jstest.TestCase{
		Name: test.Name,
		F:    body,
	}
}
