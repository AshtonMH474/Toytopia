"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const database_config_1 = __importDefault(require("./config/database.config"));
const app = (0, express_1.default)();
database_config_1.default.sync().then(() => {
    console.log('connect to db');
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
const port = 8001;
app.listen(port, () => {
    console.log(`Application running on port ${port}`);
});
