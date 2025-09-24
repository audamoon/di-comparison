package shared

import "fmt"

/*
	   ┌─────────────────┐
       │   Application   │
       │                 │
       └────────┬────────┘
			 	│
		┌───────┼─────┐
		│             │
		▼             ▼
   ┌───────────┐ ┌───────────┐
   │Dependency │ │Dependency │
   │   First   │ │  Second   │
   └─────┬─────┘ └─────┬─────┘
         │             │
         ▼             ▼
   ┌─────────────────────────┐
   │        Param            │
   │   ("some param")        │
   └─────────────────────────┘
*/

// Application is a dummy app
type Application struct {
	DependencyFirst  *SomeDependency
	DependencySecond *SomeDependency
}

func NewApplication(depFirst, depSecond *SomeDependency) *Application {
	return &Application{
		DependencyFirst:  depFirst,
		DependencySecond: depSecond,
	}
}

func (app *Application) Run() {
	app.DependencyFirst.Foo()
	app.DependencySecond.Foo()

	defer app.DependencyFirst.Boo()
	defer app.DependencySecond.Boo()
}

// SomeDependency is a dummy dependency
type SomeDependency struct {
	Param Param
}

func NewSomeDependency(param Param) *SomeDependency {
	return &SomeDependency{
		Param: param,
	}
}

func (d *SomeDependency) Foo() {
	fmt.Println("foo" + d.Param)
}

func (d *SomeDependency) Boo() {
	fmt.Println("boo" + d.Param)
}

// Param is kind of config

type Param string

func NewParam() Param {
	return "some param"
}
