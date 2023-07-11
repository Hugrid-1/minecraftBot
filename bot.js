const mineflayer = require("mineflayer");
const mineflayerViewer = require('prismarine-viewer').mineflayer

const bot = mineflayer.createBot({
     host: 'cheshuiki.aternos.me',
     port: 25781,
     username: 'GigaChadBot'
})

bot.once('spawn', function () {
    mineflayerViewer(bot, { firstPerson: true, port: 3000 })
    bot.chat('Я родился')
})

