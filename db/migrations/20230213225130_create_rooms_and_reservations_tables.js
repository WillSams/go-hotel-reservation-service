exports.up = async (knex) => {
  return knex.schema
    .createTable("rooms", (table) => {
      table.string("id").primary().unique();
      table.integer("num_beds");
      table.boolean("allow_smoking");
      table.integer("daily_rate");
      table.integer("cleaning_fee");
    })
    .createTable("reservations", (table) => {
      table.increments("id").primary();
      table.string("room_id").references("rooms.id");
      table.string("checkin_date");
      table.string("checkout_date");
      table.integer("total_charge");
      table.unique(["room_id", "checkin_date", "checkout_date"]);
    });
};

exports.down = async (knex) =>
  knex.schema.dropTable("reservations").dropTable("rooms");
