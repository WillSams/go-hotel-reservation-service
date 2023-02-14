const { readFile } = require('fs');

const jsonReader = (filepath, callback) => {
  return new Promise((resolve, reject) => {
    readFile(filepath, (error, data) => {
      if (error) reject(error);
      else {
        const obj = JSON.parse(data);
        callback(null, obj);
        resolve();
      }
    });
  });
};

const seedData = async (knex, tableName, data) => {
  await knex(tableName).del();
  return await knex(tableName).insert(data);
};

const injectTables = async (knex, tableName) => {
  await jsonReader(`./seeds/development/${tableName}.json`, (error, data) => {
    try {
      if (error) throw error;
      else inserts = data;
    } catch (ex) {
      console.log(`Json file read failed: ${ex.message}`);
    }
  });
  await seedData(knex, tableName, inserts);
};

exports.seed = async (knex) => {
  await injectTables(knex, 'rooms');
  await injectTables(knex, 'reservations');
};
