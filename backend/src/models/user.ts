import { Model,DataTypes } from "sequelize";
import { db } from "../config/database.config";
import { validator } from "sequelize/types/utils/validator-extras";

interface UserAttributes{
    id:number,
    email:string,
    firstName:string,
    lastName:string,
    username:string,
    password:string


}
class UserInstance extends Model<UserAttributes>{}


UserInstance.init(
    {
        id:{
            type:DataTypes.INTEGER,
            primaryKey: true,
            autoIncrement: true,
            allowNull:false
        },
        email:{
            type:DataTypes.STRING,
            allowNull:false,
            unique: true,
            validate:{
                len:[3,255],
                isEmail:true
            }
        },
        username:{
            type:DataTypes.STRING,
            allowNull:false,
            unique: true,
            validate: {
                len: [8, 30],
                isNotEmail(value:string) {
                    if (validator.isEmail(value)) {
                        throw new Error("Cannot be an email.");
                    }
                }
            }
        },
        firstName:{
            type:DataTypes.STRING,
            allowNull:false
        },
        lastName:{
            type:DataTypes.STRING,
            allowNull:false
        },
        password:{
            type: DataTypes.STRING,
            allowNull: false,
            validate: {
                len: [60, 60]
            }
        }

    },
    {
        sequelize:db,
        tableName:'users'
    }
)
