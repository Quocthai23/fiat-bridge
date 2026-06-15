// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";

contract EnterpriseFiatToken is ERC20, AccessControl, Pausable {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant BURNER_ROLE = keccak256("BURNER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    bytes32 public constant BLACKLISTER_ROLE = keccak256("BLACKLISTER_ROLE");

    mapping(address => bool) public isBlacklisted;
    mapping(string => bool) public processedCoreTxs;

    event FiatMinted(string coreTxId, address indexed to, uint256 amount);
    event FiatBurned(string coreTxId, address indexed from, uint256 amount);
    event Blacklisted(address indexed account);
    event UnBlacklisted(address indexed account);

    constructor() ERC20("VND Token", "VNDT") {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
        _grantRole(BURNER_ROLE, msg.sender);
        _grantRole(PAUSER_ROLE, msg.sender);
        _grantRole(BLACKLISTER_ROLE, msg.sender);
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function blacklist(address account) public onlyRole(BLACKLISTER_ROLE) {
        isBlacklisted[account] = true;
        emit Blacklisted(account);
    }

    function unblacklist(address account) public onlyRole(BLACKLISTER_ROLE) {
        isBlacklisted[account] = false;
        emit UnBlacklisted(account);
    }

    function mint(string memory _coreTxId, address _to, uint256 _amount) public onlyRole(MINTER_ROLE) whenNotPaused {
        require(!processedCoreTxs[_coreTxId], "Transaction already processed");
        require(!isBlacklisted[_to], "Recipient is blacklisted");
        
        processedCoreTxs[_coreTxId] = true;
        _mint(_to, _amount);
        
        emit FiatMinted(_coreTxId, _to, _amount);
    }

    function burn(string memory _coreTxId, address _from, uint256 _amount) public onlyRole(BURNER_ROLE) whenNotPaused {
        require(!processedCoreTxs[_coreTxId], "Transaction already processed");
        require(!isBlacklisted[_from], "Sender is blacklisted");
        
        processedCoreTxs[_coreTxId] = true;
        _burn(_from, _amount);
        
        emit FiatBurned(_coreTxId, _from, _amount);
    }

    // Override required by ERC20 to apply pausable and blacklist rules on any transfer
    function _update(address from, address to, uint256 value) internal override whenNotPaused {
        require(!isBlacklisted[from], "Sender is blacklisted");
        require(!isBlacklisted[to], "Recipient is blacklisted");
        super._update(from, to, value);
    }
}
