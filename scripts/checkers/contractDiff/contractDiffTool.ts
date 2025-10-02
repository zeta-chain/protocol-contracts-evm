import axios from "axios";
import fs from "fs";
import path from "path";

interface NetworkConfig {
    name: string;
    apiUrl: string;
    apiKeyEnv?: string;
    explorerType: "etherscan" | "blockscout";
}

const NETWORK_CONFIGS: Record<string, NetworkConfig> = {
    eth_mainnet: {
        name: "Ethereum Mainnet",
        apiUrl: "https://api.etherscan.io/v2/api?chainid=1",
        apiKeyEnv: process.env.ETHERSCAN_API_KEY!,
        explorerType: "etherscan",
    },
    bsc_mainnet: {
        name: "BNB Smart Chain",
        apiUrl: "https://api.etherscan.io/v2/api?chainid=56",
        apiKeyEnv: process.env.ETHERSCAN_API_KEY!,
        explorerType: "etherscan",
    },
    polygon_mainnet: {
        name: "Polygon",
        apiUrl: "https://api.etherscan.io/v2/api?chainid=137",
        apiKeyEnv: process.env.ETHERSCAN_API_KEY!,
        explorerType: "etherscan",
    },
    base_mainnet: {
        name: "Base",
        apiUrl: "https://api.etherscan.io/v2/api?chainid=8453",
        apiKeyEnv: process.env.ETHERSCAN_API_KEY!,
        explorerType: "etherscan",
    },
    arbitrum_mainnet: {
        name: "Arbitrum One",
        apiUrl: "https://api.etherscan.io/v2/api?chainid=42161",
        apiKeyEnv: process.env.ETHERSCAN_API_KEY!,
        explorerType: "etherscan",
    },
    avalanche_mainnet: {
        name: "Avalanche C-Chain",
        apiUrl: "https://api.etherscan.io/v2/api?chainid=43114",
        apiKeyEnv: process.env.ETHERSCAN_API_KEY!,
        explorerType: "etherscan",
    },
    zeta_mainnet: {
        name: "ZetaChain Mainnet",
        apiUrl: "https://zetachain.blockscout.com/api",
        explorerType: "blockscout",
    },
}

interface ContractSource {
    SourceCode: string;
    ABI: string;
    ContractName: string;
    CompilerVersion: string;
    OptimizationUsed: string;
    Runs: string;
    ConstructorArguments: string;
    EVMVersion: string;
    Library: string;
    LicenseType: string;
    Proxy: string;
    Implementation: string;
    SwarmSource: string;
}

interface EtherscanResponse {
    status: string;
    message: string;
    result: ContractSource[];
}

interface BlockscoutFile {
    file_path: string;
    source_code: string;
}

interface BlockscoutResponse {
    is_verified: boolean;
    name: string;
    compiler_version: string;
    optimization_enabled: boolean;
    optimization_runs: string | number;
    evm_version: string;
    constructor_args: string;
    source_code: string;
    file_path?: string;
    additional_sources?: BlockscoutFile[];
    abi?: string;
}

async function fetchFromBlockscout(
    address: string,
    apiUrl: string
): Promise<ContractSource> {
    try {
        const response = await axios.get<BlockscoutResponse>(`${apiUrl}/v2/smart-contracts/${address}`);
        const data = response.data;

        if (!data.is_verified) {
            throw new Error("Contract source code is not verified");
        }

        console.log(`   ✅ Contract: ${data.name}`);
        console.log(`   📝 Compiler: ${data.compiler_version}`);

        // Convert Blockscout response to Etherscan-compatible format
        let sourceCode = data.source_code;

        // If there are additional sources, create multi-file format
        if (data.additional_sources && data.additional_sources.length > 0) {
            const sources: Record<string, { content: string }> = {};

            // Add main file
            const mainFileName = data.file_path || `${data.name}.sol`;
            sources[mainFileName] = { content: data.source_code };

            for (const file of data.additional_sources) {
                sources[file.file_path] = { content: file.source_code };
            }

            sourceCode = JSON.stringify({ sources });
        }

        return {
            SourceCode: sourceCode,
            ABI: data.abi ? JSON.stringify(data.abi) : "[]",
            ContractName: data.name,
            CompilerVersion: data.compiler_version,
            OptimizationUsed: data.optimization_enabled ? "1" : "0",
            Runs: String(data.optimization_runs || "200"),
            ConstructorArguments: data.constructor_args || "",
            EVMVersion: data.evm_version || "",
            Library: "",
            LicenseType: "",
            Proxy: "0",
            Implementation: "",
            SwarmSource: "",
        };
    } catch (error) {
        if (axios.isAxiosError(error)) {
            if (error.response?.status === 404) {
                throw new Error("Contract not found or not verified on Blockscout");
            }
            throw new Error(`Failed to fetch contract from Blockscout: ${error.message}`);
        }
        throw error;
    }
}

