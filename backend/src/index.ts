import express from 'express'
import { db } from './config/database.config';
import config from './config';
import { Request,Response } from 'express';
const app = express();



db.authenticate()
    .then(() => {
    console.log('Database connection success! Sequelize is ready to use...');

    // Start listening for connections
    app.listen(config.port, () => {
        console.log(`Listening on port ${config.port}...`);
    });
    })
    .catch((err: Error) => {
    console.error('Database connection failure.', err);
    });


app.use(express.json())

app.get('/:id', (req:Request,res:Response) => {
    res.send({message:'TEst!',id:req.params.id})
})

app.post('/',(req:Request,res:Response) => {
    res.send({
        data:req.body
    })
})
