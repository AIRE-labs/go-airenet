package launcher

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	Bootnodes = []string{
		"enode://03c70d4597d731ef182678b7664f2a4a3add07056f23d4e01aba86f066080d18fa13abbd2e13e9d4ea762a2715a983b5ac6151162d05ee0434f1847da1a626e9@34.242.220.16:5050",
		"enode://01c64d1a9dd8a65c56f2d4e373795eb6efd27b714b2b5999363a42a0edc39d7417a431416ceb5c67b1a170983af109e8a15d0c2d44a2ac41ecfb5c23c1a1a48a@3.35.200.210:5050",
		"enode://7044c88daa5df059e2f7a2667471a8149a5cf66e68643dcb86f399d48c4ff6475b73ee91486ea830d225f7f78a2fdf955208673da51c6852230c3a90a3701c06@3.1.103.70:5050",
		"enode://594d26c2338566daca9391d73f1b1821bb0b454e6f3d48715116bf42f320924d569534c143b640feec8a8eaa137a0b822426fb62b52a90162270ea5868bdc37c@18.138.254.181:5050",
		"enode://339e331912e5239a9e13eb82b47be58ea4d3946e91caa2992103a8d4f0226c1e86f9134822d5b238f25c9cbdd473f806caa8e4f8ef1748a6c66395f4bf0dd569@54.66.206.151:5050",
		"enode://563b30428f48357f31c9d4906ca2f3d3815d663b151302c1ba9d58f3428265b554398c6fabf4b806a49525670cd9e031257c805375b9fdbcc015f60a7943e427@3.213.142.230:7946",
		"enode://8b53fe4410cde82d98d28697d56ccb793f9a67b1f8807c523eadafe96339d6e56bc82c0e702757ac5010972e966761b1abecb4935d9a86a9feed47e3e9ba27a6@3.227.34.226:7946",
		"enode://1703640d1239434dcaf010541cafeeb3c4c707be9098954c50aa705f6e97e2d0273671df13f6e447563e7d3a7c7ffc88de48318d8a3cc2cc59d196516054f17e@52.72.222.228:7946",
	}
)

func overrideParams() {
	params.MainnetBootnodes = []string{}
	params.RopstenBootnodes = []string{}
	params.RinkebyBootnodes = []string{}
	params.GoerliBootnodes = []string{}
}