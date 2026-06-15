import fs from 'fs';
import path from 'path';

const artifactPath = './artifacts/contracts/EnterpriseFiatToken.sol/EnterpriseFiatToken.json';
if (!fs.existsSync(artifactPath)) {
    console.error("Artifact not found. Compile first.");
    process.exit(1);
}

const artifact = JSON.parse(fs.readFileSync(artifactPath, 'utf8'));

// write to root of the project
fs.writeFileSync('../EnterpriseFiatToken.abi', JSON.stringify(artifact.abi));
fs.writeFileSync('../EnterpriseFiatToken.bin', artifact.bytecode.replace(/^0x/, ''));
console.log("Extracted ABI and bytecode to root directory.");
