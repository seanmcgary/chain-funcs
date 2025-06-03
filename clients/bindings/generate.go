//go:generate sh -c "cat ../contracts/FaaS.json | jq .abi | abigen --abi - --pkg bindings --type FaaS --out faas.go"
//go:generate sh -c "cat ../contracts/TaskMailbox.json | jq .abi | abigen --abi - --pkg bindings --type TaskMailbox --out taskmailbox.go"

package bindings