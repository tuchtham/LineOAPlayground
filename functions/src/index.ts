import * as functions from "firebase-functions";
const request = require("request-promise");
// import * as request from "request-promise";

const LINE_MESSAGING_API = "https://api.line.me/v2/bot/message";
const LINE_HEADER = {
  "Content-Type": "application/json",
  Authorization: `Bearer ${process.env.CHANNEL_ACCESS_TOKEN}`,
};

// Asking for refund
const inquireRefund = "ดูขั้นตอนการขอคืนเงิน";

// Ask if refund in same day
const questionSameDay =
  "คุณต้องการขอคืนเงินภายในวันเดียวกับวันที่มีการทำธุรกรรมหรือไม่?";
const choiceSameDayYes = "ใช่ ภายในวันเดียวกัน";
const choiceSameDayNo = "ไม่ใช่ วันอื่น";

// Ask if refund before 20:00
const question2000 =
  "คุณขอคืนเงินก่อนเวลา 20:00 น. ของวันที่ลูกค้าชำระเงินหรือไม่?";
const choice2000Before = "ก่อน 20:00 น.";
const choice2000After = "หลัง 20:00 น.";

// Text Response
const textRefund = "ขั้นตอนการ Refund...";
const textVoid = "ขั้นตอนการ Void...";

export const LineBot = functions.https.onRequest((req, res) => {
  if (req.body.events[0].message.type !== "text") {
    return;
  }
  switch (req.body.events[0].message.text) {
    case "ping":
      replyWithText(req.body, "pong");
      break;
    case "emoji fire":
      replyWithText(req.body, "️‍🔥");
      break;
    case "new lines":
      replyWithText(req.body, "1\n2\n3");
      break;
    case "block":
      textBlock(req.body);
      break;
    case inquireRefund:
      askSameDay(req.body);
      break;
    case choiceSameDayYes:
      ask2000(req.body);
      break;
    case choiceSameDayNo:
      replyWithText(req.body, textRefund);
      break;
    case choice2000Before:
      replyWithText(req.body, textVoid);
      break;
    case choice2000After:
      replyWithText(req.body, textRefund);
      break;
    default:
      break;
  }
});

// const echo = (bodyResponse: any): Promise<any> => {
//   return request({
//     method: `POST`,
//     uri: `${LINE_MESSAGING_API}/reply`,
//     headers: LINE_HEADER,
//     body: JSON.stringify({
//       replyToken: bodyResponse.events[0].replyToken,
//       messages: [
//         {
//           type: `text`,
//           text: bodyResponse.events[0].message.text,
//         },
//       ],
//     }),
//   });
// };

const replyWithText = (bodyResponse: any, text: string): Promise<any> => {
  return request({
    method: `POST`,
    uri: `${LINE_MESSAGING_API}/reply`,
    headers: LINE_HEADER,
    body: JSON.stringify({
      replyToken: bodyResponse.events[0].replyToken,
      messages: [
        {
          type: `text`,
          text: text,
        },
      ],
    }),
  });
};

// const ask = (
//   bodyResponse: any,
//   question: string,
//   choice1: string,
//   choice2: string
// ): Promise<any> => {
//   return request({
//     method: `POST`,
//     uri: `${LINE_MESSAGING_API}/reply`,
//     headers: LINE_HEADER,
//     body: JSON.stringify({
//       replyToken: bodyResponse.events[0].replyToken,
//       messages: [
//         {
//           type: "template",
//           altText: "this is a confirm template",
//           template: {
//             type: "confirm",
//             text: question,
//             actions: [
//               {
//                 type: "message",
//                 label: choice1,
//                 text: choice1,
//               },
//               {
//                 type: "message",
//                 label: choice2,
//                 text: choice2,
//               },
//             ],
//           },
//         },
//       ],
//     }),
//   });
// };

// const askButton = (
//   bodyResponse: any,
//   question: string,
//   choice1: string,
//   choice2: string
// ): Promise<any> => {
//   return request({
//     method: `POST`,
//     uri: `${LINE_MESSAGING_API}/reply`,
//     headers: LINE_HEADER,
//     body: JSON.stringify({
//       replyToken: bodyResponse.events[0].replyToken,
//       messages: [
//         {
//           type: "template",
//           altText: "This is a buttons template",
//           template: {
//             type: "buttons",
//             text: question,
//             actions: [
//               {
//                 type: "message",
//                 label: choice1,
//                 text: choice1,
//               },
//               {
//                 type: "message",
//                 label: choice2,
//                 text: choice2,
//               },
//             ],
//           },
//         },
//       ],
//     }),
//   });
// };

const textBlock = (bodyResponse: any): Promise<any> => {
  replyWithText(bodyResponse, "came in textBlock");
  return request({
    method: `POST`,
    uri: `${LINE_MESSAGING_API}/reply`,
    headers: LINE_HEADER,
    body: JSON.stringify({
      replyToken: bodyResponse.events[0].replyToken,
      messages: [
        {
          type: "bubble",
          header: {
            type: "box",
            layout: "vertical",
            contents: [
              {
                type: "text",
                text: "header",
              },
            ],
          },
          body: {
            type: "box",
            layout: "vertical",
            contents: [
              {
                type: "text",
                text: "body",
              },
              {
                type: "text",
                text: "body",
              },
            ],
          },
          footer: {
            type: "box",
            layout: "vertical",
            contents: [
              {
                type: "text",
                text: "footer",
              },
            ],
          },
        },
      ],
    }),
  });
};

const askQuickReply = (
  bodyResponse: any,
  question: string,
  choice1: string,
  choice2: string
): Promise<any> => {
  return request({
    method: `POST`,
    uri: `${LINE_MESSAGING_API}/reply`,
    headers: LINE_HEADER,
    body: JSON.stringify({
      replyToken: bodyResponse.events[0].replyToken,
      messages: [
        {
          type: "text", // ①
          text: question,
          quickReply: {
            // ②
            items: [
              {
                type: "action", // ③
                action: {
                  type: "message",
                  label: choice1,
                  text: choice1,
                },
              },
              {
                type: "action",
                action: {
                  type: "message",
                  label: choice2,
                  text: choice2,
                },
              },
            ],
          },
        },
      ],
    }),
  });
};

const askSameDay = (bodyResponse: any): Promise<any> => {
  return askQuickReply(
    bodyResponse,
    questionSameDay,
    choiceSameDayYes,
    choiceSameDayNo
  );
};

const ask2000 = (bodyResponse: any): Promise<any> => {
  return askQuickReply(
    bodyResponse,
    question2000,
    choice2000Before,
    choice2000After
  );
};
