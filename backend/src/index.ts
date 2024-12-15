import express from 'express'
import { db } from './config/database.config';
import { sequelizeConfig } from './config/database.config';
import { Request,Response } from 'express';
const app = express();

// db.sync().then(() => {
//     console.log('connect to db')
// })

db.authenticate()
    .then(() => {
    console.log('Database connection success! Sequelize is ready to use...');

    // Start listening for connections
    app.listen(sequelizeConfig.port, () => {
        console.log(`Listening on port ${sequelizeConfig.port}...`);
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
