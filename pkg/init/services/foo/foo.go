package foo

import (
	"fmt"

	"github.com/u-root/u-root/pkg/service"
	"github.com/u-root/u-root/pkg/service/onfail"
	"github.com/u-root/u-root/pkg/service/state"
)

// foo is a bar
type foo service.Unit

//New returns a service.Unit of foo
func New() service.Servicer {

	return &foo{
		Name:            "foo",
		Description:     "foo does all of the bar",
		Type:            service.Simple,
		OnFail:          onfail.Restart,
		Before:          []string{},
		After:           []string{},
		Requires:        []string{},
		EnvironmentFile: "",
		PIDFile:         "",
		State:           state.Unknown,
	}
}

func (f *foo) Unit() service.Unit {
	return service.Unit(*f)
}

func (f *foo) Start() error {
	fmt.Println("Hello world")

	return nil
}

func (f *foo) Stop() error {
	fmt.Println("Goodbye world")

	return nil
}

func (f *foo) Restart() error {
	f.Stop()
	f.Start()

	return nil
}

func (f *foo) Reload() error {
	f.Restart()

	return nil
}

func (f *foo) Status() error {
	fmt.Printf("Foo is %v", f.State)
	return nil

}