async function fetchFromEtherscan(
    address: string,
    apiUrl: string,
    apiKey: string
): Promise<ContractSource> {
    try {
        const response = await axios.get<EtherscanResponse>(apiUrl, {
            params: {
                module: "contract",
                action: "getsourcecode",
                address: address,
                apikey: apiKey,
            }
        });

        if (response.data.status !== "1") {
            throw new Error(`API Error: ${response.data.message}`);
        }

        if (!response.data.result || response.data.result.length === 0) {
            throw new Error("No contract source code found");
        }

        const contractData = response.data.result[0];

        if (contractData.SourceCode === "") {
            throw new Error("Contract source code is not verified");
        }

        console.log(`   ✅ Contract: ${contractData.ContractName}`);
        console.log(`   📝 Compiler: ${contractData.CompilerVersion}`);

        return contractData;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            throw new Error(`Failed to fetch contract from Etherscan: ${error.message}`);
        }
        throw error;
    }
}

async function fetchContractSource(
    address: string,
    network: string
): Promise<ContractSource> {
    const config = NETWORK_CONFIGS[network.toLowerCase()];

    if (!config) {
        throw new Error(`Unknown network: ${network}`);
    }

    if (config.explorerType === "blockscout") {
        return fetchFromBlockscout(address, config.apiUrl);
    } else {
        if (!config.apiKeyEnv) {
            throw new Error(`API key required for ${config.name}`);
        }
        return fetchFromEtherscan(address, config.apiUrl, config.apiKeyEnv);
    }
}

function extractDeclarationOrder(sourceCode: string): string[] {
    const order: string[] = [];

    // Regex to find contract, interface, library, and abstract contract declarations
    const declarationRegex = /^(?:abstract\s+)?(?:contract|interface|library)\s+(\w+)/gm;
    let match;

    while ((match = declarationRegex.exec(sourceCode)) !== null) {
        order.push(match[1]);
    }

    return order;
}

function findFileByContractName(sources: any, contractName: string): string | null {
    for (const [, fileData] of Object.entries(sources)) {
        const content = typeof fileData === "object" && fileData !== null
            ? (fileData as { content?: string }).content || ""
            : String(fileData);

        // Check if this file contains the contract/interface/library declaration
        const regex = new RegExp(`(?:abstract\\s+)?(?:contract|interface|library)\\s+${contractName}\\b`);
        if (regex.test(content)) {
            return content;
        }
    }
    return null;
}

function cleanDuplicateHeaders(code: string): string {
    const lines = code.split('\n');
    const result: string[] = [];
    let firstSpdx = true;
    let firstPragma = true;

    for (let i = 0; i < lines.length; i++) {
        const line = lines[i].trim();

        // Check for SPDX license identifier
        if (line.startsWith('// SPDX-License-Identifier:')) {
            if (firstSpdx) {
                result.push(lines[i]);
                firstSpdx = false;
            }
            continue;
        }

        // Check for pragma solidity
        if (line.startsWith('pragma solidity')) {
            if (firstPragma) {
                result.push(lines[i]);
                firstPragma = false;
            }
            continue;
        }

        result.push(lines[i]);
    }

    return result.join('\n');
}

function flattenSourceCode(sourceCode: string): string {
    if (sourceCode.startsWith("{{") || sourceCode.startsWith("{")) {
        try {
            let jsonStr = sourceCode;
            if (jsonStr.startsWith("{{")) {
                jsonStr = jsonStr.slice(1, -1);
            }

            const sourceObject = JSON.parse(jsonStr);
            const sources = sourceObject.sources || sourceObject;

            // For multi-file, just concatenate all in alphabetical order
            const sortedEntries = Object.entries(sources).sort((a, b) => a[0].localeCompare(b[0]));

            let allCode = "";
            for (const [_, fileData] of sortedEntries) {
                const content = typeof fileData === "object" && fileData !== null
                    ? (fileData as { content?: string }).content || ""
                    : String(fileData);
                allCode += content + "\n";
            }

            // Clean duplicate SPDX and pragma statements
            return cleanDuplicateHeaders(allCode);
        } catch (error) {
            return sourceCode;
        }
    }

    return sourceCode;
}

