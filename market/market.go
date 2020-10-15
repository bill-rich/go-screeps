package market

import (
	"github.com/bill-rich/go-screeps/common"
	"github.com/bill-rich/go-screeps/room"
	"github.com/gopherjs/gopherjs/js"
)

// Market is a global object representing the in-game market.
type Market struct {
	*js.Object
	Credits              int              `js:"credits"`
	IncomingTransactions []Transaction    `js:"incomingTransactions"`
	OutgoingTransactions []Transaction    `js:"outgoingTransactions"`
	Orders               map[string]Order `js:"orders"`
}

// Transaction is a record of market transactions.
type Transaction struct {
	*js.Object
	TransactionID string       `js:"transactionId"`
	Time          int          `js:"time"`
	Sender        common.Owner `js:"sender"`
	Recipient     common.Owner `js:"recipient"`
	ResourceType  string       `js:"resourceType"`
	Amount        int          `js:"amount"`
	From          string       `js:"from"`
	To            string       `js:"to"`
	Description   string       `js:"description"`
	Order         Order        `js:"order"`
}

// Order is a buy/sell order on the market.
type Order struct {
	*js.Object
	ID              string  `js:"id"`
	Created         int     `js:"created"`
	Active          bool    `js:"acitve"`
	Type            string  `js:"type"`
	ResourceType    string  `js:"resourceType"`
	RoomName        string  `js:"roomName"`
	Amount          int     `js:"amount"`
	RemainingAmount int     `js:"remainingAmount"`
	TotalAmount     int     `js:"totalAmount"`
	Price           float32 `js:"price"`
}

// ResourceHistory contains resource history data.
type ResourceHistory struct {
	*js.Object
	ResourceType string  `js:"resourceType"`
	Date         string  `js:"date"`
	Transactions int     `js:"transactions"`
	Volume       int     `js:"volume"`
	AveragePrice float32 `js:"avgPrice"`
	StdDevPrice  float32 `js:"stddevPrice"`
}

// CalcTransactionCost estimates the energy transaction cost of StructureTerminal.send and Game.market.deal methods.
func (m Market) CalcTransactionCost(amount int, room1, room2 *room.Room) int {
	return m.Call("calcTransactionCost", amount, room1.Name, room2.Name).Int()
}

// CancelOrder cancels a previously created order.
func (m Market) CancelOrder(order Order) error {
	return common.ErrT(m.Call("cancelOrder", order.ID).Int())
}

// ChangeOrderPrice changes the price of an existing order.
func (m Market) ChangeOrderPrice(order Order, price float32) error {
	return common.ErrT(m.Call("changeOrderPrice", order.ID, price).Int())
}

// CreateOrder creates a market order in your terminal.
func (m Market) CreateOrder(order Order) error {
	return common.ErrT(m.Call("createOrder", order).Int())
}

// Deal executes a trade deal from your Terminal in yourRoomName to another player's Terminal using the specified buy/sell order.
func (m Market) Deal(order Order, amount int, room room.Room) error {
	return common.ErrT(m.Call("deal", order.ID, amount, room.Name).Int())
}

// ExtendOrder adds more capacity to an existing order.
func (m Market) ExtendOrder(order Order, amount int) error {
	return common.ErrT(m.Call("extendOrder", order.ID, amount).Int())
}

// GetAllOrders gets other players' orders currently active on the market.
func (m Market) GetAllOrders(filter string) []Order {
	return m.Call("getAllOrders", filter).Interface().([]Order)
}

// GetHistory gets daily price history of the specified resource on the market for the last 14 days.
func (m Market) GetHistory(resourceType string) ResourceHistory {
	return m.Call("getHistory", resourceType).Interface().(ResourceHistory)
}

// GetOrderByID retrieves info for specific market order.
func (m Market) GetOrderByID(orderID string) Order {
	return m.Call("getOrderById", orderID).Interface().(Order)
}
