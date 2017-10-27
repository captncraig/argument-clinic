
import { Express, Request, Response, RequestHandler } from "express";
import * as bodyParser from 'body-parser';
import * as hids from 'hashids';

import { DataAccess } from './db/db'
import { CreateCommentReq, Model } from './models'



function bodyAs(cstr: new () => Model): RequestHandler {
    return async (req, res, next) => {
        var target: any = new cstr();
        var jsonObj = req.body;
        for (let prop in jsonObj) {
            target[prop] = jsonObj[prop];
        }
        var msgs = await (target as Model).Validate();
        if (msgs.length > 0) {
            console.log(msgs,"!!!!!!!!!!!!!!!")
        }
        req.body = target;
        next();
    }
}

interface response {
    data: any;
    success: boolean;
    msg: string;
}

function sendSuccess(data: any, res: Response){
    res.send({
        data: data,
        success: true,
    })
}

export default class Api {
    private hid: any;

    constructor(private db: DataAccess) { 
        var a: any = hids;
        this.hid = new a("aaa",4);
    }

    public registerRoutes(app: Express) {
        app.use(bodyParser.json());
        app.post("/api/comments", bodyAs(CreateCommentReq), this.createComment)
        app.get("/api/comments", this.getComments)
    }

    private createComment = async (req: Request, res: Response) => {
        var dat = req.body as CreateCommentReq;
        var newId = await this.db.CreateComment(dat);
        sendSuccess({id: this.hid.encode(newId) }, res);
    }

    private getComments = async(req: Request, res: Response)=>{
        var comments = await this.db.GetComments();
        comments.forEach((x)=>x.id = this.hid.encode(x.id))
        sendSuccess(comments, res)
    }

}