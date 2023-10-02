package entity

// Trabalhando para que os dados fiquem mais performáticos
// Através da HEAP, para usar, preciso implementar alguns metodos simples
// Até aqui é 1:16 da primeira aula, posso ver atrás pra comentar os demais arquivos.

type OrderQueue struct {
	Orders []*Order
}

// Pra eu trabalhar de forma performática sem precisar de métodos malucos 
// pra trabalhar baixo nivel na alocação e tudo mais, vamos usar os metodos:

// LESS -> Conta 2 valores ( > e <)
func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// SWAWP -> Inverte os dados
func (oq *OrderQueue) Swap(i, j, int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// LEN -> Tamanho dos dados que tenho
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// PUSH -> Adiciona novos caras ali com append
func (oq *OrderQueue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// POP -> Remove
func (oq *OrderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
