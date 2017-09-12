
// init project
import * as express from "express";
var app = express();
import * as sqlite from 'sqlite';
import * as fs from 'fs-extra';

async function main(){
  try{
    await fs.mkdirp(".data")
    var db = await sqlite.open(".data/comments.sqlite")
    console.log('DB Connected');
    await db.migrate({})
    console.log("Schema created")
    var listener = app.listen(process.env.PORT, function () {
      console.log('Your app is listening on port ' + listener.address().port);
    });
  }catch(e){
    console.log("DB ERROR: "+e)
  }
}
main()


