
import * as sqlite from 'sqlite';
import * as fs from 'fs-extra';
import {CreateCommentReq, Comment} from '../models';


export class DataAccess implements DataAccess{
    private db: sqlite.Database;

    public async CreateComment(req: CreateCommentReq): Promise<number>{
        var stmt = await this.db.run("INSERT INTO Comments (name,text) VALUES(?,?)", req.name, req.text)
        return stmt.lastID;
    }

    public async GetComments(): Promise<Comment[]>{
        return await this.db.all("SELECT * FROM Comments")
    }

    constructor(db: sqlite.Database){
        this.db = db;
    }
}

export default async function Connect(): Promise<DataAccess> {
    await fs.mkdirp(".data")
    var db = await sqlite.open(".data/comments.sqlite")
    console.log('DB Connected');
    await db.migrate({})
    console.log("Schema verified")
    return new DataAccess(db);
}