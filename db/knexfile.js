const defaults = {
  client: process.env.DB_CLIENT,
  migrations: {
    directory: './migrations',
    tableName: 'knex_migrations'
  },
  pool: { min: 2, max: 10 },
  debug: false,
};

const connection = {
  host: process.env.DB_HOST,
  port: process.env.DB_PORT,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWD,
};

module.exports = {
  development: {
    ...defaults,
    connection: {
      ...connection,
      database: 'hotel_development',
    },
    seeds: { directory: './seeds/development', },
  },

  test: {
    ...defaults,
    connection: {
      ...connection,
      database: 'hotel_test',
    },
    seeds: { directory: './seeds/test', },
  },
};