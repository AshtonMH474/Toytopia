import { Sequelize } from 'sequelize';
import dotenv from 'dotenv';
import { getSequelizeConfig } from './config'; // Import helper function

dotenv.config();

// Seeder Storage Configuration
const seederStorageConfig = {
  seederStorage: 'sequelize', // Applies when running seeders
};

// Initialize Sequelize
const sequelizeConfig = getSequelizeConfig();
let db: Sequelize;

// For production, use the DATABASE_URL directly
if (process.env.NODE_ENV === 'production') {
  const databaseUrl = process.env.DATABASE_URL;
  if (!databaseUrl) {
    throw new Error('DATABASE_URL is not defined in production environment');
  }

  db = new Sequelize(databaseUrl, {
    ...sequelizeConfig, // Spread the production config
    ...seederStorageConfig, // Include seeder storage configuration
  });
} else {
  db = new Sequelize(sequelizeConfig);
}

export { db, seederStorageConfig };



// Export both the db instance and seederStorageConfig for use in seeders
