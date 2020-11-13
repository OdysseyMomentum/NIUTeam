pragma solidity ^0.6.0;

import "https://github.com/OpenZeppelin/openzeppelin-solidity/contracts/access/Ownable.sol";

contract FiledgrDocument is Ownable {
    
    struct SharedDocument {
        string documentHash;
        string docSender;
        uint sizeOfMapping;
        mapping(string => uint32) receivers;
    }
    
    SharedDocument sharedDocument;
    
    constructor(string memory docSender, string memory docHash) public {
        sharedDocument.docSender= docSender;
        sharedDocument.documentHash = docHash;
    }
    
    function getDocumentHash() onlyOwner public view returns (string memory) {
        return sharedDocument.documentHash;
    }
    
    function shareDocument(string memory receiver, uint32 timestamp) public {
        sharedDocument.receivers[receiver] = timestamp;
        sharedDocument.sizeOfMapping++;
    }
    
    function getSender() public view returns (string memory) {
        return sharedDocument.docSender;
    }
    
    function getSharedAt(string memory receiver) public view returns (uint32 timestamp) {
        return sharedDocument.receivers[receiver];
    }
    
}