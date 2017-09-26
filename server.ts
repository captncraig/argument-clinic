
// init project
import * as express from "express";
var app = express();
import * as sqlite from 'sqlite';
import * as fs from 'fs-extra';
import * as crypto from 'crypto';

import {Settings} from './models'

var appSettings: Settings;
var initialLoginToken: string;

async function main(){
  try{
    await fs.mkdirp(".data")
    var db = await sqlite.open(".data/comments.sqlite")
    console.log('DB Connected');
    await db.migrate({})
    console.log("Schema verified")
    appSettings = await db.get<Settings>("SELECT * FROM Settings LIMIT 1")
    if (!appSettings.hasInitialized){
      initialLoginToken = randomHexString(16);
      console.log("!!!!!! FIRST TIME STARTUP !!!!!!");
      console.log("YOUR ADMIN TOKEN IS:");
      console.log(initialLoginToken);
      console.log("You will need to visit the home page in a browser and enter this value.");
      console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!");
    }
    var listener = app.listen(process.env.PORT || "8666", function () {
      console.log('Listening on port ' + listener.address().port);
    });
  }catch(e){
    console.log("DB ERROR: "+e)
  }
}
main()


function randomHexString(bytes: number): string {
  return crypto.randomBytes(bytes).toString("hex");
}

