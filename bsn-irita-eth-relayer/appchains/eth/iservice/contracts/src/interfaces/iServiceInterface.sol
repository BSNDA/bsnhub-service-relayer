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
