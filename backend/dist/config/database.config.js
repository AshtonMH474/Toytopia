"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.seederStorageConfig = exports.db = exports.sequelizeConfig = void 0;
const sequelize_1 = require("sequelize");
const dotenv_1 = __importDefault(require("dotenv"));
dotenv_1.default.config();
// Define environment variables
const environment = process.env.NODE_ENV || 'development';
// Load configuration based on environment
exports.sequelizeConfig = {
    environment,
    port: parseInt(process.env.PORT || '8000', 10),
    dbFile: process.env.DB_FILE || './database.sqlite',
    jwtConfig: {
        secret: process.env.JWT_SECRET || 'defaultsecret',
        expiresIn: process.env.JWT_EXPIRES_IN || '604800', // default 7 days in seconds
    },
};
// Define the database configuration for each environment
const databaseConfig = {
    development: {
        storage: exports.sequelizeConfig.dbFile,
        dialect: 'sqlite',
        logging: process.env.NODE_ENV !== 'production', // Enable logging for non-production
        typeValidation: true,
        logQueryParameters: true,
    },
    production: {
        dialect: 'postgres',
        dialectOptions: {
            ssl: {
                require: true,
                rejectUnauthorized: false, // If using Postgres with SSL
            },
        },
        define: {
            schema: process.env.SCHEMA || 'public',
        },
    },
};
// Seeder Storage Configuration
const seederStorageConfig = {
    seederStorage: 'sequelize', // This will apply only when running seeders
};
exports.seederStorageConfig = seederStorageConfig;
// Initialize Sequelize based on the environment
let db;
if (environment === 'production') {
    // Production environment: use DATABASE_URL from environment variable
    const databaseUrl = process.env.DATABASE_URL;
    if (databaseUrl) {
        exports.db = db = new sequelize_1.Sequelize(databaseUrl, databaseConfig.production);
    }
    else {
        throw new Error('DATABASE_URL is not defined in production environment');
    }
}
else {
    // Development environment: use file-based SQLite database
    exports.db = db = new sequelize_1.Sequelize(databaseConfig.development);
}
