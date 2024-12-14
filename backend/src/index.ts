import express from 'express'
import { Request,Response } from 'express';
const app = express();

app.use(express.json())

app.get('/:id', (req:Request,res:Response) => {
    res.send('TEst!')
})

app.post('/',(req:Request,res:Response) => {
    res.send({
        data:req.body
    })
})

const port:number = 8001

app.listen(port,() => {
    console.log(`Application running on port ${port}`)
})
