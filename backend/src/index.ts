import express from 'express'

const app = express();

app.use(express.json())

app.get('/', (req,res) => {
    res.send('TEst!')
})

app.post('/',(req,res) => {
    res.send({
        data:req.body
    })
})

const port:number = 8001

app.listen(port,() => {
    console.log(`Application running on port ${port}`)
})
