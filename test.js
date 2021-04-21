const fs = require('fs');
const assert = require('assert');
const { hcltojson } = require('./hcltojson');

const tf = fs.readFileSync('./example.tf', 'utf-8');
const expected = fs.readFileSync('./fixture.json', 'utf-8');

assert.equal(hcltojson(tf), expected, 'Parsed Terraform does not match expected');

console.log("OK");
