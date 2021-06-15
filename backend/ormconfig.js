// eslint-disable-next-line @typescript-eslint/no-var-requires
require("dotenv").config();

const { DB_NAME, DB_USERNAME, DB_PASSWORD } = process.env;

module.exports = {
  type: "postgres",
  host: "127.0.0.1",
  port: 5432,
  username: DB_USERNAME,
  password: DB_PASSWORD,
  database: DB_NAME,
  synchronize: true,
  logging: false,
  dropSchema: true,
  migrationsRun: true,

  entities: ["src/entities/*.ts", "build/src/entities/*.js"],
  migrations: ["src/migrations/*.ts"],
  subscribers: ["src/subscriber/*.ts"],
  cli: {
    migrationsDir: "src/migrations",
    entitiesDir: "src/entities",
  },
};
