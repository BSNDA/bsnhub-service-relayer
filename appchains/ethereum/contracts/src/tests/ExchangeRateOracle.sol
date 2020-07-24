pragma solidity ^0.6.0;

import "../interfaces/iServiceInterface.sol";

/**
 * @title Contract for exchange rate oracle powered by iService
 */
contract ExchangeRateOracle {
    string public rate; // latest exchange rate
    
    iServiceInterface iServiceContract; // iService contract address 
    
    // iService request variables
    string serviceName; // service name
    string input; // request input
    uint256 timeout; // request timeout
    
    // mapping the request id to RequestStatus
    mapping(bytes32 => RequestStatus) requests;
    
    // request status
    struct RequestStatus {
        bool sent; // request sent
        bool responded; // request responded
    }
    
    /**
     * @title Event triggered when a request is sent
     * @param _requestID Request id
     */
    event RequestSent(bytes32 _requestID);
    
    /**
     * @title Event triggered when a request is responded
     * @param _requestID Request id
     * @param _rate Exchange rate
     */
    event RequestResponded(bytes32 _requestID, string _rate);
    
    /**
     * @title Constructor
     * @param _iServiceContract Address of the iService contract
     * @param _serviceName Service name
     * @param _input Service request input
     * @param _timeout Service request timeout
     */
    constructor(
        address _iServiceContract,
        string memory _serviceName,
        string memory _input,
        uint256 _timeout
    )
        public
    {
        iServiceContract = iServiceInterface(_iServiceContract);
        
        serviceName = _serviceName;
        input = _input;
        timeout = _timeout;
    }
    
    /**
     * @title Make sure that the given request is valid
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(requests[_requestID].sent, "ExchangeRateOracle: request does not exist");
        require(!requests[_requestID].responded, "ExchangeRateOracle: request has been responded");
        
        _;
    }
    
    /**
     * @title Send iService request
     */
    function sendRequest()
        external
    {
        bytes32 requestID = iServiceContract.callService(serviceName, input, timeout, address(this), this.onResponse.selector);
        
        emit RequestSent(requestID);
        
        requests[requestID].sent = true;
    }
    
    /**
     * @title Callback function
     * @param _requestID Request id
     * @param _output Response output
     */
    function onResponse(
        bytes32 _requestID,
        string calldata _output
    )
        external
        validRequest(_requestID)
    {
        rate = _output;
        requests[_requestID].responded = true;
        
        emit RequestResponded(_requestID, rate);
    }
}