package contas

import ( "fmt"
"puppy/clientes"
)


type ContaCorrente struct {
	Titular       clientes.Titular 
	NumeroAgencia, NumeroConta int
	 
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque
	if podeSacar <= c.saldo && podeSacar > 0 {
		c.saldo -= valorDoSaque
		return "saque realizado"
	} else {
		return "nao foi possivel realizar o saque"
	}

}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	deposito := valorDeposito

	if deposito > 0 {
		c.saldo += deposito
		return "valor depositado", c.saldo
	} else {
		return "nao foi possivel fazer o deposito", c.saldo
	}

}

func (c *ContaCorrente) Transferir(valorDaTrans float64, contaDestino *ContaCorrente) bool {
	if valorDaTrans < c.saldo && valorDaTrans > 0 {
		c.saldo -= valorDaTrans
		contaDestino.Depositar(valorDaTrans)
		return true
	} else {

		fmt.Println("nao foi possivel transferir")
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64{
	return c.saldo
}
//tem que por em minusculo o nome das classes para poder usar em outro arquivo
