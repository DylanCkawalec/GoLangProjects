const functions = require("firebase-functions");
const { Telegraf } = require('telegraf')

const bot = new Telegraf("2086534811:AAFa88ZqpP82YlSt4Q8YY6IL5JUP1nl9D2E")

bot.start((ctx) => ctx.reply('hello'))
bot.help((ctx) => ctx.reply('Send me a sticker'))
bot.on('sticker', (ctx) => ctx.reply('👍'))
bot.hears('hi', (ctx) => ctx.reply('Hey there'))
bot.launch()





// // Create and Deploy Your First Cloud Functions
// // https://firebase.google.com/docs/functions/write-firebase-functions
//
exports.helloWorld = functions.https.onRequest((request, response) => {
  functions.logger.info("Hello logs!", {structuredData: true});
  response.send("Hello from Firebase!");
});
