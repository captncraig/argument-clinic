"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
// init project
const express = require("express");
var app = express();
const sqlite = require("sqlite");
async function main() {
    try {
        var db = await sqlite.open(".data/comments.sqlite");
        console.log('DB Connected');
        await db.migrate({});
        console.log("Schema created");
        var listener = app.listen(process.env.PORT, function () {
            console.log('Your app is listening on port ' + listener.address().port);
        });
    }
    catch (e) {
        console.log("DB ERROR: " + e);
    }
}
main();
