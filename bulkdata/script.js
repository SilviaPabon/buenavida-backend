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
await mongoose.connect(`mongodb://${mgUser}:${mgPassword}@${mgHost}:${mgPort}`);

// Read json file
const buffer = fs.readFileSync('./products.json');
const products = JSON.parse(buffer);
console.log(products.length)
