const main = async () => {
    const [deployer] = await hre.ethers.getSigners();
    const accountBalance = await deployer.getBalance();

    console.log("Deploying contracts with account: ", deployer.address);
    console.log("Account balance: ", accountBalance.toString());

    const simpleContractFactory = await hre.ethers.getContractFactory("SimpleContract");
    const simpleContract = await simpleContractFactory.deploy();
    await simpleContract.deployed();
};

const runMain = async () => {
    try {
        await main();
        process.exit(0);
    } catch (error){
        process.exit(1);
    }
}


runMain();
