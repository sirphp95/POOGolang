package main

import (
	"fmt"
	
	"puppy/contas"
	//botando uma palavra na frente do import eu gero um apelido exemplo poc "ppu/contas"
)
func  PagarBoleto(conta verificarConta, valordoBoleto float64) {
	conta.Sacar(valordoBoleto)
}

type verificarConta interface{
Sacar(valor float64) string
}

func main() {
	contaReno := contas.ContaCorrente{}
	contaPoc := contas.ContaPoupanca{}

	contaReno.Depositar(100)
	contaReno.Sacar(50)
	PagarBoleto(&contaPoc, 20)

	fmt.Println(contaPoc, contaReno)

}

// 	clientesa := contas.ContaCorrente{ }
// clientesa.Depositar(100)

// fmt.Println(clientesa.ObterSaldo())

// contaPoc := contas.ContaCorrente{
// 	Titular:       clientes.Titular{"pedroo","aaa", "aa" },
// 	NumeroAgencia: 0,
// 	NumeroConta:   0,
// 	Saldo:         0,

// contaPoc := contas.ContaCorrente{
// 	Titular:       "Poc" ,
// 	NumeroAgencia: 0,
// 	NumeroConta:   0,
// 	Saldo:         20,
// }
// // contaPoc.Titular = "poc"
// // contaPoc.Saldo = 111

// contaReno := contas.ContaCorrente {}
// contaReno.Titular = "reno"
// contaReno.Saldo = 200

// status := contaPoc.Transferir(500, &contaReno)

// fmt.Println(status)
// fmt.Println(contaPoc.Saldo)

// status, valor := contaPoc.depositar(20)

// fmt.Println(status, valor)

//  contaPoc :=  ContaCorrente{titular:"poc", numeroAgencia: 123, numeroConta: 332, saldo: 12.2}
// contaReno := ContaCorrente{"reno", 22, 11222, 20.0}
// contaLulu := new(ContaCorrente)
// contaLulu.titular = "lulu"
// fmt.Println(contaPoc, "---", contaReno, "---", contaLulu)

//go nao usa classe usa estrutura
