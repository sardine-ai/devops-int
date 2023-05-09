const express = require('express')
const app = express()
const port = 3000
const ethers = require('ethers')

var process = require('process')
process.on('SIGINT', () => {
    console.info("Interrupted")
    process.exit(0)
})

app.get('/', (req, res) => {
    res.send('Hello World!')
})

app.get('/eth', (req, res) => {
    const wallet = ethers.Wallet.createRandom()
    msg = `Address ${wallet.address} | mnemonic: ${wallet.mnemonic.phrase} | privatekey: ${wallet.privateKey} `
    res.send(msg)
})


app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})
