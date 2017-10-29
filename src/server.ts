

import * as express from "express";
var app = express();

import newDb from './db/db'
import Api from './api'

async function main() {
  var db = await newDb();
  var api = new Api(db);
  api.registerRoutes(app);
  app.get("/client.js", (res, resp: express.Response) => { resp.sendFile("client.js", { root: __dirname }) });
  app.get("/test", (res, resp: express.Response) => { resp.send(testpage) });
  var listener = app.listen(process.env.PORT || "8666", function () {
    console.log('Listening on port ' + listener.address().port);
  });
}
main()

const testpage = `<html>
  <head>
    <script src="client.js"></script>
  </head>
  <body>
    <h1>Argument Clinic Test Page</h1>
    <div class="ac-comments"></div>
  </body>
</html>`