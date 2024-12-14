import express from 'express'

const app = express();


app.get('/', (req,res) => {
    res.send('TEst!')
})

const port:number = 8000

app.listen(port,() => {
    console.log(`Application running on port ${port}`)
})
