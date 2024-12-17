import express from 'express';
require('dotenv').config();
import { port } from './config';
import db from './models';
const app = express();
app.use(express.json());
db.sequelize
    .authenticate()
    .then(() => {
    console.log('Database connection success!');
    return db.sequelize.sync(); // Synchronize the schema
})
    .then(() => {
    console.log('Database schema synchronized!');
    // Start listening for connections
    app.listen(port, () => {
        console.log(`Application running on port ${port}`);
    });
})
    .catch((err) => {
    console.error('Error during database initialization:', err);
});
// seeds the users in db
// createUsers()
// seeds the toys in db
// createToys()
app.get('/:id', (req, res) => {
    res.send('TEst!');
});
app.post('/', (req, res) => {
    res.send({
        data: req.body
    });
});
