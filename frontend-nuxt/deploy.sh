#!/bin/bash
echo "Run build"
npm run build

echo "Run generate"
npm run generate

echo "Deploy to vercel"
vercel --prod
