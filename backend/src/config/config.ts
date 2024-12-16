import dotenv from 'dotenv';
import { Options } from 'sequelize';

dotenv.config();

// Define the Sequelize configuration types
interface DatabaseConfig {
  development: Options;
  production: Options & { use_env_variable: string }; // Explicitly add use_env_variable for production
}

const databaseConfig: DatabaseConfig = {
  development: {
    storage: process.env.DB_FILE || './db.sqlite',
    dialect: 'sqlite',
    logging: process.env.NODE_ENV !== 'production', // Enable logging for non-production
    typeValidation: true,
    logQueryParameters: true,
  },
  production: {
    use_env_variable: 'DATABASE_URL', // Custom property for the production URL
    dialect: 'postgres',
    dialectOptions: {
      ssl: {
        require: true,
        rejectUnauthorized: false, // For Postgres SSL
      },
    },
    define: {
      schema: process.env.SCHEMA || 'public',
    },
  },
};

// Helper function to retrieve the Sequelize options
const getSequelizeConfig = (): Options => {
  const env = process.env.NODE_ENV || 'development';

  if (env === 'production') {
    const config = databaseConfig.production;
    const databaseUrl = process.env[config.use_env_variable]; // Use the custom property
    if (!databaseUrl) {
      throw new Error(`Environment variable ${config.use_env_variable} is not defined`);
    }

    return {
      ...config, // Spread the production config
      // Do not add `url` here since it's handled separately in Sequelize constructor
    };
  }

  // For development
  return databaseConfig.development;
};

export { databaseConfig, getSequelizeConfig };