function reorderMultiFileToMatchSingleFile(
    singleFileCode: string,
    multiFileCode: string
): string {
    // Extract the order from single file
    const declarationOrder = extractDeclarationOrder(singleFileCode);

    if (!multiFileCode.startsWith("{{") && !multiFileCode.startsWith("{")) {
        return multiFileCode;
    }

    try {
        let jsonStr = multiFileCode;
        if (jsonStr.startsWith("{{")) {
            jsonStr = jsonStr.slice(1, -1);
        }

        const sourceObject = JSON.parse(jsonStr);
        const sources = sourceObject.sources || sourceObject;

        // Reconstruct in the order from single file
        let reorderedCode = "";
        const usedFiles = new Set<string>();

        for (const contractName of declarationOrder) {
            const fileContent = findFileByContractName(sources, contractName);
            if (fileContent) {
                const fileKey = Object.entries(sources).find(([_, data]) => {
                    const content = typeof data === "object" && data !== null
                        ? (data as { content?: string }).content || ""
                        : String(data);
                    return content === fileContent;
                })?.[0];

                if (fileKey && !usedFiles.has(fileKey)) {
                    reorderedCode += fileContent + "\n";
                    usedFiles.add(fileKey);
                }
            }
        }

        // Add any remaining files that weren't matched
        for (const [fileName, fileData] of Object.entries(sources)) {
            if (!usedFiles.has(fileName)) {
                const content = typeof fileData === "object" && fileData !== null
                    ? (fileData as { content?: string }).content || ""
                    : String(fileData);
                reorderedCode += content + "\n";
            }
        }

        // Clean duplicate SPDX and pragma statements
        return cleanDuplicateHeaders(reorderedCode);
    } catch (error) {
        return multiFileCode;
    }
}

function saveContractToFile(
    content: string,
    fileName: string,
    outputDir: string,
): string {
    if (!fs.existsSync(outputDir)) {
        fs.mkdirSync(outputDir, { recursive: true });
    }

    const filePath = path.join(outputDir, fileName);
    fs.writeFileSync(filePath, content, "utf8");

    return filePath;
}

export async function fetchAndFlattenContract(
    oldAddress: string,
    newAddress: string,
    network: string,
) {
    const OUTPUT_DIR = "./contract-diffs"
    try {
        console.log("\n📥 Fetching OLD implementation...");
        const oldContract = await fetchContractSource(oldAddress, network);

        console.log("\n📥 Fetching NEW implementation...");
        const newContract = await fetchContractSource(newAddress, network);

        // Determine which is single-file and which is multi-file
        const oldIsMultiFile = oldContract.SourceCode.startsWith("{");
        const newIsMultiFile = newContract.SourceCode.startsWith("{");

        let oldFlattened: string;
        let newFlattened: string;

        if (oldIsMultiFile && !newIsMultiFile) {
            oldFlattened = reorderMultiFileToMatchSingleFile(newContract.SourceCode, oldContract.SourceCode);
            newFlattened = newContract.SourceCode;
        } else if (!oldIsMultiFile && newIsMultiFile) {
            oldFlattened = oldContract.SourceCode;
            newFlattened = reorderMultiFileToMatchSingleFile(oldContract.SourceCode, newContract.SourceCode);
        } else {
            oldFlattened = flattenSourceCode(oldContract.SourceCode);
            newFlattened = flattenSourceCode(newContract.SourceCode);
        }

        console.log("\n💾 Saving flattened contracts...");
        const timestamp = new Date().toISOString().replace(/[:.]/g, "-").split("T")[0];
        const oldFile = `${timestamp}_OLD_${oldContract.ContractName}_${oldAddress.slice(0, 8)}.sol`;
        const newFile = `${timestamp}_NEW_${newContract.ContractName}_${newAddress.slice(0, 8)}.sol`;

        const oldPath = saveContractToFile(oldFlattened, oldFile, OUTPUT_DIR);
        const newPath = saveContractToFile(newFlattened, newFile, OUTPUT_DIR);

        console.log(`   ✅ Old contract saved: ${oldPath}`);
        console.log(`   ✅ New contract saved: ${newPath}`);

        console.log("\n" + "=".repeat(60));
        console.log("📊 COMPARISON SUMMARY");
        console.log("=".repeat(60));
        console.log(`Old Contract: ${oldContract.ContractName}`);
        console.log(`  Address: ${oldAddress}`);
        console.log(`  Compiler: ${oldContract.CompilerVersion}`);
        console.log(`  File: ${oldFile}`);
        console.log("");
        console.log(`New Contract: ${newContract.ContractName}`);
        console.log(`  Address: ${newAddress}`);
        console.log(`  Compiler: ${newContract.CompilerVersion}`);
        console.log(`  File: ${newFile}`);
        console.log("=".repeat(60));
    } catch (error) {
        console.error("\n❌ Error:", error instanceof Error ? error.message : error);
        throw error;
    }
}
