package service

type fixedMessage string

const (
	// Asking for refund
	inquireRefund = "ดูขั้นตอนการขอคืนเงิน"

	// Ask if refund in same day
	questionSameDay  = "คุณต้องการขอคืนเงินภายในวันเดียวกับวันที่มีการทำธุรกรรมหรือไม่?"
	choiceSameDayYes = "ใช่ ภายในวันเดียวกัน"
	choiceSameDayNo  = "ไม่ใช่ วันอื่น"

	// Ask if refund before 20:00
	question2000     = "คุณขอคืนเงินก่อนเวลา 20:00 น. ของวันที่ลูกค้าชำระเงินหรือไม่?"
	choice2000Before = "ก่อน 20:00 น."
	choice2000After  = "หลัง 20:00 น."

	// Text Response
	textRefund = "ขั้นตอนการ Refund..."
	textVoid   = "ขั้นตอนการ Void..."
)
