"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const app = (0, express_1.default)();
app.use(express_1.default.json());
app.get('/:id', (req, res) => {
    res.send('TEst!');
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
