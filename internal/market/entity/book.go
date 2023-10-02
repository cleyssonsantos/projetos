package entity

import (
	"container/heap"
	"sync"
)

type Book struct {
	Order         []*Order
	Transaction   []*Transaction
	OrdersChan    chan *Order // Aqui vamos receber as ordens do Kafka
	OrdersChanOut chan *Order
	Wg            *sync.WaitGroup
}

func NewBook (orderChan chan *Order, orderChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book {
		Order: []*Order{},
		Transaction: []*Transaction{},
		OrdersChan: orderChan,
		OrdersChanOut: orderChanOut,
		Wg: wg,
	}
}

// Agora vem a complexidade do algoritmo, que é a parte das transações
// onde vai rodar em loop infinito tentando dar match nas transações.

func (b *Book) Trade() {
	buyOrders := NewOrderQueue()
	sellOrders := NewOrderQueue()

	heap.Init(buyOrders)
	heap.Init(sellOrders)

	for order := range b.OrdersChan {
		if order.OrderType == "BUY" {
			buyOrders.Push(order)
			if sellOrders.Len() > 0 && sellOrders.Orders[0].Price <= order.Price {
				sellOrder := sellOrders.Pop().(*Order)
				if sellOrder.PendingShares > 0 {
					transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
					b.AddTransaction(transaction, b.Wg)
					sellOrder.Transactions = append(sellOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)
					b.OrdersChanOut <- sellOrder // Pego o valor de sell order e jogo pro canal para que o Kafka consiga publicar isso
					b.OrdersChanOut <- order
					if sellOrder.PendingShares > 0 {
						sellOrders.Push(sellOrder)
					}
				}
			}
		} else if order.OrderType == "SELL" { // É o mesmo caso do BUY mas invertido, depois posso refatorar isso aqui
			sellOrders.Push(order)
			if buyOrders.Len() > 0 && buyOrders.Orders[0].Price >= order.Price {
				buyOrder := buyOrders.Pop().(*Order)
				if buyOrder.PendingShares > 0 {
					transaction := NewTransaction(order, buyOrder, order.Shares, buyOrder.Price)
					b.AddTransaction(transaction, b.Wg)
					buyOrder.Transactions = append(buyOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)
					b.OrdersChanOut <- buyOrder
					b.OrdersChanOut <- order
					if buyOrder.PendingShares > 0 {
						buyOrders.Push(buyOrder)
					}
				}
			}
		}
	}
}

func (b *Book) AddTransaction (transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done() // Esse comando executa todo nosso codigo, e essa linha por ultimo. Muito util, e resolveu um problemão na hr de testar...
	
	sellingShares := transaction.SellingOrder.PendingShares
	buyingShares := transaction.BuyingOrder.PendingShares

	minShares := sellingShares
	if buyingShares < minShares {
		minShares = buyingShares
	}

	transaction.SellingOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.ID, -minShares) // Trabalhar com ponteiro é bom porque agora eu consigo mexer nele in memory...
	transaction.AddSellOrderPendingShares(-minShares)

	transaction.BuyingOrder.Investor.UpdateAssetPosition(transaction.BuyingOrder.Asset.ID, minShares)
	transaction.AddBuyOrderPendingShares(-minShares)

	transaction.CalculateTotal(transaction.Shares, transaction.BuyingOrder.Price)

	transaction.CloseBuyOrder()
	transaction.CloseSellOrder()
	
	b.Transactions = append(b.Transactions, transaction)
}
