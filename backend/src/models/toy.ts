'use strict';
import {
  Model
} from 'sequelize';

interface ToyAttributes{
  id:number,
  name:string,
  productType:string,
  price:number,
  theme:string,
  releaseDate:Date,
  count:number,
  available:boolean
}

module.exports = (sequelize:any, DataTypes:any) => {
  class Toy extends Model<ToyAttributes>
  implements  ToyAttributes{
    id!:number;
    name!:string;
    productType!:string;
    price!:number;
    theme!:string;
    releaseDate!:Date;
    count!:number;
    available!:boolean;
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models:any) {
      // define association here
      Toy.belongsToMany(models.User,{
        through:'wishlists'
      })
    }
  }
  Toy.init({
    id:{
      type:DataTypes.INTEGER,
      allowNull:false,
      primaryKey:true,
      autoIncrement:true
    },
    name:{
      type:DataTypes.STRING,
      allowNull:false
    },
    productType:{
      type:DataTypes.STRING,
      allowNull:false
    },
    price: {
      type:DataTypes.FLOAT,
      allowNull:false,
      validate:{
        isNumeric:true,
        notEmpty:true,
        isDecimal:true,
      }
    },
    theme: {
      type:DataTypes.STRING,
      allowNull:false
    },
    releaseDate: {
      type:DataTypes.DATE,
      allowNull:false,
      defaultValue: new Date(),
      validate:{
        notEmpty:true,
        isDate:true
      }
    },
    count: {
      type:DataTypes.INTEGER,
      allowNull:false,
      defaultValue:0,
      validate: {
        isNumeric:true,
        notEmpty:true
      }
    },
    available: {
      type:DataTypes.BOOLEAN,
      allowNull:false,
      defaultValue:false
    },
  }, {
    sequelize,
    modelName: 'Toy',
  });
  return Toy;
};
