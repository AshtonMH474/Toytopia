import db from '../models'


const toys = [
    {
        productType: "Action Figure",
        name:"Optimus Prime G1 1984",
        price: 300.00,
        theme: "Transformers",
        releaseDate: "1984-09-17",
        count: 3,
        available: true
      },
      {
        productType: "Action Figure",
        name:"Megatron G1 1984",
        price: 300.00,
        theme: "Transformers",
        releaseDate: "1984-09-17",
        count: 80,
        available: true
      },
      {
        productType: "Action Figure",
        name:"Bumblebee G1 1984",
        price: 80.00,
        theme: "Transformers",
        releaseDate: "1984-09-17",
        count: 1,
        available: true
      },
      {
        productType: "Action Figure",
        name:"Starscream G1 1984",
        price: 100.00,
        theme: "Transformers",
        releaseDate: "1984-09-17",
        count: 0,
        available: false
      },
      {
        productType: "Action Figure",
        name:"Soundwave G1 1984",
        price: 300.00,
        theme: "Transformers",
        releaseDate: "1984-09-17",
        count: 2,
        available: true
      },
      {
        productType: "Building Set",
        name: "LEGO Star Wars Millennium Falcon",
        price: 799.99,
        theme: "Star Wars",
        releaseDate: "2017-10-01",
        count: 5,
        available: true
      },
      {
        productType: "Building Set",
        name: "LEGO Harry Potter Hogwarts Castle",
        price: 469.99,
        theme: "Harry Potter",
        releaseDate: "2018-09-01",
        count: 10,
        available: true
      },
      {
        productType: "Building Set",
        name: "LEGO Ideas Tree House",
        price: 249.99,
        theme: "LEGO Ideas",
        releaseDate: "2019-08-01",
        count: 3,
        available: true
      },
      {
        productType: "Building Set",
        name: "LEGO Technic Bugatti Chiron",
        price: 349.99,
        theme: "Technic",
        releaseDate: "2018-06-01",
        count: 0,
        available: false
      },
      {
        productType: "Building Set",
        name: "LEGO Creator Expert Roller Coaster",
        price: 379.99,
        theme: "Creator Expert",
        releaseDate: "2018-06-01",
        count: 8,
        available: true
      },
      {
        productType: "Die-Cast Car",
        name: "Hot Wheels Twin Mill",
        price: 4.99,
        theme: "Hot Wheels",
        releaseDate: "1969-01-01",
        count: 20,
        available: true
      },
      {
        productType: "RC Car",
        name: "Traxxas Slash 4x4",
        price: 299.99,
        theme: "Remote Control",
        releaseDate: "2010-03-15",
        count: 5,
        available: true
      },
      {
        productType: "Die-Cast Car",
        name: "Matchbox Ford Mustang GT",
        price: 5.99,
        theme: "Matchbox",
        releaseDate: "2020-07-10",
        count: 50,
        available: true
      },
      {
        productType: "Toy Car Set",
        name: "Hot Wheels 20-Car Gift Pack",
        price: 21.99,
        theme: "Hot Wheels",
        releaseDate: "2018-11-01",
        count: 15,
        available: true
      },
      {
        productType: "Die-Cast Car",
        name: "Tomica Nissan GT-R",
        price: 7.99,
        theme: "Tomica",
        releaseDate: "2017-04-25",
        count: 0,
        available: false
      }

]


export const createToys = () => {
    toys.map((toy) => {
        db.Toy.create(toy)
    })
}
