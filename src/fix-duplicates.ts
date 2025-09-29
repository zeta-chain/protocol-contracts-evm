import * as fs from 'fs';
import * as path from 'path';

// get the project root directory
const root = process.cwd();
// path to index.ts file
const indexPath = path.join(root, 'types', 'index.ts');
// read the file
let content = fs.readFileSync(indexPath, 'utf8');

// track exports, we've seen to avoid duplicates
const seenExports = new Set<string>();

// process the file line by line
const lines = content.split('\n');
const filteredLines = lines.filter(line => {
  // check if line is an export type statement
  const match = line.match(/export type \{ ([^}]+) \} from/);
  if (!match) return true;

  const exportName = match[1];

  // if we've seen this export before, filter it out
  if (seenExports.has(exportName)) {
    console.log(`Removing duplicate export: ${exportName}`);
    return false;
  }

  // otherwise, track it and keep the line
  seenExports.add(exportName);
  return true;
});

// write the filtered content back to the file
fs.writeFileSync(indexPath, filteredLines.join('\n'));
console.log('Duplicate exports have been removed from index.ts');