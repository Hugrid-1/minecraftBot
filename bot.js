const mineflayer = require("mineflayer");
const {pathfinder, Movements, goals: {GoalNear, GoalBlock, GoalFollow}} = require('mineflayer-pathfinder');
const mineflayerViewer = require('prismarine-viewer').mineflayer


const bot = mineflayer.createBot({
     host: 'cheshuiki.aternos.me',
     port: 25781,
     username: 'GigaChadBot',
     version: '1.20.1',
})
require('mineflayer-auto-eat').plugin(bot)
bot.loadPlugin(pathfinder)

bot.on('login',  function() {
    bot.autoEat.options = {
      priority: 'foodPoints',
      startAt: 14,
      bannedFood: []
    }
    const mcData = require('minecraft-data')(bot.version)
    const defaultMove = new Movements(bot, mcData)
    bot.pathfinder.setMovements(defaultMove)
    defaultMove.scafoldingBlocks = []
    defaultMove.allowParkour = false
  })

bot.on('health', function(){
    console.log(`food ${bot.food}`,`health ${bot.health}`)
    if (bot.food !== 20) {
        bot.autoEat.disable()
        bot.autoEat.enable()
    }
    if (bot.health !== 20) {
        bot.autoEat.enable()
    }
})

bot.once('spawn', function () {
    mineflayerViewer(bot, { firstPerson: true, port: 3000 })
    bot.chat('Я родился')
})

bot.on('chat', function (username, message) {
    const player = bot.players[username].entity
    switch (message) {
        case 'За мной':
            bot.pathfinder.setGoal(new GoalFollow(player, 1), true)
            break
        case 'Ко мне':
            bot.pathfinder.setGoal(new GoalNear(player.position.x, player.position.y, player.position.z, 1))
            break
        case 'Остановись':
            bot.pathfinder.setGoal(null, 1)
            break
    }

})