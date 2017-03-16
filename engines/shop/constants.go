package shop

const (
	// OrderStateCart cart
	OrderStateCart = "cart"
	// OrderStateAddress address
	OrderStateAddress = "address"
	// OrderStateDeliverty delivery
	OrderStateDeliverty = "delivery"
	// OrderStatePayment payment
	OrderStatePayment = "payment"
	// OrderStateConfirm confirm
	OrderStateConfirm = "confirm"
	// OrderStateComplete complete
	OrderStateComplete = "complete"

	// PaymentStateBalanceDue balance-due
	PaymentStateBalanceDue = "balance-due"
	// PaymentStatePaid paid
	PaymentStatePaid = "paid"
	// PaymentStateCreditOwed credit-owed
	PaymentStateCreditOwed = "credit-owed"
	// PaymentStateFailed failed
	PaymentStateFailed = "failed"
	// PaymentStateVoid void
	PaymentStateVoid = "void"
	// PaymentStateCheckout checkout
	PaymentStateCheckout = "checkout"
	// PaymentStatePending pending
	PaymentStatePending = "pending"
	// PaymentStateProcessing processing
	PaymentStateProcessing = "processing"
	// PaymentStateCompleted completed
	PaymentStateCompleted = "completed"

	// ShipmentStateReady ready
	ShipmentStateReady = "ready"
	// ShipmentStatePending pending
	ShipmentStatePending = "pending"
	// ShipmentStatePartial partial
	ShipmentStatePartial = "partial"
	// ShipmentStateShipped shipped
	ShipmentStateShipped = "shipped"
	// ShipmentStateBackorder backorder
	ShipmentStateBackorder = "backorder"
	// ShipmentStateCanceled canceled
	ShipmentStateCanceled = "canceled"
	// ShipmentStateAssemble assemble
	ShipmentStateAssemble = "assemble"

	// PaymentMethodTypeCash cash
	PaymentMethodTypeCash = "cash"
	// PaymentMethodTypeCredit credit
	PaymentMethodTypeCredit = "credit"
	// PaymentMethodTypeTransfer transfer
	PaymentMethodTypeTransfer = "transfer"

	// ReturnAuthorizationStateAuthorized authorized
	ReturnAuthorizationStateAuthorized = "authorized"
	// ReturnAuthorizationStateCanceled canceled
	ReturnAuthorizationStateCanceled = "canceled"

	// InventoryUnitStateBackordered backordered
	InventoryUnitStateBackordered = "backordered"
	// InventoryUnitStateOnHand on-hand
	InventoryUnitStateOnHand = "on-hand"
	// InventoryUnitStateShipped = shipped
	InventoryUnitStateShipped = "shipped"
	// InventoryUnitStateReturned returned
	InventoryUnitStateReturned = "returned"
)
