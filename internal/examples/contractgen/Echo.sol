pragma solidity ^0.8.0;

contract Echo {
    event Echoed(string message);

    function echo(string calldata message) external returns (string memory) {
        emit Echoed(message);
        return message;
    }
}