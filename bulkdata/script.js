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

const buffer2 = fs.readFileSync('./images.json');
const images = JSON.parse(buffer2);

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

const Images = new mongoose.Schema(
  {
    serial: { type: Number, index: true, unique: true },
    image: { type: String },
  },
  { versionKey: false },
);

const ProductsModel = mongoose.model('products', Product);
const ImagesModel = mongoose.model('images', Images);

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

// *** Add brands ***
products = products.map((product) => {
  // Get random brand
  const rand = Math.floor(Math.random() * brands.length);
  return { ...product, brand: brands[rand] };
});

try {
  const res = await ProductsModel.insertMany(products);
  const res2 = await ImagesModel.insertMany(images);
  console.log(`🟩 ${res.length} videos were added successfully 🟩`);
  console.log(`🟩 ${res2.length} images were added successfully 🟩`);
} catch (err) {
  console.log('🟥 Unable to insert data 🟥');
  console.error(err);
}

process.exit();
