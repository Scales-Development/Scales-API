const mongoose = require('mongoose');

const AppealSchema = mongoose.Schema({
    userID: String,
    username: String,
    appealData: {
        type: Map,
        of: String
    },
    guildID: String,
    caseNumber: String,
    date: String,
})

module.exports = mongoose.model('AppealData', AppealSchema);