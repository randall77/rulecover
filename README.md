# rulecover
Tool for producing a cover profile of SSA .rules files

Example use:
```
cd $GOROOT
cd src/cmd/compile/internal/ssa/gen
go run *.go -log
cd ../../../../../
rm -f rulelog
./all.bash
rulecover cmd/compile/internal/ssa/gen/generic.rules rulelog
```
