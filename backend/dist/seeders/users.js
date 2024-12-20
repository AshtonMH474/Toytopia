"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createUsers = void 0;
const bcryptjs_1 = __importDefault(require("bcryptjs"));
const models_1 = __importDefault(require("../models"));
const users = [
    {
        firstName: 'Demo',
        lastName: 'User',
        email: 'demo@gmail.com',
        username: 'demo1234',
        password: bcryptjs_1.default.hashSync('password')
    },
    {
        firstName: 'User1',
        lastName: 'Test',
        email: 'user1@gmail.com',
        username: 'user1user',
        password: bcryptjs_1.default.hashSync('password1')
    },
    {
        firstName: 'Jolly',
        lastName: 'Rodger',
        email: 'rodger@gmail.com',
        username: 'rodger1234',
        password: bcryptjs_1.default.hashSync('password2')
    }
];
const createUsers = () => {
    users.map((user) => {
        models_1.default.User.create(user);
    });
};
exports.createUsers = createUsers;
