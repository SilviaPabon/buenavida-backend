import mongoose from 'mongoose';
import * as dotenv from 'dotenv';
import fs from 'fs';

// Load env variables
dotenv.config();

const mgUser = process.env.MONGO_USER;
const mgPassword = process.env.MONGO_PASSWORD;
const mgHost = process.env.MONGO_HOST;
const mgPort = process.env.MONGO_PORT;

// Connect to mongo
await mongoose.connect(
  `mongodb://${mgUser}:${mgPassword}@${mgHost}:${mgPort}/buenavida?authSource=admin`,
);

// Read json file
const buffer = fs.readFileSync('./products.json');
let products = JSON.parse(buffer);

// Create mongo schemas and models
const Product = new mongoose.Schema(
  {
    serial: { type: Number, index: true, unique: true },
    name: { type: String, index: true },
    image: { type: String },
    units: { type: String },
    price: { type: Number, min: 0.0, index: true },
    discount: { type: Number, min: 0.0, default: 0.0 },
    annotations: { type: String },
    description: { type: String },
  },
  { versionKey: false },
);

const ProductsModel = mongoose.model('products', Product);

// Insert data
const brands = [
  'Natura Siberica',
  'Beter',
  'Weleda',
  'Bio Cesta',
  'Dr. Goerg',
  'LOGONA',
  'Alqvimia',
  'Terra Verda',
];

products = products.map((product) => {
  // Get random brand
  const rand = Math.floor(Math.random() * brands.length);
  return { ...product, brand: brands[rand] };
});

console.log(products);

try {
  const res = await ProductsModel.insertMany(products);
  console.log(`🟩 ${res.length} were added successfully 🟩`);
} catch (err) {
  console.log('🟥 Unable to insert data 🟥');
  console.error(err);
}

process.exit();
