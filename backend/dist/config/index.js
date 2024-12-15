"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.databaseConfig = exports.sequelizeConfig = void 0;
const dotenv_1 = __importDefault(require("dotenv"));
dotenv_1.default.config();
exports.sequelizeConfig = {
    environment: process.env.NODE_ENV || 'development',
    port: parseInt(process.env.PORT || '8000', 10),
    dbFile: process.env.DB_FILE || './db/dev.db',
    jwtConfig: {
        secret: process.env.JWT_SECRET || 'defaultsecret',
        expiresIn: process.env.JWT_EXPIRES_IN || '604800',
    },
};
exports.databaseConfig = {
    development: {
        storage: exports.sequelizeConfig.dbFile,
        dialect: 'sqlite',
        seederStorage: 'sequelize',
        logQueryParameters: true,
        typeValidation: true,
    },
    production: {
        use_env_variable: 'DATABASE_URL',
        dialect: 'postgres',
        seederStorage: 'sequelize',
        dialectOptions: {
            ssl: {
                require: true,
                rejectUnauthorized: false,
            },
        },
        define: {
            schema: process.env.SCHEMA || 'public',
        },
    },
};
