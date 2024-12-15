"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const database_config_1 = require("./config/database.config");
const database_config_2 = require("./config/database.config");
const app = (0, express_1.default)();
// db.sync().then(() => {
//     console.log('connect to db')
// })
database_config_1.db.authenticate()
    .then(() => {
    console.log('Database connection success! Sequelize is ready to use...');
    // Start listening for connections
    app.listen(database_config_2.sequelizeConfig.port, () => {
        console.log(`Listening on port ${database_config_2.sequelizeConfig.port}...`);
    });
})
    .catch((err) => {
    console.error('Database connection failure.', err);
});
app.use(express_1.default.json());
app.get('/:id', (req, res) => {
    res.send({ message: 'TEst!', id: req.params.id });
});
app.post('/', (req, res) => {
    res.send({
        data: req.body
    });
});
