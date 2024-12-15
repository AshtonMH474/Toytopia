import dotenv from 'dotenv';

dotenv.config();

export const sequelizeConfig = {
  environment: process.env.NODE_ENV || 'development',
  port: parseInt(process.env.PORT || '8000', 10),
  dbFile: process.env.DB_FILE || './db/dev.db',
  jwtConfig: {
    secret: process.env.JWT_SECRET || 'defaultsecret',
    expiresIn: process.env.JWT_EXPIRES_IN || '604800',
  },
};

export const databaseConfig = {
  development: {
    storage: sequelizeConfig.dbFile,
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
