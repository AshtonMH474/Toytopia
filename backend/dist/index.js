"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
require('dotenv').config();
const config_1 = require("./config");
const app = (0, express_1.default)();
app.use(express_1.default.json());
//  db.sequelize
//   .authenticate()
//   .then(() => {
//     console.log('Database connection success!');
//     return db.sequelize.sync(); // Synchronize the schema
//   })
//   .then(() => {
//     console.log('Database schema synchronized!');
//     try {
//         createUsers(); // seeds users
//         console.log('Users seeded successfully!');
//         createToys(); // seeds toys
//         console.log('Toys seeded successfully!');
//       } catch (err) {
//         console.error('Error during seeding:', err);
//       }
// Start listening for connections
app.listen(config_1.port, () => {
    console.log(`Application running on port ${config_1.port}`);
});
// })
// .catch((err:any) => {
//   console.error('Error during database initialization:', err);
// });
app.get('/:id', (req, res) => {
    res.send('TEst!');
});
app.post('/', (req, res) => {
    res.send({
        data: req.body
    });
});
