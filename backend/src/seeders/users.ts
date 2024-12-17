import bcrypt from 'bcryptjs';
import db from '../models';
const users = [
    {
        firstName:'Demo',
        lastName:'User',
        email:'demo@gmail.com',
        username:'demo1234',
        password:bcrypt.hashSync('password')
    },
    {
        firstName:'User1',
        lastName:'Test',
        email:'user1@gmail.com',
        username:'user1user',
        password:bcrypt.hashSync('password1')
    },
    {
        firstName:'Jolly',
        lastName:'Rodger',
        email:'rodger@gmail.com',
        username:'rodger1234',
        password:bcrypt.hashSync('password2')
    }
]


export const createUsers = () => {
    users.map((user) => {
        db.User.create(user)
    })
}
