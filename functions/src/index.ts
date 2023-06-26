import * as functions from "firebase-functions";
const request = require("request-promise");
// import * as request from "request-promise";

const LINE_MESSAGING_API = "https://api.line.me/v2/bot/message";
const LINE_HEADER = {
  "Content-Type": "application/json",
  Authorization: `Bearer ${process.env.CHANNEL_ACCESS_TOKEN}`,
};

export const LineBot = functions.https.onRequest((req, res) => {
  if (req.body.events[0].message.type !== "text") {
    return;
  }
  switch (req.body.events[0].message.text) {
    case "ask me":
      ask(req.body);
      break;

    default:
      echo(req.body);
      break;
  }
});

const echo = (bodyResponse: any): Promise<any> => {
  return request({
    method: `POST`,
    uri: `${LINE_MESSAGING_API}/reply`,
    headers: LINE_HEADER,
    body: JSON.stringify({
      replyToken: bodyResponse.events[0].replyToken,
      messages: [
        {
          type: `text`,
          text: bodyResponse.events[0].message.text,
        },
      ],
    }),
  });
};

const ask = (bodyResponse: any): Promise<any> => {
  return request({
    method: `POST`,
    uri: `${LINE_MESSAGING_API}/reply`,
    headers: LINE_HEADER,
    body: JSON.stringify({
      replyToken: bodyResponse.events[0].replyToken,
      messages: [
        {
          type: "template",
          altText: "this is a confirm template",
          template: {
            type: "confirm",
            text: "Are you sure?",
            actions: [
              {
                type: "message",
                label: "Yes",
                text: "yes",
              },
              {
                type: "message",
                label: "No",
                text: "no",
              },
            ],
          },
        },
      ],
    }),
  });
};
