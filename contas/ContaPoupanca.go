package contas

import "puppy/clientes"

type ContaPoupanca struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo float64
}


func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque
	if podeSacar <= c.saldo && podeSacar > 0 {
		c.saldo -= valorDoSaque
		return "saque realizado"
	} else {
		return "nao foi possivel realizar o saque"
	}

}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	deposito := valorDeposito

	if deposito > 0 {
		c.saldo += deposito
		return "valor depositado", c.saldo
	} else {
		return "nao foi possivel fazer o deposito", c.saldo
	}

}

func (c *ContaPoupanca) ObterSaldo() float64{
	return c.saldo
}