"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
require('dotenv').config();
const config_1 = require("./config");
const models_1 = __importDefault(require("./models"));
const app = (0, express_1.default)();
app.use(express_1.default.json());
models_1.default.sequelize
    .authenticate()
    .then(() => {
    console.log('Database connection success!');
    return models_1.default.sequelize.sync(); // Synchronize the schema
})
    .then(() => {
    console.log('Database schema synchronized!');
    // Start listening for connections
    app.listen(config_1.port, () => {
        console.log(`Application running on port ${config_1.port}`);
    });
})
    .catch((err) => {
    console.error('Error during database initialization:', err);
});
app.get('/:id', (req, res) => {
    res.send('TEst!');
});
app.post('/', (req, res) => {
    res.send({
        data: req.body
    });
});
