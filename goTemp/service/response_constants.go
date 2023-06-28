package service

const ipsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean euismod neque tincidunt, facilisis eros vel, bibendum urna. Cras dignissim ultrices facilisis. Donec nec magna at nisi pellentesque semper. Sed eu dui tortor. Sed accumsan libero ut enim imperdiet ornare. Suspendisse metus ipsum, suscipit et nisl nec, mattis consequat lectus. Nullam at dui eu ipsum laoreet vehicula nec et elit. Vestibulum tempus tortor ut orci fringilla commodo."

type fixedMessage string

const (
	// Asking for refund
	fixedMessageInquireRefund fixedMessage = "ดูขั้นตอนการขอคืนเงิน"

	// Ask if refund in same day
	fixedMessageQuestionSameDay  fixedMessage = "คุณต้องการขอคืนเงินภายในวันเดียวกับวันที่มีการทำธุรกรรมหรือไม่?"
	fixedMessageChoiceSameDayYes fixedMessage = "ใช่ ภายในวันเดียวกัน"
	fixedMessageChoiceSameDayNo  fixedMessage = "ไม่ใช่ วันอื่น"

	// Ask if refund before 20:00
	fixedMessageQuestion2000     fixedMessage = "คุณขอคืนเงินก่อนเวลา 20:00 น. ของวันที่ลูกค้าชำระเงินหรือไม่?"
	fixedMessageChoice2000Before fixedMessage = "ก่อน 20:00 น."
	fixedMessageChoice2000After  fixedMessage = "หลัง 20:00 น."

	// Text Response
	fixedMessageRefundHeader fixedMessage = "ขั้นตอนการ Refund"
	fixedMessageRefundBody   fixedMessage = "ขั้นตอนแรก... บลา ๆ ๆ " + ipsum

	fixedMessageVoidHeader fixedMessage = "ขั้นตอนการ Void"
	fixedMessageVoidBody   fixedMessage = "ขั้นตอนแรก... บลา ๆ ๆ " + ipsum

	fixedMessageDefaultFooter fixedMessage = "หากมีข้อสงสัย สอบถามได้ในช่องทางนี้เลยค่าา"
)
