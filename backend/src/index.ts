import express from 'express'
import db from './config/database.config'
import { Request,Response } from 'express';
const app = express();

db.sync().then(() => {
    console.log('connect to db')
})


app.use(express.json())

app.get('/:id', (req:Request,res:Response) => {
    res.send({message:'TEst!',id:req.params.id})
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
