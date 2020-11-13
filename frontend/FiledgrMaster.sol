pragma solidity ^0.6.0;

import "./FiledgrDocument.sol";

contract FiledgrMaster {
    
    struct DocumentMapping {
        uint sizeOfMapping;
        mapping(string => address) documentContracts;
    }
    
    DocumentMapping documents;
    
    constructor() public {
        documents.sizeOfMapping = 0;
    }
    
    function addDocument( string memory docId, string memory docHash, string memory docSender ) public returns (address) {
        FiledgrDocument filedgrDocument = new FiledgrDocument(docSender, docHash);
        documents.documentContracts[docId] = address(filedgrDocument);
        documents.sizeOfMapping++;
        return documents.documentContracts[docId];
    }
    
    function getDocumentHash(string memory docId) public view returns (string memory) {
        string memory docHash = FiledgrDocument(documents.documentContracts[docId]).getDocumentHash();
        return docHash;
    }
    
    function getDocument( string memory docId )  public view returns (address) {
        return address(documents.documentContracts[docId]);
    }
    
    function getSizeOfMapping() public view returns (uint) {
        return documents.sizeOfMapping;
    }
}

