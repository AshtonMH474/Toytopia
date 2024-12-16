import dotenv from 'dotenv';

dotenv.config();

// Type for the environment variables
type Environment = 'development' | 'production';

// Interface for JWT configuration
interface JwtConfig {
  secret: string | undefined;
  expiresIn: string | undefined;
}

// Interface for the database configuration
interface DatabaseConfig {
  environment: Environment;
  port: number;
  dbFile?: string;
  jwtConfig: JwtConfig;
}

// Export the configuration
const config: DatabaseConfig = {
  environment: (process.env.NODE_ENV as Environment) || 'development',
  port: parseInt(process.env.PORT || '8000', 10),
  dbFile: process.env.DB_FILE,
  jwtConfig: {
    secret: process.env.JWT_SECRET,
    expiresIn: process.env.JWT_EXPIRES_IN,
  },
};

export default config;
