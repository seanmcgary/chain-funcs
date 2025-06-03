//go:generate abigen --abi ../contracts/FaaS.json --pkg bindings --type FaaS --out faas.go
//go:generate abigen --abi ../contracts/TaskMailbox.json --pkg bindings --type TaskMailbox --out taskmailbox.go

package bindings