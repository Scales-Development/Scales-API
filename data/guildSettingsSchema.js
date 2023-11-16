const mongoose = require('mongoose');

const guildSchema = mongoose.Schema({
    guildID: String,
    ownerID: String,
    appealChannel: String,

}) // TODO: Finish database structure for guild schema for the backend