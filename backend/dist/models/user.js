'use strict';
import { Model } from 'sequelize';
import validator from 'validator';
module.exports = (sequelize, DataTypes) => {
    class User extends Model {
        /**
         * Helper method for defining associations.
         * This method is not a part of Sequelize lifecycle.
         * The `models/index` file will call this method automatically.
         */
        id;
        firstName;
        lastName;
        username;
        email;
        password;
        static associate(models) {
            // define association here
            User.belongsToMany(models.Toy, {
                through: 'wishlists'
            });
        }
    }
    User.init({
        id: {
            type: DataTypes.INTEGER,
            allowNull: false,
            primaryKey: true,
            autoIncrement: true
        },
        firstName: {
            type: DataTypes.STRING,
            allowNull: false
        },
        lastName: {
            type: DataTypes.STRING,
            allowNull: false
        },
        email: {
            type: DataTypes.STRING,
            allowNull: false,
            validate: {
                isEmail: true
            }
        },
        password: {
            type: DataTypes.STRING.BINARY,
            allowNull: false,
            validate: {
                len: [60, 60]
            }
        },
        username: {
            type: DataTypes.STRING,
            allowNull: false,
            validate: {
                len: [8, 30],
                isNotEmail(value) {
                    if (validator.isEmail(value)) {
                        throw new Error("Cannot be an email.");
                    }
                }
            }
        }
    }, {
        sequelize,
        modelName: 'User',
    });
    return User;
};
