const main = async () => {
    const [owner, randomPerson] = await hre.ethers.getSigners();
    const contractFactory = await hre.ethers.getContractFactory("SimpleContract");
    const simpleContract = await contractFactory.deploy();
    await simpleContract.deployed();

    console.log("Contract deployed to:", simpleContract.address);
    console.log("Contract deployed by:", owner.address);

    let messageCount;
    messageCount = await simpleContract.getTotalMessages();

    let msgTxn = await simpleContract.message();
    await msgTxn.wait();

    msgTxn = await simpleContract.connect(randomPerson).message();
    await msgTxn.wait();

    messageCount = await simpleContract.getTotalMessages();
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
