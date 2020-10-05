#!/bin/bash

echo "Build dist"
npm run generate

echo "Deploy to vercel"
vercel --prod
