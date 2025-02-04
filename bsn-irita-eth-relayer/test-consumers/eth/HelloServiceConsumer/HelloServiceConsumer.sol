//SPDX-License-Identifier: SimPL-2.0
pragma solidity ^0.6.0;

/**
 * @title iService interface
 */
interface iServiceInterface {
    /**
     * @dev Send cross chain request
     * @param _endpointInfo information of endpoint
     * @param _method Target method name
     * @param _callData Target method callData
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return requestID Request id
     */
    function sendRequest(
        string memory _endpointInfo,
        string memory _method,
        bytes memory _callData,
        address _callbackAddress,
        bytes4 _callbackFunction
    ) external returns (bytes32 requestID);

    /**
     * @dev Set the response of the specified service request
     * @param _requestID Request id
     * @param _errMsg Error message of the service invocation
     * @param _output Response output
     * @return True on success, false otherwise
     */
    function setResponse(
        bytes32 _requestID,
        string memory _errMsg,
        string memory _output
    ) external returns (bool);
}


/*
 * @title Contract for the iService core extension client
 */
contract iServiceClient {
    iServiceInterface iServiceCore; // iService Core contract address

    // mapping the request id to Request
    mapping(bytes32 => Request) requests;

    // request
    struct Request {
        address callbackAddress; // callback contract address
        bytes4 callbackFunction; // callback function selector
        bool sent; // request sent
        bool responded; // request responded
    }

    /*
     * @dev Event triggered when the iService request is sent
     * @param _requestID Request id
     */
    event RequestSent(bytes32 _requestID);

    /*
     * @dev Make sure that the given request is valid
     * @param _requestID Request id
     */
    modifier validRequest(bytes32 _requestID) {
        require(requests[_requestID].sent, "iServiceClient: request does not exist");
        require(!requests[_requestID].responded, "iServiceClient: request has been responded");

        _;
    }

    /**
     * @dev Send cross chain request
     * @param _endpointInfo information of endpoint
     * @param _method Target method name
     * @param _callData Target method callData
     * @param _callbackAddress Callback contract address
     * @param _callbackFunction Callback function selector
     * @return requestID Request id
     */
    function sendIServiceRequest(
        string memory _endpointInfo,
        string memory _method,
        bytes memory _callData,
        address _callbackAddress,
        bytes4 _callbackFunction
    )
    internal
    returns (bytes32 requestID)
    {
        requestID = iServiceCore.sendRequest(_endpointInfo, _method, _callData, address(this), this.onResponse.selector);

        Request memory request = Request(
            _callbackAddress,
            _callbackFunction,
            true,
            false
        );

        requests[requestID] = request;

        emit RequestSent(requestID);

        return requestID;
    }

    /*
     * @dev Callback function
     * @param _requestID Request id
     * @param _output Response output
     */
    function onResponse(
        bytes32 _requestID,
        string memory _output
    )
    external
    validRequest(_requestID)
    {
        address cbAddr = requests[_requestID].callbackAddress;
        bytes4 cbFunc = requests[_requestID].callbackFunction;

        (bool success,) = cbAddr.call(abi.encodeWithSelector(cbFunc, _requestID, _output));
        if (!success) revert();
    }

    /**
     * @dev Set the iService core contract address
     * @param _iServiceCore Address of the iService core contract
     */
    function setIServiceCore(address _iServiceCore) internal {
        require(_iServiceCore != address(0), "iServiceClient: iService core address can not be zero");
        iServiceCore = iServiceInterface(_iServiceCore);
    }
}

/*
 * @title Contract for inter-chain NFT minting powered by iService
 * The service is supported by price service
 */
contract ServiceConsumer is iServiceClient {
    string private endpointInfo = "{\"dest_chain_id\":\"ropsten\",\"dest_chain_type\":\"eth\",\"endpoint_address\":\"0xd41b3Afa1B4114542Eb46cDd55EcF2FffAA24504\",\"endpoint_type\":\"contract\"}";
    event Hello(bytes32 _requestID, string _helloMsg);
    string public result;
    /*
     * @notice Constructor
     * @param _iServiceContract Address of the iService contract
     * @param _defaultTimeout Default service request timeout
     */
    constructor(
        address _iServiceCore
    ) public
    {
        setIServiceCore(_iServiceCore);
    }

    /*
     * @notice Start workflow for minting nft
     * @param _method method name
     * @param _hello arguments
     */
    function helloWorld(
        string memory _method,
        string memory _hello
    )
    external
    {
        bytes memory callData;
        callData = abi.encodePacked(bytes4(keccak256(abi.encodePacked(_method, "(string)"))), abi.encode(_hello));
        sendIServiceRequest(
            endpointInfo,
            _method,
            callData,
            address(this),
            this.onHello.selector
        );
    }

    /*
     * @notice NFT service callback function
     * @param _requestID Request id
     * @param _output NFT service response output
     */
    function onHello(
        bytes32 _requestID,
        string memory _output
    )
    external
    validRequest(_requestID)
    {
        requests[_requestID].responded = true;
        result = _output;
    }
}


