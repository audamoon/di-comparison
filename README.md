# DI comparison

## Benchmarks

| Container         | Operations | ns/op   | B/op  | allocs/op |
|-------------------|------------|---------|-------|-----------|
| WireContainer     | 42416953   | 28.2    | 48    | 3         |
| GoLobbyContainer  | 1628104    | 729.2   | 856   | 18        |
| DoContainer       | 1421836    | 836.1   | 1376  | 35        |
| SarulabsContainer | 896241     | 1329.0  | 4560  | 45        |
| InDriveContainer  | 530215     | 2259.4  | 1112  | 36        |
| DigContainer      | 73129      | 14722.0 | 21246 | 242       |
| FxContainer       | 34974      | 33616.6 | 44648 | 653       |

## Output error messages

### Dig

```shell
2025/09/25 10:47:04 provide: missing dependencies for function "di-comparison/ioc/dig/builders".
BuildApplicationDig.func1 (/Users/audamoon/projects/go/personal/di-presentation/ioc/dig/builders/app.go:13):
missing type: *shared.SomeDependency[name="the-second"]
```

### Fx

```shell
[Fx] ERROR		fx.Invoke(di-comparison/ioc/fx/builders.StartApplication()) called from:
di-comparison/ioc/fx/builders.BuildFxApplication
	/Users/audamoon/projects/go/personal/di-presentation/ioc/fx/builders/app.go:23
main.main
	/Users/audamoon/projects/go/personal/di-presentation/ioc/fx/main.go:10
runtime.main
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:285
Failed: missing dependencies for function "di-comparison/ioc/fx/builders".StartApplication
	/Users/audamoon/projects/go/personal/di-presentation/ioc/fx/builders/app.go:33:
missing type:
	- *shared.SomeDependency[name="the-second"] (did you mean to Provide it?)
[Fx] ERROR		Failed to start: missing dependencies for function "di-comparison/ioc/fx/builders".StartApplication
	/Users/audamoon/projects/go/personal/di-presentation/ioc/fx/builders/app.go:33:
missing type:
	- *shared.SomeDependency[name="the-second"] (did you mean to Provide it?)
```

### Golobby

```shell
2025/09/25 10:50:50 container: no concrete found for: *shared.SomeDependency
```

### sarulabs

```shell
panic: could not build `first-dep` because the build function panicked: could not get `param` because the definition does not exist

goroutine 1 [running]:
github.com/sarulabs/di/v2.Container.Get({0x14000156240?, {0x1400014c0a0?, 0x0?, 0x0?}}, {0x102724880?, 0x102761b98?})
	/Users/audamoon/go/pkg/mod/github.com/sarulabs/di/v2@v2.5.2/containerGetter.go:156 +0xa68
di-comparison/ioc/sarulabs/builders.BuildApplicationSarulabs({0x14000156240?, {0x1400014c0a0?, 0x0?, 0x0?}})
	/Users/audamoon/projects/go/personal/di-presentation/ioc/sarulabs/builders/app.go:52 +0x44
main.main()
	/Users/audamoon/projects/go/personal/di-presentation/ioc/sarulabs/main.go:14 +0x64

```

### Wire

```shell
wire: /Users/audamoon/projects/go/personal/di-presentation/ioc/wire/builders/app.go:29:1: inject BuildWireApplication: no provider found for di-comparison/ioc/wire/builders.FirstDep
	needed by *di-comparison/shared.Application in provider "NewApplication" (/Users/audamoon/projects/go/personal/di-presentation/ioc/wire/builders/app.go:25:6)
wire: di-comparison/ioc/wire/builders: generate failed
wire: at least one generate failure
exit status 1
wire_gen.go:3: running "go": exit status 1

### Do 

```shell
panic: DI: could not find service `second-dep`, available services: `first-dep`, `*shared.Application`, `shared.Param`

goroutine 1 [running]:
github.com/samber/do.must(...)
	/Users/audamoon/go/pkg/mod/github.com/samber/do@v1.6.0/utils.go:9
github.com/samber/do.MustInvoke[...](0x14000134050?)
	/Users/audamoon/go/pkg/mod/github.com/samber/do@v1.6.0/di.go:106 +0x40
main.main()
	/Users/audamoon/projects/go/personal/di-presentation/ioc/do/main.go:20 +0x94

```