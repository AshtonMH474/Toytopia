import { Sequelize, Options } from 'sequelize';
import dotenv from 'dotenv';

dotenv.config();

// Type for environment variables (development, production, etc.)
type Environment = 'development' | 'production';

// Interface for the JWT configuration
interface JwtConfig {
  secret: string;
  expiresIn: string;
}

// Interface for the database configuration
interface DatabaseConfig {
  environment: Environment;
  port: number;
  dbFile: string;
  jwtConfig: JwtConfig;
}

// Define environment variables
const environment = (process.env.NODE_ENV as Environment) || 'development';

// Load configuration based on environment
export const sequelizeConfig: DatabaseConfig = {
  environment,
  port: parseInt(process.env.PORT || '8000', 10),
  dbFile: process.env.DB_FILE || './db.sqlite',
  jwtConfig: {
    secret: process.env.JWT_SECRET || 'defaultsecret',
    expiresIn: process.env.JWT_EXPIRES_IN || '604800', // default 7 days in seconds
  },
};

// Define the database configuration for each environment
const databaseConfig: Record<Environment, Options> = {
  development: {
    storage: sequelizeConfig.dbFile,
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

// Initialize Sequelize based on the environment
let db: Sequelize;

if (environment === 'production') {
  // Production environment: use DATABASE_URL from environment variable
  const databaseUrl = process.env.DATABASE_URL;
  if (databaseUrl) {
    db = new Sequelize(databaseUrl, databaseConfig.production);
  } else {
    throw new Error('DATABASE_URL is not defined in production environment');
  }
} else {
  // Development environment: use file-based SQLite database
  db = new Sequelize(databaseConfig.development);
}

// Export both the db instance and seederStorageConfig for use in seeders
export { db, seederStorageConfig };
