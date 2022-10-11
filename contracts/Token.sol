// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "./node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Token is ERC20 {
    constructor() ERC20("MyToken", "MTK") {
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }
    function mint(address account,uint256 amount) public virtual  {
        _mint(account,amount);

    }
}
