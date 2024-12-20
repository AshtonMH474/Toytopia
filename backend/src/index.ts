import express from 'express'
import { Request,Response } from 'express';
require('dotenv').config();
import {port} from './config'
import db from './models';
import {createUsers} from './seeders/users'
import { createToys } from './seeders/toys';

const app = express();

app.use(express.json())


 db.sequelize
  .authenticate()
  .then(() => {
    console.log('Database connection success!');

    return db.sequelize.sync(); // Synchronize the schema
  })
  .then(() => {
    console.log('Database schema synchronized!');

    try {
        createUsers(); // seeds users
        console.log('Users seeded successfully!');

        createToys(); // seeds toys
        console.log('Toys seeded successfully!');
      } catch (err) {
        console.error('Error during seeding:', err);
      }

    // Start listening for connections
    app.listen(port, () => {
      console.log(`Application running on port ${port}`);
    });
  })
  .catch((err:any) => {
    console.error('Error during database initialization:', err);
  });





app.get('/:id', (req:Request,res:Response) => {
    res.send('TEst!')
})

app.post('/',(req:Request,res:Response) => {
    res.send({
        data:req.body
    })
})
