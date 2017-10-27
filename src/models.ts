
import {validate,MinLength,MaxLength,IsNotEmpty} from "class-validator";

export class Model {

    async Validate(): Promise<string[]>{
        var errs = await validate(this)
        var msgs: string[] = [];
        if (errs.length > 0){
            for (var err of errs){
               for (var cst of Object.values(err.constraints)){
                   msgs.push(cst);
               }
            }
        }
        return msgs;
    }
}

export class CreateCommentReq extends Model {
    name: string;

    @IsNotEmpty({message: "text may not be empty"})
    @MaxLength(2048, {message: "text too long"})
    text: string;

    public Test(){
        console.log("!!!", this);
    }
}

// Comment as read from db and sent to client
export class Comment extends Model{
    name: string;
    text: string;
    id: string; // hashid of internal id. To disallow sequential scraping or reply bots.
}