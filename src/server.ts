

import * as express from "express";
var app = express();

import newDb from './db/db'
import Api from './api'

async function main() {
  var db = await newDb();
  var api = new Api(db);
  api.registerRoutes(app);
  var listener = app.listen(process.env.PORT || "8666", function () {
    console.log('Listening on port ' + listener.address().port);
  });
}
main()
