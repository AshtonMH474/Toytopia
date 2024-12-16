"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const sequelize_1 = require("sequelize");
const database_config_1 = require("../config/database.config");
const validator_extras_1 = require("sequelize/types/utils/validator-extras");
class UserInstance extends sequelize_1.Model {
}
UserInstance.init({
    id: {
        type: sequelize_1.DataTypes.INTEGER,
        primaryKey: true,
        autoIncrement: true,
        allowNull: false
    },
    email: {
        type: sequelize_1.DataTypes.STRING,
        allowNull: false,
        unique: true,
        validate: {
            len: [3, 255],
            isEmail: true
        }
    },
    username: {
        type: sequelize_1.DataTypes.STRING,
        allowNull: false,
        unique: true,
        validate: {
            len: [8, 30],
            isNotEmail(value) {
                if (validator_extras_1.validator.isEmail(value)) {
                    throw new Error("Cannot be an email.");
                }
            }
        }
    },
    firstName: {
        type: sequelize_1.DataTypes.STRING,
        allowNull: false
    },
    lastName: {
        type: sequelize_1.DataTypes.STRING,
        allowNull: false
    },
    password: {
        type: sequelize_1.DataTypes.STRING,
        allowNull: false,
        validate: {
            len: [60, 60]
        }
    }
}, {
    sequelize: database_config_1.db,
    tableName: 'users'
});
